package main

import (
	"fmt"
)

func main() {
	day := "Cuma"
	switch day {
	case "Pazartesi":
		fmt.Print("Bugün Pazartesi");
	case "Salı":
		fmt.Print("Bugün Salı");
	case "Çarşamba":
		fmt.Print("Çarşamba");
	case "Perşembe":
		fmt.Print("Bugün Perşembe");
	case "Cuma":
		fmt.Print("Bugün Cuma");
	case "Cumartesi":
		fmt.Print("Bugün Cumartesi");
	case "Pazar":
		fmt.Print("Bugün Pazar");
	default:
		fmt.Print("Bugün ne bilmiyorum.");
	}
}	