package qls_apa102c

import (
	"tinygo.org/x/drivers"
)

const (
	// Constants/addresses used for I2C.
	ADDRESS                     = 0x23
	CHANGE_ADDRESS              = 0xC7
	CHANGE_LED_LENGTH           = 0x70
	WRITE_SINGLE_LED_COLOR      = 0x71
	WRITE_ALL_LED_COLOR         = 0x72
	WRITE_RED_ARRAY             = 0x73
	WRITE_GREEN_ARRAY           = 0x74
	WRITE_BLUE_ARRAY            = 0x75
	WRITE_SINGLE_LED_BRIGHTNESS = 0x76
	WRITE_ALL_LED_BRIGHTNESS    = 0x77
	WRITE_ALL_LED_OFF           = 0x78
)

// Device wraps an I2C connection to a Sparkfun Qwiic LED Stick - APA102C.
type Device struct {
	bus     drivers.I2C
	Address uint8
}

// New creates a new LED Stick connection. The I2C bus must already be
// configured.
//
// This function only creates the Device object, it does not touch the device.
func New(bus drivers.I2C) Device {
	return Device{bus: bus, Address: ADDRESS}
}

func (d *Device) SetAllLedColor(red, green, blue uint8) bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_ALL_LED_COLOR, red, green, blue}, make([]byte, 1))
	return true
}

// set the color of a single led...indexed starting at 0
func (d *Device) SetLedColor(led_number, red, green, blue uint8) bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_SINGLE_LED_COLOR, led_number, red, green, blue}, make([]byte, 1))
	return true
}

// //Change the brightness of all LEDs, while keeping their current color
// //brightness must be a value between 0-31
// //To turn all LEDs off but remember their previous color, set brightness to 0
func (d *Device) SetAllLedBrightness(brightness uint) bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_ALL_LED_BRIGHTNESS, byte(brightness)}, make([]byte, 1))
	return true
}

func (d *Device) AllLedOff() bool {
	d.bus.Tx(ADDRESS, []byte{WRITE_ALL_LED_COLOR, 0, 0, 0}, make([]byte, 1))
	return true
}
