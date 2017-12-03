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

	inputDecoded := make([]int, 0, int(len(input)/2)) //give it some reasonable capacity
	err = json.Unmarshal(input, &inputDecoded)
	check(err)
	fmt.Println(len(inputDecoded))
	// fmt.Println(inputDecoded)
}
