package cloudflare

import (
	"encoding/json"
)

func VerifyToken() GetVerifyTokenResponse {
	url := "https://api.cloudflare.com/client/v4/user/tokens/verify"

	res, err := makeRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var response GetVerifyTokenResponse

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		panic(err)
	}

	return response
}
