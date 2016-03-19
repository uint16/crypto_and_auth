package main

import (
	"fmt"
)

func toString(counter int, arr [12]bool) {
	res := ""
	for i := 0; i < len(arr); i++ {
		if arr[i] {
			res += "1"
		} else {
			res += "0"
		}
	}
	fmt.Println(counter, ". ", res)

}

/*compute the next bit by
tapping at 12,11,10 and 4 */
func tap(arr [12]bool) bool {

	//indices have been reversed for simplicity, 12=0, 11=1. 10=2, 4=8
	firstTap := arr[0] != arr[1]
	secondTap := firstTap != arr[2]
	thirdTap := secondTap != arr[8]

	return thirdTap
}

/*
Receive a new bit from the tap fuction, move over all other bits and append the new bit
*/
func shift(arr [12]bool) [12]bool {
	newBit := tap(arr)
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	arr[len(arr)-1] = newBit
	
	return arr
}



func main() {
	arr := [12]bool{true, true, true, false, false, false, false, false, true, false, false, false}
	lfsr := [12]bool{true, true, true, false, false, false, false, false, true, false, false, false}
	counter := 1
	firstRun := true

	for arr != lfsr || firstRun {
		if firstRun {
			arr = lfsr
			firstRun = false
		}
		arr = shift(arr)
		toString(counter, arr)
		counter++
	}
}

/**
The commented out code below is the LFRS for the same problem implemented using bitwise operators provided by Golang
**/


/*
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
	lfsr = 3592 //Decimal representation of polynomial 111000001000-1
	counter := 1
	first_run := true

	for temp != lfsr || first_run {
		if counter == 1 {
			temp = lfsr
			first_run = false
		}
		lsb := temp & 1 //the output bit
		temp >>= 1      //shift register
		if lsb == 1 {
			temp ^= lfsr //toggle masks corresponding to taps
		}

		formattedPrint(counter, temp)
		counter++
	}

	fmt.Printf("%d: %b, %d\n", counter-1, temp, temp)
}
*/
