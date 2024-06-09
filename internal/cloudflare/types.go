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

type GetTunnelResponse struct {
	BaseResponse
	Result CloudflareTunnel `json:"result"`
}

type ListTunnelsResponse struct {
	BaseListResponse
	Result []CloudflareTunnel `json:"result"`
}

type CreateTunnelRequest struct {
	ConfigSrc    string `json:"config_src"`
	Name         string `json:"name"`
	TunnelSecret string `json:"tunnel_secret"`
}

type CreateTunnelResponse struct {
	BaseResponse
	Result CloudflareTunnel `json:"result"`
}

type DeleteTunnelResponse struct {
	BaseResponse
	Result CloudflareTunnel `json:"result"`
}

type GetTunneConfigResponse struct {
	BaseResponse
	Result struct {
		TunnelID  string       `json:"tunnel_id"`
		Version   int          `json:"version"`
		Config    TunnelConfig `json:"config"`
		Source    string       `json:"source"`
		CreatedAt string       `json:"created_at"`
	} `json:"result"`
}

type PutTunneConfigRequest struct {
	Config TunnelConfig `json:"config"`
}

type ListZoneRulesetsResponse struct {
	BaseListResponse
	Result []Ruleset `json:"result"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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

type TunnelConfig struct {
	Ingress []struct {
		Hostname      string        `json:"hostname"`
		OriginRequest OriginRequest `json:"originRequest"`
		Path          string        `json:"path"`
		Service       string        `json:"service"`
	} `json:"ingress"`
	OriginRequest OriginRequest `json:"originRequest"`
	WarpRouting   struct {
		Enabled bool `json:"enabled"`
	} `json:"warpRouting"`
}

type OriginRequest struct {
	Access struct {
		AudTag   []string `json:"audTag"`
		Required bool     `json:"required"`
		TeamName string   `json:"teamName"`
	} `json:"access"`
	CaPool                 string `json:"caPool"`
	ConnectTimeout         int    `json:"connectTimeout"`
	DisableChunkedEncoding bool   `json:"disableChunkedEncoding"`
	Http2Origin            bool   `json:"http2Origin"`
	HttpHostHeader         string `json:"httpHostHeader"`
	KeepAliveConnections   int    `json:"keepAliveConnections"`
	KeepAliveTimeout       int    `json:"keepAliveTimeout"`
	NoHappyEyeballs        bool   `json:"noHappyEyeballs"`
	NoTLSVerify            bool   `json:"noTLSVerify"`
	OriginServerName       string `json:"originServerName"`
	ProxyType              string `json:"proxyType"`
	TCPKeepAlive           bool   `json:"tcpKeepAlive"`
	TLSTimeout             int    `json:"tlsTimeout"`
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
