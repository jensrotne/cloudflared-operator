package cloudflare

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/config"
)

func GetTunnelConfig(id string) GetTunneConfigResponse {
	accountId := config.Get("CF_ACCOUNT_ID")

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel/%s/configurations", accountId, id)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		panic(err)
	}

	config := parseResponse[GetTunneConfigResponse](res)

	return config
}

func PutTunnelConfig(id string, tunnelConfig TunnelConfig) PutTunneConfigRequest {
	accountId := config.Get("CF_ACCOUNT_ID")

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel/%s/configurations", accountId, id)

	body := PutTunneConfigRequest{
		Config: tunnelConfig,
	}

	res, err := makeRequest("PUT", url, body, nil)

	if err != nil {
		panic(err)
	}

	config := parseResponse[PutTunneConfigRequest](res)

	return config
}
