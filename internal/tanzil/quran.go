package tanzil

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

func NewQuran(ctx context.Context, httpClient *http.Client, edition string) (Quran, error) {
	apiClient := newClient(httpClient)

	apiQuran, err := apiClient.getQuranContent(ctx, edition)
	if err != nil {
		return Quran{}, errors.Wrap(err, "BuildQuran")
	}

	parsedQuran, err := parseQuran("quran-simple.xml")
	if err != nil {
		return Quran{}, errors.Wrap(err, "BuildQuran")
	}

	return newQuranMeta(parsedQuran, apiQuran), nil
}

type (
	Quran struct {
		Suwar   []Surah
		Edition Edition
	}

	Surah struct {
		Number                 int    `json:"surahNumber"`
		Name                   string `json:"name"`
		EnglishName            string `json:"englishTransliteration"`
		EnglishNameTranslation string `json:"englishName"`
		RevelationType         string `json:"revelationType"`
		Ayaat                  []Ayah `json:"ayaat"`
	}

	Ayah struct {
		tanzilAyah
		apiAyah
	}
)

func newQuranMeta(tanzilQuran tanzilQuran, apiQuran apiQuran) Quran {
	suwar := make([]Surah, 0)
	for idx := range apiQuran.Surahs {
		suwar = append(suwar, newSuraMeta(tanzilQuran.Suraat[idx], apiQuran.Surahs[idx]))
	}

	return Quran{
		Edition: apiQuran.Edition,
		Suwar:   suwar,
	}
}

func newSuraMeta(tanzilSura tanzilSura, apiSura apisurah) Surah {
	newSura := Surah{
		Number:                 tanzilSura.Number,
		Name:                   tanzilSura.Name,
		EnglishName:            apiSura.EnglishName,
		EnglishNameTranslation: apiSura.EnglishNameTranslation,
		RevelationType:         apiSura.RevelationType,
	}

	ayaat := make([]Ayah, 0, len(tanzilSura.Ayaat))
	for idx := range tanzilSura.Ayaat {
		ayaat = append(ayaat, Ayah{
			tanzilAyah: tanzilSura.Ayaat[idx],
			apiAyah:    apiSura.Ayahs[idx],
		})
	}

	newSura.Ayaat = ayaat
	return newSura
}
