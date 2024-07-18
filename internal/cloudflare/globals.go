package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/jensrotne/cloudflared-operator/internal/config"
)

var accountId = config.Get("CF_ACCOUNT_ID")
var apiToken = config.Get("CF_API_TOKEN")
var zoneID = config.Get("CF_ZONE_ID")

func getCloudflareAPI() *cloudflare.API {
	api, err := cloudflare.NewWithAPIToken(apiToken)

	if err != nil {
		panic(err)
	}

	return api
}

func getZoneRC() *cloudflare.ResourceContainer {
	return &cloudflare.ResourceContainer{
		Identifier: zoneID,
	}
}

func getAccountRC() *cloudflare.ResourceContainer {
	return &cloudflare.ResourceContainer{
		Identifier: accountId,
		Level:      cloudflare.AccountRouteLevel,
	}
}
