package omdb_api

import (
	"encoding/json"
	"net/http"
	"omdb/internal/types"
	"strconv"
)

const omdbAPIKey = "faf7e5bb&i=tt4853102"
const omdbAPIBaseURL = "http://www.omdbapi.com/"

func GetMovieByID(id string) (*types.GetMovieByIDResponse, error) {
	url := omdbAPIBaseURL + "?apikey=" + omdbAPIKey + "&i=" + id
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result types.GetMovieByIDResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func SearchMovies(query string, page uint64) (*types.SearchMoviesResponse, error) {
	url := omdbAPIBaseURL + "?apikey=" + omdbAPIKey + "&s=" + query + "&page=" + strconv.FormatUint(page, 10)
	println(url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result types.SearchMoviesResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
