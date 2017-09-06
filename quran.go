package quran

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/jsteenb2/quran/api"
	"github.com/jsteenb2/quran/model"
	"github.com/jsteenb2/quran/parser"
)

var (
	quranBucket = []byte("quran")
)


func BuildQuranDB(db *bolt.DB) error {
	editions, err := api.GetTextTranslationEditions()
	if err != nil {
		log.Println(err)
		return err
	}

	var edition model.QuranMeta
	for idx := range editions.Editions {
		edition = BuildQuran(editions.Editions[idx].Identifier)
		check(edition.Save(db, quranBucket))
	}
	return nil
}

func BuildQuran(edition string) model.QuranMeta {
	apiQuran, err := api.GetQuranContent(edition)
	check(err)
	parsedQuran := parser.ParseQuran("quran-simple.xml")

	return model.NewQuranMeta(parsedQuran, apiQuran)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
