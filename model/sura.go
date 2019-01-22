package model

import "github.com/jsteenb2/quran/api"

type Sura struct {
	Number int    `xml:"index,attr"`
	Name   string `xml:"name,attr"`
	Ayaat  []Aya  `xml:"aya"`
}

type SuraMeta struct {
	Number                 int        `json:"surahNumber"`
	Name                   string     `json:"name"`
	EnglishName            string     `json:"englishTransliteration"`
	EnglishNameTranslation string     `json:"englishName"`
	RevelationType         string     `json:"revelationType"`
	Ayaat                  []AyahMeta `json:"ayaat"`
}

func NewSuraMeta(tanzilSura Sura, apiSura api.Surah) SuraMeta {
	newSura := SuraMeta{
		Number:                 tanzilSura.Number,
		Name:                   tanzilSura.Name,
		EnglishName:            apiSura.EnglishName,
		EnglishNameTranslation: apiSura.EnglishNameTranslation,
		RevelationType:         apiSura.RevelationType,
	}

	ayaat := make([]AyahMeta, 0)
	for idx := range tanzilSura.Ayaat {
		ayaat = append(ayaat, NewAyahMeta(tanzilSura.Ayaat[idx], apiSura.Ayahs[idx]))
	}

	newSura.Ayaat = ayaat
	return newSura
}
