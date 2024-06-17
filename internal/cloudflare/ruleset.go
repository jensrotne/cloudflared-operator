package cloudflare

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/config"
)

var rulesetApiBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/rulesets", config.Get("CF_ZONE_ID"))

func ListZoneRulesets() (*ListZoneRulesetsResponse, error) {
	res, err := makeRequest("GET", rulesetApiBaseUrl, nil, nil)

	if err != nil {
		return nil, err
	}

	rulesets := parseResponse[ListZoneRulesetsResponse](res)

	return &rulesets, nil
}
