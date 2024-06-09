package main

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/cloudflare"
)

func main() {
	newTunnel := cloudflare.CreateTunnel("test-api-tunnel", "cloudflare", "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg=")

	fmt.Println(newTunnel)

	deleteResponse := cloudflare.DeleteTunnel(newTunnel.Result.ID)

	fmt.Println(deleteResponse)

}
