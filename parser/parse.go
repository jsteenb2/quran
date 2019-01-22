package parser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/jsteenb2/quran/model"
)

func ParseXML(data []byte) (model.Quran, error) {
	var quranParsed model.Quran
	err := xml.Unmarshal(data, &quranParsed)
	return quranParsed, err
}

func ParseQuran(filename string) (model.Quran, error) {
	path := fmt.Sprintf("%s/src/github.com/jsteenb2/quran/%s", os.Getenv("GOPATH"), filename)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return model.Quran{}, errors.Wrapf(err, "reading tanzil file=%q", filename)
	}
	quran, err := ParseXML(data)
	if err != nil {
		return model.Quran{}, errors.Wrapf(err, "parsing tanzil file=%q", filename)
	}
	return quran, nil
}
