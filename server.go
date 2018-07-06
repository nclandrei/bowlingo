package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Server defines the server which we'll for our bowling game. I followed best
// practices for structuring the code and the server and its components (for more
// details, please see here: https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)
type Server struct {
	router *http.ServeMux
}

// NewServer returns a new instance of the server.
func NewServer() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

// ScoreHandler defines the handler needed to parse, calculate and return
// the current score of the user.
func (s *Server) ScoreHandler() http.Handler {
	type Request struct {
		Frames []Frame `json:"frames"`
	}
	type Response struct {
		Score int `json:"score"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not read request body: %v", err), http.StatusBadRequest)
			return
		}
		var frameRequest Request
		err = json.Unmarshal(bodyBytes, &frameRequest)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not unmarshal body to JSON: %v", err), http.StatusInternalServerError)
			return
		}
		score, err := Score(frameRequest.Frames)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not calculate total score for the frames: %v", err), http.StatusInternalServerError)
			return
		}
		resp := Response{
			Score: score,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&resp)
	})
}
