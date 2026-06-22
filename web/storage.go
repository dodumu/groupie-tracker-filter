package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetArtist() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code")
	}
	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)

	if err != nil {
		return nil, err
	}

	return artists, nil
}

func GetLocation() ([]Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code")
	}
	var locations LocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locations)

	if err != nil {
		return nil, err
	}

	return locations.Index, nil
}

func GetDates() ([]Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code")
	}
	var dates DateResponse
	err = json.NewDecoder(resp.Body).Decode(&dates)

	if err != nil {
		return nil, err
	}

	return dates.Index, nil
}

func GetRelation() ([]Relation, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code")
	}
	var relations RelationResponse
	err = json.NewDecoder(resp.Body).Decode(&relations)

	if err != nil {
		return nil, err
	}

	return relations.Index, nil
}

func SaveAllData() error {
	artist, err := GetArtist()
	if err != nil {
		return err
	}
	AllArtist = artist

	location, err := GetLocation()
	if err != nil {
		return err
	}
	AllLocation = location

	dates, err := GetDates()
	if err != nil {
		return err
	}
	AllDates = dates

	relation, err := GetRelation()
	if err != nil {
		return err
	}
	AllRelation = relation

	allartist, err := GetAllArtistPage()
	if err != nil {
		return err
	}

	AllArtistPage = allartist

	return nil
}
