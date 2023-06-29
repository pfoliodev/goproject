package main

import (
	"encoding/json"
	"fmt"
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
	getURLReq := protocols.GetUrlRequest{
		Params: "example",
	}

	reqJSON, err := json.Marshal(getURLReq)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write(reqJSON)
	if err != nil {
		log.Fatal(err)
	}

	// Exemple de réception de la réponse GetUrlResponse
	var response protocols.GetUrlRequest
	err = json.NewDecoder(conn).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	// Traitez la réponse reçue
	fmt.Println(response)
	fmt.Printf("%#v\n", response)
}
