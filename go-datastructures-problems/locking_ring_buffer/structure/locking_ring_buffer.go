package structure // Different package to scope functions/types appropriately

import "fmt"

type ringBuffer struct {
    buffer         []string
    currentPointer int
}

func NewRingBuffer(size int) ringBuffer {
    return ringBuffer{
        buffer: make([]string, size),
        currentPointer: 0,
    }
}

func (buffer ringBuffer) Print() {
    output := ""

    for key, value := range buffer.buffer {
        output += fmt.Sprintf("[%d, '%s'] ", key, value)
    }

    fmt.Println(output)
}

func (buffer *ringBuffer) Insert(word string) {

}

func (buffer *ringBuffer) Read() string {
    return "??"
}