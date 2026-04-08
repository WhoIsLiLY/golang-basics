package main

import "fmt"

func sender(ch chan<- string) {
	messages := []string{"Hello", "World", "From", "Go"}
	for _, msg := range messages {
		ch <- msg
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	
	go sender(ch)
	
	for message := range ch {
		fmt.Println("Received:", message)
	}
}