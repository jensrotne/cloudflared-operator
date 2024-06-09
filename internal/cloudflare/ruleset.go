package cloudflare

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/config"
)

func ListZoneRulesets(zoneId string) ListZoneRulesetsResponse {
	if zoneId == "" {
		zoneId = config.Get("CF_ZONE_ID")
	}

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/rulesets", zoneId)

	res, err := makeRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	rulesets := parseResponse[ListZoneRulesetsResponse](res)

	return rulesets
}
