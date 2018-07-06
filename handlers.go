package bowlingo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Scores defines the handler needed to parse, calculate and return
// the current score of the user.
func Scores() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not read request body: %v", err), http.StatusBadRequest)
			return
		}
		var frameRequest struct {
			Frames []Frame `json:"frames"`
		}
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
		resp, err := json.Marshal(struct {
			score int
		}{
			score: *score,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not marshal score into a JSON object: %v", err), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(resp)
		if err != nil {
			log.Printf("could not write score to response: %v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
