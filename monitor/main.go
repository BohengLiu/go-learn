package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// "monitor/notify"
	// "monitor/services"
	// "monitor/tasks"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Printf("%s uses %s\n", os.Getenv("NAME"), os.Getenv("EDITOR"))
	// client := services.InitClient()

	// go notify.SendEmail()

	// go tasks.SubscribeToken(client)

	// select {}
}
