package main

import (
	"log"
	"os"

	"github.com/shadowsocks/go-shadowsocks2/core"
	"github.com/shadowsocks/go-shadowsocks2/shadowsocks"
)

func main() {
	// Get port from Render's environment variable or fallback to 8388
	port := os.Getenv("PORT")
	if port == "" {
		port = "8388"
	}

	password := "your_secure_password" // CHANGE THIS
	method := "AEAD_CHACHA20_POLY1305"

	// Construct the cipher
	cipher, err := core.PickCipher(method, nil, password)
	if err != nil {
		log.Fatalf("Failed to create cipher: %v", err)
	}

	addr := ":" + port
	log.Printf("Starting Shadowsocks server on %s with method %s\n", addr, method)

	err = shadowsocks.ListenAndServe(addr, cipher)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
