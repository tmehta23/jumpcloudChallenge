// handlers.go - contains handlers for /hash and /stats
// endpoints, along with associated methods/ variables.
// See ../README.md for instructions on how to create exe.

package main

import "fmt"
import "time"
import "log"
import "net/http"
import b64 "encoding/base64"
import "crypto/sha512"
import "encoding/json"

// Initialize global variables for statistics tracking
var hashRequestTotalCount int = 0 // total number of requests made to server
var avgTimeRequest float64 = 0.0 // average time for requests to process (in microseconds)
var totalTime float64 = 0.0 // running total of time spent on server requests

// Statistics json output
type StatsJson struct {
  TotalCount map[string]int
  AverageTimeCount map[string]float64
}

// This function takes in a user inputted password and returns hashed/encoded password
// The hashing algorithm is SHA512 and encoding is Base64
func hashAndEncode(inputPassword string) string {

  hash :=sha512.New()
  hash.Write([]byte(inputPassword))
  hashedString := hash.Sum(nil)

  encodedString := b64.StdEncoding.EncodeToString([]byte(hashedString))

  return encodedString
}

// This is the handler for the /hash endpoint.
func hashHandler(writer http.ResponseWriter, request *http.Request) {

  // Takes in user's inputted password from the http request to hash & encode
  userInputPW := request.FormValue("password")

  // Start timer for measuring the time taken for this request.
  start := time.Now()

  // Add 5 second delay here so that the server does not respond immediately,
  // but rather leaves the socket open for 5 seconds before processing request.
  time.Sleep(5 * time.Second)

  // Checks if user has even provided password- if they didn't, print "None provided"
  // Increment hash request count and measure average time for all requests in order
  // to keep statistics accurate.
  if userInputPW == "" {
    fmt.Fprintln(writer,"No password provided. Try again!")
    hashRequestTotalCount++
    timeForThisRequest := time.Since(start).Seconds()
    avgTimeRequest = calculateAverageRequestTime(timeForThisRequest, hashRequestTotalCount)
    return
  }

  // Hash and encode user's provided password using the hashAndEncode method.
  // Increment hash request count and measure average time for all requests in order
  // to keep statistics accurate.
  hashedAndEncodedPW := hashAndEncode(userInputPW)
  hashRequestTotalCount++
  timeForThisRequest := time.Since(start).Seconds()
  avgTimeRequest = calculateAverageRequestTime(timeForThisRequest, hashRequestTotalCount)
  fmt.Fprintln(writer,"\n",hashedAndEncodedPW)
}

func calculateAverageRequestTime(currentRequestTime float64, totalNumberOfRequests int) float64 {
  // Takes in most recent time for request in seconds and converts to microseconds.
  currentRequestTimeInMicrosec := currentRequestTime * 1000000

  // Then, adds this currentRequestTimeInMicrosec to running total for time taken for each request
  // (totalTime).
  totalTime = totalTime + currentRequestTimeInMicrosec

  // Takes this running total and divides by number of requests made thus far in order
  // to calculate the average time taken per request.
  newAverage := totalTime / float64(totalNumberOfRequests)
  return newAverage
}

// This is the handler for the /stats enpoint.
func statsHandler(writer http.ResponseWriter, request *http.Request) {
  // Calls the getStatsJson method, which returns the total number of total
  // number of hash requests made to the server and average time per request
  // as a json text result (in the form of []byte).
  response, err := getStatsJson();
  // If there is any error in the getStatsJson method, log this.
    if err != nil {
        log.Fatal(err)
        return
    }
    // prints response as a string to console.
    fmt.Fprintf(writer, string(response))
}

// Populates json struct StatsJson with total and average, calculated within
// the hash handler.
func getStatsJson()([]byte, error) {
    // populates total with hashRequestTotalCount from hash handler
    totalRequests := make(map[string]int)
    totalRequests["total"] = hashRequestTotalCount

    // populates total with avgTimeRequest from hash handler
    averageCount := make(map[string]float64)
    averageCount["average"] = avgTimeRequest

    // creates json for statistics
    statsOutput := StatsJson{totalRequests, averageCount}

    // returns formatted json text result in form of []byte and error.
    return json.MarshalIndent(statsOutput, "", "  ")
}
