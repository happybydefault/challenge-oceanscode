package vehicle

import "errors"

var (
	ErrEngineOff = errors.New("engine is off")
	ErrEngineOn  = errors.New("engine is on")
)

type Vehicle interface {
	StartEngine() (string, error)
	StopEngine() (string, error)
	MoveForwards(meters int) (string, error)
	MoveBackwards(meters int) (string, error)
	TurnRight() (string, error)
	TurnLeft() (string, error)
}
