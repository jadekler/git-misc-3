package main

import "fmt"

func main() {
    myChan := make(chan interface{}, 3)
    myChan <- 5

    autoscale(func(chan interface{}) {

    }(myChan))

    fmt.Println("donezo")
}

type myFunc func(chan interface{})

func autoscale(f myFunc) {
    fmt.Println("autoscaling")
}