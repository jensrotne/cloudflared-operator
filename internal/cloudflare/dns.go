package cloudflare

import "fmt"

func DNSRecordExists(name string) (bool, error) {
	url := fmt.Sprintf("%s?name=%s", dnsBaseApiUrl, name)

	res, err := makeRequest("GET", url, nil, nil)

	if err != nil {
		return false, err
	}

	records, err := parseResponse[ListDNSRecordsResponse](res)

	if err != nil {
		return false, err
	}

	return len(records.Result) > 0, nil
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
