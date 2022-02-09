package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	//"text/template"
)

var tpl *template.Template

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

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tpl.ExecuteTemplate(w, "index.html", nil)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
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
	/*fmt.Println(data)
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		//log.Print(err)
	}
	fmt.Fprint(w, data)
	t.Execute(w, nil)
	//fmt.Fprint(w, "<h1>"+dataArt.artist[1]+"</h2>"+"\\n"+"<h1>"+dataArt.artist[1].Members[1]+"</h2>")
	 tpl.ExecuteTemplate(w, "index.html", data)

}*/
}

func UnmarshalLocation() {
	resp, err := http.Get(Locations)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var data location
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

	var data dates
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

}

func UnmarshalRel() {
	resp, err := http.Get(Relations)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	var data relations
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

}

func getData(w http.ResponseWriter, r *http.Request) {
	UnmarshalArtist()
	UnmarshalDates()
	UnmarshalLocation()
	UnmarshalRel()

	//tpl.ExecuteTemplate(w, "index.html",)


}

func main() {

   
   http.HandleFunc("/index.html", getData)
	//fmt.Println("Starting the server on :9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
