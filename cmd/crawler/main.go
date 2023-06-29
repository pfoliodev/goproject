package main

import (
	"encoding/json"
	"log"
	"net"

	"example.com/finalproject/pkg/protocols"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Exemple d'envoi d'une requête GetUrlRequest
	getURLRequest := protocols.GetUrlRequest{
		Params: "plus vieux",
	}
	err = json.NewEncoder(conn).Encode(getURLRequest)
	if err != nil {
		log.Fatal("Erreur d'envoi de la requête:", err)
	}

	// Lecture de la réponse GetUrlResponse
	var getURLResponse protocols.GetUrlResponse
	err = json.NewDecoder(conn).Decode(&getURLResponse)
	if err != nil {
		log.Fatal("Erreur de lecture de la réponse:", err)
	}

	// Logique métier pour la réponse GetUrlResponse

	// Exemple d'envoi d'une requête CreateUrlRequest
	createURLRequest := protocols.CreateUrlRequest{
		URL: "https://www.google.com",
	}
	err = json.NewEncoder(conn).Encode(createURLRequest)
	if err != nil {
		log.Fatal("Erreur d'envoi de la requête:", err)
	}

	// Lecture de la réponse CreateUrlResponse
	var createURLResponse protocols.CreateUrlResponse
	err = json.NewDecoder(conn).Decode(&createURLResponse)
	if err != nil {
		log.Fatal("Erreur de lecture de la réponse:", err)
	}
}
