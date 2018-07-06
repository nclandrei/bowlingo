package bowlingo

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScoreHandler(t *testing.T) {
	srv := NewServer()
	srv.Routes()
	frames := newFrameSlice(0, 10)
	frameRequest := struct {
		Frames []Frame `json:"frames"`
	}{
		Frames: frames,
	}
	reqBody, _ := json.Marshal(&frameRequest)
	req, _ := http.NewRequest("POST", "localhost:8000/api/score", bytes.NewReader(reqBody))
	rec := httptest.NewRecorder()
	srv.ScoreHandler().ServeHTTP(rec, req)
	resp := rec.Result()
	assert.Equal(t, 200, resp.StatusCode, "status code should have been 200")
}
