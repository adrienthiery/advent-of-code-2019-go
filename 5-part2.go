package main

import "fmt"
import "log"
import "reflect"
import "strconv"

var test1 = []int{1002,4,3,4,33}
var expected1 = []int{1002,4,3,4,99}
var test2 = []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}
var test3 = []int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1}
var test4 = []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}
// var expected4 = []int{30,1,1,4,2,5,6,0,99}
var program = []int{3,225,1,225,6,6,1100,1,238,225,104,0,1102,88,66,225,101,8,125,224,101,-88,224,224,4,224,1002,223,8,223,101,2,224,224,1,224,223,223,1101,87,23,225,1102,17,10,224,101,-170,224,224,4,224,102,8,223,223,101,3,224,224,1,223,224,223,1101,9,65,225,1101,57,74,225,1101,66,73,225,1101,22,37,224,101,-59,224,224,4,224,102,8,223,223,1001,224,1,224,1,223,224,223,1102,79,64,225,1001,130,82,224,101,-113,224,224,4,224,102,8,223,223,1001,224,7,224,1,223,224,223,1102,80,17,225,1101,32,31,225,1,65,40,224,1001,224,-32,224,4,224,102,8,223,223,1001,224,4,224,1,224,223,223,2,99,69,224,1001,224,-4503,224,4,224,102,8,223,223,101,6,224,224,1,223,224,223,1002,14,92,224,1001,224,-6072,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,102,33,74,224,1001,224,-2409,224,4,224,1002,223,8,223,101,7,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,107,677,677,224,1002,223,2,223,1006,224,329,101,1,223,223,108,677,677,224,1002,223,2,223,1005,224,344,101,1,223,223,1007,677,677,224,1002,223,2,223,1006,224,359,101,1,223,223,1107,226,677,224,1002,223,2,223,1006,224,374,1001,223,1,223,8,677,226,224,1002,223,2,223,1006,224,389,101,1,223,223,1108,677,677,224,1002,223,2,223,1005,224,404,1001,223,1,223,7,226,226,224,1002,223,2,223,1006,224,419,101,1,223,223,1107,677,677,224,1002,223,2,223,1005,224,434,101,1,223,223,107,226,226,224,102,2,223,223,1005,224,449,101,1,223,223,107,677,226,224,1002,223,2,223,1006,224,464,1001,223,1,223,8,226,677,224,102,2,223,223,1006,224,479,1001,223,1,223,108,677,226,224,102,2,223,223,1005,224,494,1001,223,1,223,1108,677,226,224,1002,223,2,223,1005,224,509,1001,223,1,223,1107,677,226,224,1002,223,2,223,1005,224,524,101,1,223,223,1008,226,226,224,1002,223,2,223,1006,224,539,101,1,223,223,1008,226,677,224,1002,223,2,223,1005,224,554,1001,223,1,223,7,226,677,224,1002,223,2,223,1005,224,569,101,1,223,223,1007,677,226,224,1002,223,2,223,1006,224,584,1001,223,1,223,7,677,226,224,102,2,223,223,1006,224,599,101,1,223,223,1007,226,226,224,102,2,223,223,1006,224,614,101,1,223,223,1008,677,677,224,1002,223,2,223,1006,224,629,101,1,223,223,108,226,226,224,102,2,223,223,1006,224,644,101,1,223,223,1108,226,677,224,1002,223,2,223,1005,224,659,101,1,223,223,8,226,226,224,1002,223,2,223,1005,224,674,101,1,223,223,4,223,99,226}

func executeStep(program *[]int, baseIndex int, input int, output *int) int {
	instruction := (*program)[baseIndex]
	instructionAsString := strconv.Itoa(instruction)

	if instruction < 10 {
		instructionAsString = "0000" + instructionAsString
	} else if instruction < 99 {
		instructionAsString = "0" + instructionAsString
	} else if instruction > 99 && instruction < 9999 {
		instructionAsString = "0" + instructionAsString
	}
	fmt.Println("instruction", instructionAsString)

	// split instruction to read parameter modes
	runedInstruction := []rune(instructionAsString)
	opcodeString := string(runedInstruction[len(runedInstruction) - 2]) + string(runedInstruction[len(runedInstruction) - 1])
	opcode, err := strconv.Atoi(opcodeString)

	if err != nil {
		fmt.Println("strconv.Atoi ERROR")
	}

	parametersMode := []int{}

	// fmt.Println("opcode", opcode)

	if len(runedInstruction) > 2 {
		for j := len(runedInstruction) - 3; j >= 0; j-- {
			mode, err := strconv.Atoi(string(runedInstruction[j]))
			parametersMode = append(parametersMode, mode)

			if err != nil {
				fmt.Println("strconv.Atoi ERROR")
			}
		}
		// fmt.Println("parametersMode")
		// fmt.Println(parametersMode)
	}

	returnIndex := baseIndex

	if opcode == 99 {
		// end
		return -1
	} 
	
	if opcode == 1 {
		// add
		firstValue := 0
		secondValue := 0
		if parametersMode[0] == 0 {
			firstIndex := (*program)[baseIndex + 1]
			firstValue = (*program)[firstIndex]
		} else {
			firstValue = (*program)[baseIndex + 1]
		}
		if parametersMode[1] == 0 {
			secondIndex := (*program)[baseIndex + 2]
			secondValue = (*program)[secondIndex]
		} else {
			secondValue = (*program)[baseIndex + 2]
		}
		// Can the target also be with parametersMode 1 ?
		targetIndex := (*program)[baseIndex + 3]
		result := firstValue + secondValue
		returnIndex += 4
		(*program)[targetIndex] = result
	} else if opcode == 2 {
		// multiply
		firstValue := 0
		secondValue := 0
		if parametersMode[0] == 0 {
			firstIndex := (*program)[baseIndex + 1]
			firstValue = (*program)[firstIndex]
		} else {
			firstValue = (*program)[baseIndex + 1]
		}
		if parametersMode[1] == 0 {
			secondIndex := (*program)[baseIndex + 2]
			secondValue = (*program)[secondIndex]
		} else {
			secondValue = (*program)[baseIndex + 2]
		}
		// Can the target also be with parametersMode 1 ?
		targetIndex := (*program)[baseIndex + 3]
		result := firstValue * secondValue
		(*program)[targetIndex] = result
		returnIndex += 4
	} else if opcode == 3 {
		// read input
		targetIndex := (*program)[baseIndex + 1]
		(*program)[targetIndex] = input
		returnIndex += 2
	} else if opcode == 4 {
		// write output
		value := 0
		if parametersMode[0] == 0 {
			firstIndex := (*program)[baseIndex + 1]
			value = (*program)[firstIndex]
		} else {
			value = (*program)[baseIndex + 1]
		}
		*output = value
		fmt.Println("Writing to output, value is:", value, "(output is: ", *output, ")")
		returnIndex += 2
	} else if opcode == 5 {
		// jump-if-true
		firstValue := 0
		secondValue := 0
		if parametersMode[0] == 0 {
			firstIndex := (*program)[baseIndex + 1]
			firstValue = (*program)[firstIndex]
		} else {
			firstValue = (*program)[baseIndex + 1]
		}
		if parametersMode[1] == 0 {
			secondIndex := (*program)[baseIndex + 2]
			secondValue = (*program)[secondIndex]
		} else {
			secondValue = (*program)[baseIndex + 2]
		}
		if firstValue != 0 {
			returnIndex = secondValue
		} else {
			returnIndex += 3
		}
	} else if opcode == 6 {
		// jump-if-false
		firstValue := 0
		secondValue := 0
		if parametersMode[0] == 0 {
			firstIndex := (*program)[baseIndex + 1]
			firstValue = (*program)[firstIndex]
		} else {
			firstValue = (*program)[baseIndex + 1]
		}
		if parametersMode[1] == 0 {
			secondIndex := (*program)[baseIndex + 2]
			secondValue = (*program)[secondIndex]
		} else {
			secondValue = (*program)[baseIndex + 2]
		}
		if firstValue == 0 {
			returnIndex = secondValue
		} else {
			returnIndex += 3
		}
	}  else if opcode == 7 {
		// less than
		firstValue := 0
		secondValue := 0
		if parametersMode[0] == 0 {
			firstIndex := (*program)[baseIndex + 1]
			firstValue = (*program)[firstIndex]
		} else {
			firstValue = (*program)[baseIndex + 1]
		}
		if parametersMode[1] == 0 {
			secondIndex := (*program)[baseIndex + 2]
			secondValue = (*program)[secondIndex]
		} else {
			secondValue = (*program)[baseIndex + 2]
		}
		thirdIndex := (*program)[baseIndex + 3]

		if firstValue < secondValue {
			(*program)[thirdIndex] = 1
		} else {
			(*program)[thirdIndex] = 0
		}
		returnIndex += 4
	} else if opcode == 8 {
		// less than
		firstValue := 0
		secondValue := 0
		if parametersMode[0] == 0 {
			firstIndex := (*program)[baseIndex + 1]
			firstValue = (*program)[firstIndex]
		} else {
			firstValue = (*program)[baseIndex + 1]
		}
		if parametersMode[1] == 0 {
			secondIndex := (*program)[baseIndex + 2]
			secondValue = (*program)[secondIndex]
		} else {
			secondValue = (*program)[baseIndex + 2]
		}
		thirdIndex := (*program)[baseIndex + 3]

		if firstValue == secondValue {
			(*program)[thirdIndex] = 1
		} else {
			(*program)[thirdIndex] = 0
		}
		returnIndex += 4
	} else {
		// error
		fmt.Println("opcode")
		fmt.Println(opcode)
		fmt.Println("returnIndex")
		fmt.Println(returnIndex)
		log.Fatal("Something went wrong")
	}

	return returnIndex
}

func executeProgram(program []int, input int, output *int) []int {
	fmt.Println(program)
	returnedProgram := program
	returnedIndex := 0

	for returnedIndex != -1 {
		returnedIndex = executeStep(&returnedProgram, returnedIndex, input, output)
	}

	return returnedProgram
}

func main() {
	testInput := 0
	testOutput := 99999
	// "TESTS"
	fmt.Println("Test1 passing:")
	result1 := executeProgram(test1, testInput, &testOutput)
	fmt.Println(reflect.DeepEqual(result1, expected1))
	
	fmt.Println("------")

	fmt.Println("Test2 passing:")
	inputZero := 0
	test2Output := 999999
	executeProgram(test2, inputZero, &test2Output)
	fmt.Println(reflect.DeepEqual(test2Output, 0))

	inputNonZero := 12
	test2Output2 := 999999
	executeProgram(test2, inputNonZero, &test2Output2)
	fmt.Println(reflect.DeepEqual(test2Output2, 1))
	
	fmt.Println("------")

	fmt.Println("Test3 passing:")
	test3Output := 999999
	executeProgram(test3, inputZero, &test3Output)
	fmt.Println(reflect.DeepEqual(test3Output, 0))

	test3Output2 := 999999
	executeProgram(test3, inputNonZero, &test3Output2)
	fmt.Println(reflect.DeepEqual(test3Output2, 1))
	
	fmt.Println("------")

	fmt.Println("Test4 passing:")
	input4One := 4
	test4Output := 999999
	executeProgram(test4, input4One, &test4Output)
	fmt.Println(reflect.DeepEqual(test4Output, 999))

	input4Two := 8
	test4Output2 := 999999
	executeProgram(test4, input4Two, &test4Output2)
	fmt.Println(reflect.DeepEqual(test4Output2, 1000))

	input4Three := 10
	test4Output3 := 999999
	executeProgram(test4, input4Three, &test4Output3)
	fmt.Println(reflect.DeepEqual(test4Output3, 1001))
	
	fmt.Println("------")

	testInputReal := 5
	output := 99999999
	fmt.Println("Real Deal:")
	executeProgram(program, testInputReal, &output)
	fmt.Println("Output", output)

	fmt.Println("------")
}