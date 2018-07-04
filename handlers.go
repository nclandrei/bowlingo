package bowlingo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Scores defines the handler needed to parse, calculate and return
// the current score of the user.
func Scores() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body.", http.StatusBadRequest)
			return
		}
		var frameRequest struct {
			Frames []Frame `json:"frames"`
		}
		err = json.Unmarshal(bodyBytes, &frameRequest)
		if err != nil {
			http.Error(w, "Could not unmarshal body to JSON.", http.StatusInternalServerError)
			return
		}

	})
}
