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

	inputDecoded := make([]int, 0, int(len(input)/2)+1) //give it some reasonable capacity
	err = json.Unmarshal(input, &inputDecoded)
	check(err)
	// fmt.Println(len(inputDecoded))
	// fmt.Println(cap(inputDecoded))
	// fmt.Println(inputDecoded)

	inputDecoded = append(inputDecoded, inputDecoded[0])
	// fmt.Println(len(inputDecoded))
	// fmt.Println(cap(inputDecoded))

	var sum int64
	length := len(inputDecoded)
	for i, num := range inputDecoded {
		if i < length-1 && num == inputDecoded[i+1] {
			// fmt.Printf("%d += %d\n", sum, num)
			sum += int64(num)
		}
	}

	fmt.Printf("Results: %d\n", sum)
}
