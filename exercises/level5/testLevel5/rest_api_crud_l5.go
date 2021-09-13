package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "White Pony", Artist: "Deftones", Price: 65.25},
	{ID: "2", Title: "Madball", Artist: "Madball", Price: 70.02},
	{ID: "3", Title: "Destroy Everything", Artist: "Hatebreed", Price: 36.21},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ALBUMS API")
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(albums)
}

func getAlbumById(w http.ResponseWriter, r *http.Request) {
	albumID := mux.Vars(r)["id"]

	for _, album := range albums {
		if album.ID == albumID {
			json.NewEncoder(w).Encode(album)
		}
	}
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	var newAlbum album
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Album ID, Title, Artist and Price in order to create")
	}

	json.Unmarshal(reqBody, &newAlbum)
	albums = append(albums, newAlbum)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newAlbum)
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := mux.Vars(r)["id"]
	var updatedAlbum album

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Album Title, Artist and Price in order to update")
	}
	json.Unmarshal(reqBody, &updatedAlbum)

	for i, actualAlbum := range albums {
		if actualAlbum.ID == albumID {
			actualAlbum.Title = updatedAlbum.Title
			actualAlbum.Artist = updatedAlbum.Artist
			actualAlbum.Price = updatedAlbum.Price
			albums[i] = actualAlbum
			json.NewEncoder(w).Encode(actualAlbum)
		}
	}
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	albumID := mux.Vars(r)["id"]

	for i, album := range albums {
		if album.ID == albumID {
			albums = append(albums[:i], albums[i+1:]...)
			fmt.Fprintf(w, "The album with ID %v has been deleted successfully", albumID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/album", createAlbum).Methods("POST")
	router.HandleFunc("/album", getAlbums).Methods("GET")
	router.HandleFunc("/album/{id}", getAlbumById).Methods("GET")
	router.HandleFunc("/album/{id}", updateAlbum).Methods("PATCH")
	router.HandleFunc("/album/{id}", deleteAlbum).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
