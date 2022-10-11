package main

import (
	"fmt"
)

// Assumptions
// 1 motorcycle can take a motorcycle, car or van spot
// 1 car can take a car or van spot
// 1 van can take only a van spot
// we last add or first remove from non designated spaces

type ParkingLot struct {
	CarsNr              int
	MotorcyclesNr       int
	VansNr              int
	CarSpots            int
	MotorcycleSpots     int
	VanSpots            int
	CarRealSpots        int
	MotorcycleRealSpots int
	VanRealSpots        int
}

func (p *ParkingLot) TotalSpots() int {
	return p.CarSpots + p.MotorcycleSpots + p.VanSpots
}

func (p *ParkingLot) AddMotorcycle() {
	if p.MotorcyclesNr < p.MotorcycleSpots {
		p.MotorcyclesNr++
		fmt.Println("Added a motorcycle")
	} else if p.CarsNr < p.CarSpots {
		p.MotorcyclesNr++
		fmt.Println("Added a motorcycle")

		p.CarSpots--
		p.MotorcycleSpots++
	} else if p.VansNr < p.VanSpots {
		p.MotorcyclesNr++
		fmt.Println("Added a motorcycle")

		p.VanSpots--
		p.MotorcycleSpots++
	} else {
		fmt.Println("Can't add any motorcycles! (parking lot full)")
	}
}

func (p *ParkingLot) RemoveMotorcycle() {
	if p.MotorcyclesNr > 0 {
		if p.MotorcyclesNr > p.MotorcycleRealSpots {
			if p.VanSpots < p.VanRealSpots {
				p.MotorcyclesNr--
				fmt.Println("Motorcycle removed")

				p.MotorcycleSpots--
				p.VanSpots++
			} else if p.CarSpots < p.CarRealSpots {
				p.MotorcyclesNr--
				fmt.Println("Motorcycle removed")

				p.MotorcycleSpots--
				p.CarSpots++
			}
		} else {
			p.MotorcyclesNr--
			fmt.Println("Motorcycle removed")
		}
	} else {
		fmt.Println("There are no motorcycles in the parking lot!")
	}
}

func (p *ParkingLot) AddCar() {
	if p.CarsNr < p.CarSpots {
		p.CarsNr++
		fmt.Println("Added a car")
	} else if p.VansNr < p.VanSpots {
		p.CarsNr++
		fmt.Println("Added a car")

		p.VanSpots--
		p.CarSpots++
	} else {
		fmt.Println("Can't add any cars! (parking lot full)")
	}
}

func (p *ParkingLot) RemoveCar() {
	if p.CarsNr > 0 {
		if p.CarsNr > p.CarRealSpots {
			if p.VanSpots < p.VanRealSpots {
				p.CarsNr--
				fmt.Println("Car removed")

				p.VanSpots++
				p.CarSpots--
			}
		} else {
			p.CarsNr--
			fmt.Println("Car removed")
		}

	} else {
		fmt.Println("There are no cars in the parking lot!")
	}
}

func (p *ParkingLot) AddVan() {
	if p.VansNr < p.VanSpots {
		p.VansNr++
		fmt.Println("Added a van")
	} else {
		fmt.Println("Can't add any vans! (parking lot full)")
	}
}

func (p *ParkingLot) RemoveVan() {
	if p.VansNr > 0 {
		p.VansNr--
		fmt.Println("Van removed")
	} else {
		fmt.Println("There are no vans in the parking lot!")
	}
}

func (p *ParkingLot) CheckEach() (int, int, int, int, int, int, int, int, int) {
	return p.CarsNr, p.MotorcyclesNr, p.VansNr, p.CarSpots, p.MotorcycleSpots, p.VanSpots, p.CarRealSpots, p.MotorcycleRealSpots, p.VanRealSpots
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
	lot := ParkingLot{carNr, mcycNr, vanNr, carSp, mcycSp, vanSp, carSp, mcycSp, vanSp}

	// print total car spots and initially ocupied spots
	fmt.Println("Total parking spots:", lot.TotalSpots())
	// define a string that takes a text command
	var cmd string

	// wait for a command
	for {
		fmt.Printf("Command:")
		fmt.Scanln(&cmd)

		if cmd == "total" {
			fmt.Println("Total parking spots:", lot.TotalSpots())
		}

		if cmd == "check" {
			carNr0, mcycNr0, vanNr0, carSp0, mcycSp0, vanSp0, carReSp0, mcycReSp0, vanReSp0 := lot.CheckEach()

			fmt.Println("Motorcycle numbers:", mcycNr0, " | Temporary motorcycle spots:", (mcycSp0 - mcycReSp0), " | Motorcycle real spots:", mcycReSp0)
			fmt.Println("Car numbers:       ", carNr0, " | Temporary car spots:       ", (carSp0 - carReSp0), " | Car real spots:       ", carReSp0)
			fmt.Println("Van numbers:       ", vanNr0, " | Temporary van spots:       ", (vanSp0 - vanReSp0), " | Van real spots:       ", vanReSp0)
		}

		if cmd == "addmoto" {
			lot.AddMotorcycle()
		}

		if cmd == "remmoto" {
			lot.RemoveMotorcycle()
		}

		if cmd == "addcar" {
			lot.AddCar()
		}

		if cmd == "remcar" {
			lot.RemoveCar()
		}

		if cmd == "addvan" {
			lot.AddVan()
		}

		if cmd == "remvan" {
			lot.RemoveVan()
		}

		if cmd == "exit" {
			break
		}
	}

}
