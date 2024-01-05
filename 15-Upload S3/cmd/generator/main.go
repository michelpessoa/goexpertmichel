package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 0; i <= 10; i++ {
		f, err := os.Create(fmt.Sprintf("./temp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("Hello, World!")
	}
}
