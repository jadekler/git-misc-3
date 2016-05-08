package elasticsearch

import (
    "net/http"
    "io/ioutil"
)

func ShardCount() string {
    resp, _ := http.Get("http://localhost:9200/_count")
    body, _ := ioutil.ReadAll(resp.Body)
    bodyString := string(body)
    return bodyString
}

func EmployeeCount() string {
    resp, _ := http.Get("http://localhost:9200/megacorp/employee/_search")
    body, _ := ioutil.ReadAll(resp.Body)
    bodyString := string(body)
    return bodyString
}