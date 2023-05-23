package qls_apa102c

import (
	"tinygo.org/x/drivers"
)

// Device wraps an I2C connection to a Sparkfun Qwiic LED Stick - APA102C.
type Device struct {
	bus     drivers.I2C
	Address uint8
}

// New creates a new LED Stick connection. The I2C bus must already be configured.
// This function only creates the Device object, it does not touch the device.
func New(bus drivers.I2C) Device {
	return Device{bus: bus, Address: ADDRESS}
}

// Set the color of all leds to the same RGB color.
func (d *Device) SetAllLedColor(red, green, blue uint8) bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_ALL_LED_COLOR, red, green, blue}, make([]byte, 1))
	return true
}

// Set the color of a single LED to an RGB color. LEDs are indexed starting at 0.
func (d *Device) SetLedColor(led_number, red, green, blue uint8) bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_SINGLE_LED_COLOR, led_number, red, green, blue}, make([]byte, 1))
	return true
}

// Change the brightness of all LEDs, while keeping their current color.
// Brightness must be a value between 0-31.
// To turn all LEDs off but remember their previous color, set brightness to 0.
func (d *Device) SetAllLedBrightness(brightness uint8) bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_ALL_LED_BRIGHTNESS, brightness}, make([]byte, 1))
	return true
}

// Set the brightness of a single LED. LEDs are indexed starting at 0.
// Brightness must be a value between 0-31.
func (d *Device) SetLedBrightness(led_number, brightness uint8) bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_SINGLE_LED_BRIGHTNESS, led_number, brightness}, make([]byte, 1))
	return true
}

// Turn off all LEDs.
func (d *Device) AllLedOff() bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_ALL_LED_COLOR, 0, 0, 0}, make([]byte, 1))
	return true
}
