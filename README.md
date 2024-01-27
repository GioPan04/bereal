# bereal.go

A BeReal client in go

## Docs

### Login
You have two ways to login, via phone number or using an existing session.

#### SMS auth
```go
otp_session := bereal.SendOtp("+000000000000")
if err != nil {
  panic(err)
}

// Obtain OTP from stdin
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter OTP Code: ")
otp, _ := reader.ReadString('\n')
otp = strings.TrimSuffix(otp, "\n")

// Login with given otp
session, err := bereal.VerifyOtp(otp, otp_session)
if err != nil {
  panic(err)
}
```

#### Pre-existing session
```go
src := `{"refresh_token":"<<REDACTED>>","access_token":"<<REDACTED>>","expires_in":3600}`
session, err := bereal.LoginFromFirebase(src)
if err != nil {
  panic(err)
}
```
