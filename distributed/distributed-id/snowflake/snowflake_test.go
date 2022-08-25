package snowflake

import (
	"github.com/deckarep/golang-set"
	"testing"
)

// 生成1亿个ID看有没有重复
func TestSnowFlake(t *testing.T) {
	var snowflake Snowflake

	snowflake.Init(1024)
	s := mapset.NewSet()
	for i := 0; i < 1000000; i++ {
		id, err := snowflake.GetId()
		if err != nil {
			t.FailNow()
		}
		if s.Add(id) == false {
			t.FailNow()
		}
		println(id)
	}
}
