package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)
	
	future := now.Add(24 * time.Hour)
	fmt.Println("Tomorrow:", future)
	
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted:", formatted)
	
	duration := time.Since(now)
	fmt.Println("Duration since now:", duration)
}