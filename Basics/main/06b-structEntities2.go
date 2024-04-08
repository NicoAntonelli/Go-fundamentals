package main

import "fmt"

// General interface
type Engine interface {
	kilometersLeft() uint16
}

// Gas Engine
type GasEngine struct {
	model      string
	liters     uint8
	kmPerLiter uint8
}

func (g GasEngine) kilometersLeft() uint16 {
	return uint16(g.liters) * uint16(g.kmPerLiter)
}

// Electric engine
type ElectricEngine struct {
	model    string
	kwh      uint8
	kmPerKwh uint8
}

func (e ElectricEngine) kilometersLeft() uint16 {
	return uint16(e.kwh) * uint16(e.kmPerKwh)
}

// General function
func canMakeIt(e Engine, kilometers uint16) {
	if (uint16(e.kilometersLeft()) >= kilometers) {
		fmt.Println("Can make it!")
	} else {
		fmt.Println("Definitely can't make it...")
	}
}
