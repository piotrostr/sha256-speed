package main

import (
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/sha3"
)

func LoadTheTestImage() []byte {
	file, err := os.Open("../test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return contents
}

func main() {
	input := LoadTheTestImage()

	start := time.Now()
	hash := sha3.New256()
	n, err := hash.Write(input)
	if err != nil {
		log.Fatal(err)
	}
	// this could be structured logged
	log.Printf("Sum: 0x%x", hash.Sum(nil))
	log.Printf("Took: %s", time.Since(start))
	log.Printf("Input Size: %dB", n)
}
