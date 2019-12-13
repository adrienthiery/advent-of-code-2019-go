package main

import "fmt"
import "strings"
import "reflect"
import "bufio"
import "os"

var test1 = []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
var numberOfOrbits1 = 42

var input = []string{}

func calculateNumberOfOrbits(orbitsMap []string) map[string]int {
	orbitsPoints := make(map[string]int)

	for index := 0; index < len(orbitsMap); index++ {
		orbitPart := orbitsMap[index]

		parts := strings.Split(orbitPart, ")")

		if len(parts) == 2 {
			orbitCenter := parts[0]
			orbitingEntity := parts[1]

			i, ok := orbitsPoints[orbitCenter]
			_ = i
			if !ok {
				orbitsPoints[orbitCenter] = 0
			}
			
			j, ok2 := orbitsPoints[orbitingEntity]
			_ = j
			if !ok2 {
				orbitsPoints[orbitingEntity] = orbitsPoints[orbitCenter] + 1
			} else {
				orbitsPoints[orbitingEntity] = 1 + orbitsPoints[orbitCenter]
			}
		} else {
			fmt.Println(parts)
		}
	}

	total := 0
	for key, value := range orbitsPoints {
		_ = key
		total += value
	}

	fmt.Println("total", total)
	orbitsPoints["total"] = total

	return orbitsPoints
}

func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		 err error = nil
		 line, ln []byte
		)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln),err
  }

func main() {
	// "TESTS"
	fmt.Println("Test1 passing:")
	result1 := calculateNumberOfOrbits(test1)
	fmt.Println(reflect.DeepEqual(result1["total"], numberOfOrbits1))
	fmt.Println(reflect.DeepEqual(result1["COM"], 0))
	fmt.Println(reflect.DeepEqual(result1["D"], 3))
	fmt.Println(reflect.DeepEqual(result1["L"], 7))

	file, err := os.Open("./6.txt")
    defer file.Close()

    if err != nil {
        fmt.Println(err)
    }

    // Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
    for {
		line, err =  Readln(reader)
		input = append(input, line)

        if err != nil {
            break
        }
	}
	
	result2 := calculateNumberOfOrbits(input)
	fmt.Println(result2["total"])
	
	fmt.Println("------")
}