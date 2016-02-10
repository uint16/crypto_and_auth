package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"os"
	"strings"
)

type dimensions struct {
	length, col, row int
}
type myCipher struct {
	plainText, cipherText, currentText string
	cipherTextIdx                      []int
	plainTextMatrix                    [][]string
	cipherTextMatrix                   [][]string
	d                                  dimensions
}

func (c *myCipher) encrypt(key string) {
	c.hashKey(key)
	c.cipherText = ""
	for i := 0; i < c.d.col; i++ {
		for j := 0; j < c.d.row; j++ {
			c.cipherTextMatrix[j][i] = c.plainTextMatrix[j][c.cipherTextIdx[i]]
			c.cipherText += c.cipherTextMatrix[j][i]
		}
	}
}

func (c *myCipher) decrypt(key string) {
	c.hashKey(key)
	c.plainText = ""
	for i := 0; i < c.d.col; i++ {
		for j := 0; j < c.d.row; j++ {
			c.plainTextMatrix[j][i] = c.cipherTextMatrix[j][c.cipherTextIdx[i]]
			c.plainText += c.plainTextMatrix[j][i]
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
			if l < c.d.length {
				c.plainTextMatrix[i][j] = string(s[l])
			} else {
				c.plainTextMatrix[i][j] = "*"
			}
			l++
		}
	}
	c.showMatrix(true)
}

func (c *myCipher) showMatrix(plain bool) {
	fmt.Println()
	for i := 0; i < c.d.row; i++ {
		for j := 0; j < c.d.col; j++ {
			if plain {
				fmt.Print(c.plainTextMatrix[i][j], " ")
			} else {
				fmt.Print(c.cipherTextMatrix[i][j], " ")
			}
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
	c.cipherTextIdx = make([]int, c.d.col)

	for x := range c.plainTextMatrix {
		c.plainTextMatrix[x] = make([]string, c.d.col)
		c.cipherTextMatrix[x] = make([]string, c.d.col)
	}
}

func (c *myCipher) showPlainText() {
	temp := strings.Replace(c.plainText, "*", "", -1)
	fmt.Println("Plain Text: ", temp)
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

	in := bufio.NewReader(os.Stdin)

	var userInput int
	for {
		options()

		fmt.Scanf("%d", &userInput)

		if userInput == 1 {
			fmt.Print("Enter text to encrypt: ")
			line, _, _ := in.ReadLine()
			test.plainText = string(line)
			test.createMatrix(test.plainText)
			fmt.Print("Enter encryption key: ")
			fmt.Scanln(&key)
			fmt.Println()
			test.encrypt(key)
			test.showCipherText()
			continue
		} else if userInput == 2 {
			fmt.Print("Enter text to decrypt: ")
			line, _, _ := in.ReadLine()
			test.cipherText = string(line)
			//test.createMatrix(test.cipherText)
			fmt.Print("Enter key: ")
			fmt.Scanln(&key)
			fmt.Println()
			test.encrypt(key)
			test.showPlainText()
			continue
		} else if userInput == 3 {
			if test.plainText != "" {
				test.showMatrix(true)
			} else {
				fmt.Println("ERROR: Plain text not provided")
			}
		} else if userInput == 4 {
			if test.plainText != "" {
				test.showMatrix(false)
			} else {
				fmt.Println("ERROR: Cipher text not provided")
			}
		} else if userInput == 5 {
			options()
		} else if userInput == 6 {
			break
		} else {
			options()
			continue
		}
	}

}
