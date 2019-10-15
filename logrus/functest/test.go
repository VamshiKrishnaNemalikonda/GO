package main

import (
    //"fmt"
    "functest/functions"
	logr "github.com/sirupsen/logrus"
)

func main(){
	logr = functions.GetValue()
	logr.WithFields(logr.Fields{"CloudRegion": "Invalid/Missing CloudRegion in POST request"}).Error("CreateVnfRequest bad request")
    //fmt.Println(functions.GetValue())
}