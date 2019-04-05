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

func (a *alQuranAPIClient) getTextTranslationEditions(ctx context.Context) (editionResponse, error) {
	var editions editionResponse
	err := a.c.
		Get("/Edition?format=text&type=translation").
		Success(httpc.StatusOK()).
		Decode(httpc.JSONDecode(&editions)).
		Do(ctx)
	return editions, err
}

func (a *alQuranAPIClient) getQuranContent(ctx context.Context, edition string) (quranResponse, error) {
	var quran quranResponse
	err := a.c.
		Get("/quran/" + edition).
		Success(httpc.StatusOK()).
		Decode(httpc.JSONDecode(&quran)).
		Do(ctx)
	return quran, err
}

type quranResponse struct {
	Code   int
	Status string
	Data   struct {
		Surahs  []surah
		Edition Edition
	}
}

type editionResponse struct {
	Code     int
	Status   string
	Editions []Edition `json:"data"`
}

type Edition struct {
	Identifier  string
	Language    string
	Name        string
	EnglishName string
	Format      string
	Type        string
}

type surah struct {
	Number                 int
	Name                   string
	EnglishName            string
	EnglishNameTranslation string
	RevelationType         string
	Ayahs                  []alQuranAPIAyah
}

type alQuranAPIAyah struct {
	Number      int          `json:"number"`
	Translation string       `json:"text"`
	Juz         int          `json:"juz"`
	Manzil      int          `json:"manzil"`
	Page        int          `json:"page"`
	Ruku        int          `json:"ruku"`
	HizbQuarter int          `json:"hizbQuarter"`
	Sajda       *sajdahDeets `json:"sajda"`
}

type sajdahDeets struct {
	ID          int
	Recommended bool
	Obligatory  bool
}

func (s *sajdahDeets) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}

	id, ok := m["id"].(float64)
	if !ok {
		return nil
	}

	rec, _ := m["recommended"].(bool)
	obl, _ := m["obligatory"].(bool)

	s.ID = int(id)
	s.Recommended = rec
	s.Obligatory = obl

	return nil
}

func (a alQuranAPIAyah) HasSajdah() bool {
	return a.Sajda.Recommended || a.Sajda.Obligatory
}
