package main

import "fmt"

// Step 1: Define the Handler interface
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(level string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) PassToNext(level string) {
	if b.next != nil {
		b.next.HandleRequest(level)
	} else {
		fmt.Println("No handler available for this request ❌")
	}
}

// Step 2: Create Concrete Handlers
type FrontlineHandler struct {
	BaseHandler
}

func (f *FrontlineHandler) HandleRequest(level string) {
	if level == "low" {
		fmt.Println("Frontline Support: Handling LOW level issue 🧰")
	} else {
		fmt.Println("Frontline Support: Passing to next...")
		f.PassToNext(level)
	}
}

type TechnicalHandler struct {
	BaseHandler
}

func (t *TechnicalHandler) HandleRequest(level string) {
	if level == "medium" {
		fmt.Println("Technical Support: Handling MEDIUM level issue ⚙️")
	} else {
		fmt.Println("Technical Support: Passing to next...")
		t.PassToNext(level)
	}
}

type ManagerHandler struct {
	BaseHandler
}

func (m *ManagerHandler) HandleRequest(level string) {
	if level == "high" {
		fmt.Println("Manager: Handling HIGH level issue 🚨")
	} else {
		fmt.Println("Manager: Cannot handle further.")
		m.PassToNext(level)
	}
}

// Step 3: Client code
func main() {
	frontline := &FrontlineHandler{}
	technical := &TechnicalHandler{}
	manager := &ManagerHandler{}

	// Create chain: frontline → technical → manager
	frontline.SetNext(technical)
	technical.SetNext(manager)

	fmt.Println("1️⃣ Low level request:")
	frontline.HandleRequest("low")

	fmt.Println("\n2️⃣ Medium level request:")
	frontline.HandleRequest("medium")

	fmt.Println("\n3️⃣ High level request:")
	frontline.HandleRequest("high")

	fmt.Println("\n4️⃣ Unknown level request:")
	frontline.HandleRequest("critical")
}
