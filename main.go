package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	// "strings"
	"strconv"
)

type Artist struct {
	Id                int
	Image             string
	Name              string
	Members           []string
	CreationDate      int
	FirstAlbum        string
	Locations         string
	TabLocations      []string
	ConcertDates      string
	TabConcertDates   []string
	Relations         string
	TabRelation       OneRelation
	TabIndexRelation  []string
	TabLetterRelation [][]string
	ToPrint           bool
	// TabString []string
}

type Location struct {
	ID        int
	Locations []string
}

type ConcertDate struct {
	ID    int
	Dates []string
}

type Relations struct {
	Index []OneRelation
}

type OneRelation struct {
	Id             int
	DatesLocations map[string][]string
}

// map de string d'interface

func parseArtsists() []Artist {

	res, rej := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if rej != nil {
		fmt.Println(rej)
	}

	data, err := ioutil.ReadAll(res.Body) // data stock le JSON
	if err != nil {
		fmt.Println(err)
	}

	var artists []Artist
	e := json.Unmarshal(data, &artists)
	if e != nil {
		fmt.Println("error:", e)
	}
	return artists
}

func parseLocations(url string) Location {
	res, rej := http.Get(url)
	if rej != nil {
		fmt.Println(rej)
	}

	data, err := ioutil.ReadAll(res.Body) // data stock le JSON
	if err != nil {
		fmt.Println(err)
	}

	var locations Location
	e := json.Unmarshal(data, &locations)
	if e != nil {
		fmt.Println("error:", e)
	}
	return locations
}

func parseConcertDates(url string) ConcertDate {
	res, rej := http.Get(url)
	if rej != nil {
		fmt.Println(rej)
	}

	data, err := ioutil.ReadAll(res.Body) // data stock le JSON
	if err != nil {
		fmt.Println(err)
	}

	var concertDates ConcertDate
	e := json.Unmarshal(data, &concertDates)
	if e != nil {
		fmt.Println("error:", e)
	}
	return concertDates
}

func parseRelation(url string) OneRelation {
	res, rej := http.Get(url)
	if rej != nil {
		fmt.Println(rej)
	}

	data, err := ioutil.ReadAll(res.Body) // data stock le JSON
	if err != nil {
		fmt.Println(err)
	}

	var relations OneRelation
	e := json.Unmarshal(data, &relations)
	if e != nil {
		fmt.Println("error:", e)
	}
	// fmt.Println(relations)
	return relations
}

func main() {

	fileServer := http.FileServer(http.Dir("./data"))
	http.Handle("/style.css", fileServer)
	http.Handle("/desc.css", fileServer)

	artists := parseArtsists()

	http.HandleFunc("/groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
		variable, _ := template.ParseFiles("index.html")
		IDIsma, _ := strconv.Atoi(r.FormValue("test"))

		// for i := 0; i < len(artists); i++ {
		// 	artists[i].ToPrint = true
		// 	if len(artists[i].Members) != IDIsma {
		// 		artists[i].ToPrint = true
		// 	}
		// }

		if r.FormValue("test") != "Envoyer" {
			for i := 0; i < len(artists); i++ {
				if len(artists[i].Members) == IDIsma {
					artists[i].ToPrint = true
				} else {
					artists[i].ToPrint = false
				}
			}
		}

		variable.Execute(w, artists)
	})

	http.HandleFunc("/groupie-tracker/", func(w http.ResponseWriter, r *http.Request) {
		variable, _ := template.ParseFiles("artists.html")
		ArtistPath := r.URL.Path[17:]
		IdArtist, _ := strconv.Atoi(ArtistPath)
		IdArtist--

		// Valeurs
		artists[IdArtist].TabRelation = parseRelation(artists[IdArtist].Relations)
		variable.Execute(w, artists[IdArtist])

	})

	fmt.Println("vas y le serv marche")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
