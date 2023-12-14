package formatter

import (
	"testing"

	"github.com/yuriygr/go-mlh/helpers"
)

func TestStripSpaces(t *testing.T) {
	testCases := []struct {
		name string
		got  string
		want string
	}{
		{"Empty", "", ""},
		{"Little space", " ", ""},
		{"Obi-Wan", "  Obi-Wan    Kenobi   ", "Obi-Wan Kenobi"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := StripSpaces(tc.got)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestEscapeString(t *testing.T) {
	testCases := []struct {
		name string
		got  string
		want string
	}{
		{"Empty", "", ""},
		{"HTML Classic xss", "<script>alert('hack');</script>", "&lt;script&gt;alert(&#39;hack&#39;);&lt;/script&gt;"},
		{"HTML Other", "<b><i><s><div href='#'>", "&lt;b&gt;&lt;i&gt;&lt;s&gt;&lt;div href=&#39;#&#39;&gt;"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := EscapeString(tc.got)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestMarkupURLs(t *testing.T) {
	testCases := []struct {
		name string
		got  string
		want string
	}{
		{"URL with query params", "https://www.youtube.com/watch?v=mFGq92BYmt4", `<a href="https://www.youtube.com/watch?v=mFGq92BYmt4">https://www.youtube.com/watch?v=mFGq92BYmt4</a>`},
		{"URL with other params", "http://www.reddit.com/#go-board", `<a href="http://www.reddit.com/#go-board">http://www.reddit.com/#go-board</a>`},
		{"URL with text", "https://www.reddit.com/r/golang Best subreddit", `<a href="https://www.reddit.com/r/golang">https://www.reddit.com/r/golang</a> Best subreddit`},
		{"URL with HTML", "http://www.reddit.com/?<b><i><s><div href='#'>", `<a href="http://www.reddit.com/?">http://www.reddit.com/?</a><b><i><s><div href='#'>`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MarkupURLs(tc.got)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestExtractHashtags(t *testing.T) {
	testCases := []struct {
		name string
		got  string
		want []string
	}{
		{"Star Wars", `#jedi #star #wars #hello_there`, []string{"jedi", "star", "wars", "hello_there"}},
		{"Deduplicate", `#jedi #jedi #jedi`, []string{"jedi"}},
		{"Number #1", `#1`, []string{"1"}},
		{"With numbers", `#1 #42answer`, []string{"1", "42answer"}},
		{"With slash", `#one_two_ #_three_`, []string{"one_two_", "_three_"}},
		{"Polka dot", `#polka:dot #:more_dots`, []string{"polka"}},
		{"Only _", `#____`, []string{}},
		{"Кирилл", `#почта #пришла`, []string{"почта", "пришла"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ExtractHashtags(tc.got)
			if !helpers.EqualSlices(got, tc.want) {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

// @TODO

/*func TestMarkup(t *testing.T) {
	testCases := []struct {
		name string
		got  string
		want string
	}{
		{"Bold", `**Bold**`, `<b>Bold</b>`},
		{"Italic", `*Italic*`, `<i>Italic</i>`},
		{"Bold & Italic", `***Italic***`, `<b><i>Italic</i></b>`},
		{"Strike", `~~Strike~~`, `<s>Strike</s>`},
		{"Spoiler", `\%\%Spoiler\%\%`, `<span class="markup --spoiler">Spoiler</span>`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Markup(tc.got)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}*/
