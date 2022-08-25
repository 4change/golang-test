package golang_test

import (
	"regexp"
	"testing"
)

// urlMatcher: url匹配校验器
// design-pattern string: 数据库中给定的url模式
// url string: 前端输入的url
func urlMatcher(pattern string, url string) bool {

	// match something like "/ids/{domain}" against "/ids/abc"

	// some preprocess for design-pattern
	re := regexp.MustCompile("{[^/]+}")							// 匹配括号中非斜杠的任意字符
	pattern_reg := re.ReplaceAllString(pattern, "[^/]+")		// 将括号及非斜杠的任意字符一起替换为非斜杠的任意字符

	pattern_reg_compiled := regexp.MustCompile(pattern_reg)

	return pattern_reg_compiled.MatchString(url)

}

func TestMultiTableData(t *testing.T) {
	var tests = []struct {
		dbApi string
		url string
		want  bool
	}{
		{"/v1/rules/{id}/{section}", "/v1/rules/2/input", true},
		{"/v1/rules/{id}/{section}/", "/v1/rules/2/input", true},
		{"/v1/rules/{id}/{section}", "/v1/rules/2/input/", true},
		{"/v1/rules/{id}/{section}/", "/v1/rules/2/input/", true},
	}

	for _, test := range tests {
		if got := urlMatcher(test.dbApi, test.url); got != test.want {
			t.Errorf("urlMatcher(%q, %q) = %v", test.dbApi, test.url, got)
		}
	}
}
