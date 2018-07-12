//+build linux

package main

import "errors"

var gpioMapping map[ManagerGpio]Pin
var errGpioNotInitialized = errors.New("gpio not inialized")

// LedOn turns on led
func LedOn(pin ManagerGpio) error {
	return gpioFunc(pin, func(p Pin) error {
		return p.High()
	})
}

// LedOff turns off led
func LedOff(pin ManagerGpio) error {
	return gpioFunc(pin, func(p Pin) error {
		return p.Low()
	})
}

func setup() error {

	if gpioMapping != nil {
		return nil
	}

	gpioMapping = map[ManagerGpio]Pin{
		LedUpsGreen: NewOutput(uint(LedUpsGreen), true),
		LedUpsRed:   NewOutput(uint(LedUpsRed), true),
	}

	return nil
}

func gpioFunc(gpio ManagerGpio, fn func(Pin) error) error {

	// First look if we're setup.
	err := setup()
	if err != nil {
		return err
	}

	// Try to get our pin value.
	pin, err := getPin(gpio)

	if err != nil {
		return err
	}

	// Execute our closure.
	return fn(pin)
}

func getPin(managerGpio ManagerGpio) (Pin, error) {

	if gpioMapping == nil {
		return Pin{}, errors.New("gpio mapping not initialized")
	}

	elem, ok := gpioMapping[managerGpio]

	if !ok {
		return elem, errors.New("gpio does not exsist")
	}

	return elem, nil
}
