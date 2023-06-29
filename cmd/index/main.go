package main

import (
	"encoding/json"
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

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Exemple de réception d'une requête GetUrlRequest
	var request protocols.GetUrlRequest
	err := json.NewDecoder(conn).Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}

	// Traitez la requête reçue

	// Exemple d'envoi de la réponse GetUrlResponse
	response := protocols.GetUrlRequest{
		// Définissez ici les champs de la réponse
	}

	resJSON, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = conn.Write(resJSON)
	if err != nil {
		log.Println(err)
		return
	}
}
