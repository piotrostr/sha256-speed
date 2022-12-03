package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/crypto/sha3"
)

func LoadTheTestImage() []byte {
	file, err := os.Open("../test.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return contents
}

func main() {
	input := LoadTheTestImage()

	start := time.Now()
	hash := sha3.New256()
	n, err := hash.Write(input)
	if err != nil {
		panic(err)
	}
	// this could be structured logged
	fmt.Printf("Sum: 0x%x\n", hash.Sum(nil))
	fmt.Printf("Took: %s\n", time.Since(start))
	fmt.Printf("Input Size: %dB\n", n)
}
