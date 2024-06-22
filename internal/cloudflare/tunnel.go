package cloudflare

import (
	"fmt"
)

func GetTunnel(id string) (*GetTunnelResponse, error) {
	url := fmt.Sprintf("%s/%s", tunnelApiBaseUrl, id)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		return nil, err
	}

	tunnel, err := parseResponse[GetTunnelResponse](res)

	if err != nil {
		return nil, err
	}

	return tunnel, nil
}

func ListTunnels(options map[string]string) (*ListTunnelsResponse, error) {
	res, err := makeRequest("GET", tunnelApiBaseUrl, nil, options)

	if err != nil {
		return nil, err
	}

	tunnels, err := parseResponse[ListTunnelsResponse](res)

	if err != nil {
		return nil, err
	}

	return tunnels, nil
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

	tunnel, err := parseResponse[CreateTunnelResponse](res)

	if err != nil {
		return nil, err
	}

	return tunnel, nil
}

func DeleteTunnel(id string) (*DeleteTunnelResponse, error) {
	url := fmt.Sprintf("%s/%s", tunnelApiBaseUrl, id)

	res, err := makeRequest("DELETE", url, nil, nil)

	if err != nil {
		return nil, err
	}

	tunnel, err := parseResponse[DeleteTunnelResponse](res)

	if err != nil {
		return nil, err
	}

	return tunnel, nil
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

	token, err := parseResponse[GetTunnelTokenResponse](res)

	if err != nil {
		return nil, err
	}

	return token, nil
}
