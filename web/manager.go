package web

import "fmt"

func GetArtistByID(id int) (Artist, error) {
	for _, artist := range AllArtist {
		if artist.ID == id {
			return artist, nil
		}
	}
	return Artist{}, fmt.Errorf("artist not found")
}

func GetLocationByID(id int) (Location, error) {
	for _, locations := range AllLocation {
		if locations.ID == id {
			return locations, nil
		}
	}
	return Location{}, fmt.Errorf("locations not found")
}

func GetDatesByID(id int) (Date, error) {
	for _, dates := range AllDates {
		if dates.ID == id {
			return dates, nil
		}
	}
	return Date{}, fmt.Errorf("dates for user with id not available")
}

func GetRelationByID(id int) (Relation, error) {
	for _, relations := range AllRelation {
		if relations.ID == id {
			return relations, nil
		}
	}
	return Relation{}, fmt.Errorf("relations not found")
}

func GetArtistDataByID(id int) (ArtistPage, error) {
	band, err := GetArtistByID(id)
	if err != nil {
		return ArtistPage{}, err
	}
	locations, err := GetLocationByID(id)
	if err != nil {
		return ArtistPage{}, err
	}
	date, err := GetDatesByID(id)
	if err != nil {
		return ArtistPage{}, err
	}
	relation, err := GetRelationByID(id)
	if err != nil {
		return ArtistPage{}, err
	}
	artistData := ArtistPage{
		Artist:   band,
		Location: locations.Locations,
		Date:     date.Dates,
		Relation: relation.DateLocations,
	}
	return artistData, nil
}

func GetAllArtistPage() ([]ArtistPage, error) {
	var result []ArtistPage
	for _, band := range AllArtist {

		allArtistpages, err := GetArtistDataByID(band.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, allArtistpages)
	}
	return result, nil
}

func MatchesCreationFilter(artist ArtistPage, minYear int, maxyear int) bool {
	if artist.Artist.CreationDate >= minYear && artist.Artist.CreationDate <= maxyear {
		return true
	}
	return false
}

func MatchesMemberFilter(artist ArtistPage, checks []int) bool {
	for _, num := range checks {
		if len(artist.Artist.Members) == num {
			return true
		}
	}
	return false
}

func MatchesAlbumFilter(artist ArtistPage, year string) bool {
	for _, bands := range AllArtistPage {
		var albumYear string
		album := bands.Artist.FirstAlbum
		for i := len(album) - 5; i < len(album); i++ {

			albumYear += string(album[i])
		}
		if albumYear == year {
			return true
		}
	}
	return false
}
