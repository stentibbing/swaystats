package main

import (
	// "fmt"
	"log"

	"swaystats/internal/dbus"
	// "swaystats/internal/header"
	"swaystats/internal/statusline"

	// "swaystats/internal/modules/time"
	"swaystats/internal/modules/upower"
	// "swaystats/internal/modules/volume"
)

func main() {
	dbus.Connect()
	defer dbus.Close()

	// h := header.New(header.Header{
	// 	Version: 1,
	// })

	// fmt.Print(string(h))

	sl := statusline.New()

	// register UPower block
	upower, err := upower.New()
	if err != nil {
		log.Println(err)
	} else {
		sl.Register(upower, "upower")
	}

	// sl.Register(time.New(), "time")

	sl.Render()
}
