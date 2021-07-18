package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
)

// EqualSlices - Сравнения двух срезов
func EqualSlices(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// StringInSlice - Нахождение строки в слайсе
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// CheckSubstrings - Вхождение строки из массива в другую строку(?)
// https://stackoverflow.com/questions/47131996/go-how-to-check-if-a-string-contains-multiple-substrings
func CheckSubstrings(str string, subs ...string) (bool, int) {
	matches := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches++
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}

// CheckSubstringsRegexp - Вхождение строки из массива в другую строку только на регулярках
func CheckSubstringsRegexp(str string, subs ...string) bool {
	var re = regexp.MustCompile(strings.Join(subs[:], "|"))
	return re.MatchString(str)
}

// MD5 - Ну
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	state := hex.EncodeToString(hash[:])
	return state
}
