package dbus

import (
	"log"

	"github.com/godbus/dbus/v5"
)

var Conn *dbus.Conn

func Connect() {
	var err error

	Conn, err = dbus.ConnectSystemBus()

	if err != nil {
		log.Fatalf("error connecting to dbus: %s", err)
	}
}

func Close() {
	Conn.Close()
}
