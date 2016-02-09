package main

import (
	"fmt"
	"math"
)

type myCipher struct {
	plainText, cipherText, currentText string
	cipherTextArr                      []string
	matrix                             [][]string
}

func (c *myCipher) encrypt(key string) {

}

func (c *myCipher) decrypt(index int) {
	fmt.Println(c.cipherTextArr)
}

func (c *myCipher) showMatrix(s string) {

	length, col, row := getMatrixDimensions(s)
	c.matrix = make([][]string, row)
	c.cipherTextArr = make([]string, col)
	for x := range c.matrix {
		c.matrix[x] = make([]string, col)
	}
	l := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if length > l {
				c.matrix[i][j] = string(s[l])
			} else {
				c.matrix[i][j] = "\\"
			}
			fmt.Print(c.matrix[i][j], " ")
			l++
		}
		fmt.Println()
	}
}

func getMatrixDimensions(s string) (int, int, int) {
	length := len(s)
	col := int(math.Ceil(math.Sqrt(float64(length))))
	row := int(math.Ceil(float64(length) / float64(col)))

	return length, col, row
}

func options() {
	fmt.Println("Please select an action to perform")
	fmt.Println("1. Provide text to encrypt")
	fmt.Println("2. Select text to decrypt")
	fmt.Println("3. View plain text matrix")
	fmt.Println("4. View cipher text matrix")
}

func main() {
	test := new(myCipher)
	var key string
	fmt.Println("Enter text to encrypt: ")
	fmt.Scanln(&test.plainText)
	test.showMatrix(test.plainText)
	fmt.Println("Enter encryption key: ")

	fmt.Scanln(&key)

	test.decrypt(0)

}
