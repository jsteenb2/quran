package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func GetTextTranslationEditions() (EditionResponse, error) {
	resp, err := http.Get("http://api.alquran.cloud/edition?format=text&type=translation")
	if err != nil {
		return EditionResponse{}, err
	}
	defer resp.Body.Close()

	var editions EditionResponse
	if err := json.NewDecoder(resp.Body).Decode(&editions); err != nil {
		return EditionResponse{}, errors.Wrap(err, "unmarshalling editions")
	}
	return editions, nil
}

func GetQuranContent(edition string) (QuranResponse, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.alquran.cloud/quran/%s", edition))
	if err != nil {
		return QuranResponse{}, err
	}
	defer resp.Body.Close()

	var quran QuranResponse
	if err := json.NewDecoder(resp.Body).Decode(&quran); err != nil {
		return QuranResponse{}, errors.Wrap(err, "unmarshalling quran response")
	}

	return quran, nil
}
