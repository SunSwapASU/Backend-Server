package main

import (
	"log"

	"github.com/branogarbo/sunswap_backend/prisma"
	"github.com/branogarbo/sunswap_backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	routes.Run()

	err = prisma.Client.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
}
