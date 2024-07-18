package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func GetDNSRecordIfExists(name string) (*cloudflare.DNSRecord, error) {
	api := getCloudflareAPI()

	ctx := context.Background()

	rc := getZoneRC()

	records, resultInfo, err := api.ListDNSRecords(ctx, rc, cloudflare.ListDNSRecordsParams{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	if resultInfo.Count == 0 {
		return nil, nil
	}

	record := records[0]

	return &record, nil
}

func CreateDNSCNAMERecord(name string, content string) (*cloudflare.DNSRecord, error) {
	api := getCloudflareAPI()
	proxied := true

	params := cloudflare.CreateDNSRecordParams{
		Type:    "CNAME",
		Name:    name,
		Content: content,
		Proxied: &proxied,
	}

	rc := getZoneRC()

	record, err := api.CreateDNSRecord(context.Background(), rc, params)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func DeleteDNSRecord(id string) error {
	api := getCloudflareAPI()

	rc := getZoneRC()

	err := api.DeleteDNSRecord(context.Background(), rc, id)

	if err != nil {
		return err
	}

	return nil
}
