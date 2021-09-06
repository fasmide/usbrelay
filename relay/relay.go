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
	report, err := r.GetReport(0)
	if err != nil {
		return nil, fmt.Errorf("unable to get report: %w", err)
	}

	fmt.Printf("report: (%d) %+v (len %d)\n", 0, report, len(report))

	return nil, nil
}

func (r *Relay) On(i byte) error {
	payload := []byte{0xFF, i, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	return r.SetReport(0, payload)
}

func (r *Relay) Off(i byte) error {
	payload := []byte{0xFD, i, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	return r.SetReport(0, payload)
}

func (r *Relay) AllOn() error {
	payload := []byte{0xFE, 0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	return r.SetReport(0, payload)
}

func (r *Relay) AllOff() error {
	payload := []byte{0xFC, 0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	return r.SetReport(0, payload)
}
