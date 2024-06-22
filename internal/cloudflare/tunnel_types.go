package cloudflare

type CloudflareTunnel struct {
	AccountTag  string `json:"account_tag"`
	Connections []struct {
		ClientID           string `json:"client_id"`
		ClientVersion      string `json:"client_version"`
		ColoName           string `json:"colo_name"`
		ID                 string `json:"id"`
		IsPendingReconnect bool   `json:"is_pending_reconnect"`
		OpenedAt           string `json:"opened_at"`
		OriginIP           string `json:"origin_ip"`
		UUID               string `json:"uuid"`
	} `json:"connections"`
	ConnsActiveAt   string      `json:"conns_active_at"`
	ConnsInactiveAt string      `json:"conns_inactive_at"`
	CreatedAt       string      `json:"created_at"`
	DeletedAt       string      `json:"deleted_at"`
	ID              string      `json:"id"`
	Metadata        interface{} `json:"metadata"`
	Name            string      `json:"name"`
	RemoteConfig    bool        `json:"remote_config"`
	Status          string      `json:"status"`
	TunType         string      `json:"tun_type"`
}

type GetTunnelTokenResponse struct {
	BaseResponse
	Result string `json:"result"`
}

type GetTunnelResponse struct {
	BaseResponse
	Result CloudflareTunnel `json:"result"`
}

type ListTunnelsResponse struct {
	BaseListResponse
	Result []CloudflareTunnel `json:"result"`
}

type CreateTunnelRequest struct {
	ConfigSrc    string  `json:"config_src"`
	Name         string  `json:"name"`
	TunnelSecret *string `json:"tunnel_secret,omitempty"`
}

type CreateTunnelResponse struct {
	BaseResponse
	Result CloudflareTunnel `json:"result"`
}

type DeleteTunnelResponse struct {
	BaseResponse
	Result CloudflareTunnel `json:"result"`
}
