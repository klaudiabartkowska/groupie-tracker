package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	Artists   = "https://groupietrackers.herokuapp.com/api/artists"
	Locations = "https://groupietrackers.herokuapp.com/api/locations"
	Dates     = "https://groupietrackers.herokuapp.com/api/dates"
	Relations = "https://groupietrackers.herokuapp.com/api/relation"
)

type GroupieTracker struct {
	artist    []artist
	location  []location
	dates     []dates
	relations []relations
}

type artist []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"` 
	CreationDate uint     `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type loc struct {
	Index []location `json:"index"`
}

type location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type rel struct {
	Index []relations `json:"index"`
}

type relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type dat struct {
	Index []dates `json:"index"`
}

type dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

func UnmarshalArtist() {

	resp, err := http.Get(Artists)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var data artist
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	fmt.Println(data)

	
}

func UnmarshalLocation()  {
	resp, err := http.Get(Locations)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var data loc
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

}
func UnmarshalDates() {
	resp, err := http.Get(Dates)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var data dat
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}


}

func UnmarshalRel()  {
	resp, err := http.Get(Relations)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var data rel
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

}

func main() {


}
