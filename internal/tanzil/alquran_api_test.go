package tanzil

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestGetAPIQuranContent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api tests")
	}

	apiClient := newClient(&http.Client{Timeout: 30 * time.Second})

	quranData, err := apiClient.getQuranContent(context.TODO(), "en.sahih")
	if err != nil {
		t.Fatal(err)
	}

	if len(quranData.Surahs) != 114 {
		t.Errorf("Did not get expected # of suraat, expected `%d` got `%d`", 114, len(quranData.Surahs))
	}

	if quranData.Surahs[1].EnglishNameTranslation != "The Cow" {
		t.Errorf("Did not receive surat al baqarah, expected `%s`, got `%s`", "The Cow", quranData.Surahs[1].EnglishNameTranslation)
	}

	if quranData.Surahs[0].EnglishNameTranslation != "The Opening" {
		t.Errorf("Did not receive surat al baqarah, expected `%s`, got `%s`", "The Opening", quranData.Surahs[0].EnglishNameTranslation)
	}

	if len(quranData.Surahs[1].Ayahs) != 286 {
		t.Errorf("Did not receive all surat al baqarah, expected `%d`, got `%d`", 286, len(quranData.Surahs[1].Ayahs))
	}

	bismillah := "In the name of Allah, the Entirely Merciful, the Especially Merciful."
	if quranData.Surahs[0].Ayahs[0].Translation != bismillah {
		t.Errorf("Did not recieve correct translation, expected `%s`, got `%s`", bismillah, quranData.Surahs[0].Ayahs[0].Translation)
	}
}

func TestGetTextTranslationEditions(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api tests")
	}

	expectedEnglishEdition := Edition{
		Identifier:  "en.sahih",
		Language:    "en",
		Name:        "Saheeh International",
		EnglishName: "Saheeh International",
		Format:      "text",
		Type:        "translation",
	}

	expectedIndonesEdition := Edition{
		Identifier:  "id.indonesian",
		Language:    "id",
		Name:        "Bahasa Indonesia",
		EnglishName: "Unknown",
		Format:      "text",
		Type:        "translation",
	}

	apiClient := newClient(&http.Client{Timeout: 30 * time.Second})

	editions, err := apiClient.getTextTranslationEditions(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if !contains(editions, expectedEnglishEdition) {
		t.Errorf("did not get expected Edition, expected `%v`", expectedEnglishEdition)
	}

	if !contains(editions, expectedIndonesEdition) {
		t.Errorf("did not get expected Edition, expected `%v`", expectedIndonesEdition)
	}
}

func contains(hay []Edition, needle Edition) bool {
	for _, a := range hay {
		if a == needle {
			return true
		}
	}
	return false
}
