package qls_apa102c

// Constants/addresses used for I2C.
const (
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
