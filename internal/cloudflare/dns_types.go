package cloudflare

type DNSRecord struct {
	Content    string        `json:"content"`
	Name       string        `json:"name"`
	Proxied    bool          `json:"proxied"`
	Type       string        `json:"type"`
	Comment    string        `json:"comment"`
	CreatedOn  string        `json:"created_on"`
	ID         string        `json:"id"`
	Locked     bool          `json:"locked"`
	Meta       DNSRecordMeta `json:"meta"`
	ModifiedOn string        `json:"modified_on"`
	Proxiable  bool          `json:"proxiable"`
	Tags       []string      `json:"tags"`
	TTL        int           `json:"ttl"`
	ZoneID     string        `json:"zone_id"`
	ZoneName   string        `json:"zone_name"`
}

type DNSRecordMeta struct {
	AutoAdded bool   `json:"auto_added"`
	Source    string `json:"source"`
}

type ListDNSRecordsResponse struct {
	BaseListResponse
	Result []DNSRecord `json:"result"`
}

type CreateDNSCNAMERecordRequest struct {
	Content string    `json:"content"`
	Name    string    `json:"name"`
	Proxied *bool     `json:"proxied,omitempty"`
	Type    string    `json:"type"`
	Comment *string   `json:"comment,omitempty"`
	ID      string    `json:"id"`
	Tags    *[]string `json:"tags,omitempty"`
	TTL     *int      `json:"ttl,omitempty"`
	ZoneID  *string   `json:"zone_id"`
}

type CreateDNSCNAMERecordResponse struct {
	BaseResponse
	Result *DNSRecord `json:"result"`
}

type DeleteDNSRecordResponse struct {
	BaseResponse
	Result struct {
		ID string `json:"id"`
	} `json:"result"`
}
