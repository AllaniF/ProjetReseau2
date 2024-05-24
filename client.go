package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "localhost:8080"
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Envoie un message au serveur
	message := []byte("Bonjour serveur\n")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi du message:", err)
		return
	}
	fmt.Println("Message envoyé au serveur:", string(message))

	// Lis la réponse du serveur
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}
	fmt.Println("Réponse reçue du serveur:", string(buffer[:n]))
}
