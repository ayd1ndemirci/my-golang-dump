package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/skip2/go-qrcode"
)

func main() {
	var qrText string
	fmt.Print("QR Koda dönüştürülecek metin veya URL gir: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()  
	qrText = scanner.Text()

	err := qrcode.WriteFile(qrText, qrcode.Medium, 256, "qrcode.png")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("QR Kod başarıyla oluşturuldu")
}
