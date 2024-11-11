package upower

import (
	"errors"
	"fmt"
	"log"
	"math"

	godbus "github.com/godbus/dbus/v5"

	"swaystats/internal/block"
	"swaystats/internal/dbus"
)

type UPower struct{}
type Devices map[godbus.ObjectPath]DeviceProperties

var DeviceIconBattery = []string{"󰁺", "󰁻", "󰁼", "󰁽", "󰁾", "󰁿", "󰂀", "󰂁", "󰂂", "󰁹"}
var DeviceIconCharging = "󱐋"
var DeviceIconOnline = "󰚥"

func New() (UPower, error) {
	return UPower{}, nil
}

func (u UPower) Update(c chan block.Update, n string) {
	devices := make(Devices)

	paths, err := EnumerateDevices()
	if err != nil {
		log.Println(err)
		return
	}

	for _, p := range paths {
		props, err := InspectDevice(p)
		if err != nil {
			log.Println(err)
			return
		}
		devices[p] = props
	}

	if err := dbus.Conn.AddMatchSignal(
		godbus.WithMatchSender("org.freedesktop.UPower"),
	); err != nil {
		panic(err)
	}

	signal := make(chan *godbus.Signal, 10)
	dbus.Conn.Signal(signal)

	for s := range signal {
		fmt.Printf("%s \n", string(s.Path))

		Render(devices)
	}
}

func Render(devices Devices) {
	var icon string
	var percentage int

	online := false

	for _, d := range devices {
		switch dt := d.Type; dt {
		case DeviceTypeLinePower:
			if d.Online {
				online = true
			}
		case DeviceTypeBattery:
			percentage = int(math.Abs(d.Percentage))
			if d.State == DeviceStateCharging {
				icon = DeviceIconCharging
			} else if !online {
				icon = DeviceIconBattery[int(math.Ceil(d.Percentage/10)-1)]
			}
		}
	}

	fmt.Printf("%s %v%%", icon, percentage)
}

func EnumerateDevices() ([]godbus.ObjectPath, error) {
	var paths []godbus.ObjectPath

	obj := dbus.Conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower")
	err := obj.Call("org.freedesktop.UPower.EnumerateDevices", 0).Store(&paths)

	if err != nil {
		return []godbus.ObjectPath{}, errors.New(fmt.Sprintf("error enumerating UPower devices: %s", err))
	}

	return paths, nil
}

func InspectDevice(p godbus.ObjectPath) (DeviceProperties, error) {
	var props DeviceProperties
	var err error
	var ok bool

	obj := dbus.Conn.Object("org.freedesktop.UPower", p)

	dt, err := obj.GetProperty("org.freedesktop.UPower.Device.Type")
	ds, err := obj.GetProperty("org.freedesktop.UPower.Device.State")
	dp, err := obj.GetProperty("org.freedesktop.UPower.Device.Percentage")
	do, err := obj.GetProperty("org.freedesktop.UPower.Device.Online")

	if err != nil {
		return DeviceProperties{}, err
	}

	props.Type, ok = dt.Value().(uint32)

	if !ok {
		err = errors.New(fmt.Sprintf("type assertion error on UPower Type: %v, %v \n", props.Type, ok))
		return DeviceProperties{}, err
	}

	props.State, ok = ds.Value().(uint32)

	if !ok {
		err = errors.New(fmt.Sprintf("type assertion error on UPower State: %v, %v \n", props.State, ok))
		return DeviceProperties{}, err
	}

	props.Percentage, ok = dp.Value().(float64)

	if !ok {
		err = errors.New(fmt.Sprintf("type assertion error on UPower Percentage: %v, %v \n", props.Percentage, ok))
		return DeviceProperties{}, err
	}

	props.Online, ok = do.Value().(bool)

	if !ok {
		err = errors.New(fmt.Sprintf("type assertion error on UPower Online: %v, %v \n", props.Online, ok))
		return DeviceProperties{}, err
	}

	return props, nil
}
