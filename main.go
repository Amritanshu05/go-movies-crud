// Package main defines karta hai entry point Go movies CRUD API application ke liye
package main

// Import karo sabhi zaroori packages application ke liye
import (
	"encoding/json" // JSON encoding aur decoding operations ke liye
	"fmt"           // Formatted I/O operations ke liye (console mein print karne ke liye)
	"log"           // Errors aur fatal messages log karne ke liye
	"math/rand"     // Random numbers generate karne ke liye (movie IDs ke liye use hoga)
	"net/http"      // HTTP server functionality aur request/response handling ke liye
	"strconv"       // String conversions ke liye (integers ko strings mein convert karna)

	"github.com/gorilla/mux" // Third-party router hai HTTP request routing aur URL parameter extraction ke liye
)

// Movie struct represent karta hai ek movie entity with sabhi properties
// JSON tags batate hain ki struct fields kaise serialize/deserialize honge
type Movie struct {
	ID       string    `json:"id"`       // Movie ka unique identifier
	Isbn     string    `json:"isbn"`     // Movie ka ISBN number (book/media identifier)
	Title    string    `json:"title"`    // Movie ka title/naam
	Director *Director `json:"director"` // Director struct ka pointer jo director ki information contain karta hai
}

// Director struct represent karta hai movie director with unke naam ki information
type Director struct {
	Firstname string `json:"firstname"` // Director ka first name
	Lastname  string `json:"lastname"`  // Director ka last name
}

// movies ek global slice hai jo hamara in-memory database serve karta hai
// Yeh application session ke liye sabhi movie records store karta hai
var movies []Movie

// getMovies handle karta hai GET requests sabhi movies retrieve karne ke liye
// Yeh complete list of movies JSON format mein return karta hai
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Response content type set karo JSON pe taaki client ko pata chale data format ke baare mein
	w.Header().Set("Content-Type", "application/json")
	// Poore movies slice ko JSON mein encode karo aur response mein write karo
	json.NewEncoder(w).Encode(movies)
}

// deleteMovie handle karta hai DELETE requests specific movie ko ID se remove karne ke liye
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Response content type set karo JSON pe
	w.Header().Set("Content-Type", "application/json")
	// URL parameters extract karo (specifically "id" parameter) request se
	params := mux.Vars(r)
	// Movies slice mein iterate karo matching ID wali movie find karne ke liye
	for index, item := range movies {
		// Check karo ki current movie ka ID requested ID se match karta hai
		if item.ID == params["id"] {
			// Movie ko slice se remove karo target index ke before aur after elements combine karke
			// movies[:index] target se pehle ke sabhi elements deta hai
			// movies[index+1:] target ke baad ke sabhi elements deta hai
			movies = append(movies[:index], movies[index+1:]...)
			// Loop se exit karo jab movie mil jaaye aur delete ho jaaye
			break
		}
	}
	// Updated movies list JSON response mein return karo
	json.NewEncoder(w).Encode(movies)
}

// getMovie handle karta hai GET requests ek single movie retrieve karne ke liye uske ID se
func getMovie(w http.ResponseWriter, r *http.Request) {
	// Response content type set karo JSON pe
	w.Header().Set("Content-Type", "application/json")
	// URL parameters extract karo request se
	params := mux.Vars(r)

	// Movies slice mein iterate karo matching ID wali movie find karne ke liye
	for _, item := range movies {
		// Check karo ki current movie ka ID requested ID se match karta hai
		if item.ID == params["id"] {
			// Agar mil gayi, movie ko JSON mein encode karo aur response mein send karo
			json.NewEncoder(w).Encode(item)
			// Immediately return karo taaki function ka baaki part execute na ho
			return
		}
	}
	// Agar koi movie nahi mili given ID se, empty Movie struct return karo
	json.NewEncoder(w).Encode(&Movie{})
}

// createMovie handle karta hai POST requests nai movie add karne ke liye collection mein
func createMovie(w http.ResponseWriter, r *http.Request) {
	// Response content type set karo JSON pe
	w.Header().Set("Content-Type", "application/json")
	// Naya Movie variable banao incoming data hold karne ke liye
	var movie Movie
	// Request body se JSON decode karo movie struct mein
	// Underscore (_) error return value ko ignore karta hai (production ke liye recommended nahi)
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Nai movie ke liye random ID generate karo random integer ko string mein convert karke
	// rand.Intn(10000000) 0 se 9,999,999 ke beech mein random number generate karta hai
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	// Nai movie ko movies slice mein add karo
	movies = append(movies, movie)
	// Newly created movie ko JSON response mein return karo (generated ID ke saath)
	json.NewEncoder(w).Encode(movie)
}

// updateMovie handle karta hai PUT requests existing movie update karne ke liye ID se
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Response content type set karo JSON pe
	w.Header().Set("Content-Type", "application/json")
	// URL parameters extract karo request se
	params := mux.Vars(r)
	// Movies slice mein iterate karo update karne wali movie find karne ke liye
	for index, item := range movies {
		// Check karo ki current movie ka ID requested ID se match karta hai
		if item.ID == params["id"] {
			// Existing movie ko slice se remove karo (same deletion logic jaise deleteMovie mein)
			movies = append(movies[:index], movies[index+1:]...)
			// Naya Movie variable banao updated data hold karne ke liye
			var movie Movie
			// Request body se JSON decode karo movie struct mein
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// ID set karo URL parameter se (original ID preserve karo)
			movie.ID = params["id"]
			// Updated movie ko movies slice mein wapas add karo
			movies = append(movies, movie)
			// Updated movie ko JSON response mein return karo
			json.NewEncoder(w).Encode(movie)
			// Successful update ke baad function se exit karo
			return
		}
	}
	// Agar koi movie nahi mili given ID se, empty Movie struct return karo
	json.NewEncoder(w).Encode(&Movie{})
}

// main function application ka entry point hai
func main() {
	// Gorilla Mux package use karke naya router instance banao
	// Yeh router HTTP request routing handle karega URL patterns aur methods ke base pe
	r := mux.NewRouter()

	// Movies slice ko kuch sample data se initialize karo testing ke liye
	// Pehli sample movie create aur add karo ID "1" ke saath
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	// Doosri sample movie create aur add karo ID "2" ke saath
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// API routes define karo aur unke corresponding handler functions
	// GET /movies - sabhi movies retrieve karo
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// GET /movies/{id} - specific movie retrieve karo ID se
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// POST /movies - nai movie create karo
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// PUT /movies/{id} - existing movie update karo ID se
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// DELETE /movies/{id} - movie delete karo ID se
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Console mein message print karo ki server start ho raha hai
	fmt.Println("Starting server at port 8000")
	// HTTP server start karo port 8000 pe configured router ke saath
	// log.Fatal koi bhi error log karega aur program terminate kar dega agar server start nahi hua
	log.Fatal(http.ListenAndServe(":8000", r))
}
