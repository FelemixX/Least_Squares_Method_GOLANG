package main

import (
	"fmt"
	"log"
)

type Point struct {
	x float32
	y float32
}

func main() {
	calculateLeastSquares()
}

func calculateLeastSquares() {
	pointsData := createPoints()

	n := float32(len(pointsData)) //n == size

	var sumX float32
	var sumY float32
	var sumXY float32
	var sumXX float32

	for _, p := range pointsData {
		sumX += p.x
		sumY += p.y
		sumXY += p.x * p.y
		sumXX += p.x * p.x
	}

	base := (n*sumXX - sumX*sumX)
	a := (n*sumXY - sumX*sumY) / base
	b := (sumXX*sumY - sumXY*sumX) / base

	fmt.Printf("a = %g, b = %g", a, b)
}

func createPoints() []Point {
	var pointX float32
	var pointY float32
	countX := getNumberFromInput("Count of X")
	countY := getNumberFromInput("Count of Y")

	points := make([]Point, 0, countX*countY)

	for i := 0; i < countY; i++ {
		fmt.Println("Input point Y ")

		_, errY := fmt.Scan(&pointY)
		if errY != nil {
			return nil
		}

		for i := 0; i < countX; i++ {
			fmt.Println("Input point X ")

			_, errX := fmt.Scan(&pointX)
			if errX != nil {
				return nil
			}

			points = append(points, Point{x: pointX, y: pointY})
		}
	}
	fmt.Println(points)
	return points
}

func getNumberFromInput(arrayContents string) int {
	var size int

	fmt.Printf("Input count of elements (%s) in array\n", arrayContents)
	_, err := fmt.Scan(&size)

	if err != nil {
		processError(err, "display")
	}

	return size
}

func processError(err error, action string) {
	switch action {
	case "fatal":
		log.Fatal(err)
	case "print":
		log.Println(err)
	default:
		return
	}
}
