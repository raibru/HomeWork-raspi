package rpi

import (
	"os"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	ledPinRed   = rpio.Pin(19)
	ledPinGreen = rpio.Pin(18)
	ledPinBlue  = rpio.Pin(13)
	freq        = 100000
	cycle       = 1024
)

// ApplyLedRgb set led in color red, green and blue
func ApplyLedRgb(r, g, b uint32) error {
	if err := rpio.Open(); err != nil {
		os.Exit(1)
	}
	defer rpio.Close()

	ledInit()

	//
	// Set LED color
	//
	ledColorSet(r, g, b)

	return nil
}

func ledInit() {
	ledPinRed.Mode(rpio.Pwm)
	ledPinRed.Freq(freq)
	ledPinRed.DutyCycle(0, uint32(cycle))

	ledPinGreen.Mode(rpio.Pwm)
	ledPinGreen.Freq(freq)
	ledPinGreen.DutyCycle(1, uint32(cycle))

	ledPinBlue.Mode(rpio.Pwm)
	ledPinBlue.Freq(freq)
	ledPinBlue.DutyCycle(1023, uint32(cycle))
}

func ledColorSet(redVal, greenVal, blueVal uint32) {
	// This doesn't work as expected because GPIO pins 19 & 13 are essentially linked. This
	// is also true for GPIO pins 18 & 12, but since pin 12 isn't used here pin 18, green,
	// works as expected. So, when a blueVal or redVal are provided, the same value
	// is propagated to the linked pin. So for example, if redVal is 0 and blueVal
	// is 255, both the red and blue leds will light up (purple).
	//
	// (un)comment these lines if you want to see this behavior.
	ledPinRed.Mode(rpio.Pwm)
	ledPinGreen.Mode(rpio.Pwm)
	ledPinBlue.Mode(rpio.Pwm)

	ledPinRed.DutyCycle(redVal, uint32(cycle))
	ledPinGreen.DutyCycle(greenVal, uint32(cycle))
	ledPinBlue.DutyCycle(blueVal, uint32(cycle))

	// Given the explanation above, a workaround is to set only accept a redVal of
	// 255. When this is the case the blue and green pins are changed to output mode
	// and turned off, resulting in only a red LED. If redVal isn't 255, then it's
	// disabled (like blue and green above) and the blue/green pins are reenabled
	// and set to their respective values. This is obviously a hack since the red
	// pin can only be set to 255 and can't be used in conjuction with the blue and
	// green pins.
	//
	// (un)comment these lines if you don't want to see this behavior.
	//	if redVal == 255 {
	//		ledPinRed.Mode(rpio.Pwm)
	//		ledPinRed.DutyCycle(redVal, cycle)
	//
	//		//ledPinGreen.Mode(rpio.Output)
	//		ledPinGreen.DutyCycle(greenVal, cycl)
	//		ledPinBlue.Mode(rpio.Output)
	//		ledPinBlue.Low()
	//	} else {
	//		ledPinRed.Mode(rpio.Output)
	//		ledPinRed.Low()
	//
	//		ledPinGreen.Mode(rpio.Pwm)
	//		ledPinBlue.Mode(rpio.Pwm)
	//		ledPinGreen.DutyCycle(greenVal, cycle)
	//		ledPinBlue.DutyCycle(blueVal, cycle)
	//	}
}
