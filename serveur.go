// server.go
package main

import (
	"fmt"
	"net"
)

func main() {
	// Écoute sur le port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erreur lors de l'écoute:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Serveur en écoute sur le port 8080")

	for {
		// Accepte une connexion
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
			continue
		}
		// Gère la connexion dans une goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Envoie le message "Hello" au client
	message := "Hello\n"
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du message:", err)
		return
	}
	fmt.Println("Message envoyé au client:", message)
}
