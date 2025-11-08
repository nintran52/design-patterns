package factorymethod

import (
	"fmt"

	"github.com/nintran52/design-patterns/3-factory-method/logistic"
)

func PlanDelivery(logistics logistic.Logistics) {
	transport := logistics.CreateTransport()
	transport.Deliver()
}

func main() {
	fmt.Println("Road Logistics:")
	PlanDelivery(logistic.RoadLogistics{})

	fmt.Println("\nSea Logistics:")
	PlanDelivery(logistic.SeaLogistics{})

	fmt.Println("\nAir Logistics:")
	PlanDelivery(logistic.AirLogistics{})
}
