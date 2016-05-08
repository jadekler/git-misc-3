package main

import (
    "fmt"
    "git-misc-3/go-elasticsearch/elasticsearch"
)

func main() {
    fmt.Println("hello world")
    fmt.Println(elasticsearch.ShardCount())

    elasticsearch.PutEmployee("jack", "frost")

    fmt.Println(elasticsearch.EmployeeCount())
}