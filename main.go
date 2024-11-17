package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myInt struct {
	IntValue int `json:"int_value"`
}

func main() {
	data := map[string]interface{}{
		"id":             00001,
		"we have it":     time.Now(),
		"will this work": true,
		"be nice":        "ok",
		"and we can nest it": map[string]interface{}{
			"yes":  "all eggs in the nest",
			"eggs": []int{0, 1, 2, 3, 4},
		},
		"the rest is silence": nil,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Something went wrong:%s\n", err)
		return
	}
	jsonInt, err := json.Marshal(myInt{42})
	if err != nil {
		fmt.Printf("Something went wrong:%s\n", err)
		return
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Printf("%s\n", jsonInt)
}
