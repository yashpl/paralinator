package main

import (
	"encoding/json"
)

func responseReciever() {
	for data := range responseBody { // Here `data` itself is a map.
		for index, element := range data { // Itrate over the map recieved from channel and store its value in global map `responses`
			responses[index] = element
		}

	}
}

func returnResponse() string {
	jsonString, _ := json.Marshal(responses)
	//fmt.Println(string(jsonString))
	return string(jsonString)
}
