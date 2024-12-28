package utils

import "os"

func SetZincEnv() {
	zincURL = os.Getenv("ZINC_URL")
	username = os.Getenv("ZINC_USERNAME")
	password = os.Getenv("ZINC_PASSWORD")
}
