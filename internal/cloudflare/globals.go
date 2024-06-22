package cloudflare

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/config"
)

var accountId = config.Get("CF_ACCOUNT_ID")
var zoneID = config.Get("CF_ZONE_ID")
var tunnelApiBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel", accountId)
var dnsBaseApiUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", zoneID)
var rulesetApiBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/rulesets", zoneID)
