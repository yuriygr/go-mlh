package helpers_test

import (
	"testing"

	"github.com/yuriygr/go-mlh/helpers"
)

func TestCheckSubstringsRegexp(t *testing.T) {
	testCases := []struct {
		name string
		got  struct {
			str  string
			subs []string
		}
		want bool
	}{
		{"Что-то с зеленью", struct {
			str  string
			subs []string
		}{"Поставки саженцев ели", []string{"ель", "сосна", "семена", "сажен"}}, true},
		{"Супергерои какие-то", struct {
			str  string
			subs []string
		}{"Ехал Бетмен через Готем, видит Альфреда в болоте", []string{"Бетмен", "Альфред"}}, true},
		{"Строй-материалы", struct {
			str  string
			subs []string
		}{"Поставка бетономешалки и фена строительного", []string{"бетон", "фен"}}, true},
		{"Не по фен-шую?", struct {
			str  string
			subs []string
		}{"Поставка бетономешалки и фена строительного", []string{"фура"}}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := helpers.CheckSubstringsRegexp(tc.got.str, tc.got.subs...)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestCheckSubstrings(t *testing.T) {
	testCases := []struct {
		name string
		got  struct {
			str  string
			subs []string
		}
		want int
	}{
		{"Что-то с зеленью", struct {
			str  string
			subs []string
		}{"Поставки саженцев ели", []string{"ель", "сосна", "семена", "сажен"}}, 1},
		{"Супергерои какие-то", struct {
			str  string
			subs []string
		}{"Ехал Бетмен через Готем, видит Альфреда в болоте", []string{"Бетмен", "Альфред"}}, 2},
		{"Строй-материалы", struct {
			str  string
			subs []string
		}{"Поставка бетономешалки и фена строительного", []string{"бетон", "фен"}}, 2},
		{"Не по фен-шую?", struct {
			str  string
			subs []string
		}{"Поставка бетономешалки и фена строительного", []string{"фура"}}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bl, got := helpers.CheckSubstrings(tc.got.str, tc.got.subs...)
			t.Logf("%v", bl)
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}
