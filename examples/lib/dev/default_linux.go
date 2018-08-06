package dev

import (
	"github.com/duncanfinney/ble"
	"github.com/duncanfinney/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
