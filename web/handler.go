package web

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func BaseHandler(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := []ArtistPage{}
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		maxCreation := r.FormValue("max-creation")
		minCreation := r.FormValue("min-creation")
		minAlbum := r.FormValue("min-album")
		maxAlbum := r.FormValue("max-album")

		filter := r.Form["members"]
		var num []int
		for _, number := range filter {
			n, err := strconv.Atoi(number)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			num = append(num, n)
		}
		for _, band := range AllArtistPage {
			passesCreation := true
			passesMembers := true
			passAlbum := true

			if minCreation != "" && maxCreation != "" {
				maxNum, err := strconv.Atoi(maxCreation)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				minNum, err := strconv.Atoi(minCreation)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				passesCreation = MatchesCreationFilter(band, minNum, maxNum)
			}
			if len(filter) > 0 {

				passesMembers = MatchesMemberFilter(band, num)
			}
			if minAlbum != "" && maxAlbum != "" {
				num1, err := strconv.Atoi(minAlbum)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				num2, err := strconv.Atoi(maxAlbum)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				passAlbum = MatchesAlbumFilter(band, num1, num2)
			}
			if passesCreation && passesMembers && passAlbum {
				data = append(data, band)
			}

		}

		BaseHandler(w, "home.html", data)
	}
	if r.Method == http.MethodGet {
		BaseHandler(w, "home.html", AllArtistPage)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/artist/")
	if path == "" {
		http.Error(w, "missing artist id", http.StatusNotFound)
		return
	}
	pathID, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	target, err := GetArtistDataByID(pathID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	BaseHandler(w, "artist.html", target)
}
