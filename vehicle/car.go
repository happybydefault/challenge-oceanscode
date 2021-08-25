package vehicle

import (
	"errors"
	"fmt"
)

var (
	ErrEngineOff = errors.New("engine is off")
	ErrEngineOn  = errors.New("engine is on")
)

type Car struct {
	Id          int
	Chassis     string
	Tires       string
	Engine      string
	Electronics string
	Dash        string
	Sits        string
	Windows     string
	EngineOn    bool
	TestingLog  string
	AssembleLog string
}

func (c *Car) StartEngine() (string, error) {
	if c.EngineOn {
		return "", ErrEngineOn
	}

	return "Engine started!", nil
}

func (c *Car) StopEngine() (string, error) {
	if !c.EngineOn {
		return "", ErrEngineOff
	}

	return "Engine stopped!", nil
}

func (c *Car) MoveForwards(meters int) (string, error) {
	if !c.EngineOn {
		return "", ErrEngineOff
	}

	return fmt.Sprintf("Moved forward %d meters!", meters), nil
}

func (c *Car) MoveBackwards(meters int) (string, error) {
	if !c.EngineOn {
		return "", ErrEngineOff
	}

	return fmt.Sprintf("Moved backwards %d meters!", meters), nil
}

func (c *Car) TurnRight() (string, error) {
	if !c.EngineOn {
		return "", ErrEngineOff
	}

	return "Turned right!", nil
}

func (c *Car) TurnLeft() (string, error) {
	if !c.EngineOn {
		return "", ErrEngineOff
	}

	return "Turned left!", nil
}
