package main

import (
	"encoding/json"
	"log"
	"net"

	"example.com/finalproject/pkg/protocols"
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

	// Send a request to the Index to get a site listing
	getSiteRequest := protocols.GetSiteRequest{
		ListingType: "all",
	}
	err = json.NewEncoder(conn).Encode(getSiteRequest)
	if err != nil {
		log.Fatal("Error sending GetSiteRequest:", err)
	}

	// Receive the response from the Index
	var getSiteResponse protocols.GetSiteResponse
	err = json.NewDecoder(conn).Decode(&getSiteResponse)
	if err != nil {
		log.Fatal("Error receiving GetSiteResponse:", err)
	}

	// Process the received site listing

	// Send a request to the Index to get a file listing
	getFileRequest := protocols.GetFileRequest{}
	err = json.NewEncoder(conn).Encode(getFileRequest)
	if err != nil {
		log.Fatal("Error sending GetFileRequest:", err)
	}

	// Receive the response from the Index
	var getFileResponse protocols.GetFileResponse
	err = json.NewDecoder(conn).Decode(&getFileResponse)
	if err != nil {
		log.Fatal("Error receiving GetFileResponse:", err)
	}

}
