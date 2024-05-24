package main

import (
	"fmt"
	"net"
)

func main() {
	// Écoute sur le port 8080
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Erreur lors de l'écoute:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Serveur UDP en écoute sur le port 8080")

	buffer := make([]byte, 1024)

	for {
		// Lis le message du client
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du message:", err)
			continue
		}
		fmt.Printf("Message reçu du client %s: %s\n", clientAddr, string(buffer[:n]))

		// Envoie le message "Hello" au client
		message := []byte("Hello\n")
		_, err = conn.WriteToUDP(message, clientAddr)
		if err != nil {
			fmt.Println("Erreur lors de l'envoi du message:", err)
			continue
		}
		fmt.Printf("Message envoyé au client %s: %s\n", clientAddr, message)
	}
}
