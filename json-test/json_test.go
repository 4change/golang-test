package json_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

type DOVersionType int

const(
	DOVerTypeFormal DOVersionType = 1	// 正式版本　
	DOVerTypeTrial  DOVersionType = 12	// 实验版本
	DOVerTypeNormal DOVersionType = 3	// 历史版本
	DOVerTypeDraft  DOVersionType = 4	// 草稿版本
)

var doVerTypeMap = map[int]DOVersionType {
	1: DOVerTypeFormal,
	2: DOVerTypeTrial,
	3: DOVerTypeNormal,
	4: DOVerTypeDraft,
}

func (dovt DOVersionType) Int() int {
	return int(dovt)
}

func NewDOVerType(i int) (dovt DOVersionType, err error) {
	var ok bool

	if dovt, ok = doVerTypeMap[i]; !ok {
		return
	}

	return
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday Time   `json:"birthday"`
}

func Test_Transfer_Json_And_Struct(t *testing.T) {
	now := Time(time.Now())
	t.Log("now---------------------------------------------------------------------------------------------", now)

	jsonSrc := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	person := new(Person)
	err := json.Unmarshal([]byte(jsonSrc), person)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Person struct person-------------------------------------------------------------------------", person)
	t.Log(time.Time(person.Birthday))

	js, _ := json.Marshal(person)
	t.Log(string(js))
}

func Test_Map_String_Interface_Marshal_Unmarshal(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "simon"
	m["age"] = 12
	m["addr"] = "China"
	fmt.Println("m: ", m)
	PrintMap(m)
	fmt.Println("-----------------------------------------------------------------------------------------")

	// 序列化map为字节数组
	data, err := json.Marshal(m)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	m1 := make(map[string]interface{})
	// 字节数组反序列化为对象
	err = json.Unmarshal(data, &m1)
	fmt.Println("err:", err)
	fmt.Println("m1: ", m1)
	PrintMap(m1)
	fmt.Println("=========================================================================================")

	if m1["name"] != nil {
		fmt.Println(m1["name"].(string))
	}

	if m1["type"] != nil {
		fmt.Println(m1["type"].(string))
	} else {
		fmt.Println("there is not the key of type")
	}
}

// 解析 map[string]interface{} 数据格式
func PrintMap(m map[string]interface{}) {
	for k, v := range m {
		switch value := v.(type) {
		case nil:
			fmt.Println(k, "is nil", "null")
		case string:
			fmt.Println(k, "is string", value)
		case int:
			fmt.Println(k, "is int", value)
		case float64:
			fmt.Println(k, "is float64", value)
			const test = 12
			fmt.Println("value == 12 -------------", v, reflect.TypeOf(v) , int(v.(float64)) == DOVerTypeTrial.Int(), v.(float64) == test)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range value {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			PrintMap(value)
		default:
			fmt.Println(k, "is unknown type", fmt.Sprintf("%T", v))
		}
	}
}