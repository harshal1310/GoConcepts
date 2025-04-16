package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email_"`
}

// Handler for GET request
func getHandler(w http.ResponseWriter, r *http.Request) {
	//r.URL.Path
	//r.URL.Query()
	fmt.Fprintln(w, "Welcome to the Go REST API!")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User received",
		"name":    user.Name,
		"email":   user.Email,
	})
}

func main() {
	http.HandleFunc("/get", getHandler)   // GET endpoint
	http.HandleFunc("/post", postHandler) // POST endpoint

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil) // Start server

}

/*
// Handler for GET with path parameter
func getParamHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the path parameter
    path := strings.TrimPrefix(r.URL.Path, "/get/param/")
    if path == "" {
        http.Error(w, "Path parameter is missing", http.StatusBadRequest)
        return
    }

    // Respond with the extracted path parameter
    fmt.Fprintf(w, "Path parameter received: %s\n", path)
}

// Handler for GET with query parameters
func getQueryHandler(w http.ResponseWriter, r *http.Request) {
    // Extract query parameters
    query := r.URL.Query()
    name := query.Get("name")
    email := query.Get("email")

    // Respond with the extracted query parameters
    if name == "" || email == "" {
        http.Error(w, "Query parameters 'name' and 'email' are required", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Query parameters received - Name: %s, Email: %s\n", name, email)
}*/
