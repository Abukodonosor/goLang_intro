package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type response1 struct {
	Name  string   `json:	"name"`
	Items []string `json: "items"`
	Obj   object   `json: "obj"`
}

type object struct {
	Omg    int
	Status string
}

func main() {
	//declare variable
	var resp response1
	var toJson response1

	//type of json which arrive by api
	data := []byte(`
		{
		"name":"John",
		"age":30,
		"city":"New York",
		"items":[
			"mrcko",
			"krcko",
			"smrcko"
			],
		"obj":{
			"omg": 12,
			"status":"gde si brej"
			}
		}
	`)
	fmt.Printf("Hejjj gde si bio json: %s \n", data)

	if err := json.Unmarshal(data, &resp); err != nil {
		panic(err)
		log.Fatal("Cannot parse JSON: ", err)
	}

	fmt.Println("Structure is : ", resp)

	toJson = response1{"Gibotaurus", []string{"kokain", "caj", "sex"}, object{331, "Kalabunga"}}

	data, err := json.Marshal(toJson)
	if err != nil {
		panic(err)
		log.Fatal("Cannot encode to JSON ", err)
	}

	fmt.Printf("Package object to json: %s \n", data)

	var dinamicMap map[string]interface{}

	err = json.Unmarshal(data, &dinamicMap)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Here is dinamic parsed object: \n")
	fmt.Println(": ", dinamicMap["Name"])
}
