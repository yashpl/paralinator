package main

import (
	"io/ioutil"
	"net/http"
	"sync"
)

func googleMaps(data jsonPayload) {

	var wg sync.WaitGroup
	googleMapsURL := "https://maps.googleapis.com/maps/api/geocode/json?sensor=false&key=AIzaSyBe6Nnmy1pRN2Q9WAfZxXibIrE3Ga9CErk&latlng="
	for _, payloadID := range data.Payload {

		//Add wait group instance for every go routine called
		wg.Add(1)

		go func(payloadID string, URL string) {
			var responseMap = map[string]string{}
			defer wg.Done()
			resp, err := http.Get(googleMapsURL + payloadID)
			if err != nil {
				// make a map of faulty url + error body
				responseMap[googleMapsURL+payloadID] = err.Error()
				// Send to error channel
				errorResponse <- responseMap
				println("Error - google api returned error , payload - " + payloadID)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			// make a map of <requested url > : < response from google >
			responseMap[googleMapsURL+payloadID] = string(body)
			// send to response channel
			responseBody <- responseMap

		}(payloadID, data.URL)
	}
	wg.Wait()
}
