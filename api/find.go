package api

func ArtistByID(artists []Artist, id int) *Artist {
	for i := range artists {
		if artists[i].ID == id {
			return &artists[i]
		}
	}
	return nil
}

func RelationByID(relations []Relation, id int) *Relation {
	for i := range relations {
		if relations[i].ID == id {
			return &relations[i]
		}
	}
	return nil
}
