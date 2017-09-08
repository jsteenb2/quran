package quran_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/jsteenb2/quran"
	"github.com/jsteenb2/quran/model"
)

func TestBuildQuran(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api endpoint tests")
	}
	quranComplete := quran.BuildQuran("en.sahih")

	if len(quranComplete.Suwar) != 114 {
		t.Errorf("Wrong number of suwar, expected `%d`, got `%d`", 114, len(quranComplete.Suwar))
	}

	arRahman := quranComplete.Suwar[54]
	if len(arRahman.Ayaat) != 78 {
		t.Errorf("Wrong number of ayaat for %s, expected `%d` got `%d`", arRahman.EnglishName, 78, len(arRahman.Ayaat))
	}

	expectedAyahText := "الرَّحْمَنُ"
	if string(arRahman.Ayaat[0].Text) != expectedAyahText {
		t.Errorf("Wrong text received, expected `%s`, got `%s`", expectedAyahText, arRahman.Ayaat[0].Text)
	}

	if arRahman.Ayaat[3].Translation != "[And] taught him eloquence." {
		t.Errorf("Wrong text received, expected `%s`, got `%s`", "[And] taught him eloquence.", arRahman.Ayaat[3].Translation)
	}
}

func TestBuildQuran_Indonesian(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api endpoint tests")
	}
	quranComplete := quran.BuildQuran("id.muntakhab")

	if len(quranComplete.Suwar) != 114 {
		t.Errorf("Wrong number of suwar, expected `%d`, got `%d`", 114, len(quranComplete.Suwar))
	}

	arRahman := quranComplete.Suwar[54]
	if len(arRahman.Ayaat) != 78 {
		t.Errorf("Wrong number of ayaat for %s, expected `%d` got `%d`", arRahman.EnglishName, 78, len(arRahman.Ayaat))
	}

	expectedAyahText := "الرَّحْمَنُ"
	if string(arRahman.Ayaat[0].Text) != expectedAyahText {
		t.Errorf("Wrong text received, expected `%s`, got `%s`", expectedAyahText, arRahman.Ayaat[0].Text)
	}

	expectedTranslation := "Dia menciptakan dan mengajarkan manusia kemampuan menjelaskan apa yang ada dalam dirinya, untuk membedakan dirinya dari makhluk lain."
	if arRahman.Ayaat[3].Translation != expectedTranslation {
		t.Errorf("Wrong text received, expected `%s`, got `%s`", expectedTranslation, arRahman.Ayaat[3].Translation)
	}
}

func TestBuildQuranDB(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping largest api tests")
	}

	path := fmt.Sprintf("%s/src/github.com/jsteenb2/quran", os.Getenv("GOPATH"))
	db, err := bolt.Open(path+"/quran.db", 0644, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = quran.BuildQuranDB(db)

	if err != nil {
		t.Fatal(err)
	}

	sahih, err := model.GetQuran([]byte("quran"), []byte("en.sahih"), db)

	if sahih.Identifier != "en.sahih" {
		t.Errorf("unexpected quran edition, expected `en.sahih`, got `%s`", sahih.Identifier)
	}

	expectedTitle := "Al-Faatiha"
	if sahih.Suwar[0].EnglishName != expectedTitle {
		t.Errorf("unexpected Sura returned, expected `%s`, got `%s`", expectedTitle, sahih.Suwar[0].EnglishName)
	}

	expectedAyahText := "By time,"
	if translation := sahih.Suwar[102].Ayaat[0].Translation; translation != expectedAyahText {
		t.Errorf("incorrect ayah received, expected `%s`, got `%s`", expectedAyahText, translation)
	}

	indo, err := model.GetQuran([]byte("quran"), []byte("id.muntakhab"), db)
	if indo.Identifier != "id.muntakhab" {
		t.Errorf("unexpected quran edition, expected `id.muntakhab`, got `%s`", indo.Identifier)
	}

	if suwarsLen := len(indo.Suwar); suwarsLen != 114 {
		t.Errorf("unexpected # suwar, expected `%d`, got `%d`", 114, suwarsLen)
	}
}
