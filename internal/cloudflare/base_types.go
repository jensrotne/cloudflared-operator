package cloudflare

type BaseResponse struct {
	Errors   []Error   `json:"errors"`
	Messages []Message `json:"messages"`
	Success  bool      `json:"success"`
}

type BaseListResponse struct {
	BaseResponse
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"Count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
