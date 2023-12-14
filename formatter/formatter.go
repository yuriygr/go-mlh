package formatter

import (
	"html"
	"regexp"
	"strings"
)

const (
	reNewLines       = `\r?\n`
	reReduceNewLines = `(<br(?: \/)?>\s*){3,}`
	reURL            = `(http|ftp|https):\/\/([\w\p{L}\-_]+(?:(?:\.[\w\p{L}\-_]+)+))([\w\p{L}\-\.,@?^=%&amp;:/~\+#]*[\w\p{L}\-\@?^=%&amp;/~\+#])?`
	reTags           = `#[^\s!@#$%^&*()=+.\/,\[{\]};:'"?><]{1,24}` // `#[^\W_][\w-]*` // `#[A-Za-z0-9\_]*`
	///<a\b[^>]*>|\B(#[^\W_][\w-]*)/gi
)

// EscapeString - Why not?
func EscapeString(str string) string {
	return html.EscapeString(str)
}

// StripSpaces - Remove all extra spaces and replace with one between words
func StripSpaces(str string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(str, " "))
}

// Nl2br - Change new line to br
func Nl2br(str string) string {
	re := regexp.MustCompile(reNewLines)
	return re.ReplaceAllString(str, `<br>`)
	// Hz chto lutshe
	// return strings.Replace(str, "\n", "<br />", -1)
}

// ReduceNewLines - Reduce line breaks
func ReduceNewLines(str string) string {
	re := regexp.MustCompile(reReduceNewLines)
	return re.ReplaceAllString(str, `<br><br>`)
}

// MarkupURLs - Marks up and formats links in HTML
func MarkupURLs(str string) string {
	re := regexp.MustCompile(reURL)
	return re.ReplaceAllString(str, `<a href="$0">$0</a>`)
}

// ExtractHashtags - Извлекает список всех хештегов и возвращает _уникальный_ массив тегов.
// Но каждый раз отсторированный по разному, by design
// @todo: Может добавить обработку #_hash_
func ExtractHashtags(str string) []string {
	re := regexp.MustCompile(reTags)
	tags := re.FindAllString(str, -1)

	findedTags := map[string]bool{}
	for _, tag := range tags {
		findedTags[strings.TrimLeft(tag, "#")] = true
	}

	uniqueTags := []string{}
	for tag := range findedTags {
		uniqueTags = append(uniqueTags, tag)
	}

	return uniqueTags
}

// @TODO

// Markup - Форматирование разметки
// a.k.a. корень всего зла
func Markup(str string) string {
	return str
}
