package cloudflare

import (
	"fmt"

	"github.com/jensrotne/cloudflared-operator/internal/config"
)

var accountId = config.Get("CF_ACCOUNT_ID")
var tunnelApiBaseUrl = fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/cfd_tunnel", accountId)

func GetTunnel(id string) (*GetTunnelResponse, error) {
	url := fmt.Sprintf("%s/%s", tunnelApiBaseUrl, id)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		return nil, err
	}

	tunnel := parseResponse[GetTunnelResponse](res)

	return &tunnel, nil
}

func ListTunnels(options map[string]string) (*ListTunnelsResponse, error) {
	res, err := makeRequest("GET", tunnelApiBaseUrl, nil, options)

	if err != nil {
		return nil, err
	}

	tunnels := parseResponse[ListTunnelsResponse](res)

	return &tunnels, nil
}

func CreateTunnel(name string, configSrc string, secret *string) (*CreateTunnelResponse, error) {
	body := CreateTunnelRequest{
		Name:         name,
		ConfigSrc:    configSrc,
		TunnelSecret: secret,
	}

	res, err := makeRequest("POST", tunnelApiBaseUrl, body, nil)

	if err != nil {
		return nil, err
	}

	tunnel := parseResponse[CreateTunnelResponse](res)

	return &tunnel, nil
}

func DeleteTunnel(id string) (*DeleteTunnelResponse, error) {
	url := fmt.Sprintf("%s/%s", tunnelApiBaseUrl, id)

	res, err := makeRequest("DELETE", url, nil, nil)

	if err != nil {
		return nil, err
	}

	tunnel := parseResponse[DeleteTunnelResponse](res)

	return &tunnel, nil
}

func (t *CloudflareTunnel) GetTunnelToken() (*GetTunnelTokenResponse, error) {
	return getTunnelToken(t.ID)
}

func getTunnelToken(id string) (*GetTunnelTokenResponse, error) {
	url := fmt.Sprintf("%s/%s/token", tunnelApiBaseUrl, id)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		return nil, err
	}

	token := parseResponse[GetTunnelTokenResponse](res)

	return &token, nil
}
