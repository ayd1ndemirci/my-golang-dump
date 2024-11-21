package main

import (
	"fmt"
	"time"
)

func main() {
	var status bool = false;
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		if i >= 10 {
			status = true
		}
		if status {
			fmt.Println("Counting is over.")
		}
	}
}