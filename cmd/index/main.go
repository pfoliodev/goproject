package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"example.com/finalproject/pkg/protocols"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	fmt.Println("Serveur en attente de connexions...")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Nouvelle connexion acceptée")

	for {
		// Lecture du message du client
		var request interface{}
		err := json.NewDecoder(conn).Decode(&request)
		if err != nil {
			log.Println("Erreur de lecture du message:", err)
			return
		}

		// Traitement du type de message
		switch msg := request.(type) {
		case protocols.GetUrlRequest:
			// Logique métier pour la demande d'URL
			response := protocols.GetUrlResponse{
				URL: "https://www.example.com",
			}
			err = json.NewEncoder(conn).Encode(response)
			if err != nil {
				log.Println("Erreur d'envoi de la réponse:", err)
				return
			}
		case protocols.CreateUrlRequest:
			// Logique métier pour la création d'URL
			response := protocols.CreateUrlResponse{
				Success: true,
			}
			err = json.NewEncoder(conn).Encode(response)
			if err != nil {
				log.Println("Erreur d'envoi de la réponse:", err)
				return
			}
		// ... Ajouter d'autres cas pour les autres types de messages ...
		default:
			log.Println("Type de message non pris en charge")
			return
		}
	}
}
