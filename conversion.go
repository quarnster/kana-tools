package kana

import (
	"bytes"
	"reflect"
	"strings"
	"sync"
	"unicode"
	"unsafe"
)

// ToRomaji converts hiragana and/or katakana to lowercase romaji. By default,
// the literal transliteration of づ　and ぢ are used, returnin du and di,
// respectively. Set phonetic to true to return the romaji in its correctly
// pronounced form - zu and ji.
func ToRomaji(s string, phonetic bool) string {
	s = ToRomajiCased(s, phonetic)
	s = strings.ToLower(s)

	return s
}

var pool = sync.Pool{New: func() interface{} {
	return bytes.NewBuffer(nil)
}}

func unsafeString(buf *bytes.Buffer) (ret string) {
	tmp := buf.Bytes()
	if len(tmp) == 0 {
		return ""
	}
	sh := (*reflect.StringHeader)(unsafe.Pointer(&ret))
	sh.Data = uintptr(unsafe.Pointer(&tmp[0]))
	sh.Len = len(tmp)
	return
}

// ToRomajiCased converts hiragana and/or katakana to cased romaji, where
// hiragana and katakana are presented in lowercase and uppercase respectively.
func ToRomajiCased(s string, phonetic bool) string {
	a, b := pool.Get().(*bytes.Buffer), pool.Get().(*bytes.Buffer)
	defer pool.Put(a)
	defer pool.Put(b)
	a.Reset()
	b.Reset()
	moraicNRomaji.WriteString(a, s)
	kanaToRomaji.WriteString(b, unsafeString(a))
	a.Reset()

	if phonetic {
		phoneticRomaji.WriteString(a, unsafeString(b))
	} else {
		unphoneticRomaji.WriteString(a, unsafeString(b))
	}
	b.Reset()

	s = parseRomajiDoubles([]rune(unsafeString(a)))
	postRomajiSpecial.WriteString(b, s)

	return b.String()
}

// ToHiragana converts wapuro-hepburn romaji into the equivalent hiragana.
func ToHiragana(s string) string {
	a, b := pool.Get().(*bytes.Buffer), pool.Get().(*bytes.Buffer)
	defer pool.Put(a)
	defer pool.Put(b)
	a.Reset()
	b.Reset()
	s = strings.ToLower(s)
	preHiragana.WriteString(a, s)
	romajiToHiragana.WriteString(b, unsafeString(a))
	a.Reset()
	s = strings.Map(KatakanaToHiragana, unsafeString(b))
	b.Reset()
	postHiragana.WriteString(a, s)
	postKanaSpecial.WriteString(b, unsafeString(a))
	return b.String()
}

// ToKatakana converts wapuro-hepburn romaji into the equivalent katakana.
func ToKatakana(s string) string {
	a, b := pool.Get().(*bytes.Buffer), pool.Get().(*bytes.Buffer)
	defer pool.Put(a)
	defer pool.Put(b)
	a.Reset()
	b.Reset()
	s = strings.ToUpper(s)
	preKatakana.WriteString(a, s)
	romajiToKatakana.WriteString(b, unsafeString(a))
	a.Reset()
	s = strings.Map(HiraganaToKatakana, unsafeString(b))
	b.Reset()
	postKatakana.WriteString(a, s)
	postKanaSpecial.WriteString(b, unsafeString(a))
	return b.String()
}

// ToKana converts wapuro-hepburn uppercase and lowercase romaji into
// katakana and hiragana respectively.
func ToKana(s string) string {
	a, b := pool.Get().(*bytes.Buffer), pool.Get().(*bytes.Buffer)
	defer pool.Put(a)
	defer pool.Put(b)
	a.Reset()
	b.Reset()
	preHiragana.WriteString(a, s)
	preKatakana.WriteString(b, unsafeString(a))
	a.Reset()
	romajiToHiragana.WriteString(a, unsafeString(b))
	b.Reset()
	romajiToKatakana.WriteString(b, unsafeString(a))
	a.Reset()
	postHiragana.WriteString(a, unsafeString(b))
	b.Reset()
	postKatakana.WriteString(b, unsafeString(a))
	a.Reset()
	postKanaSpecial.WriteString(a, unsafeString(b))
	return a.String()
}

// HiraganaToKatakana replaces a single hiragana character with the
// unicode equivalent katakana character.
func HiraganaToKatakana(r rune) rune {
	if (r >= 'ぁ' && r <= 'ゖ') || (r >= 'ゝ' && r <= 'ゞ') {
		return r + 0x60
	}
	return r
}

// KatakanaToHiragana replaces a single katakana character with the
// unicode equivalent hiragana character.
func KatakanaToHiragana(r rune) rune {
	if (r >= 'ァ' && r <= 'ヶ') || (r >= 'ヽ' && r <= 'ヾ') {
		return r - 0x60
	}
	return r
}

// IsKatakana returns true if every element of a string is katakana, except
// for characters indicated in sanitizeIsChecks (spaces and dashes).
func IsKatakana(s string) bool {
	s = sanitizeIsChecks.Replace(s)
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.In(r, unicode.Katakana) {
			return false
		}
	}

	return true
}

// IsHiragana returns true if every element of a string is hiragana, except
// for characters indicated in sanitizeIsChecks (spaces and dashes).
func IsHiragana(s string) bool {
	s = sanitizeIsChecks.Replace(s)
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.In(r, unicode.Hiragana) {
			return false
		}
	}

	return true
}

// IsKanji returns true if every element of a string is a kanji character,
// except for characters indicated in sanitizeIsChecksKanji (spaces).
func IsKanji(s string) bool {
	s = sanitizeIsChecksKanji.Replace(s)
	if s == "" {
		return false
	}

	for _, r := range s {
		if !unicode.In(r, unicode.Han) {
			return false
		}
	}

	return true
}

// ContainsKatakana returns true if a string contains any katakana characters.
func ContainsKatakana(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Katakana) {
			return true
		}
	}

	return false
}

// ContainsHiragana returns true if a string contains any hiragana characters.
func ContainsHiragana(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Hiragana) {
			return true
		}
	}

	return false
}

// ContainsKanji returns true if a string contains any kanji characters.
func ContainsKanji(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Han) {
			return true
		}
	}

	return false
}

// ExtractKanji returns a slice containing all kanji characters found in a
// string, in the order in which they were found. If a kanji exists multiple
// times in a string, then each instance of the kanji will be returned.
func ExtractKanji(s string) []string {
	k := []string{}
	for _, r := range s {
		if unicode.In(r, unicode.Han) {
			k = append(k, string(r))
		}
	}
	return k
}
