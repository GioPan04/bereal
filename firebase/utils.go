package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	googleApiKey     = "AIzaSyDwjfEeparokD7sXPVQli9NsTuhT6fJ6iA"
	firebaseAppToken = "54F80A258C35A916B38A3AD83CA5DDD48A44BFE2461F90831E0F97EBA4BB2EC7"
)

func setupFirebaseRequest(method string, url string, body string) *http.Request {
	req, _ := http.NewRequest(method, url, bytes.NewBufferString(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("x-client-version", "iOS/FirebaseSDK/9.6.0/FirebaseCore-iOS")
	req.Header.Set("x-ios-bundle-identifier", "AlexisBarreyat.BeReal")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("User-Agent", "FirebaseAuth.iOS/9.6.0 AlexisBarreyat.BeReal/0.31.0 iPhone/14.7.1 hw/iPhone9_1")
	req.Header.Set("x-firebase-locale", "en")
	req.Header.Set("x-firebase-gmpid", "1:405768487586:ios:28c4df089ca92b89")

	return req
}

// Basic firebase request
func sendFirebaseRequest(req *http.Request, v any) error {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		return errors.New(string(body))
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
