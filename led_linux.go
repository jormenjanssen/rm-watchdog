package main

import "errors"

var gpioMapping map[ManagerGpio]Pin
var errGpioNotInitialized = errors.New("gpio not inialized")

func LedOn(pin ManagerGpio) {
	gpioFunc(pin, func(p) {
		return p.High()
	})
}

func LedOff(pin ManagerGpio) {
	gpioFunc(pin, func(p) {
		return p.Low()
	})
}

func gpioFunc(gpio ManagerGpio, fn func(Pin) error) error {

	// First look if we're setup.
	err := setup()
	if err != nil {
		return err
	}

	// Try to get our pin value.
	pin, err := GetPin(gpio)

	if err != nil {
		return err
	}

	// Execute our closure.
	return fn(pin)
}

func setup() error {

	if !IsTargetDevice() {
		return errGpioNotInitialized
	}

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
	pin, err := GetPin(gpio)

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
