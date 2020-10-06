package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.Printf("Hello world!")
}
