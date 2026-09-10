package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	// Requesting from JSONPlaceholder open API
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	fmt.Printf("ID: %d\nTitle: %s\n", post.ID, post.Title)
}
