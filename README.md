# Kana Tools
### A Go library for Modified-Hepburn Wapuro Romaji, Katakana, and Hiragana Detection and Conversion

Kana Tools provides Romaji ←→ Kana transliteration based on a [Wāpuro rōmaji (ワープロローマ字)](https://en.wikipedia.org/wiki/Wāpuro_rōmaji) implementation of [Modified (Revised) Hepburn Romanization](https://en.wikipedia.org/wiki/Hepburn_romanization).

Where possible, the library uses static approach rather than computational approach in order to perform conversions, relying on order-of-operations to ensure the correct output and provide a higher degree of wapuro conformity and maintainability.


### Usage
```go
import "github.com/mochi-co/kana-tools"
```

```go
// Convert Hiragana and Katakana to Romaji
kana.ToRomaji("ひらがな") // -> "hiragana"
kana.ToRomaji("カタカナ") // -> "katakana"
kana.ToRomaji("ひらがな and カタカナ") // -> "hiragana and katakana"
```

```go
// Convert Hiragana and Katakana to Cased Romaji
kana.ToRomajiCased("ひらがな") // -> "hiragana"
kana.ToRomajiCased("カタカナ") // -> "KATAKANA"
kana.ToRomajiCased("ひらがな and カタカナ") // -> "hiragana and KATAKANA"
```

```go
// Convert Romaji and Katakana to Hiragana
kana.ToHiragana("hiragana") // -> "ひらがな"
kana.ToHiragana("hiragana + カタカナ") // -> "ひらがな + かたかな"
```

```go
// Convert Romaji and Hiragana to Katakana
kana.ToKatakana("katakana") // -> "カタカナ"
kana.ToKatakana("katakana + ひらがな") // -> "カタカナ + ヒラガナ"
```

```go
// Convert Romaji to Hiragana and Katakana (case sensitive romaji)
kana.ToKana("hiragana + KATAKANA") // -> "ひらがな + カタカナ"
```

```go
// String IS Hiragana
kana.IsHiragana("たべる") // -> true
kana.IsHiragana("食べる") // -> false
```

```go
// String CONTAINS Hiragana
kana.ContainsHiragana("たべる") // -> true
kana.ContainsHiragana("食べる") // -> true
kana.ContainsHiragana("カタカナ") // -> false
```

```go
// String IS Katakana
kana.IsKatakana("バナナ") // -> true
kana.IsKatakana("バナナ茶") // -> false
```

```go
// String CONTAINS Katakana
kana.ContainsKatakana("バナナ") // -> true
kana.ContainsKatakana("バナナ茶") // -> true
kana.ContainsKatakana("ひらがな") // -> false
```

```go
// String IS Kanji
kana.IsKatakana("水") // -> true
kana.IsKatakana("also 茶") // -> false
```

```go
// String CONTAINS Kanji
kana.ContainsKatakana("食べる") // -> true
kana.ContainsKatakana("also 茶") // -> true
kana.ContainsKatakana("ひらがな + カタカナ") // -> false
```

```go
// Extract Kanji from String
kana.ExtractKanji("また、平易な日本語で伝える週刊ニュースも放送します。日本語") 
// -> []string{"平", "易", "日", "本", "語", "伝", "週", "刊", "放", "送", "日", "本", "語"}
```


### Linguistic Considerations
A number of rule considerations and assumptions have been made while creating this library in order to conform to Modified-Hepburn Wapuro Romaji.

* __Long Vowels__ are indicated using using repeating characters instead of macrons/circumflexes: oo/おお instead of ō:
    * benkyou/べんきょう, not benkyō.
    * toukyou/とうきょう, not Tōkyō.
    * obaasan/おばあさん, not obāsan.
  * __Chōonpu (ー) are preferred__ for katakana and loan words, and will preserved or converted to minus-dashes.
    * セーラー, not セエラア, becoming se-ra-
    * パーティー, not パアティィ, becoming pa-ti-
* __Particles__ are always converted literally:
    * は is ha, not wa.
    * を is wo, not o.
    * へ is he, not e, etc.
* __Moraic N's are used__ to disambiguate ん and な,に,ぬ,ね,の,にゃ,にゅ,にょ:
    * かんい is kan'i
    * しにょう is shin'you
    * ぜんにん is zennin
    * ぜんいん is zen'in
    * あんない is annai
* __Long Consonants__ marked with sokuons are doubled:
    * いっしょ is issho
    * ぱっぱ is pappa
    * ざっし is zasshi
   * __However, _the Revised Hepburn intepretation of cch is used for alignment with English phonology:__ 
     * まっちゃ is matcha, not maccha
     * こっち is kotchi, not kocchi
* __la, li, lu, le, lo__ are converted to _ra, ri, ru, re, ro_ before transliteration.
* __Nihon-Shiki Romanization is used to map input-ambiguous characters:__
    * da and DA are だ and ダ
    * di and DI are ぢ and ヂ
    * du and DU are づ and ヅ
    * de and DE are で and デ
	 * do and DO are ど and ド
* __じゃ, じゅ and じょ are ja, ju, and jo,__ however, _jya, jyu, and jyo_ are also valid for a one-way romaji→kana conversion.
* __Isolated small vowel kana__ are romanized with 'x' prefixes, _if not part of a larger composite._ 
    * xa and XA are ぁ and ァ
	 * xi and XI are ぃ and ィ
	 * xu and XU are ぅ and ゥ
	 * xe and XE are ぇ and ェ
	 * xo and XO are ぉ and ォ
	 * __Dangling _x_'s__ that remain after all other transliterations are converted into っ and ッ for hiragana and katakana respectively.
 
