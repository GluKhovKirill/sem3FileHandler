package main

import (
	"LabFiles/storages/slice"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func writeFile(stringData string, filename string) bool {
	byteData := []byte(stringData)

	err := os.WriteFile(filename, byteData, 0600)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func ReadFileByChars(filename string) *slice.Slice {
	fileBuffer, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputData := string(fileBuffer)
	data := bufio.NewScanner(strings.NewReader(inputData))
	data.Split(bufio.ScanRunes)

	chars := slice.NewSlice()
	for data.Scan() {
		char := data.Text()
		chars.Add(char)
	}
	return chars
}

func DeleteFile(filename string) bool {
	err := os.Remove(filename)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func main() {
	poem := ReadFileByChars("poem.txt") // чтение
	for i := int64(0); i < poem.Len(); i++ {
		fmt.Println(poem.Get(i))
	}
	fmt.Println("Write status:", writeFile("a b\nc d", "test.txt")) // запись
	DeleteFile("test.txt")

}
