package tanzil_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jsteenb2/quran/internal/tanzil"
)

func TestGetQuranContent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api tests")
	}

	client := &http.Client{Timeout: 5 * time.Minute}

	quran, err := tanzil.NewQuran(context.TODO(), client, "en.sahih")
	if err != nil {
		t.Fatal(err)
	}

	if numSuwar := len(quran.Suwar); numSuwar != 114 {
		t.Errorf("Did not get expected # of suraat, expected `%d` got `%d`", 114, numSuwar)
	}

	if baqaraName := quran.Suwar[1].EnglishNameTranslation; baqaraName != "The Cow" {
		t.Errorf("Did not receive surat al baqarah, expected `%s`, got `%s`", "The Cow", baqaraName)
	}

	if fatihaName := quran.Suwar[0].EnglishNameTranslation; fatihaName != "The Opening" {
		t.Errorf("Did not receive surat al baqarah, expected `%s`, got `%s`", "The Opening", fatihaName)
	}

	if numAyaat := len(quran.Suwar[1].Ayaat); numAyaat != 286 {
		t.Errorf("Did not receive all surat al baqarah, expected `%d`, got `%d`", 286, numAyaat)
	}

	bismillah := "In the name of Allah, the Entirely Merciful, the Especially Merciful."
	if trans := quran.Suwar[0].Ayaat[0].Translation; trans != bismillah {
		t.Errorf("Did not recieve correct translation, expected `%s`, got `%s`", bismillah, trans)
	}
}
