package regexp_test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexpRepalceAll(t *testing.T) {
	re := regexp.MustCompile("Go(\\w+)")
	fmt.Println(re.ReplaceAllString("Hello Gopher，Hello GoLang", "Hank$1"))
	// Hello Hankpher，Hello HankLang
}

func TestRegex(t *testing.T) {

	pattern := "/v1/features/"
	url := "/v1/features"

	// some preprocess for design-pattern
	re := regexp.MustCompile("{[^/]+}")
	pattern_reg := re.ReplaceAllString(pattern, "[^/]+")
	fmt.Println("pattern_reg-------------------------------------------------------------------------", pattern_reg)

	pattern_reg_compiled := regexp.MustCompile(pattern_reg)
	fmt.Println("pattern_reg_compiled-------------------------------------------------------", pattern_reg_compiled)

	fmt.Println("match-----------------------------------------------------", pattern_reg_compiled.MatchString(url))
}

func TestStringMatch(t *testing.T) {
	text := "Hello Gopher，Hello fly"
	// \w+: \w, 匹配单词字符; +, 前一个出现字符1次或无限次
	// \w+: 单词字符出现1次或无数次
	reg := regexp.MustCompile(`\w+`)
	fmt.Println(reg.MatchString(text))

	text = "09-09-2019"
	reg = regexp.MustCompile(`^\d{1,2}-\d{1,2}-\d{4}`)
	fmt.Println(reg.MatchString(text))
}

func TestStringMatchAll(t *testing.T) {
	text := "Hello Gopher，Hello fly"
	// \w+: \w, 匹配单词字符; +, 前一个出现字符1次或无限次
	// \w+: 单词字符出现1次或无数次
	reg := regexp.MustCompile(`\w+`)
	fmt.Println(reg.FindAllString(text, -1))
}

func TestStringMatchIndex(t *testing.T) {
	text := "Hello Gopher，Hello fly"
	reg := regexp.MustCompile("llo")
	fmt.Println(reg.FindStringIndex(text))
}

func TestIsPhoneNumber(t *testing.T) {
	text := "15228717703"
	reg := regexp.MustCompile("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$")
	fmt.Println(reg.MatchString(text))
}