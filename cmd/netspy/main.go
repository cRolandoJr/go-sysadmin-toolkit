package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	url := os.Getenv("TARGET")
	if url == "" {
		url = "https://www.google.com"
	}

	fmt.Printf("[%s] Comprobando conexion a: %s\n", time.Now().Format(time.RFC3339), url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: No se pudo conectar. %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("Respuesta recibida con codigo de estado: %s\n", resp.Status)
}
