package bereal

import (
	"bytes"
	"encoding/json"
)

// var alphabet = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// func generateDeviceID() string {
// 	res := make([]rune, 16)
// 	for i := range res {
// 		res[i] = alphabet[rand.Intn(len(alphabet))]
// 	}

// 	return string(res)
// }

func encodeBody(value any) (*bytes.Buffer, error) {
	json, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(json), nil
}
