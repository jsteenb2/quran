package api

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
	Number      int
	Translation string `json:"text"`
	Juz         int
	Manzil      int
	Page        int
	Ruku        int
	HizbQuarter int
	Sajda       bool
}
