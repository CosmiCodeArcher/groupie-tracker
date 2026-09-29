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
        CreationDate int      `json:"creationDate"`
        FirstAlbum   string   `json:"firstAlbum"`
        Locations    string   `json:"locations"`
        ConcertDates string   `json:"concertDates"`
        Relations    string   `json:"relations"`
}

type Relation struct {
        ID              int                 `json:"id"`
        DatesLocations  map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
        Index []Relation `json:"index"`
}

const baseUrl = "https://groupietrackers.herokuapp.com"

func fetchArtists() ([]Artist, error) {

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(baseUrl + "/api/artists")

	if err != nil { 
		return []Artist{}, fmt.Errorf("Could not fetch artists: %w", err)
  	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Artist{}, fmt.Errorf("Server Error: %w", resp.StatusCode)
	}

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return []Artist{}, fmt.Errorf("failed to parse json: %w", err)
	}

	return artists, nil
}

func fetchRelations() (RelationIndex, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	
	resp, err := client.Get(baseUrl + "/api/relation")
	if err != nil {
		return RelationIndex{}, fmt.Errorf("Could not fetch relations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RelationIndex{}, fmt.Errorf("server error status: %w", resp.StatusCode)
	}

	var relationIndex RelationIndex
	err = json.NewDecoder(resp.Body).Decode(&relationIndex)
	if err != nil {
		return RelationIndex{}, fmt.Errorf("failed to parse json: %w", err)
	}

	return relationIndex, nil
}
