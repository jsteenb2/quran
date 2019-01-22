package quran

import (
	"github.com/boltdb/bolt"
	"github.com/jsteenb2/quran/api"
	"github.com/jsteenb2/quran/model"
	"github.com/jsteenb2/quran/parser"
	"github.com/pkg/errors"
)

var (
	quranBucket = []byte("quran")
)

func GetQuranDB(dbPath string) (*bolt.DB, error) {
	return bolt.Open(dbPath, 0666, &bolt.Options{ReadOnly: true})
}

func BuildQuranDB(db model.DBface) error {
	editions, err := api.GetTextTranslationEditions()
	if err != nil {
		return err
	}

	var edition model.QuranMeta
	for idx := range editions.Editions {
		if editions.Editions[idx].Identifier != "en.sahih" && editions.Editions[idx].Identifier != "id.muntakhab" {
			continue
		}
		edition, err = BuildQuran(editions.Editions[idx].Identifier)
		if err != nil {
			return err
		}
		if err := edition.Save(db, quranBucket); err != nil {
			return err
		}
	}
	return nil
}

func BuildQuran(edition string) (model.QuranMeta, error) {
	apiQuran, err := api.GetQuranContent(edition)
	if err != nil {
		return model.QuranMeta{}, errors.Wrap(err, "BuildQuran")
	}
	parsedQuran, err := parser.ParseQuran("quran-simple.xml")
	if err != nil {
		return model.QuranMeta{}, errors.Wrap(err, "BuildQuran")
	}

	return model.NewQuranMeta(parsedQuran, apiQuran), nil
}
