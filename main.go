// my first server building
package main
import(
	"fmt"
	"log"
	"encoding/json"
	"strconv"
	"math/rand"
	"net/http"
	"github.com/JosephKingsley/movie"
)

type movie struct{
	ID string 'json:"id"'
	Isbn string 'json:"isbn"'
	Title string 'json:"title"'
	Director *Director 'json:"director"'
}
type Director struct{
	Firstname string 'json:"firstname"'
	Lastname string 'json:"lastname"'
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params != mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params = mux.Vars(r)
	for _, item := range movies(){
		if item.ID ==params["id"]{
			json.NewEncoder(w).Encode(item)

			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movies Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updatemovie(w http.ResponseWriter, r *http.Request){
	// set json content
	w.Header().Set("Content-Type", "application/json")
	// set params
	params := mmux.Vars(r)
	// loop and range over the movies
	// delete the movie with the id you've sent
	for index, item = range movies{
		if item.ID == params["id"]{
			movie = append(movies[:index], movies[index+1:]...)
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movie = append(movies, movie)
			// add a movie - the movie that we will send in the bosy of the postman
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	
}

func main(){
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "44591", Title: "Movie One", Director: &Director(Firstname:"Oseloka", Lastname: "Obi") })
	movies = append(movies, Movie{ID: "2", Isbn: "45532", Title: "Movie Two", Director: &Director(Firstname: "Peter", Lastname: "Gregory")})

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("delete")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}