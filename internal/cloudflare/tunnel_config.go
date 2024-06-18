package cloudflare

import (
	"fmt"
)

func (t *CloudflareTunnel) GetTunnelConfig() (*GetTunnelConfigResponse, error) {
	url := fmt.Sprintf("%s/%s/configurations", tunnelApiBaseUrl, t.ID)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		return nil, err
	}

	config, err := parseResponse[GetTunnelConfigResponse](res)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (t *CloudflareTunnel) PutTunnelConfig(tunnelConfig TunnelConfig) (*PutTunnelConfigResponse, error) {
	url := fmt.Sprintf("%s/%s/configurations", tunnelApiBaseUrl, t.ID)

	body := PutTunneConfigRequest{
		Config: tunnelConfig,
	}

	res, err := makeRequest("PUT", url, body, nil)

	if err != nil {
		return nil, err
	}

	config, err := parseResponse[PutTunnelConfigResponse](res)

	if err != nil {
		return nil, err
	}

	return config, nil
}
