package model

import "github.com/jsteenb2/quran/api"

type Aya struct {
	NumberInSura int    `xml:"index,attr"`
	Bismillah    string `xml:"bismillah,attr,omitempty"`
	Text         string `xml:"text,attr"`
}

type AyahMeta struct {
	Aya
	api.Ayah
}

func NewAyahMeta(tanzilAya Aya, apiAya api.Ayah) AyahMeta {
	return AyahMeta{tanzilAya, apiAya}
}
