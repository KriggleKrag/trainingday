package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file for reading
	file, err := os.Open("todo.list")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line and print each line to the terminal
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Check for any errors during the file reading process
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
