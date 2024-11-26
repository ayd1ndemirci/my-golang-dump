package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/skip2/go-qrcode"
)

type PageVariables struct {
	Title string
}

type QRResponse struct {
	QRCode string `json:"qrCode"`
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/generate_qr", generateQRCode)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title: "QR Code Generator",
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, pageVariables)
}

func generateQRCode(w http.ResponseWriter, r *http.Request) {
	qrText := r.FormValue("text")
	if qrText != "" {
		qrCode, err := qrcode.Encode(qrText, qrcode.Medium, 256)
		if err != nil {
			http.Error(w, "Error generating QR code", http.StatusInternalServerError)
			return
		}

		encodedQRCode := base64.StdEncoding.EncodeToString(qrCode)

		response := QRResponse{
			QRCode: fmt.Sprintf("data:image/png;base64,%s", encodedQRCode),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Text for QR code is required", http.StatusBadRequest)
	}
}
