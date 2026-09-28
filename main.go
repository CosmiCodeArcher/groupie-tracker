type Artist struct {
        ID           int    `json:"id"`
        Image        string `json:"image"`
        Name         string `json:"name"`
        Members      []string `json:"members"`
        CreationDate int    `json:"creationdate`
        FirstAlbum   string `json:"firstalbum"`
        Locations    string `json:"locations"`
        ConcertDates string `json:"concertdates"`
        Relations    string `json:"relations"`
}

type Relation struct {
        ID              int                 `json:"id"`
        DatesLocations  map[string][]string `json:"datesLocations"`
}

type RelationIndex struct {
        Index []Relation `json:"index"`
}
