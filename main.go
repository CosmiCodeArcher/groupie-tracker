package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Artist struct {
        ID           int      `json:"id"`
        Image        string   `json:"image"`
        Name         string   `json:"name"`
        Members      []string `json:"members"`
        CreationDate int      `json:"creationdate"`
        FirstAlbum   string   `json:"firstalbum"`
        Locations    string   `json:"locations"`
        ConcertDates string   `json:"concertdates"`
        Relations    string   `json:"relations"`
}

type Relation struct {
        ID              int                 `json:"id"`
        DatesLocations  map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
        Index []Relation `json:"index"`
}

func fetchArtists() ([]Artist, error) {

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("/api/artists")

	if err != nil { 
		return []Artist{}, fmt.Errorf("Could not fetch artists: %v", err)
  	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Artist{}, fmt.Errorf("Server Error: %v", resp.StatusCode)
	}

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return []Artist{}, fmt.Errorf("failed to parse json: %v", err)
	}

	return artists, nil
}

func fetchRelations() (RelationIndex, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	
	resp, err := client.Get("/api/relation")
	if err != nil {
		return RelationIndex{}, fmt.Errorf("Could not fetch relations: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RelationIndex{}, fmt.Errorf("server error status: %d", resp.StatusCode)
	}

	var relationIndex RelationIndex
	err = json.NewDecoder(resp.Body).Decode(&relationIndex)
	if err != nil {
		return RelationIndex{}, fmt.Errorf("failed to parse json: %v", err)
	}

	return relationIndex, nil
}
