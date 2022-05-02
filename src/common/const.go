package common

type contextKey string

// const variables
const (
	DBAuth              = "db_auth"
	LDAPAuth            = "ldap_auth"
	UAAAuth             = "uaa_auth"
	HTTPAuth            = "http_auth"
	OIDCAuth            = "oidc_auth"
	DBCfgManager        = "db_cfg_manager"
	InMemoryCfgManager  = "in_memory_manager"
	RestCfgManager      = "rest_config_manager"
	ProCrtRestrEveryone = "everyone"
	ProCrtRestrAdmOnly  = "adminonly"
	LDAPScopeBase       = 0
	LDAPScopeOnelevel   = 1
	LDAPScopeSubtree    = 2

	OIDCName               = "oidc_name"
	OIDCEndpoint           = "oidc_endpoint"
	OIDCCLientID           = "oidc_client_id"
	OIDCClientSecret       = "oidc_client_secret"
	OIDCVerifyCert         = "oidc_verify_cert"
	OIDCAdminGroup         = "oidc_admin_group"
	OIDCGroupsClaim        = "oidc_groups_claim"
	OIDCAutoOnboard        = "oidc_auto_onboard"
	OIDCExtraRedirectParms = "oidc_extra_redirect_parms"
	OIDCScope              = "oidc_scope"
	OIDCUserClaim          = "oidc_user_claim"

	OIDCCallbackPath = "/c/oidc/callback"
	OIDCLoginPath    = "/c/oidc/login"
)
