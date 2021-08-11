package relay

import (
	"fmt"

	"github.com/zserge/hid"
)

type Relay struct {
	hid.Device

	DevNum int
	Port   string

	Product      string
	Manufacturer string
}

func (r *Relay) String() string {
	return fmt.Sprintf("%s %s devnum %d", r.Product, r.Manufacturer, r.DevNum)
}

func (r *Relay) State() ([]bool, error) {
	for i := 0; i < 256; i++ {
		report, err := r.GetReport(i)
		if err != nil {
			return nil, fmt.Errorf("unable to get report: %w", err)
		}

		fmt.Printf("report: (%d) %+v\n", i, report)

	}

	return nil, nil
}

func (r *Relay) On() error {
	payload := []byte{0xFF, 1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	return r.SetReport(0, payload)
}

func (r *Relay) Off() error {
	payload := []byte{0xFD, 1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	return r.SetReport(0, payload)
}
