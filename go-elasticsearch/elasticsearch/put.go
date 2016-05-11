package elasticsearch

import (
    "net/http"
    "strings"
    "fmt"
)

func PutEmployee(id int, firstName, surname string) {
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
    url := fmt.Sprintf("http://localhost:9200/megacorp/employee/%d", id)
    http.Post(url, "json", data)
}