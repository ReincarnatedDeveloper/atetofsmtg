package main

import (
    "log"
    "github.com/shadowsocks/shadowsocks-go/shadowsocks"
)

func main() {
    cfg := &shadowsocks.Config{
        Server:   ":8388",
        Password: "your_secure_password", // CHANGE THIS!
        Method:   "aes-256-gcm",
    }

    server, err := shadowsocks.NewServer(cfg)
    if err != nil {
        log.Fatal("Failed to create Shadowsocks server:", err)
    }

    log.Println("Starting Shadowsocks server on port 8388")
    err = server.ListenAndServe()
    if err != nil {
        log.Fatal("Server error:", err)
    }
}
