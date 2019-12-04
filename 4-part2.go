package main

import "fmt"
import "reflect"
import "strconv"

var test1 = 111111 // meets these criteria (double 11, never decreases).
var test2 = 223450 // does not meet these criteria (decreasing pair of digits 50).
var test3 = 123789 // does not meet these criteria (no double).
var test4 = 112233 // works
var test5 = 123444 // does not meet these criteria (44 is part of 444).
var test6 = 111122 // works
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
	twoAdjacentNumbersAreTheSameButNotPartOfBiggerGroup := false
	stringPass := strconv.Itoa(password)
	runes := []rune(stringPass)

	for i := 0; i < len(runes) - 1; i++ {
		if string(runes[i]) == string(runes[i + 1]) {
			fmt.Println("It seems two adjacent numbers are the same", string(runes[i]), string(runes[i + 1]))

			previousCharIsDifferent := true
			nextCharIsDifferent := true

			if i > 0 && i < len(runes) - 2 {
				// A number in the middle, check prev & next
				previousCharIsDifferent = string(runes[i]) != string(runes[i - 1])
				nextCharIsDifferent = string(runes[i + 2]) != string(runes[i])
				fmt.Println("Previous and next chars are", string(runes[i - 1]), string(runes[i + 2]))
			} else if i == 0 {
				// First char, let's check the next char only
				nextCharIsDifferent = string(runes[i + 2]) != string(runes[i])
				fmt.Println("Next char is", string(runes[i + 2]))
			} else {
				// The previous to last char, let's check before just in case
				previousCharIsDifferent = string(runes[i]) != string(runes[i - 1])
				fmt.Println("Previous char is", string(runes[i - 1]))
			}

			// the two adjacent matching digits are not part of a larger group of matching digits
			if previousCharIsDifferent && nextCharIsDifferent {
				twoAdjacentNumbersAreTheSameButNotPartOfBiggerGroup = true
				break
			}
		}
	}
	if !twoAdjacentNumbersAreTheSameButNotPartOfBiggerGroup {
		fmt.Println(password, " does not have two adjacent numbers being the same, or they are part of a bigger group")
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
	fmt.Println("Is test passing:", reflect.DeepEqual(result1, false))
	
	fmt.Println("------")

	fmt.Println("Test2")
	result2 := testPassword(test2, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result2, false))

	fmt.Println("------")

	fmt.Println("Test3")
	result3 := testPassword(test3, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result3, false))

	fmt.Println("------")

	fmt.Println("Test4")
	result4 := testPassword(test4, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result4, true))

	fmt.Println("------")

	fmt.Println("Test5")
	result5 := testPassword(test5, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result5, false))

	fmt.Println("------")

	fmt.Println("Test6")
	result6 := testPassword(test6, testRange)
	fmt.Println("Is test passing:", reflect.DeepEqual(result6, true))

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