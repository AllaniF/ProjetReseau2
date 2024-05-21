// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	//"os"
)

func main() {
	// Se connecte au serveur sur le port 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connecté au serveur sur le port 8080")

	// Lit la réponse du serveur
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		return
	}
	fmt.Print("Message du serveur: " + message)
}
