package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
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
<<<<<<< HEAD
=======
	// TabString []string
>>>>>>> 8c504fc27bf83e7b90148c5a021b79ebd6ba02cf
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

func parseArtsists() []Artist {

	res, rej := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if rej != nil {
		fmt.Println(rej)
	}

	data, err := ioutil.ReadAll(res.Body)
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

	data, err := ioutil.ReadAll(res.Body)
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

	data, err := ioutil.ReadAll(res.Body)
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

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var relations OneRelation
	e := json.Unmarshal(data, &relations)
	if e != nil {
		fmt.Println("error:", e)
	}
	return relations
}

func main() {

	fileServer := http.FileServer(http.Dir("./data"))
	http.Handle("/style.css", fileServer)
	http.Handle("/desc.css", fileServer)

	artists := parseArtsists()

	http.HandleFunc("/groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
		variable, _ := template.ParseFiles("index.html")
		var TabToPrint []Artist
		minCrea, _ := strconv.Atoi(r.FormValue("minCrea"))
		maxCrea, _ := strconv.Atoi(r.FormValue("maxCrea"))

		if r.FormValue("submit") == "Envoyer" {
			for i := 0; i < len(artists); i++ {
				if minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("OneMember") != "on" && r.FormValue("TwoMember") != "on" && r.FormValue("ThreeMember") != "on" && r.FormValue("FourMember") != "on" && r.FormValue("FiveMember") != "on" && r.FormValue("SixMember") != "on" && r.FormValue("SevenMember") != "on" || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("OneMember") == "on" && len(artists[i].Members) == 1 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("TwoMember") == "on" && len(artists[i].Members) == 2 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("ThreeMember") == "on" && len(artists[i].Members) == 3 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("FourMember") == "on" && len(artists[i].Members) == 4 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("FiveMember") == "on" && len(artists[i].Members) == 5 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("SixMember") == "on" && len(artists[i].Members) == 6 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("SevenMember") == "on" && len(artists[i].Members) == 7 {
					TabToPrint = append(TabToPrint, artists[i])
				}
			}
			if minCrea == 1958 && maxCrea == 2015 && r.FormValue("OneMember") != "on" && r.FormValue("TwoMember") != "on" && r.FormValue("ThreeMember") != "on" && r.FormValue("FourMember") != "on" && r.FormValue("FiveMember") != "on" && r.FormValue("SixMember") != "on" && r.FormValue("SevenMember") != "on" {
				variable.Execute(w, artists)
			} else {
				variable.Execute(w, TabToPrint)
=======
<<<<<<< HEAD
		for i := 0; i < len(artists); i++ {
			artists[i].ToPrint = true
			fmt.Println(artists[i].ToPrint)
		}

		//faire une fonction qui prend en parametre un tab et qui retourne un tab qui a tt sauf ce qui ne dois pas s'afficher 
		//filter avec parametre tableau et fonction qui retourne un bool, si il est faux on lenleve sinon on le remet 
=======

>>>>>>> 69d415f... filter for members OK
		variable, _ := template.ParseFiles("index.html")
		minMembers, _ := strconv.Atoi(r.FormValue("minMembers"))
		maxMembers, _ := strconv.Atoi(r.FormValue("maxMembers"))
		var TabToPrint []Artist

		for i := 0; i < len(artists); i++ {
			if len(artists[i].Members) >= minMembers && len(artists[i].Members) <= maxMembers {
				TabToPrint = append(TabToPrint, artists[i])
>>>>>>> 8c504fc27bf83e7b90148c5a021b79ebd6ba02cf
			}
		} else {
			variable.Execute(w, artists)
		}
<<<<<<< HEAD
=======

		if minMembers == 0 && maxMembers == 0 {
			variable.Execute(w, artists)
		} else {
			variable.Execute(w, TabToPrint)
		}
>>>>>>> 8c504fc27bf83e7b90148c5a021b79ebd6ba02cf
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
