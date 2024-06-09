package main

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/cloudflare"
)

func main() {
	// newTunnel := cloudflare.CreateTunnel("test-api-tunnel", "cloudflare", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg=")

	// fmt.Println(newTunnel)

	// deleteResponse := cloudflare.DeleteTunnel(newTunnel.Result.ID)

	// fmt.Println(deleteResponse)

	// config := cloudflare.GetTunnelConfig("6941f253-6743-4679-8d31-81e06c6ca0e0")

	// fmt.Println(config)

	rulesets := cloudflare.ListZoneRulesets("")

	fmt.Println(rulesets)
}
