package kana

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadmeExamples(t *testing.T) {
	require.Equal(t, "hiragana", ToRomaji("ひらがな"))
	require.Equal(t, "katakana", ToRomaji("カタカナ"))
	require.Equal(t, "hiragana and katakana", ToRomaji("ひらがな and カタカナ"))

	require.Equal(t, "hiragana", ToRomajiCased("ひらがな"))
	require.Equal(t, "KATAKANA", ToRomajiCased("カタカナ"))
	require.Equal(t, "hiragana and KATAKANA", ToRomajiCased("ひらがな and カタカナ"))

	require.Equal(t, "ひらがな", ToHiragana("hiragana"))
	require.Equal(t, "ひらがな + かたかな", ToHiragana("hiragana + カタカナ"))

	require.Equal(t, "カタカナ", ToKatakana("katakana"))
	require.Equal(t, "カタカナ + ヒラガナ", ToKatakana("katakana + ひらがな"))

	require.Equal(t, "ひらがな + カタカナ", ToKana("hiragana + KATAKANA"))

	require.Equal(t, true, IsHiragana("たべる"))
	require.Equal(t, false, IsHiragana("食べる"))
	require.Equal(t, true, ContainsHiragana("たべる"))
	require.Equal(t, true, ContainsHiragana("食べる"))
	require.Equal(t, false, ContainsHiragana("カタカナ"))

	require.Equal(t, true, IsKatakana("バナナ"))
	require.Equal(t, false, IsKatakana("バナナ茶"))
	require.Equal(t, true, ContainsKatakana("バナナ"))
	require.Equal(t, true, ContainsKatakana("バナナ茶"))
	require.Equal(t, false, ContainsKatakana("ひらがな"))

	require.Equal(t, true, IsKanji("水"))
	require.Equal(t, false, IsKanji("also 茶"))
	require.Equal(t, true, ContainsKanji("食べる"))
	require.Equal(t, true, ContainsKanji("also 茶"))
	require.Equal(t, false, ContainsKanji("ひらがな + カタカナ"))

	require.Equal(t, []string{"平", "易", "日", "本", "語", "伝", "週", "刊", "放", "送", "日", "本", "語"}, ExtractKanji("また、平易な日本語で伝える週刊ニュースも放送します。日本語"))
}

func TestIsHiragana(t *testing.T) {
	tt := []struct {
		s string
		r bool
	}{
		{"ひらがな", true},      // hiragana only
		{"ひら　がな", true},     // ignore spaces
		{"ひら がな", true},     // ignore spaces
		{"ひらーがな", true},     // ignore ー
		{"カタカナ", false},     // katakana only
		{"水カタひら", false},    // mixed kanji, katakana and hiragana
		{"水abcひらがな", false}, // mixed kanji, latin and hiragana
		{"水abcカタカナ", false}, // mixed kanji, latin and katakana
		{"ひらがな一", false},    // ichi一 is not kana
		{"", false},         // empty string
		{" 　 ", false},      // just spaces
	}

	for i, v := range tt {
		require.Equal(t, v.r, IsHiragana(v.s), "testing (%d) %s = %v", i, v.s, v.r)
	}
}

func TestIsKatakana(t *testing.T) {
	tt := []struct {
		s string
		r bool
	}{
		{"カタカナ", true},      // katakana only
		{"カタ　カナ", true},     // ignore spaces
		{"カタ カナ", true},     // ignore spaces
		{"カターカナ", true},     // ignore ー
		{"ひらがな", false},     // hiragana only
		{"水カタひら", false},    // mixed kanji, katakana and hiragana
		{"水abcひらがな", false}, // mixed kanji, latin and hiragana
		{"水abcカタカナ", false}, // mixed kanji, latin and katakana
		{"カタカナ一", false},    // ichi一 is not kana
		{"", false},         // empty string
		{" 　 ", false},      // just spaces
	}

	for i, v := range tt {
		require.Equal(t, v.r, IsKatakana(v.s), "testing (%d) %s = %v", i, v.s, v.r)
	}
}
func TestIsKanji(t *testing.T) {
	tt := []struct {
		s string
		r bool
	}{
		{"水", true},         // kanji only
		{"。", false},        // punctuation only
		{"、", false},        // punctuation only
		{"「」", false},       // punctuation only
		{"カタカナ", false},     // katakana only
		{"ひらがな", false},     // hiragana only
		{"一", true},         // ichi 一  kanji
		{"水　食", true},       // ignore spaces
		{"水 食", true},       // ignore spaces
		{"水カタひら", false},    // mixed kanji, katakana and hiragana
		{"水abcひらがな", false}, // mixed kanji, latin and hiragana
		{"水abcカタカナ", false}, // mixed kanji, latin and katakana
		{"", false},         // empty string
		{" 　 ", false},      // just spaces
	}

	for i, v := range tt {
		require.Equal(t, v.r, IsKanji(v.s), "testing (%d) %s = %v", i, v.s, v.r)
	}
}

func TestContainsHiragana(t *testing.T) {
	tt := []struct {
		s string
		r bool
	}{
		{"ひらがな", true},      // hiragana only
		{"ひら　がな", true},     // ignore spaces
		{"ひら がな", true},     // ignore spaces
		{"ひらーがな", true},     // ignore ー
		{"カタカナ", false},     // katakana only
		{"水カタひら", true},     // mixed kanji, katakana and hiragana
		{"水abcひらがな", true},  // mixed kanji, latin and hiragana
		{"水abcカタカナ", false}, // mixed kanji, latin and katakana
		{"ひらがな一", true},     // ichi一 is not kana
		{"", false},         // empty string
		{" 　 ", false},      // just spaces
	}

	for i, v := range tt {
		require.Equal(t, v.r, ContainsHiragana(v.s), "testing (%d) %s = %v", i, v.s, v.r)
	}
}

func TestContainsKatakana(t *testing.T) {
	tt := []struct {
		s string
		r bool
	}{
		{"カタカナ", true},      // katakana only
		{"カタ　カナ", true},     // ignore spaces
		{"カタ カナ", true},     // ignore spaces
		{"カターカナ", true},     // ignore ー
		{"ひらがな", false},     // hiragana only
		{"水カタひら", true},     // mixed kanji, katakana and hiragana
		{"水abcひらがな", false}, // mixed kanji, latin and hiragana
		{"水abcカタカナ", true},  // mixed kanji, latin and katakana
		{"カタカナ一", true},     // ichi一 is not kana
		{"", false},         // empty string
		{" 　 ", false},      // just spaces
	}

	for i, v := range tt {
		require.Equal(t, v.r, ContainsKatakana(v.s), "testing (%d) %s = %v", i, v.s, v.r)
	}
}

func TestContainsKanji(t *testing.T) {
	tt := []struct {
		s string
		r bool
	}{
		{"水", true},        // kanji only
		{"カタカナ", false},    // katakana only
		{"ひらがな", false},    // hiragana only
		{"一", true},        // ichi 一  kanji
		{"水　食", true},      // ignore spaces
		{"水 食", true},      // ignore spaces
		{"水カタひら", true},    // mixed kanji, katakana and hiragana
		{"水abcひらがな", true}, // mixed kanji, latin and hiragana
		{"水abcカタカナ", true}, // mixed kanji, latin and katakana
		{"", false},        // empty string
		{" 　 ", false},     // just spaces
	}

	for i, v := range tt {
		require.Equal(t, v.r, ContainsKanji(v.s), "testing (%d) %s = %v", i, v.s, v.r)
	}
}

func TestExtractKanji(t *testing.T) {
	tt := []struct {
		s string
		r []string
	}{
		{"食べる", []string{"食"}},
		{"鉛筆削り", []string{"鉛", "筆", "削"}},
		{"wakareru 分かれる ", []string{"分"}},
		{"また、平易な日本語で伝える週刊ニュースも放送します。日本語", []string{"平", "易", "日", "本", "語", "伝", "週", "刊", "放", "送", "日", "本", "語"}},
	}

	for i, v := range tt {
		require.Equal(t, v.r, ExtractKanji(v.s), "testing (%d) %s = %v", i, v.s, v.r)
	}
}

func TestToHiraganaShouldConvertXxToSmall(t *testing.T) {
	tt := [][]string{
		{"xa", "ぁ"},
		{"xi", "ぃ"},
		{"xu", "ぅ"},
		{"xe", "ぇ"},
		{"xo", "ぉ"},
		{"xaxxa", "ぁっぁ"},
		{"xx", "っっ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKatakanaShouldConvertXxToSmall(t *testing.T) {
	tt := [][]string{
		{"xa", "ァ"},
		{"xi", "ィ"},
		{"xu", "ゥ"},
		{"xe", "ェ"},
		{"xo", "ォ"},
		{"xaxxa", "ァッァ"},
		{"xx", "ッッ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKatakana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToRomajiShouldConvertSmallToXx(t *testing.T) {
	tt := [][]string{
		{"ァ", "xa"},
		{"ィ", "xi"},
		{"ゥ", "xu"},
		{"ェ", "xe"},
		{"ォ", "xo"},
		{"ァッァ", "xaxxa"},
		{"ッッ", "xx"},

		{"ぁ", "xa"},
		{"ぃ", "xi"},
		{"ぅ", "xu"},
		{"ぇ", "xe"},
		{"ぉ", "xo"},
		{"ぁっぁ", "xaxxa"},
		{"っっ", "xx"},

		// Must preserve natural sequences
		{"フォト", "foto"},
		{"ふぉと", "foto"},
		{"パーティィ", "pa-tixi"}, // must find pa-tixi not pa-texixi
		{"ぱーてぃぃ", "pa-tixi"},
		{"パーティー", "pa-ti-"},
		{"ぱーてぃー", "pa-ti-"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToRomaji(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToHiraganaShouldConvertMoraicNs(t *testing.T) {
	tt := [][]string{
		{"n'a", "んあ"},
		{"n'i", "んい"},
		{"n'u", "んう"},
		{"n'e", "んえ"},
		{"n'o", "んお"},
		{"zen'in", "ぜんいん"},
		{"zennin", "ぜんにん"},
		{"nanna", "なんな"},
		{"shin'you", "しんよう"},
		{"kan'i", "かんい"},
		{"annai", "あんない"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKatakanaShouldConvertMoraicNs(t *testing.T) {
	tt := [][]string{
		{"n'a", "ンア"},
		{"n'i", "ンイ"},
		{"n'u", "ンウ"},
		{"n'e", "ンエ"},
		{"n'o", "ンオ"},
		{"zen'in", "ゼンイン"},
		{"zennin", "ゼンニン"},
		{"nanna", "ナンナ"},
		{"shin'you", "シンヨウ"},
		{"kan'i", "カンイ"},
		{"annai", "アンナイ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKatakana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToRomajiShouldConvertMoraicNs(t *testing.T) {
	tt := [][]string{
		{"んあ", "n'a"},
		{"んい", "n'i"},
		{"んう", "n'u"},
		{"んえ", "n'e"},
		{"んお", "n'o"},
		{"んや", "n'ya"},
		{"んよ", "n'yo"},
		{"んゆ", "n'yu"},
		{"ンア", "n'a"},
		{"ンイ", "n'i"},
		{"ンウ", "n'u"},
		{"ンエ", "n'e"},
		{"ンオ", "n'o"},
		{"ンヤ", "n'ya"},
		{"ンヨ", "n'yo"},
		{"ンユ", "n'yu"},
		{"ゼンイン", "zen'in"},
		{"ゼンニン", "zennin"},
		{"ナンナ", "nanna"},
		{"シンヨウ", "shin'you"},
		{"カンイ", "kan'i"},
		{"アンナイ", "annai"},
		{"きんにく", "kinniku"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToRomaji(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToHiraganaShouldConvertDoubleConsonants(t *testing.T) {
	tt := [][]string{
		{"babba", "ばっば"},
		{"dadda", "だっだ"},
		{"faffa", "ふぁっふぁ"},
		{"gagga", "がっが"},
		{"hahha", "はっは"},
		{"jajja", "じゃっじゃ"},
		{"kakka", "かっか"},
		{"pappa", "ぱっぱ"},
		{"qaqqa", "くぁっくぁ"},
		{"lalla", "らっら"}, // rarra
		{"rarra", "らっら"},
		{"sassa", "さっさ"},
		{"tatta", "たった"},
		{"vavva", "ゔぁっゔぁ"},
		{"wawwa", "わっわ"},
		{"xaxxa", "ぁっぁ"},
		{"yayya", "やっや"},
		{"zazza", "ざっざ"},
		{"maccha", "まっちゃ"},
		{"matcha", "まっちゃ"},
		{"kotchi", "こっち"},
		{"zasshi", "ざっし"},
		{"issho", "いっしょ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKatakanaShouldConvertDoubleConsonants(t *testing.T) {
	tt := [][]string{
		{"babba", "バッバ"},
		{"dadda", "ダッダ"},
		{"faffa", "ファッファ"},
		{"gagga", "ガッガ"},
		{"hahha", "ハッハ"},
		{"jajja", "ジャッジャ"},
		{"kakka", "カッカ"},
		{"pappa", "パッパ"},
		{"qaqqa", "クァックァ"},
		{"lalla", "ラッラ"}, // rarra
		{"rarra", "ラッラ"},
		{"sassa", "サッサ"},
		{"tatta", "タッタ"},
		{"vavva", "ヴァッヴァ"},
		{"wawwa", "ワッワ"},
		{"xaxxa", "ァッァ"},
		{"yayya", "ヤッヤ"},
		{"zazza", "ザッザ"},
		{"maccha", "マッチャ"},
		{"matcha", "マッチャ"},
		{"kotchi", "コッチ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKatakana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToRomajiShouldConvertDoubleConsonants(t *testing.T) {
	tt := [][]string{
		{"ばっば", "babba"},
		{"だっだ", "dadda"},
		{"ふぁっふぁ", "faffa"},
		{"がっが", "gagga"},
		{"はっは", "hahha"},
		{"じゃっじゃ", "jajja"},
		{"かっか", "kakka"},
		{"ぱっぱ", "pappa"},
		{"くぁっくぁ", "kwakkwa"},
		{"らっら", "rarra"},
		{"さっさ", "sassa"},
		{"たった", "tatta"},
		{"ゔぁっゔぁ", "vavva"},
		{"わっわ", "wawwa"},
		{"ぁっぁ", "xaxxa"},
		{"やっや", "yayya"},
		{"ざっざ", "zazza"},
		{"まっちゃ", "matcha"},
		{"こっち", "kotchi"},
		{"ざっし", "zasshi"},
		{"いっしょ", "issho"},

		{"バッバ", "babba"},
		{"ダッダ", "dadda"},
		{"ファッファ", "faffa"},
		{"ガッガ", "gagga"},
		{"ハッハ", "hahha"},
		{"ジャッジャ", "jajja"},
		{"カッカ", "kakka"},
		{"パッパ", "pappa"},
		{"クァックァ", "kwakkwa"},
		{"ラッラ", "rarra"}, // rarra
		{"サッサ", "sassa"},
		{"タッタ", "tatta"},
		{"ヴァッヴァ", "vavva"},
		{"ワッワ", "wawwa"},
		{"ァッァ", "xaxxa"},
		{"ヤッヤ", "yayya"},
		{"ザッザ", "zazza"},
		{"マッチャ", "matcha"},
		{"コッチ", "kotchi"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToRomaji(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKatakanaShouldConvertAnyLetterCase(t *testing.T) {
	tt := [][]string{
		{"ke-susenshitibu", "ケースセンシティブ"},
		{"KE-SUSENSHITIBU", "ケースセンシティブ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKatakana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToHiraganaShouldConvertAnyLetterCase(t *testing.T) {
	tt := [][]string{
		{"ke-susenshitibu", "けーすせんしてぃぶ"},
		{"KE-SUSENSHITIBU", "けーすせんしてぃぶ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKanaCaseShouldConvertLowerToHiraAndUpperToKataRespectively(t *testing.T) {
	tt := [][]string{
		{"ke-susenshitibu", "けーすせんしてぃぶ"},
		{"KE-SUSENSHITIBU", "ケースセンシティブ"},
		{"betsu KE-SU", "べつ ケース"},
		{"kiiro BANANA", "きいろ バナナ"},
		{"OnaJI", "オなジ"},
		{"OnaJi", "オなJい"}, // incomplete kana
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToHiraganaShouldSwapToKatakana(t *testing.T) {
	tt := [][]string{
		{"バナナ tabete", "ばなな たべて"},
		{"プディング べつばら", "ぷでぃんぐ べつばら"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKatakanaShouldSwapToHiragana(t *testing.T) {
	tt := [][]string{
		{"ばなな tabete", "バナナ タベテ"},
		{"ぷでぃんぐ ベツバラ", "プディング ベツバラ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKatakana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKanaShouldNotSwapKanas(t *testing.T) {
	tt := [][]string{
		{"ばなな TABETE", "ばなな タベテ"},
		{"バナナ tabete", "バナナ たべて"},
		{"ぷでぃんぐ ベツバラ", "ぷでぃんぐ ベツバラ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}
func TestPostRomajiShouldConvertSpecialCharacters(t *testing.T) {
	tt := [][]string{
		{"ー", "-"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToRomaji(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestPostKanaShouldConvertSpecialCharacters(t *testing.T) {
	tt := [][]string{
		{"–", "ー"},
		{"-", "ー"},
		{"'", ""},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing to hiragana (%d) %s = %s", i, v[0], v[1])
		require.Equal(t, v[1], ToKatakana(v[0]), "testing to katakana (%d) %s = %s", i, v[0], v[1])
		require.Equal(t, v[1], ToKana(v[0]), "testing to kana (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToRomajiShouldPreserveNonKanaCharacters(t *testing.T) {
	tt := [][]string{
		{"～（ん）だろう", "～（n）darou"},
		{"カケル（めがねを)", "kakeru（meganewo)"},
		{"カケル（めがねを)", "kakeru（meganewo)"},
		{"ガラスの器", "garasuno器"},
		{"一番 いちばん", "一番 ichiban"},
		{"食べ物", "食be物"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToRomaji(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

var toHiraganaBasicSequences = [][]string{
	{"taberu", "たべる"},
	{"anata", "あなた"},
	{"choushoku", "ちょうしょく"},
	{"nani", "なに"},
	{"jidouhanbaiki", "じどうはんばいき"},
	{"zen'in", "ぜんいん"},
	{"kitte", "きって"},
	{"matcha", "まっちゃ"},
	{"sassato", "さっさと"},
	{"kotchi", "こっち"},
	{"eki", "えき"},
	{"shokubutsu", "しょくぶつ"},
	{"mizuumi", "みずうみ"},
}

var toKatakanaBasicSequences = [][]string{
	{"sande-", "サンデー"},
	{"baree", "バレエ"},
	{"miira", "ミイラ"},
	{"souru", "ソウル"},
	{"se-ra-", "セーラー"},
	{"takushi-", "タクシー"},
	{"konku-ru", "コンクール"},
	{"bare-bo-ru", "バレーボール"},
	{"so-ru", "ソール"},
	{"pa-tishipe-shonpuroguramu", "パーティシペーションプログラム"},
	{"metafa-", "メタファー"},
	{"purofi-ru", "プロフィール"},
	{"mi-tingu", "ミーティング"},
	{"ko-hi-", "コーヒー"},
}

func TestToRomajiShouldConvertBasicSequence(t *testing.T) {
	for i, v := range toHiraganaBasicSequences {
		require.Equal(t, v[0], ToRomaji(v[1]), "testing (%d) %s = %s", i, v[1], v[0])
	}
	for i, v := range toKatakanaBasicSequences {
		require.Equal(t, v[0], ToRomaji(v[1]), "testing (%d) %s = %s", i, v[1], v[0])
	}
}

func TestToRomajiCasedShouldConvertBasicSequence(t *testing.T) {
	for i, v := range toHiraganaBasicSequences {
		require.Equal(t, strings.ToLower(v[0]), ToRomajiCased(v[1]), "testing (%d) %s = %s", i, v[1], v[0])
	}
	for i, v := range toKatakanaBasicSequences {
		require.Equal(t, strings.ToUpper(v[0]), ToRomajiCased(v[1]), "testing (%d) %s = %s", i, v[1], v[0])
	}
}

func TestToHiraganaShouldConvertBasicSequence(t *testing.T) {
	for i, v := range toHiraganaBasicSequences {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKatakanaShouldConvertBasicSequence(t *testing.T) {
	for i, v := range toKatakanaBasicSequences {
		require.Equal(t, v[1], ToKatakana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestSwapHiraganaKatakana(t *testing.T) {
	tt := [][]string{
		{"か", "カ"},
		{"き", "キ"},
		{"く", "ク"},
		{"け", "ケ"},
		{"こ", "コ"},
		{"さ", "サ"},
		{"し", "シ"},
		{"す", "ス"},
		{"せ", "セ"},
		{"そ", "ソ"},
		{"た", "タ"},
		{"ち", "チ"},
		{"つ", "ツ"},
		{"て", "テ"},
		{"と", "ト"},
		{"な", "ナ"},
		{"に", "ニ"},
		{"ぬ", "ヌ"},
		{"ね", "ネ"},
		{"の", "ノ"},
		{"は", "ハ"},
		{"ひ", "ヒ"},
		{"ふ", "フ"},
		{"へ", "ヘ"},
		{"ほ", "ホ"},
		{"ま", "マ"},
		{"み", "ミ"},
		{"む", "ム"},
		{"め", "メ"},
		{"も", "モ"},
		{"や", "ヤ"},
		{"ゆ", "ユ"},
		{"よ", "ヨ"},
		{"ら", "ラ"},
		{"り", "リ"},
		{"る", "ル"},
		{"れ", "レ"},
		{"ろ", "ロ"},
		{"わ", "ワ"},
		{"を", "ヲ"},
		{"が", "ガ"},
		{"ぎ", "ギ"},
		{"ぐ", "グ"},
		{"げ", "ゲ"},
		{"ご", "ゴ"},
		{"ざ", "ザ"},
		{"じ", "ジ"},
		{"ず", "ズ"},
		{"ぜ", "ゼ"},
		{"ぞ", "ゾ"},
		{"だ", "ダ"},
		{"ぢ", "ヂ"},
		{"づ", "ヅ"},
		{"で", "デ"},
		{"ど", "ド"},
		{"ば", "バ"},
		{"び", "ビ"},
		{"ぶ", "ブ"},
		{"べ", "ベ"},
		{"ぼ", "ボ"},
		{"ぱ", "パ"},
		{"ぴ", "ピ"},
		{"ぷ", "プ"},
		{"ぺ", "ペ"},
		{"ぽ", "ポ"},
		{"きゃ", "キャ"},
		{"きゅ", "キュ"},
		{"きょ", "キョ"},
		{"しゃ", "シャ"},
		{"しゅ", "シュ"},
		{"しょ", "ショ"},
		{"ちゃ", "チャ"},
		{"ちゅ", "チュ"},
		{"ちょ", "チョ"},
		{"にゃ", "ニャ"},
		{"にゅ", "ニュ"},
		{"にょ", "ニョ"},
		{"ひゃ", "ヒャ"},
		{"ひゅ", "ヒュ"},
		{"ひょ", "ヒョ"},
		{"みゃ", "ミャ"},
		{"みゅ", "ミュ"},
		{"みょ", "ミョ"},
		{"りゃ", "リャ"},
		{"りゅ", "リュ"},
		{"りょ", "リョ"},
		{"ぎゃ", "ギャ"},
		{"ぎゅ", "ギュ"},
		{"ぎょ", "ギョ"},
		{"じゃ", "ジャ"},
		{"じゅ", "ジュ"},
		{"じょ", "ジョ"},
		{"ぢゃ", "ヂャ"},
		{"ぢゅ", "ヂュ"},
		{"ぢょ", "ヂョ"},
		{"びゃ", "ビャ"},
		{"びゅ", "ビュ"},
		{"びょ", "ビョ"},
		{"ぴゃ", "ピャ"},
		{"ぴゅ", "ピュ"},
		{"ぴょ", "ピョ"},
		{"いぃ", "イィ"},
		{"いぇ", "イェ"},
		{"ゐ", "ヰ"},
		{"うぅ", "ウゥ"},
		{"うぇ", "ウェ"},
		{"うゅ", "ウュ"},
		{"ゔぁ", "ヴァ"},
		{"ゔぃ", "ヴィ"},
		{"ゔ", "ヴ"},
		{"ゔぇ", "ヴェ"},
		{"ゔぉ", "ヴォ"},
		{"ゔゃ", "ヴャ"},
		{"ゔゅ", "ヴュ"},
		{"ゔぃぇ", "ヴィェ"},
		{"ゔょ", "ヴョ"},
		{"きぇ", "キェ"},
		{"ぎぇ", "ギェ"},
		{"くぁ", "クァ"},
		{"くぃ", "クィ"},
		{"くぇ", "クェ"},
		{"くぅ", "クゥ"},
		{"くぉ", "クォ"},
		{"ぐぁ", "グァ"},
		{"ぐぃ", "グィ"},
		{"ぐぇ", "グェ"},
		{"ぐぉ", "グォ"},
		{"ぐぅ", "グゥ"},
		{"しぇ", "シェ"},
		{"じぇ", "ジェ"},
		{"すぃ", "スィ"},
		{"ずぃ", "ズィ"},
		{"ちぇ", "チェ"},
		{"つぁ", "ツァ"},
		{"つぇ", "ツェ"},
		{"つぃ", "ツィ"},
		{"つぉ", "ツォ"},
		{"つゅ", "ツュ"},
		{"てぃ", "ティ"},
		{"とぅ", "トゥ"},
		{"にぇ", "ニェ"},
		{"ひぇ", "ヒェ"},
		{"びぇ", "ビェ"},
		{"ぴぇ", "ピェ"},
		{"ふぁ", "ファ"},
		{"ふぃ", "フィ"},
		{"ふぇ", "フェ"},
		{"ふぉ", "フォ"},
		{"ふゃ", "フャ"},
		{"ふゅ", "フュ"},
		{"ふょ", "フョ"},
		{"ふぅ", "フゥ"},
		{"みぇ", "ミェ"},
		{"りぇ", "リェ"},
		{"ら", "ラ"},
		{"り", "リ"},
		{"る", "ル"},
		{"れ", "レ"},
		{"くぁ", "クァ"},
		{"くぃ", "クィ"},
		{"くぇ", "クェ"},
		{"くぉ", "クォ"},
		{"くぅ", "クゥ"},
		{"あ", "ア"},
		{"い", "イ"},
		{"う", "ウ"},
		{"え", "エ"},
		{"お", "オ"},
		{"ん", "ン"},
		{"ぁ", "ァ"},
		{"ぃ", "ィ"},
		{"ぅ", "ゥ"},
		{"ぇ", "ェ"},
		{"ぉ", "ォ"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], strings.Map(HiraganaToKatakana, v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}

	for i, v := range tt {
		require.Equal(t, v[0], strings.Map(KatakanaToHiragana, v[1]), "testing (%d) %s = %s", i, v[1], v[0])
	}
}

func TestToHiraganaEquivalents(t *testing.T) {
	tt := [][]string{
		{"ka", "か"},
		{"ki", "き"},
		{"ku", "く"},
		{"ke", "け"},
		{"ko", "こ"},
		{"sa", "さ"},
		{"shi", "し"},
		{"su", "す"},
		{"se", "せ"},
		{"so", "そ"},
		{"ta", "た"},
		{"chi", "ち"},
		{"tsu", "つ"},
		{"te", "て"},
		{"to", "と"},
		{"na", "な"},
		{"ni", "に"},
		{"nu", "ぬ"},
		{"ne", "ね"},
		{"no", "の"},
		{"ha", "は"},
		{"hi", "ひ"},
		{"fu", "ふ"},
		{"he", "へ"},
		{"ho", "ほ"},
		{"ma", "ま"},
		{"mi", "み"},
		{"mu", "む"},
		{"me", "め"},
		{"mo", "も"},
		{"ya", "や"},
		{"yu", "ゆ"},
		{"yo", "よ"},
		{"ra", "ら"},
		{"ri", "り"},
		{"ru", "る"},
		{"re", "れ"},
		{"ro", "ろ"},
		{"wa", "わ"},
		{"wo", "を"},
		{"ga", "が"},
		{"gi", "ぎ"},
		{"gu", "ぐ"},
		{"ge", "げ"},
		{"go", "ご"},
		{"za", "ざ"},
		{"ji", "じ"},
		{"zu", "ず"},
		{"ze", "ぜ"},
		{"zo", "ぞ"},
		{"da", "だ"},
		{"di", "ぢ"},
		{"du", "づ"},
		{"de", "で"},
		{"do", "ど"},
		{"ba", "ば"},
		{"bi", "び"},
		{"bu", "ぶ"},
		{"be", "べ"},
		{"bo", "ぼ"},
		{"pa", "ぱ"},
		{"pi", "ぴ"},
		{"pu", "ぷ"},
		{"pe", "ぺ"},
		{"po", "ぽ"},
		{"kya", "きゃ"},
		{"kyu", "きゅ"},
		{"kyo", "きょ"},
		{"sha", "しゃ"},
		{"shu", "しゅ"},
		{"sho", "しょ"},
		{"cha", "ちゃ"},
		{"chu", "ちゅ"},
		{"cho", "ちょ"},
		{"nya", "にゃ"},
		{"nyu", "にゅ"},
		{"nyo", "にょ"},
		{"hya", "ひゃ"},
		{"hyu", "ひゅ"},
		{"hyo", "ひょ"},
		{"mya", "みゃ"},
		{"myu", "みゅ"},
		{"myo", "みょ"},
		{"rya", "りゃ"},
		{"ryu", "りゅ"},
		{"ryo", "りょ"},
		{"gya", "ぎゃ"},
		{"gyu", "ぎゅ"},
		{"gyo", "ぎょ"},
		{"ja", "じゃ"},
		{"ju", "じゅ"},
		{"jo", "じょ"},
		{"jya", "じゃ"},
		{"jyu", "じゅ"},
		{"jyo", "じょ"},
		{"dya", "ぢゃ"},
		{"dyu", "ぢゅ"},
		{"dyo", "ぢょ"},
		{"bya", "びゃ"},
		{"byu", "びゅ"},
		{"byo", "びょ"},
		{"pya", "ぴゃ"},
		{"pyu", "ぴゅ"},
		{"pyo", "ぴょ"},
		{"yi", "いぃ"},
		{"ye", "いぇ"},
		{"wi", "ゐ"},
		{"wu", "うぅ"},
		{"we", "うぇ"},
		{"wyu", "うゅ"},
		{"va", "ゔぁ"},
		{"vi", "ゔぃ"},
		{"vu", "ゔ"},
		{"ve", "ゔぇ"},
		{"vo", "ゔぉ"},
		{"vya", "ゔゃ"},
		{"vyu", "ゔゅ"},
		{"vye", "ゔぃぇ"},
		{"vyo", "ゔょ"},
		{"kye", "きぇ"},
		{"gye", "ぎぇ"},
		{"kwa", "くぁ"},
		{"kwi", "くぃ"},
		{"kwe", "くぇ"},
		{"kwu", "くぅ"},
		{"kwo", "くぉ"},
		{"gwa", "ぐぁ"},
		{"gwi", "ぐぃ"},
		{"gwe", "ぐぇ"},
		{"gwo", "ぐぉ"},
		{"gwu", "ぐぅ"},
		{"she", "しぇ"},
		{"je", "じぇ"},
		{"si", "すぃ"},
		{"zi", "ずぃ"},
		{"che", "ちぇ"},
		{"tsa", "つぁ"},
		{"tse", "つぇ"},
		{"tsi", "つぃ"},
		{"tso", "つぉ"},
		{"tsyu", "つゅ"},
		{"ti", "てぃ"},
		{"tu", "とぅ"},
		{"tyu", "ちゅ"},
		{"nye", "にぇ"},
		{"hye", "ひぇ"},
		{"bye", "びぇ"},
		{"pye", "ぴぇ"},
		{"fa", "ふぁ"},
		{"fi", "ふぃ"},
		{"fe", "ふぇ"},
		{"fo", "ふぉ"},
		{"fya", "ふゃ"},
		{"fyu", "ふゅ"},
		{"fye", "ふぇ"},
		{"fyo", "ふょ"},
		{"hu", "ふぅ"},
		{"mye", "みぇ"},
		{"rye", "りぇ"},
		{"la", "ら"},
		{"li", "り"},
		{"lu", "る"},
		{"le", "れ"},
		{"lo", "ろ"},
		{"qa", "くぁ"},
		{"qi", "くぃ"},
		{"qe", "くぇ"},
		{"qo", "くぉ"},
		{"qu", "くぅ"},
		{"a", "あ"},
		{"i", "い"},
		{"u", "う"},
		{"e", "え"},
		{"o", "お"},
		{"n", "ん"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToHiragana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToKatakanaEquivalents(t *testing.T) {
	tt := [][]string{
		{"ka", "カ"},
		{"ki", "キ"},
		{"ku", "ク"},
		{"ke", "ケ"},
		{"ko", "コ"},
		{"sa", "サ"},
		{"shi", "シ"},
		{"su", "ス"},
		{"se", "セ"},
		{"so", "ソ"},
		{"ta", "タ"},
		{"chi", "チ"},
		{"tsu", "ツ"},
		{"te", "テ"},
		{"to", "ト"},
		{"na", "ナ"},
		{"ni", "ニ"},
		{"nu", "ヌ"},
		{"ne", "ネ"},
		{"no", "ノ"},
		{"ha", "ハ"},
		{"hi", "ヒ"},
		{"fu", "フ"},
		{"he", "ヘ"},
		{"ho", "ホ"},
		{"ma", "マ"},
		{"mi", "ミ"},
		{"mu", "ム"},
		{"me", "メ"},
		{"mo", "モ"},
		{"ya", "ヤ"},
		{"yu", "ユ"},
		{"yo", "ヨ"},
		{"ra", "ラ"},
		{"ri", "リ"},
		{"ru", "ル"},
		{"re", "レ"},
		{"ro", "ロ"},
		{"wa", "ワ"},
		{"wo", "ヲ"},
		{"ga", "ガ"},
		{"gi", "ギ"},
		{"gu", "グ"},
		{"ge", "ゲ"},
		{"go", "ゴ"},
		{"za", "ザ"},
		{"ji", "ジ"},
		{"zu", "ズ"},
		{"ze", "ゼ"},
		{"zo", "ゾ"},
		{"da", "ダ"},
		{"di", "ヂ"},
		{"du", "ヅ"},
		{"de", "デ"},
		{"do", "ド"},
		{"ba", "バ"},
		{"bi", "ビ"},
		{"bu", "ブ"},
		{"be", "ベ"},
		{"bo", "ボ"},
		{"pa", "パ"},
		{"pi", "ピ"},
		{"pu", "プ"},
		{"pe", "ペ"},
		{"po", "ポ"},
		{"kya", "キャ"},
		{"kyu", "キュ"},
		{"kyo", "キョ"},
		{"sha", "シャ"},
		{"shu", "シュ"},
		{"sho", "ショ"},
		{"cha", "チャ"},
		{"chu", "チュ"},
		{"cho", "チョ"},
		{"nya", "ニャ"},
		{"nyu", "ニュ"},
		{"nyo", "ニョ"},
		{"hya", "ヒャ"},
		{"hyu", "ヒュ"},
		{"hyo", "ヒョ"},
		{"mya", "ミャ"},
		{"myu", "ミュ"},
		{"myo", "ミョ"},
		{"rya", "リャ"},
		{"ryu", "リュ"},
		{"ryo", "リョ"},
		{"gya", "ギャ"},
		{"gyu", "ギュ"},
		{"gyo", "ギョ"},
		{"ja", "ジャ"},
		{"ju", "ジュ"},
		{"jo", "ジョ"},
		{"jya", "ジャ"},
		{"jyu", "ジュ"},
		{"jyo", "ジョ"},
		{"dya", "ヂャ"},
		{"dyu", "ヂュ"},
		{"dyo", "ヂョ"},
		{"bya", "ビャ"},
		{"byu", "ビュ"},
		{"byo", "ビョ"},
		{"pya", "ピャ"},
		{"pyu", "ピュ"},
		{"pyo", "ピョ"},
		{"yi", "イィ"},
		{"ye", "イェ"},
		{"wi", "ウィ"},
		{"wu", "ウゥ"},
		{"we", "ウェ"},
		{"wyu", "ウュ"},
		{"va", "ヴァ"},
		{"vi", "ヴィ"},
		{"vu", "ヴ"},
		{"ve", "ヴェ"},
		{"vo", "ヴォ"},
		{"vya", "ヴャ"},
		{"vyu", "ヴュ"},
		{"vye", "ヴィェ"},
		{"vyo", "ヴョ"},
		{"kye", "キェ"},
		{"gye", "ギェ"},
		{"kwa", "クァ"},
		{"kwi", "クィ"},
		{"kwe", "クェ"},
		{"kwu", "クゥ"},
		{"kwo", "クォ"},
		{"gwa", "グァ"},
		{"gwi", "グィ"},
		{"gwe", "グェ"},
		{"gwo", "グォ"},
		{"gwu", "グゥ"},
		{"she", "シェ"},
		{"je", "ジェ"},
		{"si", "スィ"},
		{"zi", "ズィ"},
		{"che", "チェ"},
		{"tsa", "ツァ"},
		{"tse", "ツェ"},
		{"tsi", "ツィ"},
		{"tso", "ツォ"},
		{"tsyu", "ツュ"},
		{"ti", "ティ"},
		{"tu", "トゥ"},
		{"tyu", "テュ"},
		{"nye", "ニェ"},
		{"hye", "ヒェ"},
		{"bye", "ビェ"},
		{"pye", "ピェ"},
		{"fa", "ファ"},
		{"fi", "フィ"},
		{"fe", "フェ"},
		{"fo", "フォ"},
		{"fya", "フャ"},
		{"fyu", "フュ"},
		{"fye", "フィェ"},
		{"fyo", "フョ"},
		{"hu", "ホゥ"},
		{"mye", "ミェ"},
		{"rye", "リェ"},
		{"la", "ラ"},
		{"li", "リ"},
		{"lu", "ル"},
		{"le", "レ"},
		{"lo", "ロ"},
		{"qa", "クァ"},
		{"qi", "クィ"},
		{"qe", "クェ"},
		{"qo", "クォ"},
		{"qu", "クヮ"},
		{"a", "ア"},
		{"i", "イ"},
		{"u", "ウ"},
		{"e", "エ"},
		{"o", "オ"},
		{"n", "ン"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToKatakana(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}

func TestToRomajiEquivalents(t *testing.T) {
	tt := [][]string{
		{"か", "ka"},
		{"き", "ki"},
		{"く", "ku"},
		{"け", "ke"},
		{"こ", "ko"},
		{"さ", "sa"},
		{"し", "shi"},
		{"す", "su"},
		{"せ", "se"},
		{"そ", "so"},
		{"た", "ta"},
		{"ち", "chi"},
		{"つ", "tsu"},
		{"て", "te"},
		{"と", "to"},
		{"な", "na"},
		{"に", "ni"},
		{"ぬ", "nu"},
		{"ね", "ne"},
		{"の", "no"},
		{"は", "ha"},
		{"ひ", "hi"},
		{"ふ", "fu"},
		{"へ", "he"},
		{"ほ", "ho"},
		{"ま", "ma"},
		{"み", "mi"},
		{"む", "mu"},
		{"め", "me"},
		{"も", "mo"},
		{"や", "ya"},
		{"ゆ", "yu"},
		{"よ", "yo"},
		{"ら", "ra"},
		{"り", "ri"},
		{"る", "ru"},
		{"れ", "re"},
		{"ろ", "ro"},
		{"わ", "wa"},
		{"を", "wo"},
		{"が", "ga"},
		{"ぎ", "gi"},
		{"ぐ", "gu"},
		{"げ", "ge"},
		{"ご", "go"},
		{"ざ", "za"},
		{"じ", "ji"},
		{"ず", "zu"},
		{"ぜ", "ze"},
		{"ぞ", "zo"},
		{"だ", "da"},
		{"ぢ", "di"},
		{"づ", "du"},
		{"で", "de"},
		{"ど", "do"},
		{"ば", "ba"},
		{"び", "bi"},
		{"ぶ", "bu"},
		{"べ", "be"},
		{"ぼ", "bo"},
		{"ぱ", "pa"},
		{"ぴ", "pi"},
		{"ぷ", "pu"},
		{"ぺ", "pe"},
		{"ぽ", "po"},
		{"きゃ", "kya"},
		{"きゅ", "kyu"},
		{"きょ", "kyo"},
		{"しゃ", "sha"},
		{"しゅ", "shu"},
		{"しょ", "sho"},
		{"ちゃ", "cha"},
		{"ちゅ", "chu"},
		{"ちょ", "cho"},
		{"にゃ", "nya"},
		{"にゅ", "nyu"},
		{"にょ", "nyo"},
		{"ひゃ", "hya"},
		{"ひゅ", "hyu"},
		{"ひょ", "hyo"},
		{"みゃ", "mya"},
		{"みゅ", "myu"},
		{"みょ", "myo"},
		{"りゃ", "rya"},
		{"りゅ", "ryu"},
		{"りょ", "ryo"},
		{"ぎゃ", "gya"},
		{"ぎゅ", "gyu"},
		{"ぎょ", "gyo"},
		{"じゃ", "ja"},
		{"じゅ", "ju"},
		{"じょ", "jo"},
		{"ぢゃ", "dya"},
		{"ぢゅ", "dyu"},
		{"ぢょ", "dyo"},
		{"びゃ", "bya"},
		{"びゅ", "byu"},
		{"びょ", "byo"},
		{"ぴゃ", "pya"},
		{"ぴゅ", "pyu"},
		{"ぴょ", "pyo"},
		{"いぃ", "yi"},
		{"いぇ", "ye"},
		{"ゐ", "wi"},
		{"うぅ", "wu"},
		{"うぇ", "we"},
		{"うゅ", "wyu"},
		{"ゔぁ", "va"},
		{"ゔぃ", "vi"},
		{"ゔ", "vu"},
		{"ゔぇ", "ve"},
		{"ゔぉ", "vo"},
		{"ゔゃ", "vya"},
		{"ゔゅ", "vyu"},
		{"ゔぃぇ", "vye"},
		{"ゔょ", "vyo"},
		{"きぇ", "kye"},
		{"ぎぇ", "gye"},
		{"くぁ", "kwa"},
		{"くぃ", "kwi"},
		{"くぇ", "kwe"},
		{"くぅ", "kwu"},
		{"くぉ", "kwo"},
		{"ぐぁ", "gwa"},
		{"ぐぃ", "gwi"},
		{"ぐぇ", "gwe"},
		{"ぐぉ", "gwo"},
		{"ぐぅ", "gwu"},
		{"しぇ", "she"},
		{"じぇ", "je"},
		{"すぃ", "si"},
		{"ずぃ", "zi"},
		{"ちぇ", "che"},
		{"つぁ", "tsa"},
		{"つぇ", "tse"},
		{"つぃ", "tsi"},
		{"つぉ", "tso"},
		{"つゅ", "tsyu"},
		{"てぃ", "ti"},
		{"とぅ", "tu"},
		{"にぇ", "nye"},
		{"ひぇ", "hye"},
		{"びぇ", "bye"},
		{"ぴぇ", "pye"},
		{"ふぁ", "fa"},
		{"ふぃ", "fi"},
		{"ふぇ", "fe"},
		{"ふぉ", "fo"},
		{"ふゃ", "fya"},
		{"ふゅ", "fyu"},
		{"ふょ", "fyo"},
		{"ふぅ", "hu"},
		{"みぇ", "mye"},
		{"りぇ", "rye"},
		{"あ", "a"},
		{"い", "i"},
		{"う", "u"},
		{"え", "e"},
		{"お", "o"},
		{"ん", "n"},
		{"ぁ", "xa"},
		{"ぃ", "xi"},
		{"ぅ", "xu"},
		{"ぇ", "xe"},
		{"ぉ", "xo"},

		{"キャ", "kya"},
		{"キュ", "kyu"},
		{"キョ", "kyo"},
		{"シャ", "sha"},
		{"シュ", "shu"},
		{"ショ", "sho"},
		{"チャ", "cha"},
		{"チュ", "chu"},
		{"チョ", "cho"},
		{"ニャ", "nya"},
		{"ニュ", "nyu"},
		{"ニョ", "nyo"},
		{"ヒャ", "hya"},
		{"ヒュ", "hyu"},
		{"ヒョ", "hyo"},
		{"ミャ", "mya"},
		{"ミュ", "myu"},
		{"ミョ", "myo"},
		{"リャ", "rya"},
		{"リュ", "ryu"},
		{"リョ", "ryo"},
		{"ギャ", "gya"},
		{"ギュ", "gyu"},
		{"ギョ", "gyo"},
		{"ジャ", "ja"},
		{"ジュ", "ju"},
		{"ジョ", "jo"},
		{"ヂャ", "dya"},
		{"ヂュ", "dyu"},
		{"ヂョ", "dyo"},
		{"ビャ", "bya"},
		{"ビュ", "byu"},
		{"ビョ", "byo"},
		{"ピャ", "pya"},
		{"ピュ", "pyu"},
		{"ピョ", "pyo"},
		{"イィ", "yi"},
		{"イェ", "ye"},
		{"ウゥ", "wu"},
		{"ウェ", "we"},
		{"ウュ", "wyu"},
		{"ヴァ", "va"},
		{"ヴィ", "vi"},
		{"ヴ", "vu"},
		{"ヴェ", "ve"},
		{"ヴォ", "vo"},
		{"ヴャ", "vya"},
		{"ヴュ", "vyu"},
		{"ヴィェ", "vye"},
		{"ヴョ", "vyo"},
		{"キェ", "kye"},
		{"ギェ", "gye"},
		{"クァ", "kwa"},
		{"クィ", "kwi"},
		{"クェ", "kwe"},
		{"クゥ", "kwu"},
		{"クォ", "kwo"},
		{"グァ", "gwa"},
		{"グィ", "gwi"},
		{"グェ", "gwe"},
		{"グォ", "gwo"},
		{"グゥ", "gwu"},
		{"シェ", "she"},
		{"ジェ", "je"},
		{"スィ", "si"},
		{"ズィ", "zi"},
		{"チェ", "che"},
		{"ツァ", "tsa"},
		{"ツェ", "tse"},
		{"ツィ", "tsi"},
		{"ツォ", "tso"},
		{"ツュ", "tsyu"},
		{"ティ", "ti"},
		{"トゥ", "tu"},
		{"ニェ", "nye"},
		{"ヒェ", "hye"},
		{"ビェ", "bye"},
		{"ピェ", "pye"},
		{"ファ", "fa"},
		{"フィ", "fi"},
		{"フェ", "fe"},
		{"フォ", "fo"},
		{"フャ", "fya"},
		{"フュ", "fyu"},
		{"フョ", "fyo"},
		{"ホゥ", "hu"},
		{"ミェ", "mye"},
		{"リェ", "rye"},
		{"カ", "ka"},
		{"キ", "ki"},
		{"ク", "ku"},
		{"ケ", "ke"},
		{"コ", "ko"},
		{"サ", "sa"},
		{"シ", "shi"},
		{"ス", "su"},
		{"セ", "se"},
		{"ソ", "so"},
		{"タ", "ta"},
		{"チ", "chi"},
		{"ツ", "tsu"},
		{"テ", "te"},
		{"ト", "to"},
		{"ナ", "na"},
		{"ニ", "ni"},
		{"ヌ", "nu"},
		{"ネ", "ne"},
		{"ノ", "no"},
		{"ハ", "ha"},
		{"ヒ", "hi"},
		{"フ", "fu"},
		{"ヘ", "he"},
		{"ホ", "ho"},
		{"マ", "ma"},
		{"ミ", "mi"},
		{"ム", "mu"},
		{"メ", "me"},
		{"モ", "mo"},
		{"ヤ", "ya"},
		{"ユ", "yu"},
		{"ヨ", "yo"},
		{"ラ", "ra"},
		{"リ", "ri"},
		{"ル", "ru"},
		{"レ", "re"},
		{"ロ", "ro"},
		{"ワ", "wa"},
		{"ヲ", "wo"},
		{"ガ", "ga"},
		{"ギ", "gi"},
		{"グ", "gu"},
		{"ゲ", "ge"},
		{"ゴ", "go"},
		{"ザ", "za"},
		{"ジ", "ji"},
		{"ズ", "zu"},
		{"ゼ", "ze"},
		{"ゾ", "zo"},
		{"ダ", "da"},
		{"ヂ", "di"},
		{"ヅ", "du"},
		{"デ", "de"},
		{"ド", "do"},
		{"バ", "ba"},
		{"ビ", "bi"},
		{"ブ", "bu"},
		{"ベ", "be"},
		{"ボ", "bo"},
		{"パ", "pa"},
		{"ピ", "pi"},
		{"プ", "pu"},
		{"ペ", "pe"},
		{"ポ", "po"},
		{"ウィ", "wi"},
		{"ア", "a"},
		{"イ", "i"},
		{"ウ", "u"},
		{"エ", "e"},
		{"オ", "o"},
		{"ン", "n"},
		{"ァ", "xa"},
		{"ィ", "xi"},
		{"ゥ", "xu"},
		{"ェ", "xe"},
		{"ォ", "xo"},
	}

	for i, v := range tt {
		require.Equal(t, v[1], ToRomaji(v[0]), "testing (%d) %s = %s", i, v[0], v[1])
	}
}
