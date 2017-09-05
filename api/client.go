package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetTextTranslationEditions() (EditionResponse, error) {
	body, err := getEditions("text", "translation")
	check(err)

	var editions EditionResponse
	err = json.Unmarshal(body, &editions)
	check(err)
	return editions, err
}

func getEditions(format, textType string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.alquran.cloud/edition?format=%s&type=%s", format, textType))
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return body, err
}

func GetQuranContent(edition string) (QuranResponse, error) {
	body, err := getQuran(edition)
	check(err)

	var quran QuranResponse
	err = json.Unmarshal(body, &quran)
	check(err)

	return quran, err
}

func getQuran(edition string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://api.alquran.cloud/quran/%s", edition))
	check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	return body, err
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
