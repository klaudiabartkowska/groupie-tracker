package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	"strconv"
	"strings"
)

var tpl *template.Template

var tplArtists *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/index.html"))
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
	// Locations []string 
	// DatesLocations map[string][]string 
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

	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}

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



func getArtists(w http.ResponseWriter, r *http.Request) {

	
	if r.URL.Path != "/artist"{
		http.NotFound(w, r)
		return
	}
	
	var bandInfo GroupieTracker

	bandInfo.Location = unmarshalLocation()
	bandInfo.Artist = unmarshalArtist()
	bandInfo.Dates = unmarshalDates()
	bandInfo.Relations = unmarshalRel()

	if err := r.ParseForm(); err !=nil{
		http.Error(w, "404", 400)
	}

	artistName := r.FormValue("infoArtist")
	if artistName == ""{
		http.Error(w, " 400 Bad request", http.StatusBadRequest)
		return 
	}
	numArtist, err := strconv.Atoi(artistName)
	if err != nil {
		// handle error
		http.Error(w, " 400 Bad request", http.StatusBadRequest)
		return 
	
	}


	fmt.Fprintln(w, "<title>"+bandInfo.Artist[numArtist-1].Name+"</title>")

									// Artist Name//
	fmt.Fprintln(w, "<h6>"+bandInfo.Artist[numArtist-1].Name+"</h6>")

	   						//Image// 
	fmt.Fprint(w, "<h3>"+"<img src="+bandInfo.Artist[numArtist-1].Image+">"+"</h3>")

	  							// Members//
	fmt.Fprint(w, "<h1>"+"Band Members"+"</h1>")
	fmt.Fprintln(w, "<h2>"+strings.Join(bandInfo.Artist[numArtist-1].Members, " "+"<br>")+"</h2>")
                      // First Album //
	fmt.Fprintln(w, "<h1>"+"First Album "+"</h1>"+"<h2>"+bandInfo.Artist[numArtist-1].FirstAlbum+"</h2>")

								
	fmt.Fprintln(w, "<h1>"+"Concerts"+"</h1>")

							//Relations//


			
	for _, m := range bandInfo.Relations {
		if numArtist == m.ID  {
		for k, v := range m.DatesLocations {


			fmt.Fprint(w,"<h5>",strings.ToUpper(k),"</h5>")
			fmt.Fprintln(w,"<h2>",strings.Join(v,"<br>"+ " "),"</h2>")
			
		}
	}
	    
}

	tplArtists.ExecuteTemplate(w, "artist.html", bandInfo)

}

func main() {

	tplArtists = template.Must(template.ParseFiles("templates/artist.html"))

	http.HandleFunc("/artist", getArtists)
	http.HandleFunc("/", getData)

	http.HandleFunc("/index.html", getData)
	fmt.Println("Starting the server on :9090...")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	log.Fatal(http.ListenAndServe(":9090", nil))

}



