package main

import "fmt"

type Vehicle interface {
	Move() string
	Stop() string
	SpeedChange(speedControl int) (string, int, string)
}

type Passengers interface {
	PickUpPassengers(passengers int) (string, int, string)
	DropOffPassengers() string
}

type Car struct{}

type Plane struct{}

type Train struct{}

type Tour struct {
	Vehicle    []Vehicle
	Passengers []Passengers
}

func (c Car) Move() string {
	return "Car start moving!"
}

func (p Plane) Move() string {
	return "Plane start moving!"
}

func (t Train) Move() string {
	return "Train start moving!"
}

func (c Car) Stop() string {
	return "Car stopped!"
}

func (p Plane) Stop() string {
	return "Plane stopped!"
}

func (p Train) Stop() string {
	return "Plane stopped!"
}

func (c Car) SpeedChange(speedControl int) (string, int, string) {
	fmt.Println("Choose speed:")
	fmt.Scan(&speedControl)
	return "Car is moving with speed", speedControl, "km/h"
}

func (p Plane) SpeedChange(speedControl int) (string, int, string) {
	fmt.Println("Choose speed:")
	fmt.Scan(&speedControl)
	return "Plane is moving with speed", speedControl, "km/h"
}

func (t Train) SpeedChange(speedControl int) (string, int, string) {
	fmt.Println("Choose speed:")
	fmt.Scan(&speedControl)
	return "Train is moving with speed", speedControl, "km/h"

}

func (c Car) PickUpPassengers(passengers int) (string, int, string) {
	fmt.Println("Choose amount of passengers for car:")
	fmt.Scan(&passengers)
	return "Car pick up", passengers, "passengers!"
}

func (p Plane) PickUpPassengers(passengers int) (string, int, string) {
	fmt.Println("Choose amount of passengers for plane:")
	fmt.Scan(&passengers)
	return "Plane pick up", passengers, "passengers!"
}

func (t Train) PickUpPassengers(passengers int) (string, int, string) {
	fmt.Println("Choose amount of passengers for train:")
	fmt.Scan(&passengers)
	return "Train pick up", passengers, "passengers!"
}

func (t Train) DropOffPassengers() string {
	return "Train drop off passengers!"
}

func (c Car) DropOffPassengers() string {
	return "Car drop off passengers!"
}

func (p Plane) DropOffPassengers() string {
	return "Plane drop off passengers!"
}

func StartTour(t Tour) {
	fmt.Println("Tour is starting!")
	var passengers int
	var speedControl int
	for _, value := range t.Passengers {
		fmt.Println(value.PickUpPassengers(passengers))
	}
	for _, value := range t.Vehicle {
		fmt.Println(value.Move())
		fmt.Println(value.SpeedChange(speedControl))
		fmt.Println(value.Stop())
	}
	for _, value := range t.Vehicle {
		fmt.Println(value.Stop())
	}
	for _, value := range t.Passengers {
		fmt.Println(value.DropOffPassengers())
	}
}

func main() {
	car := Car{}
	plane := Plane{}
	train := Train{}
	tour := Tour{
		Vehicle:    []Vehicle{car, plane, train},
		Passengers: []Passengers{car, plane, train},
	}

	StartTour(tour)

}
