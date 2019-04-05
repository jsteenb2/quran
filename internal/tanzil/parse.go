package tanzil

import (
	"encoding/xml"
	"os"
	"strings"
)

type tanzilQuran struct {
	Suraat []tanzilSura `xml:"sura"`
}

type tanzilSura struct {
	Number int          `xml:"index,attr"`
	Name   string       `xml:"name,attr"`
	Ayaat  []tanzilAyah `xml:"aya"`
}

type tanzilAyah struct {
	NumberInSura int    `json:"numberInSurah" xml:"index,attr"`
	Bismillah    string `json:"bismillah" xml:"bismillah,attr,omitempty"`
	Text         string `json:"arabicText" xml:"text,attr"`
}

type ArabicText string

func (arTxt *ArabicText) UnmarshalXMLAttr(attr xml.Attr) error {
	r := strings.NewReplacer("للَّه", "لله", "لِلَّهِ", "لِلهِ")
	*arTxt = ArabicText(r.Replace(attr.Value))
	return nil
}

func parseQuran(filepath string) (tanzilQuran, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return tanzilQuran{}, err
	}

	var quran tanzilQuran
	if err := xml.NewDecoder(f).Decode(&quran); err != nil {
		return tanzilQuran{}, err
	}
	return quran, nil
}
