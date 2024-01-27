package firebase

type FirebaseClientVerify struct {
	Receipt string `json:"receipt"`
}

func VerifyClient() (string, error) {
	url := "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyClient?key=" + googleApiKey
	reqBody := `{"appToken": "` + firebaseAppToken + `"}`

	req := setupFirebaseRequest("POST", url, reqBody)

	var parsed FirebaseClientVerify
	err := sendFirebaseRequest(req, &parsed)
	if err != nil {
		return "", err
	}

	return parsed.Receipt, nil
}

type FirebaseOtpSession struct {
	SessionInfo string `json:"sessionInfo"`
}

// Send an OTP to the desired phone number.
// This function requires a `receipt` returned from `VerifyClient`
func SendOtp(phone string, receipt string) (string, error) {
	url := "https://www.googleapis.com/identitytoolkit/v3/relyingparty/sendVerificationCode?key=" + googleApiKey
	reqBody := `{"phoneNumber": "` + phone + `","iosReceipt":"` + receipt + `"}`

	req := setupFirebaseRequest("POST", url, reqBody)

	var parsed FirebaseOtpSession
	err := sendFirebaseRequest(req, &parsed)
	if err != nil {
		return "", err
	}

	return parsed.SessionInfo, nil
}

type FirebaseVerifyOtp struct {
	RefreshToken string `json:"refreshToken"`
}

func VerifyOtp(otp string, session string) (string, error) {
	url := "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPhoneNumber?key=" + googleApiKey
	reqBody := `{"code": "` + otp + `","sessionInfo":"` + session + `","operation":"SIGN_UP_OR_IN"}`

	req := setupFirebaseRequest("POST", url, reqBody)

	var parsed FirebaseVerifyOtp
	err := sendFirebaseRequest(req, &parsed)
	if err != nil {
		return "", err
	}

	return parsed.RefreshToken, nil
}
