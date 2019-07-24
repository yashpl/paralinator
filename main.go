package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Todo -
// error handling ?

var responseBody = make(chan map[string]string)
var errorResponse = make(chan map[string]string)
var responses = map[string]string{}

type jsonPayload struct {
	URL     string   `json:"url"`
	Payload []string `json:"payload"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Populate data for testing.
	//testData := `{"url" : "maps.google.com", "payload" : ["13.8856449,14.677982" , "12.8359006,12.7306095" , "12.8,13.9"]}`

	// Make an interface for de-serialization
	data := jsonPayload{}

	// Deserialize with Unmarshal method
	err := json.Unmarshal([]byte(request.Body), &data)

	if err == nil {
		fmt.Print("Parsed Successfully ")

	} else {
		fmt.Print(err)
		// TODO return a 5xx error
		return events.APIGatewayProxyResponse{Body: "Not able to decode json", StatusCode: 500}, nil
	}

	// Use WaitGroup to wait for all go routines to finish before exiting main thread.
	if strings.Contains(data.URL, "maps.google.com") {
		// start reciever
		go responseReciever()
		googleMaps(data)
		return events.APIGatewayProxyResponse{Body: returnResponse(), StatusCode: 200}, nil

	}
	return events.APIGatewayProxyResponse{Body: "Undefined URL sent", StatusCode: 422}, nil

}

func main() {
	lambda.Start(handleRequest)

}
