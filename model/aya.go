package model

import (
	"encoding/xml"
	"strings"

	"github.com/jsteenb2/quran/api"
)

type Aya struct {
	NumberInSura int        `xml:"index,attr"`
	Bismillah    string     `xml:"bismillah,attr,omitempty"`
	Text         ArabicText `xml:"text,attr"`
}

type ArabicText string

func (arTxt *ArabicText) UnmarshalXMLAttr(attr xml.Attr) error {
	r := strings.NewReplacer("للَّه", "لله", "لِلَّهِ", "لِلهِ")
	*arTxt = ArabicText(r.Replace(attr.Value))
	return nil
}

type AyahMeta struct {
	Aya
	api.Ayah
}

func NewAyahMeta(tanzilAya Aya, apiAya api.Ayah) AyahMeta {
	return AyahMeta{tanzilAya, apiAya}
}
