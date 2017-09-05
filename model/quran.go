package model

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/boltdb/bolt"
	"github.com/jsteenb2/quran/api"
)

type Quran struct {
	Suraat []Sura `xml:"sura"`
}

type QuranMeta struct {
	Suwar []SuraMeta
	api.Edition
}

func (q QuranMeta) Save(db *bolt.DB, bucket []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		data, err := q.gobEncode()
		if err != nil {
			log.Println(err)
		}

		var editionID = []byte(q.Edition.Identifier)
		return b.Put(editionID, data)
	})
}

func GetQuran(bucket, quranEdition []byte, db *bolt.DB) (QuranMeta, error) {
	var quran QuranMeta
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get(quranEdition)

		var decodeErr error
		quran, decodeErr = gobDecode(v)
		if decodeErr != nil {
			log.Println(decodeErr)
			return decodeErr
		}
		return nil
	})
	return quran, err
}

func (q QuranMeta) gobEncode() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(q)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gobDecode(data []byte) (QuranMeta, error) {
	var q QuranMeta
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&q)
	if err != nil {
		return QuranMeta{}, err
	}
	return q, nil
}

func NewQuranMeta(tanzilQuran Quran, apiQuran api.QuranResponse) QuranMeta {
	suwar := make([]SuraMeta, 0)
	for idx := range apiQuran.Data.Surahs {
		suwar = append(suwar, NewSuraMeta(tanzilQuran.Suraat[idx], apiQuran.Data.Surahs[idx]))
	}

	return QuranMeta{
		Edition: apiQuran.Data.Edition,
		Suwar:   suwar,
	}
}
