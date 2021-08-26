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
		t.Fatalf("wanted chassis %q, got %q", wanted, got)
	}

	got = assembled.Tires
	if got != wanted {
		t.Fatalf("wanted tires %q, got %q", wanted, got)
	}

	got = assembled.Engine
	if got != wanted {
		t.Fatalf("wanted engine %q, got %q", wanted, got)
	}

	got = assembled.Electronics
	if got != wanted {
		t.Fatalf("wanted electronics %q, got %q", wanted, got)
	}

	got = assembled.Dash
	if got != wanted {
		t.Fatalf("wanted dash %q, got %q", wanted, got)
	}

	got = assembled.Sits
	if got != wanted {
		t.Fatalf("wanted sits %q, got %q", wanted, got)
	}

	got = assembled.Windows
	if got != wanted {
		t.Fatalf("wanted windows %q, got %q", wanted, got)
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
