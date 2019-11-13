package mikan

import (
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

const (
	defaultRuneWidth = 80

	joshis        = "でなければ|について|かしら|くらい|けれど|なのか|ばかり|ながら|ことよ|こそ|こと|さえ|しか|した|たり|だけ|だに|だの|つつ|ても|てよ|でも|とも|から|など|なり|ので|のに|ほど|まで|もの|やら|より|って|で|と|な|に|ね|の|も|は|ば|へ|や|わ|を|か|が|さ|し|ぞ|て"
	keywords      = "[\\(（「『]+.*?[\\)）」』]|[\\s　]+|[a-zA-Z0-9]+\\.[a-z]{2,}|[一-龠々〆ヵヶゝ・'’`]+|[ぁ-んゝ・'’`]+|[ァ-ヴー・'’`]+|[ｦ-ﾟ・'’`]+|[a-zA-Z0-9À-ÿ・'’`]+|[ａ-ｚＡ-Ｚ０-９・'’`]+"
	periods       = "[\\.\\,．。、！\\!？\\?&＆]+"
	bracketsBegin = "[〈《「『｢（(\\[【〔〚〖〘❮❬❪❨(<{❲❰｛❴]"
	bracketsEnd   = "[〉》」』｣)）\\]】〕〗〙〛}>\\)❩❫❭❯❱❳❵｝]"
	spaces        = "[\\s　]+"
	hiraganas     = "[ぁ-んゝ]+"

	typeKeywords      = "keywords"
	typePeriods       = "periods"
	typeBracketsBegin = "bracketsBegin"
	typeBracketsEnd   = "bracketsEnd"
	typeSpaces        = "spaces"
)

// Option sets the options specified
type Option func(*Mikan)

// RuneWidth sets rune width of line
func RuneWidth(runeWidth int) Option {
	return func(m *Mikan) {
		m.RuneWidth = runeWidth
	}
}

// Mikan is core struct
type Mikan struct {
	RuneWidth int
}

// NewMikan create Mikan instance
// if you want Mikan's field, set options
func NewMikan(options ...Option) *Mikan {
	m := &Mikan{
		RuneWidth: defaultRuneWidth,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

// Split returns strings that has been split to fit the runeWidth
func (m *Mikan) Split(str string) []string {
	result := make([]string, 0)
	words := Analyze(str)

	line := ""
	for _, word := range words {
		if runewidth.StringWidth(line+word) > m.RuneWidth {
			result = append(result, line)
			line = word
			continue
		}
		line += word
	}
	if line != "" {
		result = append(result, line)
	}

	return result
}

// Analyze returns the sentence divided into words
func Analyze(str string) []string {
	rules := []string{
		keywords,
		periods,
		bracketsBegin,
		bracketsEnd,
		spaces,
	}
	rep := regexp.MustCompile(strings.Join(rules, "|"))
	words := rep.FindAllString(str, -1)

	result := make([]string, 0)
	prevType := ""
	prevWord := ""
	for _, word := range words {
		token := getToken(word)

		spacesRep := regexp.MustCompile(spaces)
		if spacesRep.MatchString(word) {
			result = append(result, word)
			prevType = typeSpaces
			prevWord = word
			continue
		}

		bracketsBeginRep := regexp.MustCompile(bracketsBegin)
		if bracketsBeginRep.MatchString(word) {
			prevType = typeBracketsBegin
			prevWord = word
			continue
		}

		bracketsEndRep := regexp.MustCompile(bracketsEnd)
		if bracketsEndRep.MatchString(word) {
			result[len(result)-1] += word
			prevType = typeBracketsEnd
			prevWord = word
			continue
		}

		if prevType == typeBracketsBegin {
			word = prevWord + word
			prevType = ""
		}

		// すでに文字が入っている上で助詞 or Periods or Spacesが続く場合は結合する
		if len(result) > 0 && len(token) > 0 && prevType == "" {
			result[len(result)-1] += word
			prevType = typeKeywords
			prevWord = word
			continue
		}

		// 単語のあとの文字がひらがななら結合する
		hiraganaRep := regexp.MustCompile(hiraganas)
		if len(result) > 1 && len(token) > 0 || (prevType == typeKeywords && hiraganaRep.MatchString(word)) {
			result[len(result)-1] += word
			prevType = ""
			prevWord = word
			continue
		}

		result = append(result, word)
		prevType = typeKeywords
		prevWord = word
	}

	return result
}

func getToken(word string) []string {
	periodRep := regexp.MustCompile(periods)
	if token := periodRep.FindAllString(word, -1); token != nil {
		return token
	}
	joshiRep := regexp.MustCompile(joshis)
	if token := joshiRep.FindAllString(word, -1); token != nil {
		return token
	}
	return nil
}
