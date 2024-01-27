package bereal

import "encoding/json"

type BeRealSession struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Expiration   int64  `json:"expires_in"`
}

func NewBeRealSession(src string) (*BeRealSession, error) {
	var parsed BeRealSession
	err := json.Unmarshal([]byte(src), &parsed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
