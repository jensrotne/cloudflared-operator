package cloudflare

func ListZoneRulesets() (*ListZoneRulesetsResponse, error) {
	res, err := makeRequest("GET", rulesetApiBaseUrl, nil, nil)

	if err != nil {
		return nil, err
	}

	rulesets, err := parseResponse[ListZoneRulesetsResponse](res)

	if err != nil {
		return nil, err
	}

	return rulesets, nil
}
