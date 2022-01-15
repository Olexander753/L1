package main

import "fmt"

// ChargerBrick - charger brick for charging some devices
type ChargerBrick struct{}

// PlugIn - plug in usbA cabel into brick via chargerBrick interface
func (brick *ChargerBrick) PlugIn(cabel chargerBrick) {
	cabel.PlugInTypeA()
	fmt.Println("Cabel inserted")
}

// chargerBrick interface for connecting a USB-typeA cable
type chargerBrick interface {
	PlugInTypeA()
}

// USB-typeA--Lightning charger cabel
type UsbTypeALightning struct{}

// PlugInTypeA - connects charger brick and cabel typeA
func (c *UsbTypeALightning) PlugInTypeA() {
	fmt.Println("USB-typeA -- Lightning cabel plugged into charger brick")
}

// USB-typeC--Lightning charger cabel
type UsbTypeCLightning struct{}

// PlugInTypeC - connects charger brick and cabel typeC
func (c *UsbTypeCLightning) PlugInTypeC() {
	fmt.Println("USB-typeC -- Lightning cabel plugged into charger brick")
}

// UsbA2UsbC - adapter afford to connect usbC cabel to usbA interface
type UsbC2UsbA struct {
	typeC *UsbTypeCLightning
}

func (adapter *UsbC2UsbA) PlugInTypeA() {
	fmt.Println("Usb-typeC plugged in via UsbC2UsbA adapter")
	adapter.typeC.PlugInTypeC()
}

func main() {
	typeA := UsbTypeALightning{}
	typeC := UsbTypeCLightning{}

	brick := ChargerBrick{}

	adapter := UsbC2UsbA{
		typeC: &typeC,
	}

	// connect typeA cabel
	brick.PlugIn(&typeA)
	// connect typeC cabel
	brick.PlugIn(&adapter)

}
