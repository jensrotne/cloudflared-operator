package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func GetTunnel(id string) (*cloudflare.Tunnel, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	tunnel, err := api.GetTunnel(context.Background(), rc, id)

	if err != nil {
		return nil, err
	}

	return &tunnel, nil
}

func ListTunnels(params cloudflare.TunnelListParams) (*[]cloudflare.Tunnel, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	tunnels, resultInfo, err := api.ListTunnels(context.Background(), rc, params)

	if err != nil {
		return nil, err
	}

	if resultInfo.Count == 0 {
		return nil, nil
	}

	return &tunnels, nil
}

func CreateTunnel(name string, configSrc string, secret *string) (*cloudflare.Tunnel, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	params := cloudflare.TunnelCreateParams{
		Name:      name,
		ConfigSrc: configSrc,
		Secret:    *secret,
	}

	tunnel, err := api.CreateTunnel(context.Background(), rc, params)

	if err != nil {
		return nil, err
	}

	return &tunnel, nil
}

func DeleteTunnel(id string) error {
	api := getCloudflareAPI()

	rc := getAccountRC()

	err := api.DeleteTunnel(context.Background(), rc, id)

	if err != nil {
		return err
	}

	return nil
}

func GetTunnelToken(id string) (*string, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	token, err := api.GetTunnelToken(context.Background(), rc, id)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
