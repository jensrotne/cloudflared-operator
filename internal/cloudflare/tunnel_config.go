package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func GetTunnelConfig(id string) (*cloudflare.TunnelConfiguration, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	res, err := api.GetTunnelConfiguration(context.Background(), rc, id)

	if err != nil {
		return nil, err
	}

	return &res.Config, nil
}

func UpdateTunnelConfig(id string, config cloudflare.TunnelConfiguration) (*cloudflare.TunnelConfiguration, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	params := cloudflare.TunnelConfigurationParams{
		TunnelID: id,
		Config:   config,
	}

	res, err := api.UpdateTunnelConfiguration(context.Background(), rc, params)

	if err != nil {
		return nil, err
	}

	return &res.Config, nil
}
