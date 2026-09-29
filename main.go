package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/smarthome/smarthome"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !pairDevice(scanner) {
			return
		}

		fmt.Print("Add another device? (y/n): ")
		if !scanner.Scan() || !strings.EqualFold(strings.TrimSpace(scanner.Text()), "y") {
			return
		}
	}
}

func pairDevice(scanner *bufio.Scanner) bool {
	fmt.Print("Enter Device ID: ")
	if !scanner.Scan() {
		fmt.Println("Invalid device ID")
		return false
	}
	id := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Device Type: ")
	if !scanner.Scan() {
		fmt.Println("Invalid device type")
		return false
	}
	typeName := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter Room: ")
	if !scanner.Scan() {
		fmt.Println("Invalid room")
		return false
	}
	location := strings.TrimSpace(scanner.Text())

	fmt.Print("Power State (on/off): ")
	if !scanner.Scan() {
		fmt.Println("Invalid power state (on/off)")
		return false
	}
	var powerState bool
	switch strings.ToLower(strings.TrimSpace(scanner.Text())) {
	case "on":
		powerState = true
	case "off":
		powerState = false
	default:
		fmt.Println("Invalid power state (on/off)")
		return false
	}

	fmt.Print("Enter Battery Level (0-100): ")
	if !scanner.Scan() {
		fmt.Println("Invalid battery level (0-100)")
		return false
	}
	battery, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || battery < 0 || battery > 100 {
		fmt.Println("Invalid battery level (0-100)")
		return false
	}

	device := smarthome.Device{
		DeviceID:     id,
		DeviceType:   typeName,
		RoomLocation: location,
		PowerState:   powerState,
		BatteryLevel: battery,
	}

	if err := smarthome.PairDevice(device); err != nil {
		fmt.Println("Error:", err)
		return false
	}

	got, err := smarthome.GetDevice(id)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	fmt.Println("\nDevice paired successfully!")
	fmt.Println("Device ID:", got.DeviceID)
	fmt.Println("Device Type:", got.DeviceType)
	fmt.Println("Room:", got.RoomLocation)
	fmt.Println("Power:", got.PowerState)
	fmt.Println("Battery:", got.BatteryLevel, "%")

	file, err := os.OpenFile(
		"device_history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644,
	)
	if err != nil {
		fmt.Println("File error:", err)
		return false
	}
	defer file.Close()

	_, err = fmt.Fprintf(file,
		"Date: %s | ID: %s | Type: %s | Room: %s | Power: %t | Battery: %d%%\n",
		time.Now().Format("2006-01-02 15:04:05"),
		got.DeviceID,
		got.DeviceType,
		got.RoomLocation,
		got.PowerState,
		got.BatteryLevel,
	)
	if err != nil {
		fmt.Println("Error saving history:", err)
		return false
	}

	fmt.Println("History saved in device_history.txt")
	return true
}
