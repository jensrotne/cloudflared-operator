package cloudflare

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/config"
)

func GetTunnel(id string) GetTunnelResponse {
	accountId := config.Get("CF_ACCOUNT_ID")

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel/%s", accountId, id)

	res, err := makeRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	tunnel := parseResponse[GetTunnelResponse](res)

	return tunnel
}

func ListTunnels() ListTunnelsResponse {
	accountId := config.Get("CF_ACCOUNT_ID")

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel", accountId)

	res, err := makeRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	tunnels := parseResponse[ListTunnelsResponse](res)

	return tunnels
}

func CreateTunnel(name string, configSrc string, secret string) CreateTunnelResponse {
	accountId := config.Get("CF_ACCOUNT_ID")

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel", accountId)

	body := CreateTunnelRequest{
		Name:         name,
		ConfigSrc:    configSrc,
		TunnelSecret: secret,
	}

	res, err := makeRequest("POST", url, body)

	if err != nil {
		panic(err)
	}

	tunnel := parseResponse[CreateTunnelResponse](res)

	return tunnel
}

func DeleteTunnel(id string) DeleteTunnelResponse {
	accountId := config.Get("CF_ACCOUNT_ID")

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel/%s", accountId, id)

	res, err := makeRequest("DELETE", url, nil)

	if err != nil {
		panic(err)
	}

	tunnel := parseResponse[DeleteTunnelResponse](res)

	return tunnel
}
