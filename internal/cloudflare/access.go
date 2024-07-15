package cloudflare

func ListAccessApplications() ([]CloudflareAccessApplication, error) {
	res, err := makeRequest("GET", accessApplicationsApiBaseUrl, nil, nil)

	if err != nil {
		return nil, err
	}

	apps, err := parseResponse[ListAccessApplicationsResponse](res)

	if err != nil {
		return nil, err
	}

	return apps.Result, nil
}
