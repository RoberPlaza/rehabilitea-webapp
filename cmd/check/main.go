package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)

	err := bcrypt.CompareHashAndPassword(pass, []byte("password"))
	if err != nil {
		log.Fatal("UNMATCHED")
	}

	log.Println("Mached")
}
