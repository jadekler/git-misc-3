package elasticsearch

import (
    "net/http"
    "strings"
    "fmt"
)

func PutEmployee(firstName, surname string) {
    requestBody := fmt.Sprintf(`
    {
        "first_name": "%s",
        "last_name": "%s",
        "age": 25,
        "about": "I love to go rock climbing",
        "interests": ["climbing", "basketball"]
    }
    `, firstName, surname)
    data := strings.NewReader(requestBody)
    http.Post("http://localhost:9200/megacorp/employee/5", "json", data)
}