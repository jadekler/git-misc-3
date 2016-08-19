package main

import "fmt"

func main() {
    fmt.Println("Starting")

    var queue = make(chan int, 1)
    queue <- 0

    for {
        select {
        case item := <-queue:
            fmt.Println(item)

            if item < 9 {
                queue <- item + 1
            }
        default:
            fmt.Println("Ending")
            return
        }
    }

    fmt.Println("Well... this should never have happened")
}