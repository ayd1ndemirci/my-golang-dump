package main

import (
	"fmt"
)

func main() {
	var num1, num2 float32;
	var operator string;

	fmt.Println("----- Hesap Makinesi -----");

	fmt.Print("1. sayıyı gir: ");
	_, err := fmt.Scanln(&num1);
	if err != nil {
		fmt.Println("Hata:", err);
		return
	}

	fmt.Print("2. sayıyı gir: ");
	_, err = fmt.Scanln(&num2);
	if err != nil {
		fmt.Println("Hata:", err);
		return
	}

	fmt.Print("Operatörü gir (+, -, *, /): ");
	_, err = fmt.Scanln(&operator);
	if err != nil {
		fmt.Println("Hata:", err);
		return
	}

	result, err := calculate(num1, num2, operator);
	if err != nil {
		fmt.Println("Hata:", err);
	} else {
		fmt.Printf("Sonuç: %.2f\n", result);
	}
}

func calculate(num1, num2 float32, operator string) (float32, error) {
	switch operator {
	case "+":
		return num1 + num2, nil;
	case "-":
		return num1 - num2, nil;
	case "*":
		return num1 * num2, nil;
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Bir sayı sıfıra bölünemez");
		}
		return num1 / num2, nil;
	default:
		return 0, fmt.Errorf("Geçersiz operatör");
	}
}
