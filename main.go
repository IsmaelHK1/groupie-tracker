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
		var TabToPrint []Artist

		test1, _ := strconv.Atoi(r.FormValue("minCrea"))
		test2, _ := strconv.Atoi(r.FormValue("maxCrea"))

		// for i := 0; i < len(artists); i++ {
		// 	if r.FormValue("OneMember") == "on" && len(artists[i].Members) == 1 || r.FormValue("TwoMember") == "on" && len(artists[i].Members) == 2 || r.FormValue("ThreeMember") == "on" && len(artists[i].Members) == 3 || r.FormValue("FourMember") == "on" && len(artists[i].Members) == 4 || r.FormValue("FiveMember") == "on" && len(artists[i].Members) == 5 || r.FormValue("SixMember") == "on" && len(artists[i].Members) == 6 || r.FormValue("SevenMember") == "on" && len(artists[i].Members) == 7  {
		// 		TabToPrint = append(TabToPrint, artists[i])
		// 	}
		// }

		// if r.FormValue("OneMember") == "on" || r.FormValue("TwoMember") == "on" || r.FormValue("ThreeMember") == "on" || r.FormValue("FourMember") == "on" || r.FormValue("FiveMember") == "on" || r.FormValue("SixMember") == "on" || r.FormValue("SevenMember") == "on" {
		// 	variable.Execute(w, TabToPrint)
		// } else {
		// variable.Execute(w, artists)
		// }

		// for i := 0; i < len(artists); i++ {
		// 	if test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate {
		// 		TabToPrint = append(TabToPrint, artists[i])
		// 	}
		// }

		if r.FormValue("submit") == "Envoyer" {
			for i := 0; i < len(artists); i++ {
				if test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate && r.FormValue("OneMember") == "on" && len(artists[i].Members) == 1 || test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate && r.FormValue("TwoMember") == "on" && len(artists[i].Members) == 2 || test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate && r.FormValue("ThreeMember") == "on" && len(artists[i].Members) == 3 || test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate && r.FormValue("FourMember") == "on" && len(artists[i].Members) == 4 || test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate && r.FormValue("FiveMember") == "on" && len(artists[i].Members) == 5 || test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate && r.FormValue("SixMember") == "on" && len(artists[i].Members) == 6 || test1 <= artists[i].CreationDate && test2 >= artists[i].CreationDate && r.FormValue("SevenMember") == "on" && len(artists[i].Members) == 7  {
					TabToPrint = append(TabToPrint, artists[i])
				}
			}
			variable.Execute(w, TabToPrint)
		} 
		if r.FormValue("submit") == "Envoyer" && test1 <= 1958 && test2 >= 2015 {
			fmt.Println("zefroihzerouirze")
			variable.Execute(w, artists)
		} else {
			variable.Execute(w, artists)
		}

		// if r.FormValue("submit") == "Envoyer" && test1 == 1958 && test2 == 2015 && r.FormValue("OneMember") != "on" && r.FormValue("TwoMember") != "on" && r.FormValue("ThreeMember") != "on" && r.FormValue("FourMember") != "on" && r.FormValue("FiveMember") != "on" && r.FormValue("SixMember") != "on" && r.FormValue("SevenMember") != "on" {
		// 	fmt.Println("zefroihzerouirze")
		// 	variable.Execute(w, artists)
		// }

		// else if r.FormValue("OneMember") != "on" && r.FormValue("TwoMember") != "on" && r.FormValue("ThreeMember") != "on" && r.FormValue("FourMember") != "on" && r.FormValue("FiveMember") != "on" && r.FormValue("SixMember") != "on" && r.FormValue("SevenMember") != "on" {
		// 	variable.Execute(w, artists)
		// }


		// if r.FormValue("OneMember") != "on" && r.FormValue("TwoMember") != "on" && r.FormValue("ThreeMember") != "on" && r.FormValue("FourMember") != "on" && r.FormValue("FiveMember") != "on" && r.FormValue("SixMember") != "on" && r.FormValue("SevenMember") != "on" {
		// 	variable.Execute(w, artists)
		// }

		// if test1 >= 1 || test2 <= 2500 {
		// 	fmt.Println(TabToPrint)
		// 	variable.Execute(w, TabToPrint)
		// } else {
		// 	variable.Execute(w, artists)
		// 	}


		// add creationdate en mode negatif
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
