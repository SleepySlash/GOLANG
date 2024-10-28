package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Email string1

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ToLower(string(e)))
}

func (e *Email) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*e = Email(s)
	return nil
}

type Person struct {
	Name  string `json:name`
	Email Email  `json:email`
}

func main() {

	p := Person{
		Name:  "Jon",
		Email: "JON@eMail.com",
	}

	jsonData, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println(string(jsonData))

	var m Person
	errr := json.Unmarshal(jsonData, &m)
	if err != nil {
		fmt.Println("Error : ", errr)
		return
	}
	fmt.Printf("%+v\n", m)
}
