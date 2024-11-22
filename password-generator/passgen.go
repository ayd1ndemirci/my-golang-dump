package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var length int;
	fmt.Print("Şifrenizin uzunluğunu girin: ");
	fmt.Scan(&length);

	rand.Seed(time.Now().UnixNano());

	charset := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm01234567890@#$%^&*()_+-=[]{}|;':\",./<>?";

	password := make([]byte, length);
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))];
	}

	fmt.Print("Oluşturulan şifre: ", string(password));
}