package cloudflare

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
)

type IPRanges struct {
	IP IP `json:"ip"`
}

type IP struct {
	IP string `json:"ip"`
}

func GetAccessApplication(id string) (*cloudflare.AccessApplication, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	res, err := api.GetAccessApplication(context.Background(), rc, id)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func ListAccessApplications() (*[]cloudflare.AccessApplication, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	res, _, err := api.ListAccessApplications(context.Background(), rc, cloudflare.ListAccessApplicationsParams{})

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func CreateAccessApplication(name string, domain string, ip string) (*cloudflare.AccessApplication, error) {
	api := getCloudflareAPI()

	rc := getAccountRC()

	appLauncherVisible := false

	params := cloudflare.CreateAccessApplicationParams{
		Name:               name,
		Domain:             domain,
		Type:               "self_hosted",
		AppLauncherVisible: &appLauncherVisible,
	}

	res, err := api.CreateAccessApplication(context.Background(), rc, params)

	if err != nil {
		return nil, err
	}

	ipInclude := IPRanges{
		IP: IP{
			IP: fmt.Sprintf("%s/32", ip),
		},
	}

	policyParams := cloudflare.CreateAccessPolicyParams{
		Decision: "bypass",
		Include: []interface{}{
			ipInclude,
		},
		Name:          fmt.Sprintf("%s-bypass-policy", name),
		ApplicationID: res.ID,
	}

	_, err = api.CreateAccessPolicy(context.Background(), rc, policyParams)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func DeleteAccessApplication(id string) error {
	api := getCloudflareAPI()

	rc := getAccountRC()

	err := api.DeleteAccessApplication(context.Background(), rc, id)

	if err != nil {
		return err
	}

	return nil
}
