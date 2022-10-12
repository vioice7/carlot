package main

import (
	"fmt"
	"testing"
)

// ------------------------------------------------------------------
// Parking lot structure
// ------------------------------------------------------------------
// Cars Number
// Motorcycle NUmber
// Vans Number
// Cars spots (can take motorcycles or cars)
// Motorcycles spots (can take motorcycles only)
// Vans spots (can take motorcycles, cars or vans)
// Cars real spots (the number of designated cars spots)
// Motorcycles real spots (the number of designated cars motorcycles)
// Vans real spots (the number of designated cars vans)
// ------------------------------------------------------------------

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

func TestParking(t *testing.T) {
	type test struct {
		lot  ParkingLot
		want []int
	}

	// ------------------------------------------------------------------
	// Parking lot structure
	// ------------------------------------------------------------------
	// Cars Number
	// Motorcycle NUmber
	// Vans Number
	// Cars spots (can take motorcycles or cars)
	// Motorcycles spots (can take motorcycles only)
	// Vans spots (can take motorcycles, cars or vans)
	// Cars real spots (the number of designated cars spots)
	// Motorcycles real spots (the number of designated cars motorcycles)
	// Vans real spots (the number of designated cars vans)
	// ------------------------------------------------------------------

	// want: []int {CarsNr, MotorcyclesNr, VanNr}
	tests := []test{
		{lot: ParkingLot{1, 1, 1, 2, 2, 2, 2, 2, 2}, want: []int{2, 0, 1}},
		{lot: ParkingLot{1, 2, 2, 2, 4, 6, 2, 4, 6}, want: []int{2, 1, 2}},
		{lot: ParkingLot{1, 3, 3, 3, 5, 7, 3, 5, 7}, want: []int{2, 2, 3}},
		{lot: ParkingLot{1, 0, 3, 3, 5, 7, 3, 5, 7}, want: []int{2, 0, 3}},
	}

	for v, tc := range tests {

		fmt.Println("--------------------- Test nr:", v+1, "---------------------")

		// we add a car
		tc.lot.AddCar()
		// we remove a motorcycle
		tc.lot.RemoveMotorcycle()
		// we add a van
		tc.lot.AddVan()
		// we remove a van
		tc.lot.RemoveVan()

		// checking car, moto and van number values
		gotCar, gotMoto, gotVan := tc.lot.CarsNr, tc.lot.MotorcyclesNr, tc.lot.VansNr

		if gotCar != tc.want[0] {
			t.Errorf("got %v cars want %v cars", gotCar, tc.want[0])
		}

		if gotMoto != tc.want[1] {
			t.Errorf("got %v motorcycles  want %v motorcycles", gotMoto, tc.want[1])
		}

		if gotVan != tc.want[2] {
			t.Errorf("got %v vans want %v vans", gotVan, tc.want[2])
		}
	}

}
