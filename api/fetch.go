package api

import "fmt"

const baseURL = "https://groupietrackers.herokuapp.com/api"

func FetchArtists() ([]Artist, error) {
	var artists []Artist
	if err := fetchJSON(baseURL+"/artists", &artists); err != nil {
		return nil, fmt.Errorf("fetch artists: %w", err)
	}
	return artists, nil
}

func FetchRelations() (RelationIndex, error) {
	var relationIndex RelationIndex
	if err := fetchJSON(baseURL+"/relation", &relationIndex); err != nil {
		return RelationIndex{}, fmt.Errorf("fetch relations: %w", err)
	}
	return relationIndex, nil
}
