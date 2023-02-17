package main

import (
	"log"
	"os"
)

func main() {
	file := openFile("my-file.txt", os.O_WRONLY | os.O_APPEND, os.FileMode(0600))
	defer closeFile(file)
	writeToFile(file, "\nAmazing!\n")
}

func openFile(fileName string, perm int, mode os.FileMode) *os.File {
	file, err := os.OpenFile(fileName, perm, mode)
	check(err)
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	check(err)
}

func writeToFile(file *os.File, text string) {
	_, err := file.Write([]byte(text))
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

