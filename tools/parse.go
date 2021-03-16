package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"groupie-tracker/structure"
)

//retrieveJSON reads json from api link and returns json in []byte
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

//parseJSONArtists unmarshals the json and returns it in []structure.Artist
func ParseJSONArtsists(url string) []structure.Artist {
	var artists []structure.Artist
	relations := ParseJSONRelation("https://groupietrackers.herokuapp.com/api/relation")
	data := retrieveJSON(url)
	err := json.Unmarshal(data, &artists)
	if err != nil {
		fmt.Println("error while unmarshal artist:", err)
	}
	for i := 0; i < len(artists); i++ {
		artists[i].TabRelation = relations.Index[i].DatesLocations
	}
	return artists
}

//parseJSONRelation unmarshals the json and returns it in OneRelation
func ParseJSONRelation(url string) structure.Relations {
	var relations structure.Relations
	data := retrieveJSON(url)
	err := json.Unmarshal(data, &relations)
	if err != nil {
		fmt.Println("error:", err)
	}
	return relations
}
