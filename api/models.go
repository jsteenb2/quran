package api

import "encoding/json"

type QuranResponse struct {
	Code   int
	Status string
	Data   struct {
		Surahs  []Surah
		Edition Edition
	}
}

type EditionResponse struct {
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

type Surah struct {
	Number                 int
	Name                   string
	EnglishName            string
	EnglishNameTranslation string
	RevelationType         string
	Ayahs                  []Ayah
}

type Ayah struct {
	Number      int          `json:"number"`
	Translation string       `json:"text"`
	Juz         int          `json:"juz"`
	Manzil      int          `json:"manzil"`
	Page        int          `json:"page"`
	Ruku        int          `json:"ruku"`
	HizbQuarter int          `json:"hizbQuarter"`
	Sajda       *SajdahDeets `json:"sajda"`
}

type SajdahDeets struct {
	ID          int
	Recommended bool
	Obligatory  bool
}

func (s *SajdahDeets) UnmarshalJSON(b []byte) error {
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

func (a Ayah) HasSajdah() bool {
	return a.Sajda.Recommended || a.Sajda.Obligatory
}
