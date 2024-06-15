package main

import (
	"log"
	"os"
)

func test() {

	CreateFile("learn/io_learn/test.txt")
}

func CreateFile(fileName string) {
	_, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("create file err: %v", err)
	}
}

func WriteString(content string) {
	file, err := os.Open("learn/io_learn/test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(content)

}
