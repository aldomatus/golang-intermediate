package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Variables
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// Error Handling
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// Maps
	myMap := map[string]int{"seven": 7}
	fmt.Printf("myMap['seven'] = %d\n", myMap["seven"])

	m := make(map[string]int)
	m["six"] = 6
	fmt.Printf("m['six'] = %d\n", m["six"])

	// Slices
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Printf("index: %d - value: %d\n", index, value)
	}

	s = append(s, 4)
	for index, value := range s {
		fmt.Printf("index: %d - value: %d\n", index, value)
	}
}
