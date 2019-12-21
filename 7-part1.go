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
// var input = 5
// var program = []int{3,225,1,225,6,6,1100,1,238,225,104,0,1102,88,66,225,101,8,125,224,101,-88,224,224,4,224,1002,223,8,223,101,2,224,224,1,224,223,223,1101,87,23,225,1102,17,10,224,101,-170,224,224,4,224,102,8,223,223,101,3,224,224,1,223,224,223,1101,9,65,225,1101,57,74,225,1101,66,73,225,1101,22,37,224,101,-59,224,224,4,224,102,8,223,223,1001,224,1,224,1,223,224,223,1102,79,64,225,1001,130,82,224,101,-113,224,224,4,224,102,8,223,223,1001,224,7,224,1,223,224,223,1102,80,17,225,1101,32,31,225,1,65,40,224,1001,224,-32,224,4,224,102,8,223,223,1001,224,4,224,1,224,223,223,2,99,69,224,1001,224,-4503,224,4,224,102,8,223,223,101,6,224,224,1,223,224,223,1002,14,92,224,1001,224,-6072,224,4,224,102,8,223,223,101,5,224,224,1,223,224,223,102,33,74,224,1001,224,-2409,224,4,224,1002,223,8,223,101,7,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,107,677,677,224,1002,223,2,223,1006,224,329,101,1,223,223,108,677,677,224,1002,223,2,223,1005,224,344,101,1,223,223,1007,677,677,224,1002,223,2,223,1006,224,359,101,1,223,223,1107,226,677,224,1002,223,2,223,1006,224,374,1001,223,1,223,8,677,226,224,1002,223,2,223,1006,224,389,101,1,223,223,1108,677,677,224,1002,223,2,223,1005,224,404,1001,223,1,223,7,226,226,224,1002,223,2,223,1006,224,419,101,1,223,223,1107,677,677,224,1002,223,2,223,1005,224,434,101,1,223,223,107,226,226,224,102,2,223,223,1005,224,449,101,1,223,223,107,677,226,224,1002,223,2,223,1006,224,464,1001,223,1,223,8,226,677,224,102,2,223,223,1006,224,479,1001,223,1,223,108,677,226,224,102,2,223,223,1005,224,494,1001,223,1,223,1108,677,226,224,1002,223,2,223,1005,224,509,1001,223,1,223,1107,677,226,224,1002,223,2,223,1005,224,524,101,1,223,223,1008,226,226,224,1002,223,2,223,1006,224,539,101,1,223,223,1008,226,677,224,1002,223,2,223,1005,224,554,1001,223,1,223,7,226,677,224,1002,223,2,223,1005,224,569,101,1,223,223,1007,677,226,224,1002,223,2,223,1006,224,584,1001,223,1,223,7,677,226,224,102,2,223,223,1006,224,599,101,1,223,223,1007,226,226,224,102,2,223,223,1006,224,614,101,1,223,223,1008,677,677,224,1002,223,2,223,1006,224,629,101,1,223,223,108,226,226,224,102,2,223,223,1006,224,644,101,1,223,223,1108,226,677,224,1002,223,2,223,1005,224,659,101,1,223,223,8,226,226,224,1002,223,2,223,1005,224,674,101,1,223,223,4,223,99,226}
var program = []int{3,8,1001,8,10,8,105,1,0,0,21,34,43,60,81,94,175,256,337,418,99999,3,9,101,2,9,9,102,4,9,9,4,9,99,3,9,102,2,9,9,4,9,99,3,9,102,4,9,9,1001,9,4,9,102,3,9,9,4,9,99,3,9,102,4,9,9,1001,9,2,9,1002,9,3,9,101,4,9,9,4,9,99,3,9,1001,9,4,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,99}


func getValue(program *[]int, baseIndex int, parameterMode int, increment int) int {
	if parameterMode == 0 {
		index := (*program)[baseIndex + increment]
		return (*program)[index]
	} else {
		return (*program)[baseIndex + increment]
	}
}

func executeStep(program *[]int, baseIndex int, inputs []int, inputPointer *int, output *int, debug bool) int {
	instruction := (*program)[baseIndex]
	instructionAsString := strconv.Itoa(instruction)

	if instruction < 10 {
		instructionAsString = "0000" + instructionAsString
	} else if instruction < 99 {
		instructionAsString = "0" + instructionAsString
	} else if instruction > 99 && instruction < 9999 {
		instructionAsString = "0" + instructionAsString
	}

	if debug {
		fmt.Println("instruction", instructionAsString)
	}

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
		// fmt.Println("parametersMode", parametersMode)
	}

	returnIndex := baseIndex

	if opcode == 99 {
		// end
		return -1
	} 
	
	if opcode == 1 {
		// add
		firstValue := getValue(program, baseIndex, parametersMode[0], 1)
		secondValue := getValue(program, baseIndex, parametersMode[1], 2)
		targetIndex := (*program)[baseIndex + 3]
		result := firstValue + secondValue
		returnIndex += 4
		(*program)[targetIndex] = result
	} else if opcode == 2 {
		// multiply
		firstValue := getValue(program, baseIndex, parametersMode[0], 1)
		secondValue := getValue(program, baseIndex, parametersMode[1], 2)
		// Can the target also be with parametersMode 1 ?
		targetIndex := (*program)[baseIndex + 3]
		result := firstValue * secondValue
		(*program)[targetIndex] = result
		returnIndex += 4
	} else if opcode == 3 {
		// read input
		targetIndex := (*program)[baseIndex + 1]
		(*program)[targetIndex] = inputs[*inputPointer]
		*inputPointer += 1
		returnIndex += 2
	} else if opcode == 4 {
		// write output
		*output = getValue(program, baseIndex, parametersMode[0], 1)
		// fmt.Println("Writing to output, value is:", *output)
		returnIndex += 2
	} else if opcode == 5 {
		// jump-if-true
		firstValue := getValue(program, baseIndex, parametersMode[0], 1)
		secondValue := getValue(program, baseIndex, parametersMode[1], 2)
		if firstValue != 0 {
			returnIndex = secondValue
		} else {
			returnIndex += 3
		}
	} else if opcode == 6 {
		// jump-if-false
		firstValue := getValue(program, baseIndex, parametersMode[0], 1)
		secondValue := getValue(program, baseIndex, parametersMode[1], 2)
		if firstValue == 0 {
			returnIndex = secondValue
		} else {
			returnIndex += 3
		}
	}  else if opcode == 7 {
		// less than
		firstValue := getValue(program, baseIndex, parametersMode[0], 1)
		secondValue := getValue(program, baseIndex, parametersMode[1], 2)
		thirdIndex := (*program)[baseIndex + 3]

		if firstValue < secondValue {
			(*program)[thirdIndex] = 1
		} else {
			(*program)[thirdIndex] = 0
		}
		returnIndex += 4
	} else if opcode == 8 {
		// less than
		firstValue := getValue(program, baseIndex, parametersMode[0], 1)
		secondValue := getValue(program, baseIndex, parametersMode[1], 2)
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

func executeProgram(program []int, inputs []int, output *int, debug bool) []int {
	if debug {
		fmt.Println(program)
	}
	returnedProgram := program
	returnedIndex := 0
	inputPointer := 0

	for returnedIndex != -1 {
		returnedIndex = executeStep(&returnedProgram, returnedIndex, inputs, &inputPointer, output, debug)
	}

	return returnedProgram
}

type PhaseSet = [5]int

func main() {
	testInput := 0
	testOutput := 99999
	// "TESTS"
	fmt.Println("Test1:")
	result1 := executeProgram(test1, []int{testInput}, &testOutput, false)
	fmt.Println(" passing: ", reflect.DeepEqual(result1, expected1))
	
	fmt.Println("------")

	fmt.Println("Test2:")
	inputZero := 0
	test2Output := 999999
	executeProgram(test2, []int{inputZero}, &test2Output, false)
	fmt.Println(" passing: ", reflect.DeepEqual(test2Output, 0))

	inputNonZero := 12
	test2Output2 := 999999
	executeProgram(test2, []int{inputNonZero}, &test2Output2, false)
	fmt.Println(" passing: ", reflect.DeepEqual(test2Output2, 1))
	
	fmt.Println("------")

	fmt.Println("Test3:")
	test3Output := 999999
	executeProgram(test3, []int{inputZero}, &test3Output, false)
	fmt.Println(" passing: ", reflect.DeepEqual(test3Output, 0))

	test3Output2 := 999999
	executeProgram(test3, []int{inputNonZero}, &test3Output2, false)
	fmt.Println(" passing: ", reflect.DeepEqual(test3Output2, 1))
	
	fmt.Println("------")

	fmt.Println("Test4:")
	input4One := 4
	test4Output := 999999
	executeProgram(test4, []int{input4One}, &test4Output, false)
	fmt.Println(" passing: ", reflect.DeepEqual(test4Output, 999))

	input4Two := 8
	test4Output2 := 999999
	executeProgram(test4, []int{input4Two}, &test4Output2, false)
	fmt.Println(" passing: ", reflect.DeepEqual(test4Output2, 1000))

	input4Three := 10
	test4Output3 := 999999
	executeProgram(test4, []int{input4Three}, &test4Output3, false)
	fmt.Println(" passing: ", reflect.DeepEqual(test4Output3, 1001))
	
	fmt.Println("------")

	fmt.Println("Real Deal:")
	testSets := []PhaseSet{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					for m := 0; m < 5; m++ {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							testSets = append(testSets, PhaseSet{i, j, k, l, m})
						}
					}
				}
			}
		}
	}

	// if len(testSets) < 256 {
	// 	panic("Not enough sets")
	// }

	max := 0

	for index := 0; index < len(testSets); index++ {
		phases := testSets[index]
		fmt.Println("phases", phases)
		output := 99999999	
		output1 := 99999999
		output2 := 99999999
		output3 := 99999999
		output4 := 99999999
		
		// Amplifier 1
		inputs := []int{phases[0], 0}
		executeProgram(program, inputs, &output1, false)
		// Amplifier 2
		inputs = []int{phases[1], output1}
		executeProgram(program, inputs, &output2, false)
		// Amplifier 3
		inputs = []int{phases[2], output2}
		executeProgram(program, inputs, &output3, false)
		// Amplifier 4
		inputs = []int{phases[3], output3}
		executeProgram(program, inputs, &output4, false)
		// Amplifier 5
		inputs = []int{phases[4], output4}
		executeProgram(program, inputs, &output, false)

		if output > max {
			max = output
		}

		fmt.Println("Output", output)
	}

	fmt.Println("------")
	fmt.Println("Max", max)
}