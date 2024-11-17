package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type address struct {
	Street     string `json:"street"`
	Country    string `json:"country"`
	PostalCode int    `json:"postal code"`
}

type person struct {
	ID          int       `json:"id"`
	AdultBool   bool      `json:"adult"`
	Age         *int      `json:"age,omitempty"`
	Name        string    `json:"name"`
	NickName    *string   `json:"nick name"`
	Address     address   `json:"address"`
	FriendsList []*person `json:"friends"`
}

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

	person1 := person{ID: 1337, Name: "jannes", AdultBool: true, Age: func() *int { val := 29; return &val }(), Address: address{Street: "langestraße 22", Country: "nirgendland", PostalCode: 3394}, FriendsList: []*person{}}
	person2 := person{ID: 1338, Name: "horaloranot", AdultBool: false, NickName: func() *string { nick := "hori"; return &nick }(), Address: address{Street: "kurzestraße 11", Country: "nirgendland", PostalCode: 3394}, FriendsList: []*person{&person1}}
	// person1.FriendsList = append(person1.FriendsList, &person2)

	jsonPerson2, err := json.Marshal(person2)

	if err != nil {
		fmt.Printf("Something went wrong:%s\n", err)
		return
	}
	fmt.Printf("%s\n", jsonPerson2)
}
