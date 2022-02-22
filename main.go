package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

var (
	Artists   = "https://groupietrackers.herokuapp.com/api/artists"
	Locations = "https://groupietrackers.herokuapp.com/api/locations"
	Dates     = "https://groupietrackers.herokuapp.com/api/dates"
	Relations = "https://groupietrackers.herokuapp.com/api/relation"
)

type GroupieTracker struct {
	Artist    []artist
	Location  []locations
	Dates     []dates
	Relations []relations
}

type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type loc struct {
	Index []locations `json:"index"`
}

type locations struct {
	Id        int      `json:"id"`
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

func unmarshalArtist() []artist {

	resp, err := http.Get(Artists)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body) // response body is []byte
	if err2 != nil {
		panic(err.Error())
	}

	var data []artist
	if err3 := json.Unmarshal(body, &data); err3 != nil { // Parse []byte to the go struct pointer
		log.Println(err)
	}
	return data

}

func unmarshalLocation() []locations {
	resp, err := http.Get(Locations)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body) // response body is []byte
	if err2 != nil {
		panic(err.Error())
	}

	var data loc
	if err3 := json.Unmarshal(body, &data); err3 != nil { // Parse []byte to the go struct pointer
		log.Println(err)
	}
	return data.Index

}

func unmarshalDates() []dates {
	resp, err := http.Get(Dates)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body) // response body is []byte
	if err2 != nil {
		panic(err.Error())
	}

	var data dat
	if err3 := json.Unmarshal(body, &data); err3 != nil { // Parse []byte to the go struct pointer
		log.Println(err)
	}
	return data.Index
}

func unmarshalRel() []relations {
	resp, err := http.Get(Relations)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body) // response body is []byte
	if err2 != nil {
		panic(err.Error())
	}

	var data rel
	if err3 := json.Unmarshal(body, &data); err3 != nil { // Parse []byte to the go struct pointer
		log.Println(err)
	}
	return data.Index

}

func getData(w http.ResponseWriter, r *http.Request) {


	var bandInfo GroupieTracker

	bandInfo.Location = unmarshalLocation()
	bandInfo.Artist = unmarshalArtist()
	bandInfo.Dates = unmarshalDates()
	bandInfo.Relations = unmarshalRel()
	

	/*for i := range bandInfo.Artist {
		fmt.Println(bandInfo.Artist[i])
		fmt.Println(bandInfo.Location[i])
		fmt.Println(bandInfo.Dates[i])
		fmt.Println(bandInfo.Relations[i])
		fmt.Println("-=-=-=-=-=-=-=-")
*/
tpl.ExecuteTemplate(w, "index.html", bandInfo)
	}


	


func main() {

	http.HandleFunc("/index.html", getData)
	fmt.Println("Starting the server on :9090...")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	log.Fatal(http.ListenAndServe(":9090", nil))

}
