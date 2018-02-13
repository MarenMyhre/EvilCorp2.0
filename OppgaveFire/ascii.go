package ascii

import "fmt"


const Ascii = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f" + "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8A\x8B\x8C\x8D\x8E\x8F\x90\x91\x92" +
	"\x93\x94\x95\x96\x97\x98\x99\x9A\x9C\x9D\x9E\x9F\xA0\xA1\xA2\xA3\xA4\xA5\xA6" +
	"\xA7\xA8\xA9\xAA\xAB\xAC\xAD\xAE\xAF\xB0\xB1\xB2\xB3\xB4\xB5\xB6\xB7\xB8\xB9" +
	"\xBA\xBB\xBC\xBD\xBE\xBF\xC0\xC1\xC2\xC3\xC4\xC5\xC6\xC7\xC8\xC9\xCA\xCB\xCC" +
	"\xCD\xCE\xCF\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8\xD8\xD9\xDA\xDB\xDC\xDD\xDE" +
	"\xDF\xE0\xE1\xE2\xE3\xE4\xE5\xE6\xE7\xE8\xE9\xEA\xEB\xEC\xED\xEE\xEF\xF0\xF1" +
	"\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9\xFA\xFB\xFC\xFD\xFE\xFF"

func  IterateOverExtendedASCIIStringLiteral( sl string) {
	for i := 0; i < len(sl) ; i++{
		fmt.Printf("%X %c %b \n", sl[i], sl[i], sl[i])

	}

}

func cp1252ToUTF8(cp string) string {
	r := make([]rune, len(cp))
	for i := 0; i < len(cp); i++ {
		r[i] = cp1252[cp[i]]
	}
	return string(r)
}

func ExtendedASCIIText() {
	cp := "\x80\xf7\xbe dollar"
	str := cp1252ToUTF8(cp)
	fmt.Printf("%q\n", str)
}

func init() {
	for i, r := range cp1252 {
		if r == 0 {
			cp1252[i] = rune(i)
		}
	}
}

// cp1252 to Unicode table
// ftp://ftp.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/WINDOWS/CP1252.TXT
var cp1252 = [256]rune{
	0x80: '\u20AC', // EURO SIGN
	0x81: '\uFFFD', // UNDEFINED
	0x82: '\u201A', // SINGLE LOW-9 QUOTATION MARK
	0x83: '\u0192', // LATIN SMALL LETTER F WITH HOOK
	0x84: '\u201E', // DOUBLE LOW-9 QUOTATION MARK
	0x85: '\u2026', // HORIZONTAL ELLIPSIS
	0x86: '\u2020', // DAGGER
	0x87: '\u2021', // DOUBLE DAGGER
	0x88: '\u02C6', // MODIFIER LETTER CIRCUMFLEX ACCENT
	0x89: '\u2030', // PER MILLE SIGN
	0x8A: '\u0160', // LATIN CAPITAL LETTER S WITH CARON
	0x8B: '\u2039', // SINGLE LEFT-POINTING ANGLE QUOTATION MARK
	0x8C: '\u0152', // LATIN CAPITAL LIGATURE OE
	0x8D: '\uFFFD', // UNDEFINED
	0x8E: '\u017D', // LATIN CAPITAL LETTER Z WITH CARON
	0x8F: '\uFFFD', // UNDEFINED
	0x90: '\uFFFD', // UNDEFINED
	0x91: '\u2018', // LEFT SINGLE QUOTATION MARK
	0x92: '\u2019', // RIGHT SINGLE QUOTATION MARK
	0x93: '\u201C', // LEFT DOUBLE QUOTATION MARK
	0x94: '\u201D', // RIGHT DOUBLE QUOTATION MARK
	0x95: '\u2022', // BULLET
	0x96: '\u2013', // EN DASH
	0x97: '\u2014', // EM DASH
	0x98: '\u02DC', // SMALL TILDE
	0x99: '\u2122', // TRADE MARK SIGN
	0x9A: '\u0161', // LATIN SMALL LETTER S WITH CARON
	0x9B: '\u203A', // SINGLE RIGHT-POINTING ANGLE QUOTATION MARK
	0x9C: '\u0153', // LATIN SMALL LIGATURE OE
	0x9D: '\uFFFD', // UNDEFINED
	0x9E: '\u017E', // LATIN SMALL LETTER Z WITH CARON
	0x9F: '\u0178', // LATIN CAPITAL LETTER Y WITH DIAERESIS
}


