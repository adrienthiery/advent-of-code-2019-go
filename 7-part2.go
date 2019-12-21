package main

import "fmt"
import "log"
import "reflect"
import "strconv"

var program = []int{3,8,1001,8,10,8,105,1,0,0,21,34,43,60,81,94,175,256,337,418,99999,3,9,101,2,9,9,102,4,9,9,4,9,99,3,9,102,2,9,9,4,9,99,3,9,102,4,9,9,1001,9,4,9,102,3,9,9,4,9,99,3,9,102,4,9,9,1001,9,2,9,1002,9,3,9,101,4,9,9,4,9,99,3,9,1001,9,4,9,102,2,9,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1001,9,1,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,99}

func getValue(program *[]int, baseIndex int, parameterMode int, increment int) int {
	if parameterMode == 0 {
		index := (*program)[baseIndex + increment]
		return (*program)[index]
	} else {
		return (*program)[baseIndex + increment]
	}
}

type StepReturn struct{ 
	returnedIndex int
	exitCode int
} 

func executeStep(program *[]int, baseIndex int, inputs *[]int, inputPointer *int, output *int, debug bool) StepReturn {
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
		return StepReturn{ returnedIndex: -1, exitCode: 99 }
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
		(*program)[targetIndex] = (*inputs)[*inputPointer]
		(*inputPointer)++
		returnIndex += 2
	} else if opcode == 4 {
		// write output
		*output = getValue(program, baseIndex, parametersMode[0], 1)
		if debug {
			fmt.Println("Writing to output, value is:", *output)
		}
		returnIndex += 2
		return StepReturn{ returnedIndex: returnIndex, exitCode: 1 }
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
		// equal
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

	return StepReturn{ returnedIndex: returnIndex, exitCode: 0 }
}

type ProgramReturn struct{ 
	returnedIndex int
	returnedProgram []int
} 

func executeProgram(program []int, inputsP *[]int, outputP *int, startAtIndex int, debug bool) ProgramReturn {
	if debug {
		fmt.Println(program)
	}
	returnedProgram := program
	inputPointer := 0
	returnFromStep := StepReturn{ returnedIndex: startAtIndex, exitCode: 0 }

	for returnFromStep.exitCode == 0 {
		returnFromStep = executeStep(&returnedProgram, returnFromStep.returnedIndex, inputsP, &inputPointer, outputP, debug)
	}

	return ProgramReturn{ returnedIndex: returnFromStep.returnedIndex, returnedProgram: returnedProgram }
}

func runAmplifiers(phases PhaseSet, program []int, loop bool, debug bool) int {
	fmt.Println("phases", phases)
	output := 99999999
	loopEnded := false
	
	// Initialization program 1
	inputs1 := []int{phases[0], 0}
	program1 := program
	returnProgram1 := ProgramReturn{ returnedIndex: 0, returnedProgram: program1 }
	// Initialization program 2
	inputs2 := []int{phases[1], 0}
	program2 := program
	returnProgram2 := ProgramReturn{ returnedIndex: 0, returnedProgram: program2 }
	// Initialization program 3
	inputs3 := []int{phases[2], 0}
	program3 := program
	returnProgram3 := ProgramReturn{ returnedIndex: 0, returnedProgram: program3 }
	// Initialization program 4
	inputs4 := []int{phases[3], 0}
	program4 := program
	returnProgram4 := ProgramReturn{ returnedIndex: 0, returnedProgram: program4 }
	// Initialization program 5
	inputs5 := []int{phases[4], 0}
	program5 := program
	returnProgram5 := ProgramReturn{ returnedIndex: 0, returnedProgram: program5 }

	i := 0
	
	for loop && loopEnded == false || !loop && i < 1 {
		i++
		if debug {
			fmt.Println("Loop", i)
		}
		// Amplifier 1
		returnProgram1 = executeProgram(returnProgram1.returnedProgram, &inputs1, &output, returnProgram1.returnedIndex, debug)
		if debug {
			fmt.Println("Output Amplifier 1", output)
		}
		if i > 1 {
			inputs2 = []int{output}
		} else {
			inputs2[1] = output
		}
		// Amplifier 2
		returnProgram2 = executeProgram(returnProgram1.returnedProgram, &inputs2, &output, returnProgram2.returnedIndex, debug)
		if debug {
			fmt.Println("Output Amplifier 2", output)
		}
		if i > 1 {
			inputs3 = []int{output}
		} else {
			inputs3[1] = output
		}
		// Amplifier 3
		returnProgram3 = executeProgram(returnProgram1.returnedProgram, &inputs3, &output, returnProgram3.returnedIndex, debug)
		if debug {
			fmt.Println("Output Amplifier 3", output)
		}
		if i > 1 {
			inputs4 = []int{output}
		} else {
			inputs4[1] = output
		}
		// Amplifier 4
		returnProgram4 = executeProgram(returnProgram1.returnedProgram, &inputs4, &output, returnProgram4.returnedIndex, debug)
		if debug {
			fmt.Println("Output Amplifier 4", output)
		}
		if i > 1 {
			inputs5 = []int{output}
		} else {
			inputs5[1] = output
		}
		// Amplifier 5
		returnProgram5 = executeProgram(returnProgram1.returnedProgram, &inputs5, &output, returnProgram5.returnedIndex, debug)
		if debug {
			fmt.Println("Output Amplifier 5", output)
		}
		inputs1 = []int{output}

		// fmt.Println("Loop", i, "output", output)

		if returnProgram5.returnedIndex == -1 {
			loopEnded = true
		}
	}

	return output
}

func runTestsAllPossiblePhases (program []int, basePhaseIndex int, loop bool) int {
	testSets := []PhaseSet{}
	for i := basePhaseIndex; i < basePhaseIndex + 5; i++ {
		for j := basePhaseIndex; j < basePhaseIndex + 5; j++ {
			for k := basePhaseIndex; k < basePhaseIndex + 5; k++ {
				for l := basePhaseIndex; l < basePhaseIndex + 5; l++ {
					for m := basePhaseIndex; m < basePhaseIndex + 5; m++ {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							testSets = append(testSets, PhaseSet{i, j, k, l, m})
						}
					}
				}
			}
		}
	}

	max := 0

	for index := 0; index < len(testSets); index++ {
		phases := testSets[index]
	
		amplifiersOutput := runAmplifiers(phases, program, loop, false)

		if amplifiersOutput > max {
			max = amplifiersOutput
		}
	}

	return max
}

type PhaseSet = [5]int

func main() {
	debugProgram := false
	debugTests := true
	testInput := 0
	testOutput := 99999

	// "TESTS"
	fmt.Println("Test1:")
	var test1 = []int{1002,4,3,4,33}
	var expected1 = []int{1002,4,3,4,99}
	result1 := executeProgram(test1, &[]int{testInput}, &testOutput, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", result1.returnedProgram, "=", expected1)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(result1.returnedProgram, expected1))
	
	fmt.Println("------")

	fmt.Println("Test2:")
	var test2 = []int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}
	inputZero := 0
	test2Output := 999999
	executeProgram(test2, &[]int{inputZero}, &test2Output, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", test2Output, "=", 0)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(test2Output, 0))

	inputNonZero := 12
	test2Output2 := 999999
	executeProgram(test2, &[]int{inputNonZero}, &test2Output2, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", test2Output2, "=", 1)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(test2Output2, 1))
	
	fmt.Println("------")

	fmt.Println("Test3:")
	var test3 = []int{3,3,1105,-1,9,1101,0,0,12,4,12,99,1}
	test3Output := 999999
	executeProgram(test3, &[]int{inputZero}, &test3Output, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", test3Output, "=", 0)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(test3Output, 0))

	test3Output2 := 999999
	executeProgram(test3, &[]int{inputNonZero}, &test3Output2, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", test3Output2, "=", 1)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(test3Output2, 1))
	
	fmt.Println("------")

	fmt.Println("Test4:")
	var test4 = []int{3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99}
	input4One := 4
	test4Output := 999999
	executeProgram(test4, &[]int{input4One}, &test4Output, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", test4Output, "=", 999)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(test4Output, 999))

	input4Two := 8
	test4Output2 := 999999
	executeProgram(test4, &[]int{input4Two}, &test4Output2, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", test4Output2, "=", 1000)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(test4Output2, 1000))

	input4Three := 10
	test4Output3 := 999999
	executeProgram(test4, &[]int{input4Three}, &test4Output3, 0, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", test4Output3, "=", 1001)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(test4Output3, 1001))
	
	fmt.Println("------")

	fmt.Println("Non-Looping Tests")
	fmt.Println("Test 5:")
	var test5 = []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}
	testOutput5 := runAmplifiers(PhaseSet{4,3,2,1,0}, test5, false, true)
	if debugTests {
		fmt.Println(" test is: ", testOutput5, "=", 43210)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(testOutput5, 43210))

	fmt.Println("------")

	fmt.Println("Test 6:")
	var test6 = []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}
	testOutput6 := runAmplifiers(PhaseSet{0,1,2,3,4}, test6, false, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", testOutput6, "=", 54321)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(testOutput6, 54321))

	fmt.Println("------")

	fmt.Println("Test 7:")
	var test7 = []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}
	testOutput7 := runAmplifiers(PhaseSet{1,0,4,3,2}, test7, false, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", testOutput7, "=", 65210)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(testOutput7, 65210))

	fmt.Println("------")

	fmt.Println("Test Part 1:")
	max1 := runTestsAllPossiblePhases(program, 0, false)
	if debugTests {
		fmt.Println(" test is: ", max1, "=", 11828)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(max1, 11828))

	fmt.Println("------")
	
	fmt.Println("Looping Tests")
	fmt.Println("Test 8:")
	var test8 = []int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5}
	testOutput8 := runAmplifiers(PhaseSet{9, 8, 7, 6, 5}, test8, true, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", testOutput8, "=", 139629729)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(testOutput8, 139629729))

	fmt.Println("------")

	fmt.Println("Test 9:")
	var test9 = []int{3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10}
	testOutput9 := runAmplifiers(PhaseSet{9, 7, 8, 5, 6}, test9, true, debugProgram)
	if debugTests {
		fmt.Println(" test is: ", testOutput9, "=", 18216)
	}
	fmt.Println(" passing: ", reflect.DeepEqual(testOutput9, 18216))

	fmt.Println("------")

	fmt.Println("Real Deal:")

	max := runTestsAllPossiblePhases(program, 5, true)

	fmt.Println("------")
	fmt.Println("Max", max)
}