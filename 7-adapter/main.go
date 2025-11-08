package main

import "fmt"

// Step 1 Define Target Interface
type USBTypeC interface {
	ChargeWithTypeC()
}

// Step 2 Define Adaptee (Lightning Charger)
type LightningCharger struct{}

func (l *LightningCharger) ChargeWithLightning() {
	fmt.Println("Charging with Lightning connector ⚡")
}

// Step 3. Create the Adapter
type LightningToTypeCAdapter struct {
	lightningDevice *LightningCharger
}

func (a *LightningToTypeCAdapter) ChargeWithTypeC() {
	fmt.Println("Adapter converts Type-C request → Lightning port 🔌")
	a.lightningDevice.ChargeWithLightning()
}

// Step 4. Client Code
func main() {
	// Old charger (Lightning)
	lightning := &LightningCharger{}

	// Adapter that converts to Type-C
	adapter := &LightningToTypeCAdapter{
		lightningDevice: lightning,
	}

	// Client expects USB-C interface
	var typeC USBTypeC = adapter
	fmt.Println("Connecting Android phone with Lightning charger...")
	typeC.ChargeWithTypeC()
}
