package main

import "fmt"
import "reflect"
import "strconv"

var test1 = 111111 // meets these criteria (double 11, never decreases).
var test2 = 223450 // does not meet these criteria (decreasing pair of digits 50).
var test3 = 123789 // does not meet these criteria (no double).
var testRange = []int{99999, 333333}
var inputRange = []int{387638, 919123} // Puzzle input

func testPassword(password int, theRange []int) bool {
	fmt.Println("Testing ", password)

	// It is a six-digit number.
	if password <= 99999 {
		fmt.Println(password, " not a 5 digit number")
		return false
	}
	// The value is within the range given in your puzzle input.
	if password < theRange[0] || password > theRange[1] {
		fmt.Println(password, " not in range")
		return false
	}
	// Two adjacent digits are the same (like 22 in 122345).
	twoAdjacentNumbersAreTheSame := false
	stringPass := strconv.Itoa(password)
	runes := []rune(stringPass)

	for i := 0; i < len(runes) - 1; i++ {
		if string(runes[i]) == string(runes[i + 1]) {
			fmt.Println("It seems two adjacent numbers the same", string(runes[i]), string(runes[i + 1]))
			twoAdjacentNumbersAreTheSame = true
			break
		}
	}
	if !twoAdjacentNumbersAreTheSame {
		fmt.Println(password, " does not have two adjacent numbers being the same")
		return false
	}
	// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	for j := 0; j < len(runes) - 1; j++ {
		if string(runes[j]) > string(runes[j + 1]) {
			fmt.Println(password, "decreases")
			return false
		}
	}

	return true
}

func main() {
	// "TESTS"
	fmt.Println("Test1")
	result1 := testPassword(test1, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result1, true))
	
	fmt.Println("------")

	fmt.Println("Test2")
	result2 := testPassword(test2, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result2, false))

	fmt.Println("------")

	fmt.Println("Test3")
	result3 := testPassword(test3, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result3, false))

	fmt.Println("------")

	fmt.Println("For real")
	count := 0

	for k := 0; k < inputRange[1] - inputRange[0]; k++ {
		isPassing := testPassword(inputRange[0] + k, inputRange)

		if isPassing {
			count++
		}
	}
	fmt.Println("Number of possibilities passing:", count)
}