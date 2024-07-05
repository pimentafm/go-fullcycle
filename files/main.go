package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	//size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Wrinting bytes to a file\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes\n", size)

	f.Close()

	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Reading chunks

	f2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f2)
	buffer := make([]byte, 2)

	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:bytesRead]))
	}

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
