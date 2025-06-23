package day4

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RWFile() {
	//read file
	file, err := os.OpenFile("/home/kagerou/Development/FullExerciseOfGO/sample.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	//write file
	data2 := []byte("\nI can write this file\nand I can delete this")
	_, er := file.WriteString(string(data2))
	if er != nil {
		fmt.Println("Error writing file ", er)
		return
	}

	fmt.Println("Writing file sample.txt is successfull")
}

func TestRWFile(t *testing.T) {
	RWFile()
	assert.True(t, true)
}
