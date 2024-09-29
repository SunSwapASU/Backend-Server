package main

import (
	"fmt"

	"github.com/branogarbo/sunswap_backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	routes.Run()
}
