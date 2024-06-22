package cloudflare

type ListZoneRulesetsResponse struct {
	BaseListResponse
	Result []Ruleset `json:"result"`
}

type Ruleset struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	LastUpdated string `json:"last_updated"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Kind        string `json:"kind"`
	Phase       string `json:"phase"`
}
