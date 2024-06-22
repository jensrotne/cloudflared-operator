package cloudflare

type GetVerifyTokenResponse struct {
	BaseResponse
	Result struct {
		ExpiresOn string `json:"expires_on"`
		ID        string `json:"id"`
		NotBefore string `json:"not_before"`
		Status    string `json:"status"`
	} `json:"result"`
}
