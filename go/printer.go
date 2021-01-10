package main

import (
	"periph.io/x/host/v3"
)

type PrinterOptions struct {
	LED    LEDOptions
	Button ButtonOptions
}

type Printer struct {
	LED    *LEDPin
	Button *ButtonPin
}

func NewPrinter(opts PrinterOptions) (*Printer, error) {
	_, err := host.Init()
	if err != nil {
		return nil, err
	}
	return &Printer{
		LED:    NewLEDPin(opts.LED),
		Button: NewButtonPin(opts.Button),
	}, nil
}

func (p *Printer) Shutdown() error {
	var err error
	err = p.LED.Off()
	if err != nil {
		return err
	}
	return nil
}
