package dev

import (
	"github.com/duncanfinney/ble"
	"github.com/duncanfinney/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
