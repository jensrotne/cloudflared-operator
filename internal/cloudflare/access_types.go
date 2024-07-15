package cloudflare

type ListAccessApplicationsResponse struct {
	Result []CloudflareAccessApplication `json:"result"`
	BaseListResponse
}

type CloudflareAccessApplication struct {
	Aud                        string      `json:"aud"`
	CreatedAd                  string      `json:"created_at"`
	ID                         string      `json:"id"`
	ScimConfig                 ScimConfig  `json:"scim_config"`
	UpdatedAt                  string      `json:"updated_at"`
	AllowAuthenticationViaWarp bool        `json:"allow_authentication_via_warp"`
	AllowedIdps                []string    `json:"allowed_idps"`
	AppLauncherVisible         bool        `json:"app_launcher_visible"`
	AutoRedirectToIdentity     bool        `json:"auto_redirect_to_identity"`
	CorsHeaders                CorsHeaders `json:"cors_headers"`
	CustomDenyMessage          string      `json:"custom_deny_message"`
	CustomDenyUrl              string      `json:"custom_deny_url"`
	CustomNonIdentityDenyUrl   string      `json:"custom_non_identity_deny_url"`
	CustomPages                []string    `json:"custom_pages"`
	Domain                     string      `json:"domain"`
	EnableBindingCookie        bool        `json:"enable_binding_cookie"`
	HttpOnlyCookieAttribute    bool        `json:"http_only_cookie_attribute"`
	LogoUrl                    string      `json:"logo_url"`
	Name                       string      `json:"name"`
	OptionsPreflightBypass     bool        `json:"options_preflight_bypass"`
	PathCookieAttribute        bool        `json:"path_cookie_attribute"`
	SameSiteCookieAttribute    string      `json:"same_site_cookie_attribute"`
	SelfHostedDomains          []string    `json:"self_hosted_domains"`
	ServiceAuth401Redirect     bool        `json:"service_auth_401_redirect"`
	SessionDuration            string      `json:"session_duration"`
	SkipInterstitial           bool        `json:"skip_interstitial"`
	Tags                       []string    `json:"tags"`
}

type ScimConfig struct {
	Authentication     Authentication `json:"authentication"`
	DeactivateOnDelete bool           `json:"deactivate_on_delete"`
	Enabled            bool           `json:"enabled"`
	IDPUid             string         `json:"idp_uid"`
	Mappings           []Mapping      `json:"mappings"`
	RemoteUri          string         `json:"remote_uri"`
}

type Authentication struct {
	Scheme string `json:"scheme"`
	HttpBasicScheme
	OAuthBearerTokenScheme
	OAuth2Scheme
}

type Mapping struct {
	Enabled          *bool      `json:"enabled,omitempty"`
	Filter           *string    `json:"filter,omitempty"`
	Operations       Operations `json:"operations"`
	Schema           string     `json:"schema"`
	TransformJsonata *string    `json:"transform_jsonata,omitempty"`
}

type Operations struct {
	Create *bool `json:"create,omitempty"`
	Delete *bool `json:"delete,omitempty"`
	Update *bool `json:"update,omitempty"`
}

type HttpBasicScheme struct {
	Password *string `json:"password,omitempty"`
	User     *string `json:"user,omitempty"`
}

type OAuthBearerTokenScheme struct {
	Token *string `json:"token,omitempty"`
}

type OAuth2Scheme struct {
	AuthorizationUrl *string  `json:"authorization_url,omitempty"`
	ClientID         *string  `json:"client_id,omitempty"`
	ClientSecret     *string  `json:"client_secret,omitempty"`
	Scopes           []string `json:"scopes,omitempty"`
	TokenUrl         *string  `json:"token_url,omitempty"`
}

type CorsHeaders struct {
	AllowAllHeaders  bool     `json:"allow_all_headers"`
	AllowAllMethods  bool     `json:"allow_all_methods"`
	AllowAllOrigins  bool     `json:"allow_all_origins"`
	AllowCredentials bool     `json:"allow_credentials"`
	AllowedHeaders   []string `json:"allowed_headers"`
	AllowedMethods   []string `json:"allowed_methods"`
	AllowedOrigins   []string `json:"allowed_origins"`
	MaxAge           int      `json:"max_age"`
}

type Policy struct {
	Precedence                   int                `json:"precedence"`
	ApprovalGroups               []ApprovalGroup    `json:"approval_groups"`
	ApprovalRequired             bool               `json:"approval_required"`
	CreatedAt                    string             `json:"created_at"`
	Decision                     string             `json:"decision"`
	Exclude                      []PolicyFilterType `json:"exclude"`
	ID                           string             `json:"id"`
	Include                      []PolicyFilterType `json:"include"`
	IsolationRequired            bool               `json:"isolation_required"`
	Name                         string             `json:"name"`
	PurposeJustificationPrompt   string             `json:"purpose_justification_prompt"`
	PurposeJustificationRequired bool               `json:"purpose_justification_required"`
	Require                      []PolicyFilterType `json:"require"`
	SessionDuration              string             `json:"session_duration"`
	UpdatedAt                    string             `json:"updated_at"`
}

type ApprovalGroup struct {
	ApprovalsNeeded int    `json:"approvals_needed"`
	EmailAddresses  string `json:"email_addresses"`
	EmailListUuid   string `json:"email_list_uuid"`
}

type PolicyFilterType struct {
	PolicyFilterEmailType
	PolicyFilterEmailListType
	PolicyFilterEmailDomainType
	PolicyFilterEveryoneType
	PolicyFilterIpRangesType
	PolicyFilterIpListType
	PolicyFilerValidCertificateType
	PolicyFilterAccessGroupType
	PolicyFilterAzureGroupType
	PolicyFilterGithubOrganizationType
	PolicyFilterGoogleWorkspaceGroupType
	PolictyFilterOktaGroupType
	PolicyFilterSamlGroupType
	PolicyFilterServiceTokenType
	PolicyFilterAnyValidServiceTokenType
	PolicyFilterExternalEvaluationType
	PolicyFilterGeoType
	PolicyFilterAuthenticationMethodType
	PolicyFilterDevicePostureType
}

type PolicyFilterEmailType struct {
	Email struct {
		Email string `json:"email"`
	} `json:"email"`
}

type PolicyFilterEmailListType struct {
	EmailList struct {
		ID string `json:"id"`
	} `json:"email_list"`
}

type PolicyFilterEmailDomainType struct {
	EmailDomain struct {
		Domain string `json:"domain"`
	} `json:"email_domain"`
}

type PolicyFilterEveryoneType struct {
	Everyone interface{} `json:"everyone"`
}

type PolicyFilterIpRangesType struct {
	IP struct {
		IP string `json:"ip"`
	} `json:"ip"`
}

type PolicyFilterIpListType struct {
	IPList struct {
		ID string `json:"id"`
	} `json:"ip_list"`
}

type PolicyFilerValidCertificateType struct {
	Certificate interface{} `json:"certificate"`
}

type PolicyFilterAccessGroupType struct {
	Group struct {
		ID string `json:"id"`
	} `json:"group"`
}

type PolicyFilterAzureGroupType struct {
	AzureAD struct {
		ConnectionID string `json:"connection_id"`
		ID           string `json:"id"`
	} `json:"azureAD"`
}

type PolicyFilterGithubOrganizationType struct {
	GithubOrganization struct {
		ConnectionID string `json:"connection_id"`
		ID           string `json:"id"`
	} `json:"github-organization"`
}

type PolicyFilterGoogleWorkspaceGroupType struct {
	GSuite struct {
		ConnectionID string `json:"connection_id"`
		Email        string `json:"email"`
	} `json:"gsuite"`
}

type PolictyFilterOktaGroupType struct {
	Okta struct {
		ConnectionID string `json:"connection_id"`
		Email        string `json:"email"`
	} `json:"okta"`
}

type PolicyFilterSamlGroupType struct {
	Saml struct {
		AttributeName  string `json:"attribute_name"`
		AttributeValue string `json:"attribute_value"`
	} `json:"saml"`
}

type PolicyFilterServiceTokenType struct {
	ServiceToken struct {
		TokenID string `json:"token_id"`
	} `json:"service_token"`
}

type PolicyFilterAnyValidServiceTokenType struct {
	AnyValidServiceToken interface{} `json:"any_valid_service_token"`
}

type PolicyFilterExternalEvaluationType struct {
	ExternalEvaluation struct {
		EvaluateUrl string `json:"evaluate_url"`
		KeysUrl     string `json:"keys_url"`
	} `json:"external_evaluation"`
}

type PolicyFilterGeoType struct {
	Geo struct {
		CountryCode string `json:"country_code"`
	} `json:"geo"`
}

type PolicyFilterAuthenticationMethodType struct {
	AuthMethod struct {
		AuthMethod string `json:"auth_method"`
	} `json:"auth_method"`
}

type PolicyFilterDevicePostureType struct {
	DevicePosture struct {
		IntegrationUid string `json:"integration_uid"`
	} `json:"device_posture"`
}
