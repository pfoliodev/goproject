package main

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"example.com/finalproject/pkg/protocols"
)

func main() {
	for {
		conn, err := net.Dial("tcp", "localhost:8000")
		if err != nil {
			log.Fatal(err)
		}

		// Send a request to the Index to get a site to visit
		getURLRequest := protocols.GetUrlRequest{}
		err = json.NewEncoder(conn).Encode(getURLRequest)
		if err != nil {
			log.Fatal("Error sending GetUrlRequest:", err)
		}

		// Receive the response from the Index
		var getURLResponse protocols.GetUrlResponse
		err = json.NewDecoder(conn).Decode(&getURLResponse)
		if err != nil {
			log.Fatal("Error receiving GetUrlResponse:", err)
		}

		// Check if there is a site to visit
		if getURLResponse.URL == "" {
			// No site to visit, wait for 5 seconds before retrying
			time.Sleep(5 * time.Second)
			continue
		}

		// Visit the site
		visitSite(getURLResponse.URL)

		// Send multiple GetFileRequest / CreateFileRequest or UpdateFileRequest to the Index
		sendFileRequests(conn)

		// Receive the response from the Index
		receiveResponses(conn)

		// Send an UpdateSiteRequest to the Index
		updateSite(conn)

		// Receive the response from the Index
		receiveResponses(conn)

		// Close the connection to the Index
		conn.Close()
	}
}

func visitSite(url string) {
	// Placeholder for visiting the site
	log.Printf("Visiting site: %s\n", url)

	// Implement your logic to visit the site here
}

func sendFileRequests(conn net.Conn) {
	// Placeholder for sending file requests to the Index

	// Create and send multiple file requests to the Index
	getFileRequest := protocols.GetFileRequest{}
	err := json.NewEncoder(conn).Encode(getFileRequest)
	if err != nil {
		log.Fatal("Error sending GetFileRequest:", err)
	}

	// Create and send create file requests to the Index
	createFileRequest := protocols.CreateFileRequest{}
	err = json.NewEncoder(conn).Encode(createFileRequest)
	if err != nil {
		log.Fatal("Error sending CreateFileRequest:", err)
	}

	// Create and send update file requests to the Index
	updateFileRequest := protocols.UpdateFileRequest{}
	err = json.NewEncoder(conn).Encode(updateFileRequest)
	if err != nil {
		log.Fatal("Error sending UpdateFileRequest:", err)
	}

	// Implement your logic to send file requests here
}

func updateSite(conn net.Conn) {
	// Placeholder for updating the site

	// Create and send an update site request to the Index
	updateSiteRequest := protocols.UpdateSiteRequest{}
	err := json.NewEncoder(conn).Encode(updateSiteRequest)
	if err != nil {
		log.Fatal("Error sending UpdateSiteRequest:", err)
	}

	// Implement your logic to update the site here
}

func receiveResponses(conn net.Conn) {
	// Placeholder for receiving responses from the Index

	// Read and decode the response from the Index
	var response protocols.Response
	err := json.NewDecoder(conn).Decode(&response)
	if err != nil {
		log.Fatal("Error receiving response:", err)
	}

	// Handle the received response based on its type
	switch response.Type {
	case protocols.ResponseTypeGetFile:
		getFileResponse := response.GetFileResponse
		// Process the GetFileResponse
		for _, file := range getFileResponse.Files {
			// Access individual file information
			fileID := file.ID
			fileName := file.Name
			fileURL := file.URL
			// Perform necessary actions with the file information
			// ...
		}
	case protocols.ResponseTypeCreateFile:
		createFileResponse := response.CreateFileResponse
		// Process the CreateFileResponse
		createdFileID := createFileResponse.ID
		createdFileName := createFileResponse.Name
		createdFileURL := createFileResponse.URL
		// Perform necessary actions with the created file information
		// ...
	case protocols.ResponseTypeUpdateFile:
		updateFileResponse := response.UpdateFileResponse
		// Process the UpdateFileResponse
		updatedFileID := updateFileResponse.ID
		updatedFileName := updateFileResponse.Name
		updatedFileURL := updateFileResponse.URL
		// Perform necessary actions with the updated file information
		// ...
	case protocols.ResponseTypeGetSite:
		getSiteResponse := response.GetSiteResponse
		// Process the GetSiteResponse
		siteID := getSiteResponse.ID
		siteURL := getSiteResponse.URL
		// Perform necessary actions with the site information
		// ...
	default:
		log.Println("Unknown response type:", response.Type)
	}

	// Implement your logic to process the received responses here
}
