package bereal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/GioPan04/bereal/firebase"
)

// Send an One Time Password (6 code digit) to the phone number requested.
// It returns the otp session token, that must be used to verify the code.
func SendOtp(phone string) (string, error) {
	verifyRes, err := firebase.VerifyClient()
	if err != nil {
		return "", nil
	}

	session, err := firebase.SendOtp(phone, verifyRes)
	if err != nil {
		return "", err
	}

	return session, nil
}

func VerifyOtp(otp string, sessionToken string) (*BeRealSession, error) {
	refreshToken, err := firebase.VerifyOtp(otp, sessionToken)
	if err != nil {
		fmt.Println("Error: The OTP may be invalid")
		return nil, err
	}

	firebaseAuth, err := firebase.RefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	j, _ := json.Marshal(firebaseAuth)
	fmt.Println(string(j))

	return login(*firebaseAuth)
}

func LoginFromFirebase(src string) (*BeRealSession, error) {
	session, err := firebase.NewFirebaseAuthSessionFromJson(src)
	if err != nil {
		return nil, err
	}

	return login(*session)
}

func login(fbSession firebase.FirebaseAuthSession) (*BeRealSession, error) {
	url := "https://auth.bereal.team/token?grant_type=firebase"
	reqBody, _ := encodeBody(map[string]string{
		"grant_type":    "firebase",
		"client_id":     "ios",
		"client_secret": "962D357B-B134-4AB6-8F53-BEA2B7255420",
		"token":         fbSession.AccessToken,
	})

	req, _ := http.NewRequest("POST", url, reqBody)
	req.Header.Set("user-agent", "BeReal/1.0.1 (AlexisBarreyat.BeReal; build:9513; iOS 16.0.2) 1.0.0/BRApriKit")
	req.Header.Set("x-ios-bundle-identifier", "AlexisBarreyat.BeReal")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		return nil, errors.New(string(body))
	}

	return NewBeRealSession(string(body))
}
