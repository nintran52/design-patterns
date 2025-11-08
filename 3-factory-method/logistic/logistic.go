package logistic

import "fmt"

// Product interface
type Transport interface {
	Deliver()
}

// Concrete Product 1
type Truck struct{}

func (Truck) Deliver() {
	fmt.Println("Deliver by land in a box.")
}

// Concrete Product 2
type Ship struct{}

func (Ship) Deliver() {
	fmt.Println("Deliver by sea in a container.")
}

type Airplane struct{}

func (Airplane) Deliver() {
	fmt.Println("Deliver by air in a parcel.")
}

// Creator (Factory)
type Logistics interface {
	CreateTransport() Transport
}

// Concrete Creator 1
type RoadLogistics struct{}

func (RoadLogistics) CreateTransport() Transport {
	return Truck{}
}

// Concrete Creator 2
type SeaLogistics struct{}

func (SeaLogistics) CreateTransport() Transport {
	return Ship{}
}

// Concrete Creator 3
type AirLogistics struct{}

func (AirLogistics) CreateTransport() Transport {
	return Airplane{}
}
