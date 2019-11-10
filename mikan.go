package mikan

import (
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

const (
	runeWidth = 24
)

const (
	joshis        = "でなければ|について|かしら|くらい|けれど|なのか|ばかり|ながら|ことよ|こそ|こと|さえ|しか|した|たり|だけ|だに|だの|つつ|ても|てよ|でも|とも|から|など|なり|ので|のに|ほど|まで|もの|やら|より|って|で|と|な|に|ね|の|も|は|ば|へ|や|わ|を|か|が|さ|し|ぞ|て"
	keywords      = "[\\(（「『]+.*?[\\)）」』]|\\s+|[a-zA-Z0-9]+\\.[a-z]{2,}|[一-龠々〆ヵヶゝ]+|[ぁ-んゝ]+|[ァ-ヴー]+|[a-zA-Z0-9]+|[ａ-ｚＡ-Ｚ０-９]+"
	periods       = "[\\.\\,。、！\\!？\\?]+"
	bracketsBegin = "[〈《「『｢（(\\[【〔〚〖〘❮❬❪❨(<{❲❰｛❴]"
	bracketsEnd   = "[〉》」』｣)）\\]】〕〗〙〛}>\\)❩❫❭❯❱❳❵｝]"
	spaces        = "\\s"
	hiraganas     = "[ぁ-んゝ]+"
)

const (
	TypeKeywords      = "keywords"
	TypePeriods       = "periods"
	TypeBracketsBegin = "bracketsBegin"
	TypeBracketsEnd   = "bracketsEnd"
	TypeSpaces        = "spaces"
)

func Mikan(str string) []string {
	result := make([]string, 0)

	words := Split(str)

	line := ""
	for _, word := range words {
		if runewidth.StringWidth(line+word) > runeWidth {
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

func Split(str string) []string {
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
			prevType = TypeSpaces
			prevWord = word
			continue
		}

		bracketsBeginRep := regexp.MustCompile(bracketsBegin)
		if bracketsBeginRep.MatchString(word) {
			prevType = TypeBracketsBegin
			prevWord = word
			continue
		}

		bracketsEndRep := regexp.MustCompile(bracketsEnd)
		if bracketsEndRep.MatchString(word) {
			result[len(result)-1] += word
			prevType = TypeBracketsEnd
			prevWord = word
			continue
		}

		if prevType == TypeBracketsBegin {
			word = prevWord + word
			prevWord = ""
			prevType = ""
		}

		// すでに文字が入っている上で助詞 or Periods or Spacesが続く場合は結合する
		if len(result) > 0 && len(token) > 0 && prevType == "" {
			result[len(result)-1] += word
			prevType = TypeKeywords
			prevWord = word
			continue
		}

		// 単語のあとの文字がひらがななら結合する
		hiraganaRep := regexp.MustCompile(hiraganas)
		if len(result) > 1 && len(token) > 0 || (prevType == TypeKeywords && hiraganaRep.MatchString(word)) {
			result[len(result)-1] += word
			prevType = ""
			prevWord = word
			continue
		}

		result = append(result, word)
		prevType = TypeKeywords
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
