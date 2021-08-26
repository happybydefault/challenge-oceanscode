package assemblyspot

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/happybydefault/challenge-oceanscode/vehicle"
)

const (
	Assembled  = "Assembled"
	timeLayout = "2006-01-02 15:04:05.000"
)

type AssemblySpot struct {
	vehicle *vehicle.Car
	log     string
}

func (s *AssemblySpot) SetVehicle(v *vehicle.Car) {
	s.vehicle = v
}

func (s *AssemblySpot) Vehicle() *vehicle.Car {
	return s.vehicle
}

func (s *AssemblySpot) Log() string {
	return s.log
}

// hint: improve this function to execute this process concurrenlty
func (s *AssemblySpot) Assemble() (*vehicle.Car, error) {
	if s.vehicle == nil {
		return nil, errors.New("no vehicle set to start assembling")
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.assembleChassis()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.assembleTires()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.assembleEngine()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.assembleElectronics()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.assembleDash()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.assembleSeats()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.assembleWindows()
	}()

	return s.vehicle, nil
}

func (s *AssemblySpot) assembleChassis() {
	s.vehicle.Chassis = Assembled
	time.Sleep(1 * time.Second)
	s.log += fmt.Sprintf("Chassis at [%s], ", time.Now().Format(timeLayout))
}

func (s *AssemblySpot) assembleTires() {
	s.vehicle.Tires = Assembled
	time.Sleep(1 * time.Second)
	s.log += fmt.Sprintf("Tires at [%s], ", time.Now().Format(timeLayout))
}

func (s *AssemblySpot) assembleEngine() {
	s.vehicle.Engine = Assembled
	time.Sleep(1 * time.Second)
	s.log += fmt.Sprintf("Engine at [%s], ", time.Now().Format(timeLayout))
}

func (s *AssemblySpot) assembleElectronics() {
	s.vehicle.Electronics = Assembled
	time.Sleep(1 * time.Second)
	s.log += fmt.Sprintf("Electronics at [%s], ", time.Now().Format(timeLayout))
}

func (s *AssemblySpot) assembleDash() {
	s.vehicle.Dash = Assembled
	time.Sleep(1 * time.Second)
	s.log += fmt.Sprintf("Dash at [%s], ", time.Now().Format(timeLayout))
}

func (s *AssemblySpot) assembleSeats() {
	s.vehicle.Sits = Assembled
	time.Sleep(1 * time.Second)
	s.log += fmt.Sprintf("Sits at [%s], ", time.Now().Format(timeLayout))
}

func (s *AssemblySpot) assembleWindows() {
	s.vehicle.Windows = Assembled
	time.Sleep(1 * time.Second)
	s.log += fmt.Sprintf("Windows at [%s], ", time.Now().Format(timeLayout))
}
