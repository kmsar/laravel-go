package Str

import (
	"errors"
	"fmt"
	"html"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// PadTypes for StrPad
const (
	StrPadRight = iota
	StrPadLeft
	StrPadBoth
)

// RuneMatchFunc is function to check if a rune match some condition
type RuneMatchFunc func(rune) bool

// Substr returns the portion of string specified by the start and length parameters
//
// The behaviour of this function is mostly the same as the PHP mb_substr function,
//
// see http://php.net/manual/en/function.mb-substr.php
//
// except that:
//
// 1) If start or length is invalid, empty string will be return
//
// 2) If length is 0, the substring starting from start until the end of the string will be returned
func Substr(str string, start, length int) string {

	rs := []rune(str)
	rl := len(rs)

	if rl == 0 {
		return ""
	}

	if start < 0 {
		start = rl + start
	}
	if start < 0 {
		start = 0
	}
	if start > rl-1 {
		return ""
	}

	end := rl

	if length < 0 {
		end = rl + length
	} else if length > 0 {
		end = start + length
	}

	if end < 0 || start >= end {
		return ""
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// Strpos find position of first occurrence of string in a string
//
// It's multi-byte safe. return -1 if can not find the substring
func Strpos(haystack, needle string) int {

	pos := strings.Index(haystack, needle)

	if pos < 0 {
		return pos
	}

	rs := []rune(haystack[0:pos])

	return len(rs)
}

// Strrpos find the position of the last occurrence of a substring in a string
func Strrpos(haystack, needle string) int {

	pos := strings.LastIndex(haystack, needle)

	if pos < 0 {
		return pos
	}

	rs := []rune(haystack[0:pos])

	return len(rs)
}

// Stripos find position of the first occurrence of a case-insensitive substring in a string
func Stripos(haystack, needle string) int {
	return Strpos(strings.ToLower(haystack), strings.ToLower(needle))
}

// Strripos find the position of the last occurrence of a case-insensitive substring in a string
func Strripos(haystack, needle string) int {
	return Strrpos(strings.ToLower(haystack), strings.ToLower(needle))
}

// buildReplaceSlice is a helper function for Replace and Ireplace
func buildReplaceSlice(search, replace any) ([]string, []string, error) {

	var aSearch, aReplace []string

	switch s := search.(type) {
	case int:
		aSearch = append(aSearch, strconv.Itoa(s))
	case rune:
		aSearch = append(aSearch, string(s))
	case string:
		aSearch = append(aSearch, s)
	case []string:
		aSearch = s
	default:
		return aSearch, aReplace, errors.New("unsupported type of search")
	}

	switch r := replace.(type) {
	case int:
		aReplace = append(aReplace, strconv.Itoa(r))
	case rune:
		aReplace = append(aReplace, string(r))
	case string:
		aReplace = append(aReplace, r)
	case []string:
		aReplace = r
	default:
		return aSearch, aReplace, errors.New("unsupported type of replace")
	}

	return aSearch, aReplace, nil
}

// Replace all occurrences of the search string with the replacement string
//
// This function is an implement of PHP's str_replace
//
// see http://php.net/manual/en/function.str-replace.php
func Replace(search, replace any, subject string) string {

	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		subject = strings.Replace(subject, s, r, -1)
	}

	return subject
}

// Ireplace is case-insensitive version of Replace()
func Ireplace(search, replace any, subject string) string {

	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	var reg *regexp.Regexp
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		reg = regexp.MustCompile("(?i:" + s + ")")
		subject = reg.ReplaceAllString(subject, r)
	}

	return subject
}

// Addslashes quote string with slashes
//
// The characters to be escaped are single quote ('), double quote (") and backslash (\)
func Addslashes(str string) string {
	return Replace([]string{"\\", "'", "\""}, []string{"\\\\", "\\'", "\\\""}, str)
}

// Stripslashes un-quotes a quoted string
func Stripslashes(str string) string {
	return Replace([]string{"\\\\", "\\'", "\\\""}, []string{"\\", "'", "\""}, str)
}

// Chr returns a one-character string containing the character specified by ascii
func Chr(ascii int) string {
	for ascii < 0 {
		ascii += 256
	}
	ascii %= 256
	return string(rune(ascii))
}

// Explode returns an slice of strings, each of which is a substring of str
// formed by splitting it on boundaries formed by the string delimiter
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

// Implode returns a string containing a string representation of all the slice
// elements in the same order, with the glue string between each element
func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// Lcfirst make a string's first character lowercase
func Lcfirst(str string) string {
	return strings.ToLower(Substr(str, 0, 1)) + Substr(str, 1, 0)
}

// Ucfirst make a string's first character uppercase
func Ucfirst(str string) string {
	return strings.ToUpper(Substr(str, 0, 1)) + Substr(str, 1, 0)
}

// Strstr find the first occurrence of a string
func Strstr(haystack, needle string) string {
	pos := Strpos(haystack, needle)
	if pos < 0 {
		return ""
	}
	return Substr(haystack, pos, 0)
}

// Stristr is case-insensitive Strstr()
func Stristr(haystack, needle string) string {
	pos := Stripos(strings.ToLower(haystack), strings.ToLower(needle))
	if pos < 0 {
		return ""
	}
	return Substr(haystack, pos, 0)
}

// Trim strips whitespace (or other characters) from the beginning and end of a string
func Trim(str string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimSpace(str)
	}
	var cutset = ""
	for _, c := range chars {
		cutset += c
	}
	return strings.Trim(str, cutset)
}

// Ltrim strips whitespace (or other characters) from the beginning of a string
func Ltrim(str string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	var cutset = ""
	for _, c := range chars {
		cutset += c
	}
	return strings.TrimLeft(str, cutset)
}

// Rtrim strips whitespace (or other characters) from the end of a string
func Rtrim(str string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	var cutset = ""
	for _, c := range chars {
		cutset += c
	}
	return strings.TrimRight(str, cutset)
}

// HTMLSpecialchars converts special characters to HTML entities
func HTMLSpecialchars(str string) string {
	return html.EscapeString(str)
}

// HTMLSpecialcharsDecode converts special HTML entities back to characters
func HTMLSpecialcharsDecode(str string) string {
	return html.UnescapeString(str)
}

// StrWordCount returns an array containing all the words found inside the string
func StrWordCount(str string) []string {
	return strings.Fields(str)
}

// NumberFormat formats a number with grouped thousands
//
// decimals: sets the number of decimal points
//
// decPoint: sets the separator for the decimal point
//
// thousandsSep: sets the thousands separator
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	if decimals < 0 {
		decimals = 0
	}
	var isNegative bool
	num := Round(number, decimals)
	if num < 0 {
		isNegative = true
		num = math.Abs(num)
	}

	str := fmt.Sprintf("%."+strconv.Itoa(decimals)+"F", num)

	var integerPart, decimalPart string

	if decimals > 0 {
		integerPart = str[:len(str)-(decimals+1)]
		decimalPart = str[len(str)-decimals:]
	} else {
		integerPart = str
	}

	var s string

	length := len(integerPart)
	i := length % 3
	if i > 0 {
		s = integerPart[0:i]
	}

	for i < length {
		if s != "" {
			s += thousandsSep
		}
		j := i + 3
		s += integerPart[i:j]
		i = j
	}

	if decimals > 0 {
		s += decPoint + decimalPart
	}

	if isNegative {
		s = "-" + s
	}

	return s
}

// DefaultNumberFormat is default NumberFormat for english notation with thousands separator
func DefaultNumberFormat(number float64, decimals int) string {
	return NumberFormat(number, decimals, ".", ",")
}

// NumberFormat2 number_format()
// decimals: Sets the number of decimal points.
// decPoint: Sets the separator for the decimal point.
// thousandsSep: Sets the thousands' separator.
func NumberFormat2(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// StrPad pads a string to a certain length with another string
func StrPad(input string, padLength int, padString string, padType int) string {
	// if the value of padLength is less than or equal to the length of the input string,
	// no padding takes place, and input will be returned.
	if padLength <= len(input) {
		return input
	}
	// default padType is StrPadRight
	if padType > StrPadBoth {
		padType = StrPadRight
	}
	// default padString is space
	if padString == "" {
		padString = " "
	}

	var s string
	l := padLength - len(input)
	switch padType {
	case StrPadRight:
		s = input
		i := 0
		for i < l {
			for j := 0; j < len(padString); j++ {
				s += padString[j : j+1]
				i++
				if i >= l {
					break
				}
			}
		}
	case StrPadLeft:
		i := 0
		for i < l {
			for j := 0; j < len(padString); j++ {
				s += padString[j : j+1]
				i++
				if i >= l {
					break
				}
			}
		}
		s += input
	case StrPadBoth:
		l1 := l / 2
		s = StrPad(StrPad(input, len(input)+l1, padString, StrPadLeft), padLength, padString, StrPadRight)
	}

	return s
}

// Strtoupper makes a string uppercase
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}

// Strtolower makes a string lowercase
func Strtolower(str string) string {
	return strings.ToLower(str)
}

// Ucwords uppercases the first character of each word in a string
func Ucwords(str string, fs ...RuneMatchFunc) string {
	if len(fs) == 0 {
		fs = []RuneMatchFunc{unicode.IsSpace}
	}

	var s = make([]rune, len(str))
	var isFirst = true
	for i, r := range str {
		var match = false
		for _, f := range fs {
			match = f(r)
			if match {
				break
			}
		}
		if match {
			isFirst = true
			s[i] = r
		} else if isFirst {
			v := []rune(strings.ToUpper(string(r)))
			s[i] = v[0]
			isFirst = false
		} else {
			s[i] = r
		}
	}
	return string(s)
}

// Ucname uppercases names
func Ucname(str string) string {
	f := func(r rune) bool {
		return r == '-' || r == '\''
	}
	return Ucwords(Strtolower(str), unicode.IsSpace, f)
}
