package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Least squares method")
	var arrN = fillNArray()
	var xMatrix = fillXMatrix()

	fmt.Println(arrN)
	fmt.Println(xMatrix)
}

// количество экспериментов
func fillNArray() []float64 {
	var size = getArraySizeFromInput("array of N")

	var arrN = make([]float64, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Input N idx %d\n", i)

		_, err := fmt.Scan(&arrN[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	return arrN
}

// Fill x1, x2 like matrix
func fillXMatrix() [][]float64 {
	fmt.Println("Input count of elements (X) in array")

	var size = getArraySizeFromInput("multidimensional array of X")
	var value float64
	var matrix = make([][]float64, size)
	var err error

	for i := 0; i < size; i++ {
		fmt.Println("Array ", i+1, ":")
		for j := 0; j < size; j++ {
			fmt.Println("Array ", j+1, " число:")

			_, err = fmt.Scan(&value)
			if err != nil {
				log.Fatal(err)
			}

			matrix[i] = append(matrix[i], value)
		}
	}

	return matrix
}

func getArraySizeFromInput(arrayName string) int {
	var size int

	fmt.Println("Input count of elements (%s) in array", arrayName)
	_, err := fmt.Scan(&size)

	if err != nil {
		log.Fatal(err)
	}

	return size
}

/*func solveLeastSquares() {

}*/
