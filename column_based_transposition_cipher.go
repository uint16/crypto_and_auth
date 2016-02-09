package main

import (
	"fmt"
	"math"
)

func encrypt() {

}

func decrypt() {

}

func showMatrix(s string) {
	length := len(s)
	col := int(math.Ceil(math.Sqrt(float64(length))))
	row := int(math.Ceil(float64(length) / float64(col)))
	matrix := make([][]string, row)
	for x := range matrix {
		matrix[x] = make([]string, col)
	}
	l := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if length > l {
				matrix[i][j] = string(s[l])
			} else {
				matrix[i][j] = "\\"
			}
			fmt.Print(matrix[i][j], " ")
			l++
		}
		fmt.Println()
	}
}

func main() {
	var plainText string
	//	var cipherText, key string
	fmt.Println("Enter text to encrypt: ")
	fmt.Scanln(&plainText)

	showMatrix(plainText)
}
