package main

import (
	"log"

	"github.com/happybydefault/challenge-oceanscode/factory"
)

const carsAmount = 100

func main() {
	factory := factory.New()

	// Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	// each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	cars := factory.StartAssemblingProcess(carsAmount)

	for car := range cars {
		log.Printf("car %d: testing log: %s\n", car.Id, car.TestingLog)
		log.Printf("car %d: assemble log: %s\n", car.Id, car.AssembleLog)
	}
}
