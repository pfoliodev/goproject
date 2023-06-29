package main

import (
	"encoding/json"
	"log"
	"net"
)

type SearchResult struct {
	// Définissez ici les champs du résultat de recherche
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Exemple d'envoi d'une requête de recherche
	searchReq := "example"

	_, err = conn.Write([]byte(searchReq))
	if err != nil {
		log.Fatal(err)
	}

	// Exemple de réception des résultats de recherche
	var results []SearchResult
	err = json.NewDecoder(conn).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}

	// Traitez les résultats de recherche reçus
}
