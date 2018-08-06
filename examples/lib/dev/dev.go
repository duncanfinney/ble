package dev

import "github.com/duncanfinney/ble"

// NewDevice ...
func NewDevice(impl string) (d ble.Device, err error) {
	return DefaultDevice()
}
