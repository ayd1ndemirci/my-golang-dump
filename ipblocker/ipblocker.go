package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var blockedIPs = []string{
	"192.168.1.1",
	"203.0.113.5",
	"192.168.2.67",
}

func ipBlocker(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIP := strings.Split(r.RemoteAddr, ":")[0]
		for _, blocked := range blockedIPs {
			if userIP == blocked {
				log.Printf("Bloklanmış IP: %s", userIP)
				http.Error(w, "Erişim reddedildi.", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.Header.Get("X-Forwarded-For")
		if clientIP == "" {
			clientIP = r.RemoteAddr
		}
		fmt.Fprintf(w, "Hoş geldiniz! IP adresiniz: %s", clientIP)
	})

	port := ":8080"
	log.Printf("Sunucu %s portunda çalışıyor...", port)
	err := http.ListenAndServe(port, ipBlocker(http.DefaultServeMux))
	if err != nil {
		log.Fatal("Sunucu başlatılamadı:", err)
	}
}
