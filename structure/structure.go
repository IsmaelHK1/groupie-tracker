package structure

//Artist is the struct for parsing the api artists
type Artist struct {
	//-- data from json --\\
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	URLRelations string   `json:"relations"`

	//-- data created --\\
	//store Relations' informations in TabRelation
	TabRelation map[string][]string
}

//ConcertDate is the struct for parsing the api dates
type ConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

//Relations is []OneRelation
type Relations struct {
	Index []OneRelation
}

//OneRelation is the struct for parsing the api relation
type OneRelation struct {
	ID             int
	DatesLocations map[string][]string
}
