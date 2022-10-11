package main

import (
	"fmt"
)

// Assumptions
// 1 motorcycle can take a motorcycle, car or van spot
// 1 car can take a car or van spot
// 1 van can take only a van spot
// we last add or first remove from non designated spaces

type parkingLot struct {
	carsNr              int
	motorcyclesNr       int
	vansNr              int
	carSpots            int
	motorcycleSpots     int
	vanSpots            int
	carRealSpots        int
	motorcycleRealSpots int
	vanRealSpots        int
}

func (p *parkingLot) totalSlots() int {
	return p.carSpots + p.motorcycleSpots + p.vanSpots
}

func (p *parkingLot) addMotorcycle() {
	if p.motorcyclesNr < p.motorcycleSpots {
		p.motorcyclesNr++
		fmt.Println("Added a motorcycle")
	} else if p.carsNr < p.carSpots {
		p.motorcyclesNr++
		fmt.Println("Added a motorcycle")

		p.carSpots--
		p.motorcycleSpots++
	} else if p.vansNr < p.vanSpots {
		p.motorcyclesNr++
		fmt.Println("Added a motorcycle")

		p.vanSpots--
		p.motorcycleSpots++
	} else {
		fmt.Println("Can't add any motorcycles! (parking lot full)")
	}
}

func (p *parkingLot) removeMotorcycle() {
	if p.motorcyclesNr > 0 {
		if p.motorcyclesNr > p.motorcycleRealSpots {
			if p.vanSpots < p.vanRealSpots {
				p.motorcyclesNr--
				fmt.Println("Motorcycle removed")

				p.motorcycleSpots--
				p.vanSpots++
			} else if p.carSpots < p.carRealSpots {
				p.motorcyclesNr--
				fmt.Println("Motorcycle removed")

				p.motorcycleSpots--
				p.carSpots++
			}
		} else {
			p.motorcyclesNr--
			fmt.Println("Motorcycle removed")
		}
	} else {
		fmt.Println("There are no motorcycles in the parking lot!")
	}
}

func (p *parkingLot) addCar() {
	if p.carsNr < p.carSpots {
		p.carsNr++
		fmt.Println("Added a car")
	} else if p.vansNr < p.vanSpots {
		p.carsNr++
		fmt.Println("Added a car")

		p.vanSpots--
		p.carSpots++
	} else {
		fmt.Println("Can't add any cars! (parking lot full)")
	}
}

func (p *parkingLot) removeCar() {
	if p.carsNr > 0 {
		if p.carsNr > p.carRealSpots {
			if p.vanSpots < p.vanRealSpots {
				p.carsNr--
				fmt.Println("Car removed")

				p.vanSpots++
				p.carSpots--
			}
		} else {
			p.carsNr--
			fmt.Println("Car removed")
		}

	} else {
		fmt.Println("There are no cars in the parking lot!")
	}
}

func (p *parkingLot) addVan() {
	if p.vansNr < p.vanSpots {
		p.vansNr++
		fmt.Println("Added a van")
	} else {
		fmt.Println("Can't add any vans! (parking lot full)")
	}
}

func (p *parkingLot) removeVan() {
	if p.vansNr > 0 {
		p.vansNr--
		fmt.Println("Van removed")
	} else {
		fmt.Println("There are no vans in the parking lot!")
	}
}

func (p *parkingLot) checkEach() (int, int, int, int, int, int, int, int, int) {
	return p.carsNr, p.motorcyclesNr, p.vansNr, p.carSpots, p.motorcycleSpots, p.vanSpots, p.carRealSpots, p.motorcycleRealSpots, p.vanRealSpots
}

func main() {

	var carNr, mcycNr, vanNr, carSp, mcycSp, vanSp int

	fmt.Printf("Initial car numbers: ")
	fmt.Scanln(&carNr)

	fmt.Printf("Initial car spots: ")
	fmt.Scanln(&carSp)

	if carSp < carNr {
		panic("Excedded total car spots!")
	}

	fmt.Printf("Initial motorcycle numbers: ")
	fmt.Scanln(&mcycNr)

	fmt.Printf("Initial motorcycle spots: ")
	fmt.Scanln(&mcycSp)

	if mcycSp < mcycNr {
		panic("Excedded total motorcycle spots!")
	}

	fmt.Printf("Initial van numbers: ")
	fmt.Scanln(&vanNr)

	fmt.Printf("Initial van spots: ")
	fmt.Scanln(&vanSp)

	if vanSp < vanNr {
		panic("Excedded total van spots!")
	}

	// define the lot
	lot := parkingLot{carNr, mcycNr, vanNr, carSp, mcycSp, vanSp, carSp, mcycSp, vanSp}

	// print total car spots and initially ocupied spots
	fmt.Println("Total parking slots:", lot.totalSlots())
	// define a string that takes a text command
	var cmd string

	// wait for a command
	for {
		fmt.Printf("Command:")
		fmt.Scanln(&cmd)

		if cmd == "total" {
			fmt.Println("Total parking slots:", lot.totalSlots())
		}

		if cmd == "check" {
			carNr0, mcycNr0, vanNr0, carSp0, mcycSp0, vanSp0, carReSp0, mcycReSp0, vanReSp0 := lot.checkEach()

			fmt.Println("Motorcycle numbers:", mcycNr0, "Temporary spots that take motorcycles:", mcycSp0, "Motorcycle real spots:", mcycReSp0)
			fmt.Println("Car numbers:       ", carNr0, "Temporary spots that take cars:       ", carSp0, "Car real spots:       ", carReSp0)
			fmt.Println("Van numbers:       ", vanNr0, "Temporary spots that take vans:       ", vanSp0, "Van real spots:       ", vanReSp0)
		}

		if cmd == "addmoto" {
			lot.addMotorcycle()
		}

		if cmd == "remmoto" {
			lot.removeMotorcycle()
		}

		if cmd == "addcar" {
			lot.addCar()
		}

		if cmd == "remcar" {
			lot.removeCar()
		}

		if cmd == "addvan" {
			lot.addVan()
		}

		if cmd == "remvan" {
			lot.removeVan()
		}

		if cmd == "exit" {
			break
		}
	}

}
