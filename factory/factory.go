package factory

import (
	"fmt"

	"github.com/happybydefault/challenge-oceanscode/assemblyspot"
	"github.com/happybydefault/challenge-oceanscode/vehicle"
)

const assemblySpots int = 5

type Factory struct {
	AssemblingSpots chan *assemblyspot.AssemblySpot
}

func New() *Factory {
	factory := &Factory{
		AssemblingSpots: make(chan *assemblyspot.AssemblySpot, assemblySpots),
	}

	for i := 0; i < assemblySpots; i++ {
		factory.AssemblingSpots <- &assemblyspot.AssemblySpot{}
	}

	return factory
}

// HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
// (Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) {
	vehicleList := f.generateVehicleLots(amountOfVehicles)

	for _, vehicle := range vehicleList {
		fmt.Println("Assembling vehicle...")

		idleSpot := <-f.AssemblingSpots
		idleSpot.SetVehicle(&vehicle)
		vehicle, err := idleSpot.Assemble()

		if err != nil {
			continue
		}

		vehicle.TestingLog = f.testCar(vehicle)
		vehicle.AssembleLog = idleSpot.Log()

		idleSpot.SetVehicle(nil)
		f.AssemblingSpots <- idleSpot
	}
}

func (Factory) generateVehicleLots(amount int) []vehicle.Car {
	vehicles := make([]vehicle.Car, amount)
	for i := range vehicles {
		vehicles[i] = vehicle.Car{
			Id:          i,
			Chassis:     "NotSet",
			Tires:       "NotSet",
			Engine:      "NotSet",
			Electronics: "NotSet",
			Dash:        "NotSet",
			Sits:        "NotSet",
			Windows:     "NotSet",
			EngineOn:    false,
		}
	}
	return vehicles
}

func (f *Factory) testCar(car *vehicle.Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
