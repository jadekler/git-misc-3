package main

import "fmt"

func main() {
    buffer := newRingBuffer(3)

    buffer.print()

    buffer.queueWord("hello")
    buffer.queueWord("world")
    buffer.queueWord("this")
    buffer.queueWord("is")

    buffer.print()
}

type ringBuffer struct {
    buffer         []string
    currentPointer int
    maxSize        int
}

func newRingBuffer(size int) ringBuffer {
    return ringBuffer{
        buffer: make([]string, size),
        currentPointer: 0,
        maxSize: size,
    }
}

func (buffer ringBuffer) print() {
    output := ""

    for key, value := range buffer.buffer {
        output += fmt.Sprintf("[%d, '%s'] ", key, value)
    }

    fmt.Println(output)
}

func (buffer *ringBuffer) queueWord(word string) {
    buffer.buffer[buffer.currentPointer] = word

    if buffer.currentPointer + 1 == buffer.maxSize {
        buffer.currentPointer = 0
    } else {
        buffer.currentPointer++
    }
}