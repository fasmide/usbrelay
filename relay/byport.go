package relay

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/zserge/hid"
)

func ByPort() (map[string]*Relay, error) {
	relays := make(map[string]*Relay)

	err := filepath.Walk("/sys/bus/usb/devices", func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		product, err := ioutil.ReadFile(path.Join(p, "idProduct"))
		if err != nil {
			return nil
		}
		if string(product[0:4]) != "05df" {
			return nil
		}

		vendor, err := ioutil.ReadFile(path.Join(p, "idVendor"))
		if err != nil {
			return nil
		}
		if string(vendor[0:4]) != "16c0" {
			return nil
		}

		deviceID, err := ioutil.ReadFile(path.Join(p, "devnum"))
		if err != nil {
			return err
		}

		intID, err := strconv.Atoi(string(deviceID[:len(deviceID)-1]))
		if err != nil {
			return err
		}

		productName, err := ioutil.ReadFile(path.Join(p, "product"))
		if err != nil {
			return err
		}

		manufacturer, err := ioutil.ReadFile(path.Join(p, "manufacturer"))
		if err != nil {
			return err
		}

		port := path.Base(p)

		relays[port] = &Relay{
			DevNum:       intID,
			Product:      string(productName[:len(productName)-1]),
			Port:         port,
			Manufacturer: string(manufacturer[:len(manufacturer)-1]),
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("unable to walk relays: %w", err)
	}

	// populate hid.Device
	hid.UsbWalk(func(d hid.Device) {
		info := d.Info()
		for _, relay := range relays {
			if info.Device == relay.DevNum {
				relay.Device = d
			}
		}
	})

	return relays, nil
}
