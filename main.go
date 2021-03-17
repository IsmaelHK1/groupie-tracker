package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie-tracker/structure"
	"groupie-tracker/tools"
)

//serveFile makes files available for the website
func serveFile() {
	fileServer := http.FileServer(http.Dir("./data"))
	http.Handle("/main.css", fileServer)
	http.Handle("/index.css", fileServer)
	http.Handle("/artists.css", fileServer)
	http.Handle("/header.css", fileServer)
	http.Handle("/logo.png", fileServer)
	http.Handle("/filtre.png", fileServer)
}

func filters(artists, TabToPrint []structure.Artist, variable *template.Template, w http.ResponseWriter, r *http.Request) {
	minCrea, _ := strconv.Atoi(r.FormValue("minCrea"))
	maxCrea, _ := strconv.Atoi(r.FormValue("maxCrea"))
	for i := 0; i < len(artists); i++ {
		if minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("OneMember") != "on" && r.FormValue("TwoMember") != "on" && r.FormValue("ThreeMember") != "on" && r.FormValue("FourMember") != "on" && r.FormValue("FiveMember") != "on" && r.FormValue("SixMember") != "on" && r.FormValue("SevenMember") != "on" || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("OneMember") == "on" && len(artists[i].Members) == 1 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("TwoMember") == "on" && len(artists[i].Members) == 2 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("ThreeMember") == "on" && len(artists[i].Members) == 3 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("FourMember") == "on" && len(artists[i].Members) == 4 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("FiveMember") == "on" && len(artists[i].Members) == 5 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("SixMember") == "on" && len(artists[i].Members) == 6 || minCrea <= artists[i].CreationDate && maxCrea >= artists[i].CreationDate && r.FormValue("SevenMember") == "on" && len(artists[i].Members) == 7 {
			TabToPrint = append(TabToPrint, artists[i])
		}
	}
	if minCrea == 1958 && maxCrea == 2015 && r.FormValue("OneMember") != "on" && r.FormValue("TwoMember") != "on" && r.FormValue("ThreeMember") != "on" && r.FormValue("FourMember") != "on" && r.FormValue("FiveMember") != "on" && r.FormValue("SixMember") != "on" && r.FormValue("SevenMember") != "on" {
		variable.Execute(w, artists)
	} else {
		variable.Execute(w, TabToPrint)
	}
}

func searchBar(artists, TabToPrint []structure.Artist, variable *template.Template, w http.ResponseWriter, r *http.Request) {
	filter := r.FormValue("search")
	for i := 0; i < len(artists); i++ {
		for _, value := range artists[i].Members {
			if strings.ToUpper(strings.Join(strings.Split(filter, " -Members"), "")) == strings.ToUpper(value) {
				TabToPrint = append(TabToPrint, artists[i])
			}
		}
		for index, value := range artists[i].TabRelation {
			for _, value2 := range value {
				if value2 == strings.Join(strings.Split(filter, " -ConcertDate"), "") {
					TabToPrint = append(TabToPrint, artists[i])
				}
			}
			if index == strings.Join(strings.Split(filter, " -Location"), "") {
				TabToPrint = append(TabToPrint, artists[i])
			}
		}

		if strings.ToUpper(strings.Join(strings.Split(filter, " -Name"), "")) == strings.ToUpper(artists[i].Name) || strings.ToUpper(strings.Join(strings.Split(filter, " -FirstAlbum"), "")) == strings.ToUpper(artists[i].FirstAlbum) || strings.Join(strings.Split(filter, " -CreationDate"), "") == strconv.Itoa(artists[i].CreationDate) {
			TabToPrint = append(TabToPrint, artists[i])
		}
	}
	variable.Execute(w, TabToPrint)
}

//handleGroupieTracker is the handle function for the main page (index.html)
func handleGroupieTracker(artists []structure.Artist) {
	http.HandleFunc("/groupie-tracker", func(w http.ResponseWriter, r *http.Request) {

		variable, _ := template.ParseFiles("index.html")
		var TabToPrint []structure.Artist

		if r.FormValue("submit") == "Envoyer" {
			filters(artists, TabToPrint, variable, w, r)
		} else if r.FormValue("search") != "" {
			searchBar(artists, TabToPrint, variable, w, r)
		} else {
			variable.Execute(w, artists)
		}

	})
}

//handleArtist is the handle function for the artist page (artists.html)
func handleArtist(artists []structure.Artist) {
	http.HandleFunc("/groupie-tracker/", func(w http.ResponseWriter, r *http.Request) {
		variable, _ := template.ParseFiles("artists.html")
		ArtistPath := r.URL.Path[17:]
		IDArtist, _ := strconv.Atoi(ArtistPath)
		IDArtist--

		// Values
		variable.Execute(w, artists[IDArtist])
	})
}

//runServer sets the listenandserve port to 8080
func runServer() {
	fmt.Println("server is runing")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serveFile()
	jsonArtist := tools.ParseJSONArtsists("https://groupietrackers.herokuapp.com/api/artists")
	handleGroupieTracker(jsonArtist)
	handleArtist(jsonArtist)
	runServer()
}
