package cloudflare

var tokenApiBaseUrl = "https://api.cloudflare.com/client/v4/user/tokens"

func VerifyToken() (*GetVerifyTokenResponse, error) {
	res, err := makeRequest("GET", tokenApiBaseUrl, nil, nil)

	if err != nil {
		return nil, err
	}

	tokenResponse, err := parseResponse[GetVerifyTokenResponse](res)

	if err != nil {
		return nil, err
	}

	return tokenResponse, nil
}
