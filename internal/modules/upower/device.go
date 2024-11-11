package upower

const (
	DeviceTypeUnknown = iota
	DeviceTypeLinePower
	DeviceTypeBattery
	DeviceTypeUps
	DeviceTypeMonitor
	DeviceTypeMouse
	DeviceTypeKeyboard
	DeviceTypePda
	DeviceTypePhone
	DeviceTypeMediaPlayer
	DeviceTypeTablet
	DeviceTypeComputer
	DeviceTypeGamingInput
	DeviceTypeGamingPen
	DeviceTypeGamingTouchpad
	DeviceTypeGamingModem
	DeviceTypeGamingNetwork
	DeviceTypeGamingHeadset
	DeviceTypeGamingSpeakers
	DeviceTypeGamingHeadphones
	DeviceTypeGamingVideo
	DeviceTypeGamingOtherAudio
	DeviceTypeGamingRemoteControl
	DeviceTypeGamingPrinter
	DeviceTypeGamingScanner
	DeviceTypeGamingCamera
	DeviceTypeGamingWearable
	DeviceTypeGamingToy
	DeviceTypeGamingBluetoothGenreic
)

const (
	DeviceStateUnknown = iota
	DeviceStateCharging
	DeviceStateDischarging
	DeviceStateEmpty
	DeviceStateFullycharged
	DeviceStatePendingcharge
	DeviceStatePendingdischarge
)

type DeviceProperties struct {
	NativePath               string
	Vendor                   string
	Model                    string
	Serial                   string
	UpdateTime               uint64
	Type                     uint32
	PowerSupply              bool
	HasHistory               bool
	HasStatistics            bool
	Online                   bool
	Energy                   float64
	EnergyEmpty              float64
	EnergyFull               float64
	EnergyFullDesign         float64
	EnergyRate               float64
	Voltage                  float64
	ChargeCycles             int32
	Luminosity               float64
	TimeToEmpty              int64
	TimeToFull               int64
	Percentage               float64
	Temperature              float64
	IsPresent                bool
	State                    uint32
	IsRechargeable           bool
	Capacity                 float64
	Technology               uint32
	WarningLevel             uint32
	BatteryLevel             uint32
	IconName                 string
	ChargeStartThreshold     uint32
	ChargeEndThreshold       uint32
	ChargeThresholdEnabled   bool
	ChargeThresholdSupported bool
}
