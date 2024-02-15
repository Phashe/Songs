package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Song struct to represent the data format
type Song struct {
	ID          string `json:"id,omitempty"`
	Artist      string `json:"artist,omitempty"`
	SongName    string `json:"songName,omitempty"`
	SpotifyLink string `json:"spotifyLink,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Description string `json:"description,omitempty"`
}

var songs []Song

// GetSongs returns the list of all songs
func GetSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}

// GetSong returns a specific song by ID
func GetSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range songs {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Song not found")
}

// CreateSong adds a new song to the list
func CreateSong(w http.ResponseWriter, r *http.Request) {
	var song Song
	_ = json.NewDecoder(r.Body).Decode(&song)
	song.ID = fmt.Sprintf("%d", len(songs)+1)
	songs = append(songs, song)
	w.WriteHeader(http.StatusCreated)
}

// UpdateSong updates an existing song by ID
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range songs {
		if item.ID == params["id"] {
			var updatedSong Song
			_ = json.NewDecoder(r.Body).Decode(&updatedSong)
			updatedSong.ID = item.ID
			songs[index] = updatedSong
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Song not found")
}

// DeleteSong removes a song by ID
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range songs {
		if item.ID == params["id"] {
			songs = append(songs[:index], songs[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Song not found")
}

func main() {
	router := mux.NewRouter()

	// Sample data for demonstration purposes
	songs = append(songs, Song{ID: "1", Artist: "Sonic Eigma", SongName: "Our Kingdom", SpotifyLink: "https://open.spotify.com/track/7bV0OSwRSf3BfAKB9pIT7t?si=15e2b8ac41b74bd8", Genre: "Melodic House/Techno", Description: "Our Kingdom is a melodic house song. Its about staying true to your kingdom and fighting to protect it"})
	songs = append(songs, Song{ID: "2", Artist: "Sonic Eigma", SongName: "Only now", SpotifyLink: "https://open.spotify.com/track/674DR9xXGpqsER2PcOIsoU?si=c5111ce5f7934def", Genre: "Melodic House/Techno", Description: "Only Now is a dance electronic jam."})

	// Define API endpoints
	router.HandleFunc("/songs", GetSongs).Methods("GET")
	router.HandleFunc("/songs/{id}", GetSong).Methods("GET")
	router.HandleFunc("/songs", CreateSong).Methods("POST")
	router.HandleFunc("/songs/{id}", UpdateSong).Methods("PUT")
	router.HandleFunc("/songs/{id}", DeleteSong).Methods("DELETE")

	fmt.Println("Sonic Enigma is a South African electronic music producer.")
	fmt.Println("The Sonic Enigma server provides endpoints for managing a repository containing Sonic Enigma songs")
	fmt.Println("===================================================")
	fmt.Println("The Sonic Enigma server is running on :8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}
