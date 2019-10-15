package main

import (  
    "fmt"
    "time"
	"os"
	"os/signal"
	"syscall"
)

func hello(done chan bool) {  
    fmt.Println("hello go routine is going to sleep")
    time.Sleep(4 * time.Second)
    fmt.Println("hello go routine awake and going to write to done")
    done <- true
}
func main() {  
    done := make(chan bool)
    fmt.Println("Main going to call hello go goroutine")
    go hello(done)
    <-done
    fmt.Println("Main received data")
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		for {
		  fmt.Println("for loop")
		  time.Sleep(4 * time.Second)
		 
		}
	}()
	
	fmt.Println("awaiting signal")
	sig := <-c
	fmt.Println(sig)
	fmt.Println("exiting")
}