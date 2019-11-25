package rpi

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// ApplyBlink let raspberry pi blinking
func ApplyBlink(i uint, n uint) error {
	fmt.Printf("Apply Blink of Raspberry PI with parameter: interval=%d ntimes=%d\n\n", i, n)

	err := rpio.Open()
	if err != nil {
		fmt.Printf("unable to open gpio %s", err.Error())
		return err
	}

	defer rpio.Close()

	pin := rpio.Pin(18)
	pin.Output()
	pin.Low()

	for x := uint(0); x < n; x++ {
		fmt.Printf("%d rpio high pin(18)...\n", x)
		pin.Toggle()
		time.Sleep(time.Millisecond * time.Duration(i))
		pin.Toggle()
		time.Sleep(time.Millisecond * time.Duration(i))
	}

	pin.Low()
	return nil
}
