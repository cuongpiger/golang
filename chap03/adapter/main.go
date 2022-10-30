package main

import "fmt"

type (
	Computer interface {
		InsertIntoLightningPort()
	}

	Client struct{}

	Mac struct{}

	Windows struct{}

	WindowsAdapter struct {
		windowsMachine *Windows
	}
)

// Client's collection of methods
func (s *Client) InsertLightningConnectorIntoComputer(computer Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	computer.InsertIntoLightningPort()
}

// Mac's collection of methods
func (s *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Windows's collection of methods
func (s *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}

// WindowsAdapter's collection of methods
func (s *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	s.windowsMachine.insertIntoUSBPort()
}

// main function
func main() {
	client := new(Client)
	mac := new(Mac)
	windowsMachine := new(Windows)
	windowsMachineAdapter := &WindowsAdapter{windowsMachine: windowsMachine}

	client.InsertLightningConnectorIntoComputer(mac)
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}