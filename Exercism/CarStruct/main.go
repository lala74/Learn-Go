package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func CreateCar(speed, batteryDrain int) *Car {
	return &Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

func (car *Car) Drive() {
	car.battery -= car.batteryDrain
	if car.battery < 0 {
		car.battery = 0
	}
	car.distance += car.speed
}

func (car *Car) CanFinish(track Track) bool {
	for {
		if car.distance >= track.distance {
			return true
		}
		if car.battery == 0 {
			return false
		}
		car.Drive()
	}
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

type Track struct {
	distance int
}

func CreateTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

func main() {
	speed := 5
	batteryDrain := 2
	car := CreateCar(speed, batteryDrain)

	distance := 100
	raceTrack := CreateTrack(distance)

	finish := car.CanFinish(*raceTrack)
	fmt.Println("can finish: ", finish)
	fmt.Println(car.DisplayBattery())
	fmt.Println(car.DisplayDistance())
}
