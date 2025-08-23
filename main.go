package main

import (
	"fmt"
)

type Timer struct {
	id    string
	value int
}

func (t *Timer) tick() {
	t.value++
	fmt.Println("value: ", t.value)
}

func main() {
	// //read file
	// file, err := os.Open("/home/kagerou/Development/FullExerciseOfGO/sample.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("read %d bytes: %q\n", count, data[:count])

	// //write file
	// data2 := []byte("this is new data")
	// if er := os.WriteFile("sample.txt", data2, 0644); er != nil {
	// 	fmt.Println("Error writing file ", er)
	// 	return
	// }

	// fmt.Println("Writing file sample.txt is successfull")

	// var sb strings.Builder
	// sb.WriteString("my name is kagerou13 and my code is kag13")
	// fmt.Println(sb.String())

	var x int

	fmt.Scan(&x)

	t := Timer{"timer1", 0}

	for i := 0; i < x; i++ {
		t.tick()
	}

}
