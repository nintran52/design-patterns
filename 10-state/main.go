package main

import "fmt"

// Step 1: Define the State interface
type State interface {
	InsertCoin()
	SelectItem()
	DispenseItem()
}

type VendingMachine struct {
	noCoinState     State
	hasCoinState    State
	dispensingState State
	currentState    State
	itemCount       int
}

// Step 2: Define concrete states
func NewVendingMachine(itemCount int) *VendingMachine {
	m := &VendingMachine{itemCount: itemCount}
	noCoin := &NoCoinState{machine: m}
	hasCoin := &HasCoinState{machine: m}
	dispensing := &DispensingState{machine: m}

	m.noCoinState = noCoin
	m.hasCoinState = hasCoin
	m.dispensingState = dispensing
	m.currentState = noCoin

	return m
}

func (m *VendingMachine) SetState(state State) {
	m.currentState = state
}

func (m *VendingMachine) InsertCoin() {
	m.currentState.InsertCoin()
}

func (m *VendingMachine) SelectItem() {
	m.currentState.SelectItem()
}

func (m *VendingMachine) DispenseItem() {
	m.currentState.DispenseItem()
}

// Step 3: Implement concrete states
type NoCoinState struct {
	machine *VendingMachine
}

func (s *NoCoinState) InsertCoin() {
	fmt.Println("💰 Coin inserted.")
	s.machine.SetState(s.machine.hasCoinState)
}

func (s *NoCoinState) SelectItem() {
	fmt.Println("Please insert a coin first!")
}

func (s *NoCoinState) DispenseItem() {
	fmt.Println("You need to pay first.")
}

type HasCoinState struct {
	machine *VendingMachine
}

func (s *HasCoinState) InsertCoin() {
	fmt.Println("Coin already inserted!")
}

func (s *HasCoinState) SelectItem() {
	fmt.Println("🛒 Item selected.")
	s.machine.SetState(s.machine.dispensingState)
	s.machine.DispenseItem()
}

func (s *HasCoinState) DispenseItem() {
	fmt.Println("Please select an item first.")
}

type DispensingState struct {
	machine *VendingMachine
}

func (s *DispensingState) InsertCoin() {
	fmt.Println("Please wait, dispensing in progress...")
}

func (s *DispensingState) SelectItem() {
	fmt.Println("Dispensing now, can’t select another item.")
}

func (s *DispensingState) DispenseItem() {
	if s.machine.itemCount > 0 {
		fmt.Println("🎁 Dispensing item...")
		s.machine.itemCount--
		if s.machine.itemCount > 0 {
			s.machine.SetState(s.machine.noCoinState)
		} else {
			fmt.Println("⚠️ Machine out of stock!")
		}
	} else {
		fmt.Println("No items left!")
		s.machine.SetState(s.machine.noCoinState)
	}
}

// Step 4: Client code
func main() {
	machine := NewVendingMachine(2)

	machine.SelectItem()
	machine.InsertCoin()
	machine.SelectItem()

	fmt.Println("\nNext Customer:")
	machine.InsertCoin()
	machine.SelectItem()

	fmt.Println("\nOut of stock test:")
	machine.InsertCoin()
	machine.SelectItem()
}
