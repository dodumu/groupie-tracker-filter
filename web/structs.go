package web

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"location"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relation"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID            int                 `json:"id"`
	DateLocations map[string][]string `json:"datesLocations"`
}

type LocationResponse struct {
	Index []Location `json:"index"`
}

type DateResponse struct {
	Index []Date `json:"index"`
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}

type ArtistPage struct {
	Artist   Artist
	Location []string
	Date     []string
	Relation map[string][]string
}
