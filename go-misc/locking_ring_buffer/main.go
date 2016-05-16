package main

import (
    "git-misc-3/go-misc/locking_ring_buffer/structure"
    "time"
    "fmt"
)

func main() {
    buffer := structure.NewRingBuffer(3)

    buffer.Print()

    go func() {
        fmt.Println("Writing")
        buffer.Insert("hello")
        buffer.Print()

        fmt.Println("Writing")
        buffer.Insert("world")
        buffer.Print()

        fmt.Println("Writing")
        buffer.Insert("this")
        buffer.Print()

        fmt.Println("Writing")
        buffer.Insert("is")
        buffer.Print()

        fmt.Println("Writing")
        buffer.Insert("jean")
        buffer.Print()
    }()

    oneSecond, _ := time.ParseDuration("1s")

    time.Sleep(oneSecond)
    fmt.Println("Reading")
    fmt.Println(buffer.Read())
    buffer.Print()

    time.Sleep(oneSecond)
    fmt.Println("Reading")
    fmt.Println(buffer.Read())
    buffer.Print()
}
