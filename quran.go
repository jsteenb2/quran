package quran

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/jsteenb2/quran/api"
	"github.com/jsteenb2/quran/model"
	"github.com/jsteenb2/quran/parser"
)

var (
	quranBucket = []byte("quran")
)

func GetQuranDB() *bolt.DB {
	path := fmt.Sprintf("%s/src/github.com/jsteenb2/quran", os.Getenv("GOPATH"))
	db, err := bolt.Open(path+"/quran.db", 0666, &bolt.Options{ReadOnly: true})
	check(err)
	return db
}

func BuildQuranDB(db model.DBface) error {
	editions, err := api.GetTextTranslationEditions()
	if err != nil {
		log.Println(err)
		return err
	}

	var edition model.QuranMeta
	for idx := range editions.Editions {
		if editions.Editions[idx].Identifier == "en.sahih" || editions.Editions[idx].Identifier == "id.muntakhab" {
			edition = BuildQuran(editions.Editions[idx].Identifier)
			check(edition.Save(db, quranBucket))
		}
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
