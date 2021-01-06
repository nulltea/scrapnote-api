package meta

type AuthClaims struct {
	UniqueID  string `json:"nameid,omitempty"`
	Username  string `json:"unique_name,omitempty"`
	Email     string `json:"sub,omitempty"`
	Role      string `json:"role,omitempty"`
}