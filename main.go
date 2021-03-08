package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

//verifier que les noms de variables collent bien avec le html, y a eu des modif

//Artist is the struct to parse the api artists
type Artist struct {
	//-- data from json --\\
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	URLRelations string   `json:"relations"`

	//-- data from json, not used --\\
	//UrlLocations 		string		`json:"locations"`
	//UrlConcertDate		string 		`json:"concertDates"`

	//-- data created --\\
	TabRelation OneRelation
}

//ConcertDate is the struct to parse the api dates
type ConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

//Relations is []OneRelation
type Relations struct {
	Index []OneRelation
}

//OneRelation is the struct to parse the api relation
type OneRelation struct {
	ID             int
	DatesLocations map[string][]string
}

//retrieveJSON read json from api link and return json in []byte
func retrieveJSON(url string) []byte {
	//-- read json from url --\\
	response, reject := http.Get(url)
	if reject != nil {
		fmt.Println(reject)
	}
	//-- return the json --\\
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

//parseJSONArtists unmarshal the json and return it in []Artist
func parseJSONArtsists(url string) []Artist {
	var artists []Artist
	data := retrieveJSON(url)
	err := json.Unmarshal(data, &artists)
	if err != nil {
		fmt.Println("error while unmarshal artist:", err)
	}
	return artists
}

//parseJSONRelation unmarshal the json and return it in OneRelation
func parseJSONRelation(url string) OneRelation {
	var relations OneRelation
	data := retrieveJSON(url)
	err := json.Unmarshal(data, &relations)
	if err != nil {
		fmt.Println("error:", err)
	}
	return relations
}

//serveFile make files available for the website
func serveFile() {
	fileServer := http.FileServer(http.Dir("./data"))
	http.Handle("/style.css", fileServer)
	http.Handle("/desc.css", fileServer)
}

func handleGroupieTracker(artists []Artist) {
	http.HandleFunc("/groupie-tracker", func(w http.ResponseWriter, r *http.Request) {
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
			}
		} else if r.FormValue("search") != "" {
			filter := r.FormValue("search")
			for i := 0; i < len(artists); i++ {
				for _, value := range artists[i].Members {
					if strings.ToUpper(filter) == strings.ToUpper(value) {
						TabToPrint = append(TabToPrint, artists[i])
					}
				}
				artists[i].TabRelation = parseJSONRelation(artists[i].URLRelations)
				for index, value := range artists[i].TabRelation.DatesLocations {
					for _, value2 := range value {
						if value2 == filter {
							TabToPrint = append(TabToPrint, artists[i])
						}
					}
					if index == filter {
						TabToPrint = append(TabToPrint, artists[i])
					}
				}

				if strings.ToUpper(filter) == strings.ToUpper(artists[i].Name) || strings.ToUpper(filter) == strings.ToUpper(artists[i].FirstAlbum) || filter == strconv.Itoa(artists[i].CreationDate) {
					TabToPrint = append(TabToPrint, artists[i])
				}
			}
			variable.Execute(w, TabToPrint)
		} else {
			variable.Execute(w, artists)
		}
	})
}

func handleArtist(artists []Artist) {
	http.HandleFunc("/groupie-tracker/", func(w http.ResponseWriter, r *http.Request) {
		variable, _ := template.ParseFiles("artists.html")
		ArtistPath := r.URL.Path[17:]
		IDArtist, _ := strconv.Atoi(ArtistPath)
		IDArtist--

		// Valeurs
		artists[IDArtist].TabRelation = parseJSONRelation(artists[IDArtist].URLRelations)
		variable.Execute(w, artists[IDArtist])
	})
}

func runServer() {
	fmt.Println("server is runing")
	if err := http.ListenAndServe(":7070", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {

	serveFile()
	handleGroupieTracker(parseJSONArtsists("https://groupietrackers.herokuapp.com/api/artists"))
	handleArtist(parseJSONArtsists("https://groupietrackers.herokuapp.com/api/artists"))
	runServer()

}
