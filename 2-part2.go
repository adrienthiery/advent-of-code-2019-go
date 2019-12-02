package main

import "fmt"
import "log"
import "reflect"

var exitProgram = []int{0}

var test1 = []int{1,0,0,0,99}
var expected1 = []int{2,0,0,0,99}
var test2 = []int{2,3,0,3,99}
var expected2 = []int{2,3,0,6,99}
var test3 = []int{2,4,4,5,99,0}
var expected3 = []int{2,4,4,5,99,9801}
var test4 = []int{1,1,1,4,99,5,6,0,99}
var expected4 = []int{30,1,1,4,2,5,6,0,99}

func executeStep(program []int, stepNumber int) []int {
	baseIndex := stepNumber * 4
	opcode := program[baseIndex]

	if opcode == 99 {
		// end
		return exitProgram
	} 

	firstIndex := program[baseIndex + 1]
	secondIndex := program[baseIndex + 2]
	targetIndex := program[baseIndex + 3]

	firstValue := program[firstIndex]
	secondValue := program[secondIndex]
	
	if opcode == 1 {
		// add
		result := firstValue + secondValue
		program[targetIndex] = result
	} else if opcode == 2 {
		// multiply
		result := firstValue * secondValue
		program[targetIndex] = result
	} else {
		// error
		fmt.Println("opcode")
		fmt.Println(opcode)
		fmt.Println("stepNumber")
		fmt.Println(stepNumber)
		log.Fatal("Something went wrong")
	}

	return program
}

func executeProgram(program []int) []int {
	returnCode := program
	returnedProgram := program
	i := 0

	for len(returnCode) != 1 {
		returnCode = executeStep(program, i)

		if len(returnCode) != 1 {
			returnedProgram = returnCode
		}

		i++
	}

	return returnedProgram
}

func main() {
	// "TESTS"
	fmt.Println("Test1 passing:")
	result1 := executeProgram(test1)
	fmt.Println(reflect.DeepEqual(result1, expected1))
	
	fmt.Println("------")

	fmt.Println("Test2 passing:")
	result2 := executeProgram(test2)
	fmt.Println(reflect.DeepEqual(result2, expected2))

	fmt.Println("------")

	fmt.Println("Test3 passing:")
	result3 := executeProgram(test3)
	fmt.Println(reflect.DeepEqual(result3, expected3))

	fmt.Println("------")

	fmt.Println("Test4 passing:")
	result4 := executeProgram(test4)
	fmt.Println(reflect.DeepEqual(result4, expected4))

	fmt.Println("------")
	fmt.Println("Real deal")
	fmt.Println("------")

	output := 19690720
	potentialOutput := 0

	finalNoun := 0
	finalVerb := 0
	
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			newInput := []int{1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,1,19,6,23,2,23,13,27,1,27,5,31,2,31,10,35,1,9,35,39,1,39,9,43,2,9,43,47,1,5,47,51,2,13,51,55,1,55,9,59,2,6,59,63,1,63,5,67,1,10,67,71,1,71,10,75,2,75,13,79,2,79,13,83,1,5,83,87,1,87,6,91,2,91,13,95,1,5,95,99,1,99,2,103,1,103,6,0,99,2,14,0,0}
			newInput[1] = noun
			newInput[2] = verb

			result := executeProgram(newInput)
			potentialOutput = result[0]
			
			fmt.Println(noun, verb, potentialOutput, "(", output, ")")

			if potentialOutput == output {
				finalNoun = noun
				finalVerb = verb
			}
		}
	}

	fmt.Println("Valid inputs are: (noun, verb) ")
	fmt.Println(finalNoun)
	fmt.Println(finalVerb)

	fmt.Println("100 * noun + verb is:")
	fmt.Println(100 * finalNoun + finalVerb)
}