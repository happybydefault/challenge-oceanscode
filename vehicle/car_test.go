package vehicle

import (
	"errors"
	"testing"
)

func TestCar_StartEngine(t *testing.T) {
	t.Run("with engine off", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: false}
		_, err := car.StartEngine()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("with engine already on", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: true}
		_, err := car.StartEngine()
		if err == nil {
			t.Fatal("should have gotten an error")
		}
	})
}

func TestCar_StopEngine(t *testing.T) {
	t.Run("with engine on", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: true}
		_, err := car.StopEngine()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("with engine already off", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: false}
		_, err := car.StopEngine()
		if err == nil {
			t.Fatal("should have gotten an error")
		}
	})
}

func TestCar_MoveForwards(t *testing.T) {
	t.Run("17 meters, with engine on", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: true}
		got, err := car.MoveForwards(17)
		if err != nil {
			t.Fatal(err)
		}

		wanted := "Moved forward 17 meters!"
		if got != wanted {
			t.Fatalf("wanted %v, got %v", wanted, got)
		}
	})

	t.Run("17 meters, with engine off", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: false}
		_, err := car.MoveForwards(17)
		if !errors.Is(err, ErrEngineOff) {
			t.Fatal("should have gotten ErrEngineOff")
		}
	})
}

func TestCar_MoveBackwards(t *testing.T) {
	t.Run("59 meters, with engine on", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: true}
		got, err := car.MoveBackwards(59)
		if err != nil {
			t.Fatal(err)
		}

		wanted := "Moved backwards 59 meters!"
		if got != wanted {
			t.Fatalf("wanted %v, got %v", wanted, got)
		}
	})

	t.Run("59 meters, with engine off", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: false}
		_, err := car.MoveBackwards(59)
		if !errors.Is(err, ErrEngineOff) {
			t.Fatal("should have gotten ErrEngineOff")
		}
	})
}

func TestCar_TurnRight(t *testing.T) {
	t.Run("with engine on", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: true}
		got, err := car.TurnRight()
		if err != nil {
			t.Fatal(err)
		}

		wanted := "Turned right!"
		if got != wanted {
			t.Fatalf("wanted %v, got %v", wanted, got)
		}
	})

	t.Run("with engine off", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: false}
		_, err := car.TurnRight()
		if !errors.Is(err, ErrEngineOff) {
			t.Fatal("should have gotten ErrEngineOff")
		}
	})
}

func TestCar_TurnLeft(t *testing.T) {
	t.Run("with engine on", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: true}
		got, err := car.TurnLeft()
		if err != nil {
			t.Fatal(err)
		}

		wanted := "Turned left!"
		if got != wanted {
			t.Fatalf("wanted %v, got %v", wanted, got)
		}
	})

	t.Run("with engine off", func(t *testing.T) {
		t.Parallel()

		car := Car{EngineOn: false}
		_, err := car.TurnLeft()
		if !errors.Is(err, ErrEngineOff) {
			t.Fatal("should have gotten ErrEngineOff")
		}
	})
}
