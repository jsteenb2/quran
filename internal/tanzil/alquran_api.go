package tanzil

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jsteenb2/httpc"
)

type alQuranAPIClient struct {
	c *httpc.Client
}

func newClient(client *http.Client) *alQuranAPIClient {
	return &alQuranAPIClient{
		c: httpc.New(client, httpc.WithBaseURL("http://api.alquran.cloud")),
	}
}

func (a *alQuranAPIClient) getTextTranslationEditions(ctx context.Context) ([]Edition, error) {
	var editions struct {
		Editions []Edition `json:"data"`
	}
	err := a.c.
		Get("/edition?format=text&type=translation").
		Success(httpc.StatusOK()).
		Decode(httpc.JSONDecode(&editions)).
		Do(ctx)
	return editions.Editions, err
}

func (a *alQuranAPIClient) getQuranContent(ctx context.Context, edition string) (apiQuran, error) {
	var quran struct {
		Data apiQuran `json:"data"`
	}
	err := a.c.
		Get("/quran/" + edition).
		Success(httpc.StatusOK()).
		Decode(httpc.JSONDecode(&quran)).
		Do(ctx)
	return quran.Data, err
}

type (
	apiQuran struct {
		Surahs  []apisurah
		Edition Edition
	}

	Edition struct {
		Identifier  string
		Language    string
		Name        string
		EnglishName string
		Format      string
		Type        string
	}

	apisurah struct {
		Number                 int
		Name                   string
		EnglishName            string
		EnglishNameTranslation string
		RevelationType         string
		Ayahs                  []apiAyah
	}
)

type apiAyah struct {
	Number      int          `json:"number"`
	Translation string       `json:"text"`
	Juz         int          `json:"juz"`
	Manzil      int          `json:"manzil"`
	Page        int          `json:"page"`
	Ruku        int          `json:"ruku"`
	HizbQuarter int          `json:"hizbQuarter"`
	Sajda       *sajdahDeets `json:"sajda"`
}

func (a apiAyah) HasSajdah() bool {
	return a.Sajda.Recommended || a.Sajda.Obligatory
}

type sajdahDeets struct {
	ID          int
	Recommended bool
	Obligatory  bool
}

func (s *sajdahDeets) UnmarshalJSON(b []byte) error {
	var m struct {
		ID          *int `json:"id"`
		Recommended bool `json:"recommended"`
		Obligatory  bool `json:"obligatory"`
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}

	if m.ID == nil {
		return nil
	}

	s.ID = *m.ID
	s.Recommended = m.Recommended
	s.Obligatory = m.Obligatory

	return nil
}
