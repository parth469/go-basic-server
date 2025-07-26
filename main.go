package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	PORT = ":3030"
	ID   = 1
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Users []User

type ApiResponse struct {
	Timestamp string      `json:"timestamp"`
	Method    string      `json:"method"`
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func SaveData(users Users) error {
	byteValue, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}
	return os.WriteFile("db.json", byteValue, 0644)
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	var body User
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersPtr, err := LoadData()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	// Dereference the pointer to work with the slice
	users := *usersPtr

	// Generate new ID (assuming you want auto-increment)
	if len(users) > 0 {
		body.ID = users[len(users)-1].ID + 1
	} else {
		body.ID = 1
	}

	users = append(users, body)

	// Save the updated data
	err = SaveData(users)
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := ApiResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Method:    req.Method,
		Status:    http.StatusCreated,
		Message:   "User created successfully",
		Data:      body,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s \n", req.Method, req.URL.Path)
	fmt.Fprintln(w, "Delete User")

}
func ReadUser(w http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s\n", req.Method, req.URL.Path)

	users, err := LoadData()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	reponce := (ApiResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Method:    req.Method,
		Status:    http.StatusOK,
		Message:   "Operation succeeded",
		Data:      users,
	})

	if err := json.NewEncoder(w).Encode(reponce); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func LoadData() (*Users, error) {
	byteValue, err := os.ReadFile("db.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON file: %v", err)
	}

	var users Users
	if err := json.Unmarshal(byteValue, &users); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &users, nil
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	log.Printf("[%s] %s \n", req.Method, req.URL.Path)
	fmt.Fprintln(w, "Update User")

}

func UserHander(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	switch method {
	case "GET":
		ReadUser(w, req)
	case "DELETE":
		DeleteUser(w, req)
	case "PUT":
		UpdateUser(w, req)
	case "POST":
		CreateUser(w, req)
	default:
		w.Header().Set("Allow", "GET, POST, PUT, DELETE")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	// This logs on EVERY request
	log.Printf("[%s] %s %s\n", time.Now().Format(time.RFC3339), req.Method, req.URL.Path)
	fmt.Fprintf(w, "hello\n")
}
func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user", UserHander)

	log.Printf("Server starting on port %s", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
