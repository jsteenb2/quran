package api_test

import (
	"log"
	"testing"

	"github.com/jsteenb2/quran/api"
)

func TestGetQuranContent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api tests")
	}

	quranData, err := api.GetQuranContent("en.sahih")

	if err != nil {
		log.Println(err)
	}

	if len(quranData.Data.Surahs) != 114 {
		t.Errorf("Did not get expected # of suraat, expected `%d` got `%d`", 114, len(quranData.Data.Surahs))
	}

	if quranData.Data.Surahs[1].EnglishNameTranslation != "The Cow" {
		t.Errorf("Did not receive surat al baqarah, expected `%s`, got `%s`", "The Cow", quranData.Data.Surahs[1].EnglishNameTranslation)
	}

	if quranData.Data.Surahs[0].EnglishNameTranslation != "The Opening" {
		t.Errorf("Did not receive surat al baqarah, expected `%s`, got `%s`", "The Opening", quranData.Data.Surahs[0].EnglishNameTranslation)
	}

	if len(quranData.Data.Surahs[1].Ayahs) != 286 {
		t.Errorf("Did not receive all surat al baqarah, expected `%d`, got `%d`", 286, len(quranData.Data.Surahs[1].Ayahs))
	}

	bismillah := "In the name of Allah, the Entirely Merciful, the Especially Merciful."
	if quranData.Data.Surahs[0].Ayahs[0].Translation != bismillah {
		t.Errorf("Did not recieve correct translation, expected `%s`, got `%s`", bismillah, quranData.Data.Surahs[0].Ayahs[0].Translation)
	}
}

func TestGetTextTranslationEditions(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api tests")
	}

	expectedEnglishEdition := api.Edition{
		Identifier:  "en.sahih",
		Language:    "en",
		Name:        "Saheeh International",
		EnglishName: "Saheeh International",
		Format:      "text",
		Type:        "translation",
	}

	expectedIndonesEdition := api.Edition{
		Identifier:  "id.indonesian",
		Language:    "id",
		Name:        "Bahasa Indonesia",
		EnglishName: "Unknown",
		Format:      "text",
		Type:        "translation",
	}

	editions, err := api.GetTextTranslationEditions()

	if err != nil {
		t.Error(err)
	}

	if !contains(editions.Editions, expectedEnglishEdition) {
		t.Errorf("did not get expected edition, expected `%v`", expectedEnglishEdition)
	}

	if !contains(editions.Editions, expectedIndonesEdition) {
		t.Errorf("did not get expected edition, expected `%v`", expectedIndonesEdition)
	}
}

func contains(s []api.Edition, e api.Edition) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
