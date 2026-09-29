package smarthome

import "testing"

func TestPairAndGetDevice(t *testing.T) {
	dev := Device{
		DeviceID:     "D101",
		DeviceType:   "Fan",
		RoomLocation: "Bedroom",
		BatteryLevel: 80,
	}

	if err := PairDevice(dev); err != nil {
		t.Fatalf("PairDevice returned error: %v", err)
	}

	stored, err := GetDevice("D101")
	if err != nil {
		t.Fatalf("GetDevice returned error: %v", err)
	}

	if stored.DeviceID != "D101" || stored.RoomLocation != "Bedroom" {
		t.Fatalf("unexpected stored device: %+v", stored)
	}
}
