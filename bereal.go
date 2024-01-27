package bereal

type BeRealSession struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Expiration   int64  `json:"expires_in"`
}
