package parser_test

import (
	"testing"

	"github.com/jsteenb2/quran/parser"
)

var suraatInfo map[int]int = map[int]int{
	1:   7,
	2:   286,
	3:   200,
	4:   176,
	5:   120,
	6:   165,
	7:   206,
	8:   75,
	9:   129,
	10:  109,
	11:  123,
	12:  111,
	13:  43,
	14:  52,
	15:  99,
	16:  128,
	17:  111,
	18:  110,
	19:  98,
	20:  135,
	21:  112,
	22:  78,
	23:  118,
	24:  64,
	25:  77,
	26:  227,
	27:  93,
	28:  88,
	29:  69,
	30:  60,
	31:  34,
	32:  30,
	33:  73,
	34:  54,
	35:  45,
	36:  83,
	37:  182,
	38:  88,
	39:  75,
	40:  85,
	41:  54,
	42:  53,
	43:  89,
	44:  59,
	45:  37,
	46:  35,
	47:  38,
	48:  29,
	49:  18,
	50:  45,
	51:  60,
	52:  49,
	53:  62,
	54:  55,
	55:  78,
	56:  96,
	57:  29,
	58:  22,
	59:  24,
	60:  13,
	61:  14,
	62:  11,
	63:  11,
	64:  18,
	65:  12,
	66:  12,
	67:  30,
	68:  52,
	69:  52,
	70:  44,
	71:  28,
	72:  28,
	73:  20,
	74:  56,
	75:  40,
	76:  31,
	77:  50,
	78:  40,
	79:  46,
	80:  42,
	81:  29,
	82:  19,
	83:  36,
	84:  25,
	85:  22,
	86:  17,
	87:  19,
	88:  26,
	89:  30,
	90:  20,
	91:  15,
	92:  21,
	93:  11,
	94:  8,
	95:  8,
	96:  19,
	97:  5,
	98:  8,
	99:  8,
	100: 11,
	101: 11,
	102: 8,
	103: 3,
	104: 9,
	105: 5,
	106: 4,
	107: 7,
	108: 3,
	109: 6,
	110: 3,
	111: 5,
	112: 4,
	113: 5,
	114: 6,
}

func TestParseXML(t *testing.T) {
	quranData := `
		<quran>
			<sura index="1" name="الفاتحة">
				<aya index="1" text="بِسْمِ ٱللَّهِ ٱلرَّحْمَٰنِ ٱلرَّحِيمِ" />
				<aya index="2" text="ٱلْحَمْدُ لِلَّهِ رَبِّ ٱلْعَٰلَمِينَ" />
				<aya index="3" text="ٱلرَّحْمَٰنِ ٱلرَّحِيمِ" />
				<aya index="4" text="مَٰلِكِ يَوْمِ ٱلدِّينِ" />
				<aya index="5" text="إِيَّاكَ نَعْبُدُ وَإِيَّاكَ نَسْتَعِينُ" />
				<aya index="6" text="ٱهْدِنَا ٱلصِّرَٰطَ ٱلْمُسْتَقِيمَ" />
				<aya index="7" text="صِرَٰطَ ٱلَّذِينَ أَنْعَمْتَ عَلَيْهِمْ غَيْرِ ٱلْمَغْضُوبِ عَلَيْهِمْ وَلَا ٱلضَّآلِّينَ" />
			</sura>
			<sura index="113" name="الفلق">
				<aya index="1" text="قُلْ أَعُوذُ بِرَبِّ ٱلْفَلَقِ" bismillah="بِسْمِ ٱللَّهِ ٱلرَّحْمَٰنِ ٱلرَّحِيمِ" />
				<aya index="2" text="مِن شَرِّ مَا خَلَقَ" />
				<aya index="3" text="وَمِن شَرِّ غَاسِقٍ إِذَا وَقَبَ" />
				<aya index="4" text="وَمِن شَرِّ ٱلنَّفَّٰثَٰتِ فِى ٱلْعُقَدِ" />
				<aya index="5" text="وَمِن شَرِّ حَاسِدٍ إِذَا حَسَدَ" />
			</sura>
		</quran>
	`

	quran, err := parser.ParseXML([]byte(quranData))
	if err != nil {
		t.Fatal(err)
	}

	if quran.Suraat[0].Number != 1 {
		t.Errorf("Expected `%d` got `%d`", 1, quran.Suraat[0].Number)
	}

	if quran.Suraat[1].Name != "الفلق" {
		t.Errorf("Expected `%s` got `%s`", "الفلق", quran.Suraat[1].Name)
	}

	bismillah := "بِسْمِ ٱللهِ ٱلرَّحْمَٰنِ ٱلرَّحِيمِ"
	if string(quran.Suraat[0].Ayaat[0].Text) != bismillah {
		t.Errorf("Expected `%s` got `%s`", bismillah, quran.Suraat[0].Ayaat[0].Text)
	}

	if quran.Suraat[0].Ayaat[1].NumberInSura != 2 {
		t.Errorf("Expected `%d` got `%d`", 2, quran.Suraat[0].Ayaat[1].NumberInSura)
	}

	if string(quran.Suraat[1].Ayaat[4].Text) != "وَمِن شَرِّ حَاسِدٍ إِذَا حَسَدَ" {
		t.Errorf("Expected `%s` got `%s`", "وَمِن شَرِّ حَاسِدٍ إِذَا حَسَدَ", quran.Suraat[1].Ayaat[4].Text)
	}

	if quran.Suraat[1].Ayaat[0].Bismillah != "بِسْمِ ٱللَّهِ ٱلرَّحْمَٰنِ ٱلرَّحِيمِ" {
		t.Errorf("Expected `%s` got `%s`", "بِسْمِ ٱللَّهِ ٱلرَّحْمَٰنِ ٱلرَّحِيمِ", quran.Suraat[0].Ayaat[0].Bismillah)
	}
}

func TestParseQuran(t *testing.T) {
	parsedQuran, err := parser.ParseQuran("quran-simple.xml")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 114; i++ {
		if parsedQuran.Suraat[i].Number != i+1 {
			t.Errorf("expected `%d` got `%d`", i+1, parsedQuran.Suraat[i].Number)
		}
		if len(parsedQuran.Suraat[i].Ayaat) != suraatInfo[parsedQuran.Suraat[i].Number] {
			t.Errorf("sura length, expected `%d` got `%d`", suraatInfo[parsedQuran.Suraat[i].Number], len(parsedQuran.Suraat[i].Ayaat))
		}
	}
}
