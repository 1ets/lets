package loader

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var EnvFile string

// Loading .env environment variable into memory.
func Environment() {
	fmt.Println("LOAD ENV")
	if EnvFile != "" {
		err := godotenv.Load(EnvFile)
		if err != nil {
			log.Fatalln("Error loading .env file")
		}
		return
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
