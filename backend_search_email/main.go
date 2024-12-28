package main

import (
	"bufio"
	"chellenge/email"
	"chellenge/shared.go"
	"chellenge/utils"
	"log"
	"net/http"
	"os"
	"strings"
)

func LoadEnvFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo .env: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("Formato inválido en .env: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error leyendo el archivo .env: %v", err)
	}
}



func main() {
    LoadEnvFile(".env") // Cargar el archivo .env
	utils.SetZincEnv()

	port := ":3333"
    router := email.InitRoutes()

    corsRouter := shared.Cors(router)

    log.Printf("Server is running on http://localhost%s\n", port)
    if err := http.ListenAndServe(port, corsRouter); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

