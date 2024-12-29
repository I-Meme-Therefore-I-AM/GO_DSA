package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	packt := "packt"

	jsonPackt, ok := json.Marshal(packt)

	if ok != nil {
		panic("could not marshal")
	} else {
		fmt.Println(string(jsonPackt))
	}
	Jmain()
	jsonUnmarshall()
	MapInterface()
}

type MyObject struct {
	Number int    `json:"number"`
	Word   string `json:"word"`
}

func Jmain() {
	object := MyObject{5, "Packt"}

	oJson, _ := json.Marshal(object)

	fmt.Printf("%s\n", oJson)
}

func jsonUnmarshall() {
	jsonBytes := []byte(`{"number":5, "word": "Packt"}`)

	var object MyObject

	err := json.Unmarshal(jsonBytes, &object)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Number is %d, word is %s\n", object.Number, object.Word)
}

func MapInterface() {
	jsonBytes := []byte(`{"number":5, "word": "Packt"}`)

	var dangerousObject map[string]interface{}

	err := json.Unmarshal(jsonBytes, &dangerousObject)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number is %f, ", dangerousObject["number"])
	fmt.Printf("Word is %s\n", dangerousObject["word"])
	fmt.Printf("Error reference is %v\n",
		dangerousObject["nothing"])
	// fmt.Println(dangerousObject["number"])
	fmt.Println("my name")
}
