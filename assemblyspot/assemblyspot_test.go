package assemblyspot

import (
	"testing"

	"github.com/happybydefault/challenge-oceanscode/vehicle"
)

func TestAssemblySpot_Assemble(t *testing.T) {
	v := &vehicle.Car{}
	spot := AssemblySpot{}
	spot.SetVehicle(v)

	assembled, err := spot.Assemble()
	if err != nil {
		t.Fatal(err)
	}

	wanted := Assembled

	got := assembled.Chassis
	if got != wanted {
		t.Fatalf("wanted chassis %v, got %v", wanted, got)
	}

	got = assembled.Tires
	if got != wanted {
		t.Fatalf("wanted tires %v, got %v", wanted, got)
	}

	got = assembled.Engine
	if got != wanted {
		t.Fatalf("wanted engine %v, got %v", wanted, got)
	}

	got = assembled.Electronics
	if got != wanted {
		t.Fatalf("wanted electronics %v, got %v", wanted, got)
	}

	got = assembled.Dash
	if got != wanted {
		t.Fatalf("wanted dash %v, got %v", wanted, got)
	}

	got = assembled.Sits
	if got != wanted {
		t.Fatalf("wanted sits %v, got %v", wanted, got)
	}

	got = assembled.Windows
	if got != wanted {
		t.Fatalf("wanted windows %v, got %v", wanted, got)
	}
}

func TestAssemblySpot_Vehicle(t *testing.T) {
	v := &vehicle.Car{}
	spot := AssemblySpot{vehicle: v}

	wanted := v
	got := spot.Vehicle()
	if got != wanted {
		t.Fatalf("wanted %p, got %p", wanted, got)
	}
}

func TestAssemblySpot_SetVehicle(t *testing.T) {
	v := &vehicle.Car{}
	spot := AssemblySpot{}
	spot.SetVehicle(v)

	wanted := v
	got := spot.vehicle
	if got != wanted {
		t.Fatalf("wanted %p, got %p", wanted, got)
	}
}
