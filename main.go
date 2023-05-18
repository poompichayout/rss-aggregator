package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("can not load .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:", portString)
}
