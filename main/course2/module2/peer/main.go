package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc float64, vel float64, inD float64) func(float64) float64 {

	return func(time float64) float64 {
		var s float64
		s = float64((0.5)*acc)*math.Pow(time, 2) + float64(vel)*time + float64(inD)
		return s
	}
}

func main() {

	var calcD func(float64) float64
	var acc, vel, inD float64

	fmt.Printf("Input your Acceleration (a) (float64) : ")
	fmt.Scan(&acc)

	fmt.Printf("Input your Initial Velocity (Vo) (float64) : ")
	fmt.Scan(&vel)

	fmt.Printf("Input your Initial Displacement (So) (float64) : ")
	fmt.Scan(&inD)

	calcD = GenDisplaceFn(acc, vel, inD)

	fmt.Println("Enter value for time")

	var time float64

	fmt.Scanf("%f", &time)

	fmt.Println(calcD(time))

}
