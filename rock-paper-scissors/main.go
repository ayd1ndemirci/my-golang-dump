package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func handleUserChoice(userChoice string) {
	switch strings.ToLower(userChoice) {
	case "taş", "kağıt", "makas":
		playGame(strings.Title(userChoice))
	default:
		fmt.Printf("Geçersiz seçim: '%s'. Lütfen taş, kağıt veya makas giriniz.\n", userChoice)
	}
}

func playGame(userChoice string) {
	choices := []string{"Taş", "Kağıt", "Makas"}

	computerChoice := choices[rand.Intn(len(choices))]

	fmt.Printf("Senin seçimin: %s | Bilgisayarın seçimi: %s\n", userChoice, computerChoice)

	switch {
	case userChoice == computerChoice:
		fmt.Println("Sonuç: Berabere!")
	case (userChoice == "Taş" && computerChoice == "Makas") ||
		(userChoice == "Kağıt" && computerChoice == "Taş") ||
		(userChoice == "Makas" && computerChoice == "Kağıt"):
		fmt.Println("Sonuç: Kazandın!")
	default:
		fmt.Println("Sonuç: Kaybettin!")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Taş Kağıt Makas Oyunu ===")
	fmt.Println("Lütfen bir seçim yapınız: taş, kağıt veya makas.")

	var userChoice string
	fmt.Print("Seçiminiz: ")
	fmt.Scanln(&userChoice)

	handleUserChoice(userChoice)
}
