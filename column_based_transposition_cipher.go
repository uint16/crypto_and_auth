package main

import (
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
)

type dimensions struct {
	length, col, row int
}
type myCipher struct {
	plainText, cipherText, currentText string
	cipherTextArr                      []string
	cipherTextIdx                      []int
	plainTextMatrix                    [][]string
	cipherTextMatrix                   [][]string
	d                                  dimensions
}

func (c *myCipher) encrypt(key string) {
	c.hashKey(key)
	for i := 0; i < c.d.row; i++ {
		for j := 0; j < len(c.cipherTextIdx); j++ {
			c.cipherTextMatrix[i][j] = c.plainTextMatrix[i][c.cipherTextIdx[j]]
			c.cipherText += c.cipherTextMatrix[i][j]
		}
		c.cipherText += " "
	}
}

func (c *myCipher) decrypt(key string) {
	c.hashKey(key)
	for i := 0; i < len(c.cipherTextIdx); i++ {
		for j := 0; j < c.d.row; j++ {
			c.cipherTextMatrix[i][j] = c.plainTextMatrix[i][c.cipherTextIdx[j]]
			c.cipherText += c.plainTextMatrix[i][j]
		}
	}
}

func (c *myCipher) hashKey(key string) {
	k := []byte(key)
	hash := fnv.New64a()
	hash.Write(k)
	rand.Seed(int64(hash.Sum64()))
	c.cipherTextIdx = rand.Perm(c.d.col)
}

func (c *myCipher) createMatrix(s string) {
	c.init(s)
	l := 0
	for i := 0; i < c.d.row; i++ {
		for j := 0; j < c.d.col; j++ {
			if c.d.length > l {
				c.plainTextMatrix[i][j] = string(s[l])
			} else {
				c.plainTextMatrix[i][j] = "\\"
			}
			l++
		}
	}
	c.showMatrix()
}

func (c *myCipher) showMatrix() {
	fmt.Println()
	for i := 0; i < c.d.row; i++ {
		for j := 0; j < c.d.col; j++ {
			fmt.Print(c.plainTextMatrix[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (c *myCipher) init(s string) {
	c.d.length = len(s)
	c.d.col = int(math.Ceil(math.Sqrt(float64(c.d.length))))
	c.d.row = int(math.Ceil(float64(c.d.length) / float64(c.d.col)))
	c.plainTextMatrix = make([][]string, c.d.row)
	c.cipherTextMatrix = make([][]string, c.d.row)
	c.cipherTextArr = make([]string, c.d.col)
	c.cipherTextIdx = make([]int, c.d.col)

	for x := range c.plainTextMatrix {
		c.plainTextMatrix[x] = make([]string, c.d.col)
		c.cipherTextMatrix[x] = make([]string, c.d.col)
	}
}

func (c *myCipher) showPlainText() {
	fmt.Println("Plain Text: ", c.plainText)
	fmt.Println()
}

func (c *myCipher) showCipherText() {
	fmt.Println("Cipher Text: ", c.cipherText)
	fmt.Println()
}

func options() {
	fmt.Println()
	fmt.Println("Please select an action to perform")
	fmt.Println("1. Provide text to encrypt")
	fmt.Println("2. Select text to decrypt")
	fmt.Println("3. View plain text matrix")
	fmt.Println("4. View cipher text matrix")
	fmt.Println("5. View options")
	fmt.Println("6. Exit")
	fmt.Println()
}

func main() {
	test := new(myCipher)
	var key string
	fmt.Println("Select action: ")

	var userInput int
	for {
		options()

		fmt.Scanf("%d", &userInput)

		if userInput == 1 {
			fmt.Println("Enter text to encrypt: ")
			fmt.Scanln(&test.plainText)
			test.createMatrix(test.plainText)
			fmt.Println("Enter encryption key: ")
			fmt.Scanln(&key)
			fmt.Println()
			test.encrypt(key)
			test.showCipherText()
			continue
		} else if userInput == 2 {
			fmt.Println("Enter text to decrypt: ")
			fmt.Scanln(&test.cipherText)
			test.createMatrix(test.cipherText)
			fmt.Println("Enter key: ")
			fmt.Scanln(&key)
			fmt.Println()
			test.encrypt(key)
			test.showCipherText()
			continue
		} else if userInput == 3 {

		} else if userInput == 4 {

		} else if userInput == 5 {

		} else if userInput == 6 {
			break
		} else {
			options()
			continue
		}
	}

}
