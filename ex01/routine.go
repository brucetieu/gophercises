package main

import (
	"fmt"
	"time"
)

func main() {
	process("hello")
}

func process(item string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(item)
	}
}
