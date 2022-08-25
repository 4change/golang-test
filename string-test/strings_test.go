package string_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStr(t *testing.T) {
	str1 := "123,"
	println(strings.HasSuffix(str1, ","))
	str1 += "," + "34"
	println(str1)

	//str2 := "123"
	str3 := "123,"
 	str2Arr := strings.Split(str3, ",")
 	for i := 0; i < len(str2Arr); i++ {
 		println(str2Arr[i])
	}
}

func TestRemoveDuplicateElement(t *testing.T) {
	s := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}
	fmt.Println(removeDuplicateElement(s))
}

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

func TestTrimRight(t *testing.T) {
	str := "test trim,"

	str	= strings.TrimRight(str, ",")

	fmt.Println(str)
}

func TestReplace(t *testing.T) {
	newIdStr := "23"
	str := "0000000000"

	i := len(str) - len(newIdStr)

	str = str[0:i] + newIdStr

	fmt.Println(str)
}
