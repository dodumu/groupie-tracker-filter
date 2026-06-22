package main

import (
	"fmt"
	"groupie/web"
	"log"
	"net/http"
)

func main() {

	fmt.Println("server is running on http://localhost:8095")

	err := web.SaveAllData()
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/artist/", web.ArtistHandler)

	http.ListenAndServe(":8095", nil)
}
