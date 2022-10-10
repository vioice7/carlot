package main

import (
	"fmt"
)

type parkingLot struct {
	totalSpots   int
	ocupiedSpots int
}

func (p *parkingLot) total() int {
	return p.totalSpots
}

func (p *parkingLot) ocupied() int {
	return p.ocupiedSpots
}

func (p *parkingLot) addCar(nr int) {
	p.ocupiedSpots += nr
}

func (p *parkingLot) removeCar(nr int) {
	p.ocupiedSpots -= nr
}

func (p *parkingLot) emptyFullCheck() {
	// check if the lot is full
	if p.totalSpots <= p.ocupiedSpots {
		fmt.Println("Parking lot is full! We can't add anymore cars!")

		// equalise the total slots and the occupied slots
		p.ocupiedSpots = p.totalSpots
	}
	// reset ocupied slots if we remove more tahn total number of cars
	if p.ocupiedSpots < 0 {
		fmt.Println("We can't remove more cars than the ocupied spots!")
		p.ocupiedSpots = 0
	}
}

func main() {

	var totalNumber int
	var occupiedNumber int

	fmt.Printf("Total car spots: ")
	fmt.Scanln(&totalNumber)

	fmt.Printf("Initial occupied spots: ")
	fmt.Scanln(&occupiedNumber)

	if totalNumber < occupiedNumber {
		panic("Excedded total spots!")
	}

	// define the lot
	lot := parkingLot{totalNumber, occupiedNumber}

	// print total car spots and initially ocupied spots

	fmt.Println("---")
	fmt.Println("Total car spots:", lot.total())
	fmt.Println("Occupied spots:", lot.ocupied())
	fmt.Println("---")

	// Add cars

	var carAddNr int
	fmt.Printf("How many cars would you like to add? :")
	fmt.Scanln(&carAddNr)
	lot.addCar(carAddNr)

	lot.emptyFullCheck()

	fmt.Println("---")
	fmt.Println("Total car spots:", lot.total())
	fmt.Println("Occupied spots:", lot.ocupied())
	fmt.Println("---")

	// Remove cars
	var carRemoveNr int
	fmt.Printf("How many cars would you like to remove? :")
	fmt.Scanln(&carRemoveNr)
	lot.removeCar(carRemoveNr)

	lot.emptyFullCheck()

	fmt.Println("---")
	fmt.Println("Total car spots:", lot.total())
	fmt.Println("Occupied spots:", lot.ocupied())
	fmt.Println("---")

}

/*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          */
