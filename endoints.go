package bereal

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/GioPan04/bereal/models"
)

const apiUrl = "https://mobile.bereal.com/api"

func (s *BeRealSession) GetMemories() (*models.BeRealMemoriesEndpoint, error) {
	url := apiUrl + "/feeds/memories"
	req, err := s.GetHttpClient("GET", url)
	if err != nil {
		return nil, err
	}

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

	var parsed models.BeRealMemoriesEndpoint
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
