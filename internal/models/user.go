// internal/models/user.go
package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var Users = []User{
    {ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
    {ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com"},
}
