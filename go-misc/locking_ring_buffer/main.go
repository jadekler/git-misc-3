package main

import "git-misc-3/go-misc/locking_ring_buffer/structure"

func main() {
    buffer := structure.NewRingBuffer(3)

    buffer.Print()

    buffer.Insert("hello")
    buffer.Insert("world")
    buffer.Insert("this")
    buffer.Insert("is")
    buffer.Insert("jean")

    buffer.Print()
}
