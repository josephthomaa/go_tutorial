package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type CovidData struct {
    date      int    `json:"date"`
    State     string `json:"state"`
    Positive  int    `json:"positive"`
    Negative  int    `json:"negative"`
    Recovered int    `json:"recovered"`
    Deaths    int    `json:"death"`
}

func main() {
    url := "https://api.covidtracking.com/v1/states/daily.json"
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    var data []CovidData
    err = json.NewDecoder(resp.Body).Decode(&data)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("First element of response:")
    fmt.Printf("%+v\n", data[0])
}

