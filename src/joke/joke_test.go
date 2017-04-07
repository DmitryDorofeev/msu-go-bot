package joke

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJoke(t *testing.T) {

	resp := `{"type":"success","value":{"id":0,"joke":"The best joke"}}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	}))

	joke := GetJoke(ts.URL)

	t.Log("Joke: ", joke)

	if joke != "The best joke" {
		t.Error("Joke is not the best :(")
	}

	resp = "Bad joke"

	joke = GetJoke(ts.URL)

	if joke != "Joke error" {
		t.Error("Joke must be bad")
	}
}


