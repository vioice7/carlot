package main

import (
	"testing"
)

func TestAddCar(t *testing.T) {
	lot := ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}

	lot.AddCar()
	got := lot.CarsNr
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRemoveCar(t *testing.T) {
	lot := ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}

	lot.RemoveCar()
	got := lot.CarsNr
	want := 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestAddMotorcycle(t *testing.T) {
	lot := ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}

	lot.AddMotorcycle()
	got := lot.MotorcyclesNr
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRemoveMotorcycle(t *testing.T) {
	lot := ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}

	lot.RemoveMotorcycle()
	got := lot.MotorcyclesNr
	want := 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAddVan(t *testing.T) {
	lot := ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}

	lot.AddVan()
	got := lot.VansNr
	want := 2

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestRemoveVan(t *testing.T) {
	lot := ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}

	lot.RemoveVan()
	got := lot.VansNr
	want := 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestTotalSpots(t *testing.T) {
	lot := ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}

	got := lot.TotalSpots()
	want := 6

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
