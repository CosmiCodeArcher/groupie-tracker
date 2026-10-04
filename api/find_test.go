package api

import "testing"

func TestArtistByID(t *testing.T) {
	validID := 7
	invalidID := 999
	testArtists := []Artist{
		{ID: 1},
		{ID: 7},
		{ID: 9},
	}

	testValidArtist := ArtistByID(testArtists, validID)
	if testValidArtist == nil {
		t.Fatalf("ArtistByID(%d) = nil, want artist %d", validID, validID)
	}
	if testValidArtist.ID != validID {
		t.Errorf("ArtistByID = %d, want ArtistByID = %d", testValidArtist.ID, validID)
	}

	testInvalidArtist := ArtistByID(testArtists, invalidID)
	if testInvalidArtist != nil {
		t.Errorf("ArtistByID = %d, want nil", testInvalidArtist.ID)
	}

	testEmptyArtists := ArtistByID([]Artist{}, validID)
	if testEmptyArtists != nil {
		t.Fatalf("ArtistByID = %v, want nil", testEmptyArtists)
	}
}

func TestRelationByID(t *testing.T) {
	validID := 7
	invalidID := 999
	testRelations := []Relation{
		{ID: 1},
		{ID: 7},
		{ID: 9},
	}

	testValidRelation := RelationByID(testRelations, validID)
	if testValidRelation == nil {
		t.Fatalf("RelationByID(%d) = nil, want relation %d", validID, validID)
	}
	if testValidRelation.ID != validID {
		t.Errorf("RelationByID = %d, want RelationByID = %d", testValidRelation.ID, validID)
	}

	testInvalidRelation := RelationByID(testRelations, invalidID)
	if testInvalidRelation != nil {
		t.Errorf("RelationByID = %d, want nil", testInvalidRelation.ID)
	}

	testEmptyRelations := RelationByID([]Relation{}, validID)
	if testEmptyRelations != nil {
		t.Fatalf("RelationByID = %v, want nil", testEmptyRelations)
	}
}
