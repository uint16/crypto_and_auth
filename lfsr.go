package main

import (
	"fmt"
	"math"
)

func formattedPrint(counter int, value uint) {
	temp := ""
	if value < 2048 {
		for i := 0; i <= 10; i++ {
			if value < uint(math.Pow(float64(2), float64(i))) {
				temp += "0"
			}
		}

		fmt.Printf("%d: %s%b, %d\n", counter, temp, value, value)
	}
}

func main() {
	fmt.Println("LFSR for a 12 bit state, with taps x^12 + x^11 + x10 + x^4 + 1")
	var temp, lfsr uint
	lfsr = 3592 						//Decimal representation of plynomial 111000001000-1
	counter := 1
	first_run := true

	for temp != lfsr || first_run{
		if counter == 1 {
			temp = lfsr
			first_run = false
		}
		lsb := temp & 1					//the output bit
		temp >>= 1						//shift register
		if lsb == 1 {
			temp ^= lfsr 				//toggle masks corresponding to taps
		}

		formattedPrint(counter, temp)
		counter++
	}
	
	fmt.Printf("%d: %b, %d\n", counter-1, temp, temp)
}
