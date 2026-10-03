package main

import (
	"fmt"
	"groupie-tracker/api"
	"os"
)

func main() {
	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	relations, err := api.FetchRelations()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	id := 1
	artist := api.ArtistByID(artists, id)
	if artist == nil {
		fmt.Fprintf(os.Stderr, "artist %d not found", id)
		return
	}
	relation := api.RelationByID(relations.Index, artist.ID)
	if relation == nil {
		fmt.Fprintln(os.Stderr, "artist 1's relation not found")
		return
	}

	fmt.Println(artist.Name)
	for location, dates := range relation.DatesLocations {
		fmt.Println(location, dates)
	}
}
