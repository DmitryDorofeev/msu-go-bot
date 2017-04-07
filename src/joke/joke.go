package joke

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

func GetJoke(url string) string {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return "jokes API not responding"
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	joke := JokeResponse{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "Joke error"
	}

	return joke.Value.Joke
}
