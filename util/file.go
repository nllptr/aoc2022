package util

import (
	"log"
	"os"
)

func ReadFile(filename string) []byte {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("could not read file")
	}
	return bytes
}
