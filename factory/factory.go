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
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) <-chan *vehicle.Car {
	if amountOfVehicles <= 0 {
		return nil
	}

	ch := make(chan *vehicle.Car)
	go func() {
		defer close(ch)

		vehicles := f.generateVehicleLots(amountOfVehicles)
		for _, v := range vehicles {
			fmt.Println("Assembling vehicle...")

			idleSpot := <-f.AssemblingSpots
			idleSpot.SetVehicle(&v)

			assembled, err := idleSpot.Assemble()

			if err != nil {
				continue
			}

			assembled.TestingLog = f.testCar(assembled)
			assembled.AssembleLog = idleSpot.Log()

			vehicle := &vehicle.Car{}
			*vehicle = *assembled
			ch <- vehicle

			idleSpot.SetVehicle(nil)
			f.AssemblingSpots <- idleSpot
		}
	}()
	return ch
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
