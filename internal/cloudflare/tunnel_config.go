package cloudflare

import (
	"fmt"
)

func (t *CloudflareTunnel) GetTunnelConfig() (*GetTunneConfigResponse, error) {
	url := fmt.Sprintf("%s/%s/configurations", tunnelApiBaseUrl, t.ID)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		return nil, err
	}

	config := parseResponse[GetTunneConfigResponse](res)

	return &config, nil
}

func (t *CloudflareTunnel) PutTunnelConfig(tunnelConfig TunnelConfig) (*PutTunneConfigRequest, error) {
	url := fmt.Sprintf("%s/%s/configurations", tunnelApiBaseUrl, t.ID)

	body := PutTunneConfigRequest{
		Config: tunnelConfig,
	}

	res, err := makeRequest("PUT", url, body, nil)

	if err != nil {
		return nil, err
	}

	config := parseResponse[PutTunneConfigRequest](res)

	return &config, nil
}
