package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    APIURL := "http://localhost:8080/actuator/health"
    req, err := http.NewRequest(http.MethodGet, APIURL, nil)
    if err != nil {
        panic(err)
    }
	
	client := http.DefaultClient
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
	
	defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	
	fmt.Printf("%v", string(body))
}	
	