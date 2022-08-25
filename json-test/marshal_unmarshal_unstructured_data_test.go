package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Unmarshal_Unstructured_Data(t *testing.T) {
	birdJson := `{
					"birds": 
						{
							"pigeon":"likes to perch on rocks",
							"eagle":"bird of prey"
						},
					"animals":"none"
				}`
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(birdJson), &result); err != nil {
		fmt.Println(err)
	}
	birds := result["birds"].(map[string]interface{})
	for key, value := range birds {
		fmt.Println(key, "-----", value.(string))
	}
}

type Human struct {
	ManSlice 		[]interface{}					`json:"manSlice"`			// json标签名定制
	ManMap			map[string]interface{}			`json:"manMap"`
	WomanSlice		[]interface{}					`json:"womanSlice"`
	WomanMap		map[string]interface{}			`json:"womanMap"`
}

type Man struct {
	Id 			int									`json:"manId,omitempty"`	// omitempty: value为空的json字段不进行序列化
	Name 		string								`json:"manName,omitempty"`
	Age 		float64								`json:"manAge,omitempty"`
}

type Woman struct {
	Id 			int									`json:"womanId,omitempty"`
	Name  		string								`json:"womanName,omitempty"`
	Age  		float64								`json:"womanAge,omitempty"`
}

func Test_Marshal_UnstructuredData(t *testing.T) {
	var err error
	var humanBytes []byte

	human := &Human{
		ManMap: make(map[string]interface{}),
		WomanMap: make(map[string]interface{}),
	}

	man1 := &Man{
		Id: 1,
		Name: "tom",
		Age: 1,
	}
	man2 := &Man{
		Id: 2,
		Name: "jack",
	}

	woman1 := &Woman{
		Id: 1,
		Name: "tom",
		Age: 1,
	}
	woman2 := &Woman{
		Id: 2,
		Name: "jack",
	}

	// 向 []interface{} 添加不同类型的数据
	var manSlice, womanSlice []interface{}
	manSlice = append(manSlice, "1")
	manSlice = append(manSlice, 2)
	manSlice = append(manSlice, man1)
	manSlice = append(manSlice, man2)
	human.ManSlice = manSlice

	womanSlice = append(womanSlice, "11")
	womanSlice = append(womanSlice, 22)
	womanSlice = append(womanSlice, woman1)
	womanSlice = append(womanSlice, woman2)
	human.WomanSlice = womanSlice

	// 向 map[string]interface{} 中添加不同类型的数据
	human.ManMap["1"] = 1
	human.ManMap["two"] = "two"
	human.ManMap["man1"] = man1
	human.ManMap["man2"] = man2

	human.WomanMap["1"] = 1
	human.WomanMap["two"] = "two"
	human.WomanMap["woman1"] = woman1
	human.WomanMap["woman2"] = woman2

	// 序列化
	if humanBytes, err = json.Marshal(human); err == nil {
		humanJson := string(humanBytes)			// []byte转json
		fmt.Println(humanJson)
	}
}
