package smarthome

import "fmt"

type Device struct {
	DeviceID     string
	DeviceType   string
	RoomLocation string
	PowerState   bool
	BatteryLevel int
}

var devices = make(map[string]Device)

// Pair a new device
func PairDevice(device Device) error {
	if device.DeviceID == "" {
		return fmt.Errorf("device ID cannot be empty")
	}

	if device.DeviceType == "" {
		return fmt.Errorf("device type cannot be empty")
	}

	if device.RoomLocation == "" {
		return fmt.Errorf("room location cannot be empty")
	}

	if device.BatteryLevel < 0 || device.BatteryLevel > 100 {
		return fmt.Errorf("invalid battery level")
	}

	if _, exists := devices[device.DeviceID]; exists {
		return fmt.Errorf("device already exists")
	}

	devices[device.DeviceID] = device
	return nil
}

// Get a paired device
func GetDevice(id string) (Device, error) {
	device, exists := devices[id]

	if !exists {
		return Device{}, fmt.Errorf("device not found")
	}

	return device, nil
}
