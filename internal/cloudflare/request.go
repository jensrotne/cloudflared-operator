package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jensrotne/cloudflared-operator/internal/config"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}

func makeRequest(method string, url string, body interface{}, queryParams map[string]string) (*http.Response, error) {
	apiToken := config.Get("CF_API_TOKEN")

	var req *http.Request
	var err error

	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			panic(err)
		}

		buffer := bytes.NewBuffer(b)

		req, err = http.NewRequest(method, url, buffer)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		panic(err)
	}

	if queryParams != nil {
		q := req.URL.Query()
		for key, value := range queryParams {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	res, err := client.Do(req)

	return res, err
}

func parseResponse[T interface{}](res *http.Response) T {
	defer res.Body.Close()

	var response T

	err := json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		panic(err)
	}

	return response
}
