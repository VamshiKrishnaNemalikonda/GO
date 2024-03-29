package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
	"bytes"
)

func main() {
    url := "http://localhost:8080/people"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"firstName": "Frodo", "lastName": "Baggins"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}