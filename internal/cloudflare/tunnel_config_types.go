package cloudflare

type TunnelConfig struct {
	Ingress       []TunnelConfigIngress      `json:"ingress"`
	OriginRequest *TunnelConfigOriginRequest `json:"originRequest,omitempty"`
	WarpRouting   *TunnelConfigWarpRouting   `json:"warp-routing,omitempty"`
}

type TunnelConfigWarpRouting struct {
	Enabled bool `json:"enabled"`
}

type TunnelConfigIngress struct {
	Hostname      *string                    `json:"hostname,omitempty"`
	OriginRequest *TunnelConfigOriginRequest `json:"originRequest,omitempty"`
	Path          *string                    `json:"path,omitempty"`
	Service       string                     `json:"service"`
}

type TunnelConfigOriginRequest struct {
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

type GetTunnelConfigResponse struct {
	BaseResponse
	Result struct {
		TunnelID  string        `json:"tunnel_id"`
		Version   int           `json:"version"`
		Config    *TunnelConfig `json:"config"`
		Source    string        `json:"source"`
		CreatedAt string        `json:"created_at"`
	} `json:"result"`
}

type PutTunnelConfigResponse struct {
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
