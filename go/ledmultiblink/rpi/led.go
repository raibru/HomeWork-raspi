package rpi

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// ApplyMultiBlink let raspberry pi blinking
func ApplyMultiBlink(i uint, n uint) error {
	fmt.Printf("Apply Multi Led Blink of Raspberry PI with parameter: interval=%d ntimes=%d\n\n", i, n)

	err := rpio.Open()
	if err != nil {
		fmt.Printf("unable to open gpio %s", err.Error())
		return err
	}

	defer rpio.Close()

	pin17 := rpio.Pin(17)
	pin17.Output()
	pin17.Low()

	pin18 := rpio.Pin(18)
	pin18.Output()
	pin18.Low()

	for x := uint(0); x < n; x++ {
		fmt.Printf("%d rpio high pin(17)...\n", x)
		pin17.Toggle()
		time.Sleep(time.Millisecond * time.Duration(i))
		pin17.Toggle()
		time.Sleep(time.Millisecond * time.Duration(i))

		fmt.Printf("%d rpio high pin(18)...\n", x)
		pin18.Toggle()
		time.Sleep(time.Millisecond * time.Duration(i))
		pin18.Toggle()
		time.Sleep(time.Millisecond * time.Duration(i))
	}

	pin18.Low()
	return nil
}
