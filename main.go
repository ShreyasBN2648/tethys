package main

import (
	"log"

	"github.com/ShreyasBN2648/tethys/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if present (dev convenience; production is fine without it)
	_ = godotenv.Load()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
