package main

import (
	"fmt"
	"reflect"
	//corev1 "k8s.io/api/core/v1"
)

type PodInfoToAAI struct{
	VserverName 			string
	VserverName2 			string
	ProvStatus 			string
	I3InterfaceIPv4Address		string
	I3InterfaceIPvPrefixLength 	string
}

type InstanceRequest struct {
        RBName      string            `json:"rb-name"`
        RBVersion   string            `json:"rb-version"`
        ProfileName string            `json:"profile-name"`
        CloudRegion string            `json:"cloud-region"`
        Labels      map[string]string `json:"labels"`
}

// PodStatus defines the observed state of ResourceBundleState
type PodStatus struct {
        Name        string           `json:"name"`
        Namespace   string           `json:"namespace"`
        Ready       bool             `json:"ready"`
        //Status      corev1.PodStatus `json:"status,omitempty"`
        IPAddresses []string         `json:"ipaddresses"`
}



type InstanceStatus struct {
        Request         InstanceRequest  `json:"request"`
        Ready           bool             `json:"ready"`
        ResourceCount   int32            `json:"resourceCount"`
        PodStatuses     []PodStatus      `json:"podStatuses"`
        //ServiceStatuses []corev1.Service `json:"serviceStatuses"`
}


func main() {

instanceRequest := InstanceRequest {"rb_name","rb_version","profile_name","c_region",map[string]string{"hello":"world"}}
//podStatus := PodStatus {"a","b",true,[]string{"aa","bb"}} 
instanceStatus := InstanceStatus {instanceRequest ,true,12,[]PodStatus{PodStatus{"pod_name","b",true,[]string{"aa","bb"}} ,PodStatus{"pod_name","namespace",true,[]string{"aa","bb"}} }}

//fmt.Println("InstanceRequest =", instanceRequest )
//fmt.Println("podStatus =", podStatus )
//fmt.Println("instanceStatus =", instanceStatus )

//sa := reflect.ValueOf(&instanceRequest).Elem()
//typeOf := sa.Type()
//for i := 0; i < sa.NumField(); i++ {
//    f := sa.Field(i)
//    fmt.Printf("%d: %s %s = %v\n", i,
//        typeOf.Field(i).Name, f.Type(), f.Interface())
//}

//sa = reflect.ValueOf(&podStatus ).Elem()
//typeOf = sa.Type()
//for i := 0; i < sa.NumField(); i++ {
//    f := sa.Field(i)
//    fmt.Printf("%d: %s %s = %v\n", i,
//        typeOf.Field(i).Name, f.Type(), f.Interface())
//}

pita := ProcessInstanceRequest(instanceStatus)
fmt.Println(pita)

}

func ProcessInstanceRequest(instanceStatus InstanceStatus) PodInfoToAAI {

var pita PodInfoToAAI
sa := reflect.ValueOf(&instanceStatus).Elem()
typeOf := sa.Type()
for i := 0; i < sa.NumField(); i++ {
    f := sa.Field(i)
    fmt.Printf("%d: %s %s = %v\n", i,
        typeOf.Field(i).Name, f.Type(), f.Interface())
    if typeOf.Field(i).Name == "Request"{
        fmt.Printf("it's a Request object \n")
        request := f.Interface()
        if ireq, ok := request.(InstanceRequest); ok {
            fmt.Printf("it's a InstanceRequest \n")
            fmt.Printf(ireq.RBName + "\n")
            fmt.Printf(ireq.RBVersion + "\n")
	    fmt.Printf(ireq.ProfileName + "\n")
	    pita.VserverName2 = ireq.ProfileName
	    fmt.Printf(ireq.CloudRegion + "\n")
        } else {
            fmt.Printf("it's not a InstanceRequest \n")
        }
    }

    if typeOf.Field(i).Name == "Ready"{
        fmt.Printf("it's a bool parameter \n")
        ready := f.Interface()
        if red, ok := ready.(bool); ok {
            fmt.Printf("it's a bool \n")
            fmt.Printf("%v\n",red)
            fmt.Printf("%v\n",ready.(bool))
            
        } else {
            fmt.Printf("it's not a InstanceRequest \n")
        }
    }

    if typeOf.Field(i).Name == "ResourceCount"{
        fmt.Printf("it's a int32 parameter \n")
        ready := f.Interface()
        if rc, ok := ready.(int32); ok {
            fmt.Printf("it's a int32 \n")
            fmt.Printf("%d\n",rc)
            
        } else {
            fmt.Printf("it's not a InstanceRequest \n")
        }
    }

    if typeOf.Field(i).Name == "PodStatuses"{
        fmt.Printf("it's a []PodStatus parameter \n")
        ready := f.Interface()
        if pss, ok := ready.([]PodStatus); ok {
            fmt.Printf("it's a []PodStatus \n")
            fmt.Printf("%v\n",pss)
            for i, ps := range pss {
        	fmt.Printf("%v\n",i)
        	fmt.Printf("%v\n",ps)
	        fmt.Printf("%v\n",ps.Name)
	        pita.VserverName = ps.Name
	        fmt.Printf("%v\n",ps.Namespace)
	        pita.ProvStatus = ps.Namespace
	        fmt.Printf("%v\n",ps.Ready)
	        fmt.Printf("%v\n",ps.IPAddresses)

   	    }
            
        } else {
            fmt.Printf("it's not a InstanceRequest \n")
        }
    }
}

return pita 

}