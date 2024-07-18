package cloudflare

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/jensrotne/cloudflared-operator/internal/config"
)

var accountId = config.Get("CF_ACCOUNT_ID")
var apiToken = config.Get("CF_API_TOKEN")
var zoneID = config.Get("CF_ZONE_ID")
var tunnelApiBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel", accountId)
var rulesetApiBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/rulesets", zoneID)
var accessApplicationsApiBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/access/apps", accountId)

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
