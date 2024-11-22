package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	for {
		
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		currentTime := time.Now()
		fmt.Println("Saat ve Tarih: ", currentTime.Format("02.01.2006 15:04:05"))
		time.Sleep(1 * time.Second)
	}
}
