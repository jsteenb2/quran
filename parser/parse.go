package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jsteenb2/quran/model"
)

func ParseXML(data []byte) (model.Quran, error) {
	var quranParsed model.Quran
	err := xml.Unmarshal(data, &quranParsed)
	return quranParsed, err
}

func ParseQuran() model.Quran {
	path := fmt.Sprintf("%s/src/github.com/jsteenb2/quran/quran-uthmani.xml", os.Getenv("GOPATH"))
	data, err := ioutil.ReadFile(path)
	check(err)
	quran, err := ParseXML(data)
	check(err)
	return quran
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
