package firebase

import "encoding/json"

type FirebaseAuthSession struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"id_token"`
	ExpiresIn    string `json:"expires_in"`
	UserId       string `json:"user_id"`
}

func NewFirebaseAuthSessionFromJson(src string) (*FirebaseAuthSession, error) {
	var session FirebaseAuthSession
	err := json.Unmarshal([]byte(src), &session)

	return &session, err
}

func RefreshToken(refresh string) (*FirebaseAuthSession, error) {
	url := "https://securetoken.googleapis.com/v1/token?key=" + googleApiKey
	body := `{"grantType":"refresh_token","refreshToken":"` + refresh + `"}`

	req := setupFirebaseRequest("POST", url, body)
	req.Header.Set("User-Agent", "BeReal/1.0.1 (AlexisBarreyat.BeReal; build:9513; iOS 16.0.2) 1.0.0/BRApriKit")

	var parsed FirebaseAuthSession
	err := sendFirebaseRequest(req, &parsed)

	return &parsed, err
}
