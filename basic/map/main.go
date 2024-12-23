package main

import (
	"fmt"
)

// map[keyType]valueType
var m map[string]int

func createScore() map[string]int {
	newMap := make(map[string]int)
	return newMap
}

func main() {
	// Declaring and Initializing Maps
	// Using make
	scores := make(map[string]int)
	scores["Alice"] = 23

	fmt.Println(scores["Alice"]) // 23

	// Using composite literals - pre populating the map
	ages := map[string]int{
		"Alice": 25,
	}

	fmt.Println(ages["Alice"]) // 25

	// Specifying the capacity of the map
	// scoresSizeMax := make(map[string]int, 100)

	// Accessing and Modifying Map Elements
	// Retrieve a value and check its existence

	table := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	value1 := table["key1"]
	value3 := table["key3"]

	fmt.Println(value1) // value1
	// 값이 없을 경우 해당 타입의 zero value를 반환
	fmt.Println(value3) // ""

	// comma ok idiom
	val, ok := table["key4"]

	if ok {
		fmt.Println("val", val)
	} else {
		fmt.Println("key4 does not exist")
	}

	// Update map
	ages["Alice"] = 26

	// Delete a key. delete(map, key)
	// delete(ages, "Alice")
	//
	// Iterating Over a Map
	// iterate over key-value pairs
	// iteration order는 랜덤
	for name, score := range scores {
		fmt.Println(name, score)
	}

	// Maps as Function Parameters and Return Values
	// map은 reference type이다.
	score2 := createScore()
	score2["James"] = 30

	// Map of Structs

}
