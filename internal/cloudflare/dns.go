package cloudflare

import "fmt"

func GetDNSRecordIfExists(name string) (*DNSRecord, error) {
	url := fmt.Sprintf("%s?name=%s", dnsBaseApiUrl, name)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		return nil, err
	}

	records, err := parseResponse[ListDNSRecordsResponse](res)

	if err != nil {
		return nil, err
	}

	if len(records.Result) == 0 {
		return nil, nil
	}

	record := records.Result[0]

	return &record, nil
}

func CreateDNSCNAMERecord(name string, content string) (*DNSRecord, error) {
	proxied := true

	body := CreateDNSCNAMERecordRequest{
		Name:    name,
		Content: content,
		Type:    "CNAME",
		Proxied: &proxied,
	}

	res, err := makeRequest("POST", dnsBaseApiUrl, body, nil)

	if err != nil {
		return nil, err
	}

	createResponse, err := parseResponse[CreateDNSCNAMERecordResponse](res)

	if err != nil {
		return nil, err
	}

	return createResponse.Result, nil
}

func DeleteDNSRecord(id string) error {
	url := fmt.Sprintf("%s/%s", dnsBaseApiUrl, id)

	res, err := makeRequest("DELETE", url, nil, nil)

	if err != nil {
		return err
	}

	deleteResponse, err := parseResponse[DeleteDNSRecordResponse](res)

	if err != nil {
		return err
	}

	if deleteResponse.Success {
		return nil
	}

	return fmt.Errorf("failed to delete DNS record with ID %s", id)
}
