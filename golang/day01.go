package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := ioutil.ReadFile("day01-input.json")
	check(err)
	fmt.Println(len(input))

	inputDecoded := make([]int, 0, int(len(input))) //give it some reasonable capacity
	err = json.Unmarshal(input, &inputDecoded)
	check(err)
	// fmt.Println(len(inputDecoded))
	// fmt.Println(cap(inputDecoded))
	// fmt.Println(inputDecoded)

	length := len(inputDecoded)
	inputDecoded = append(inputDecoded, inputDecoded[0])
	// fmt.Println(len(inputDecoded))
	// fmt.Println(cap(inputDecoded))

	var sum int64
	for i, num := range inputDecoded {
		if i < length && num == inputDecoded[i+1] {
			// fmt.Printf("%d += %d\n", sum, num)
			sum += int64(num)
		}
	}

	fmt.Printf("Results 1: %d\n", sum)

	inputDecoded = append(inputDecoded[:len(inputDecoded)-1], inputDecoded[:len(inputDecoded)/2]...)
	step := length / 2
	// newLength := len(inputDecoded)
	sum = 0 //re-zero this for reuse

	for i, num := range inputDecoded {
		if i < length && num == inputDecoded[i+step] {
			sum += int64(num)
		}
	}

	fmt.Printf("Results %d: %d\n", step, sum)
}
