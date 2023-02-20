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

// experiments count
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

// Fill x1, x2, y like matrix
func fillXMatrix() [][][]float64 {
	fmt.Println("Input count of elements (X) in array")

	var size = getArraySizeFromInput("multidimensional array of X")
	var matrix = make([][][]float64, size)

	for i := range matrix {
		matrix[i] = make([][]float64, size)
		for j := range matrix {
			matrix[i][j] = make([]float64, size)
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				fmt.Println("Array x1-x2-y values:")

				_, err := fmt.Scan(&matrix[i][j][k])
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	return matrix
}

func getArraySizeFromInput(arrayName string) int {
	var size int

	fmt.Printf("Input count of elements (%s) in array\n", arrayName)
	_, err := fmt.Scan(&size)

	if err != nil {
		log.Fatal(err)
	}

	return size
}

/*func solveLeastSquares() {

}*/
