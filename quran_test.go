package quran_test

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/jsteenb2/quran"
	"github.com/jsteenb2/quran/model"
)

func TestBuildQuran(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api endpoint tests")
	}
	quranComplete, err := quran.BuildQuran("en.sahih")
	if err != nil {
		t.Fatal(err)
	}

	if len(quranComplete.Suwar) != 114 {
		t.Errorf("Wrong number of suwar, expected `%d`, got `%d`", 114, len(quranComplete.Suwar))
	}

	bismillah := "بِسْمِ اللَّهِ الرَّحْمَنِ الرَّحِيمِ"
	if string(quranComplete.Suwar[0].Ayaat[0].Text) != bismillah {
		t.Errorf("Expected `%s` got `%s`", bismillah, quranComplete.Suwar[0].Ayaat[0].Text)
	}

	alfaatiha := quranComplete.Suwar[0]
	if alfaatiha.Ayaat[0].Number != 1 {
		t.Errorf("Wrong ayaat number %s, expected `%d` got `%d`", alfaatiha.EnglishName, 1, alfaatiha.Ayaat[0].Number)
	}

	arRahman := quranComplete.Suwar[54]
	if len(arRahman.Ayaat) != 78 {
		t.Errorf("Wrong number of ayaat for %s, expected `%d` got `%d`", arRahman.EnglishName, 78, len(arRahman.Ayaat))
	}

	expectedAyahText := "الرَّحْمَنُ"
	if string(arRahman.Ayaat[0].Text) != expectedAyahText {
		t.Errorf("Wrong text received, expected `%s`, got `%s`", expectedAyahText, arRahman.Ayaat[0].Text)
	}

	expectedHizb := 213
	if expectedHizb != arRahman.Ayaat[0].HizbQuarter {
		t.Errorf("Wrong hizb received, expected `%d`, got `%d`", expectedHizb, arRahman.Ayaat[0].HizbQuarter)
	}

	if arRahman.Ayaat[3].Translation != "[And] taught him eloquence." {
		t.Errorf("Wrong text received, expected `%s`, got `%s`", "[And] taught him eloquence.", arRahman.Ayaat[3].Translation)
	}

	var sajdaat int
	for _, surah := range quranComplete.Suwar {
		for _, aya := range surah.Ayaat {
			if ok := aya.HasSajdah(); ok {
				sajdaat++
			}
		}
	}

	if 15 != sajdaat {
		t.Errorf("expected `15` sajdaat, got `%d`", sajdaat)
	}
}

func TestBuildQuran_Indonesian(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping api endpoint tests")
	}
	quranComplete, err := quran.BuildQuran("id.muntakhab")
	if err != nil {
		t.Fatal(err)
	}

	if suwaarCount := 114; len(quranComplete.Suwar) != suwaarCount {
		t.Errorf("Wrong number of suwar, expected `%d`, got `%d`", suwaarCount, len(quranComplete.Suwar))
	}

	bismillah := "بِسْمِ اللَّهِ الرَّحْمَنِ الرَّحِيمِ"
	if string(quranComplete.Suwar[0].Ayaat[0].Text) != bismillah {
		t.Errorf("Expected `%s` got `%s`", bismillah, quranComplete.Suwar[0].Ayaat[0].Text)
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

	var sajdaat int
	for _, surah := range quranComplete.Suwar {
		for _, aya := range surah.Ayaat {
			if ok := aya.HasSajdah(); ok {
				sajdaat++
			}
		}
	}

	if 15 != sajdaat {
		t.Errorf("expected `15` sajdaat, got `%d`", sajdaat)
	}
}

func TestBuildQuranJSON(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping largest api tests")
	}

	sahih, err := quran.BuildQuran("en.sahih")
	if err != nil {
		t.Fatal(err)
	}

	if sahih.Identifier != "en.sahih" {
		t.Errorf("unexpected quran edition, expected `en.sahih`, got %q", sahih.Identifier)
	}

	expectedTitle := "Al-Faatiha"
	if sahih.Suwar[0].EnglishName != expectedTitle {
		t.Errorf("unexpected Sura returned, expected %q, got %q", expectedTitle, sahih.Suwar[0].EnglishName)
	}

	expectedAyahText := "By time,"
	if translation := sahih.Suwar[102].Ayaat[0].Translation; translation != expectedAyahText {
		t.Errorf("incorrect ayah received, expected %q, got %q", expectedAyahText, translation)
	}

	if len(sahih.Suwar) != 114 {
		t.Errorf("incorrect suwar, expected `114`, got `%d`", len(sahih.Suwar))
	}

	quranPath := path.Join(os.Getenv("GOPATH"), "src/github.com/jsteenb2/quran", "sahih.json")
	f, err := os.Create(quranPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	if err := enc.Encode(sahih); err != nil {
		t.Fatal(err)
	}
}

func TestBuildQuranDB(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping largest api tests")
	}

	dbPath := path.Join(os.Getenv("GOPATH"), "src/github.com/jsteenb2/quran", "quran.db")
	db, err := bolt.Open(dbPath, 0644, nil)
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
