package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type ChallengeResponse struct {
    Follow  string `json:"follow"`
    Message string `json:"message"`
}

var prevFollow = ""
var requestCounter = 0

func main() {
    url := "https://www.letsrevolutionizetesting.com/challenge"
    processChallenge(url)
}

func processChallenge(url string) {
    requestCounter++
    fmt.Printf("Request #%d to %s\n", requestCounter, url)

    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating HTTP request:", err)
        return
    }

    req.Header.Set("Accept", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making HTTP request:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("HTTP request returned a non-OK status code:", resp.Status)
        return
    }

    var response ChallengeResponse
    decoder := json.NewDecoder(resp.Body)
    if err := decoder.Decode(&response); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Check if Follow/ Message is empty
    if response.Follow != "" {
        fmt.Printf("Follow: %s\n", response.Follow)
        prevFollow = response.Follow
    } else {
        fmt.Println("Change in response format, stopping recursion.")
        fmt.Println("---------------")
        fmt.Println("Message: ", response.Message)
        return
    }

    processChallenge(response.Follow)
}
