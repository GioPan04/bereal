package bereal

import (
	"encoding/json"
	"net/http"
)

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

func (s *BeRealSession) GetHttpClient(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "BeReal/1.0.1 (AlexisBarreyat.BeReal; build:9513; iOS 16.0.2) 1.0.0/BRApriKit")
	req.Header.Set("x-ios-bundle-identifier", "AlexisBarreyat.BeReal")
	req.Header.Set("Authorization", "Bearer "+s.AccessToken)
	req.Header.Set("bereal-app-version-code", "14549")
	req.Header.Set("bereal-signature", "berealsignature")
	req.Header.Set("bereal-device-id", "berealdeviceid")

	return req, nil
}
