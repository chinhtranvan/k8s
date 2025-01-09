package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies(t *testing.T) {
	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", "/movies", nil)
	assert.NoError(t, err)

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies", GetMovies)
	router.ServeHTTP(rr, req)

	// Assert the response code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code 200")

	// Assert the response body
	var responseMovies []Movie
	err = json.NewDecoder(rr.Body).Decode(&responseMovies)
	assert.NoError(t, err, "Response body should be parsable JSON")
	assert.Equal(t, len(movies), len(responseMovies), "Response should match the number of movies in memory")
}

func TestCreateMovie(t *testing.T) {
	// Create a new movie
	newMovie := Movie{
		Isbn:     "12345",
		Title:    "Test Movie",
		Director: Director{Firstname: "Jane", Lastname: "Doe"},
	}
	body, _ := json.Marshal(newMovie)

	req, err := http.NewRequest("POST", "/movies", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies", CreateMovie).Methods("POST")
	router.ServeHTTP(rr, req)

	// Assert the response code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code 200")

	// Assert the movie is created and returned
	var createdMovie Movie
	err = json.NewDecoder(rr.Body).Decode(&createdMovie)
	assert.NoError(t, err)
	assert.Equal(t, newMovie.Title, createdMovie.Title, "Created movie title should match")
	assert.Equal(t, newMovie.Director.Firstname, createdMovie.Director.Firstname, "Created movie director's firstname should match")
}

func TestDeleteMovie(t *testing.T) {
	// Delete an existing movie
	req, err := http.NewRequest("DELETE", "/movies/1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")
	router.ServeHTTP(rr, req)

	// Assert the response code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code, "Expected response code 200")

	// Assert the movie is deleted
	for _, movie := range movies {
		assert.NotEqual(t, "1", movie.ID, "Movie with ID 1 should be deleted")
	}
}
