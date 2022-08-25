package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"testing"
	"time"
)

func GetNetworkEdge(driver neo4j.Driver) ([]string, error) {
	var network interface{}
	var err error
	var session neo4j.Session

	// 获取neo4j session, 切记关闭session
	if session, err = driver.Session(neo4j.AccessModeRead); err != nil {
		return nil, err
	}
	defer session.Close()

	network, err = session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string
		var result neo4j.Result

		CYPHER := "MATCH (n)-[r]->(m) " +
			" RETURN distinct {id: id(r), rel: r,  name: type(r), source: id(n), target: id(m)}, id(r) AS rID " +
			" ORDER BY rID"
		if result, err = tx.Run(CYPHER, nil); err != nil {
			return nil, err
		}

		var r neo4j.Record
		for result.Next() {
			//recordMap := result.Record().GetByIndex(0).(map[string]interface{})
			//fmt.Println(recordMap["id"], recordMap["source"], recordMap["target"], recordMap["name"])
			r = result.Record()
			prop := r.Values()[0].(map[string]interface{})["rel"]
			//fmt.Println(prop)
			relationships := prop.(neo4j.Relationship).Props()
			//fmt.Println(relationships)
			fmt.Println(relationships["insert_date"].(time.Time))
			fmt.Println(relationships["update_date"].(time.Time))
			// 输出结果集中的记录
			//r = result.Record()
			//prop := r.Values()[0].(map[string]interface{})["properties"]
			////prop["id"]
			//fmt.Println(prop)

			//i := r.GetByIndex(0).(map[string]interface{})
			//fmt.Println(i["id"], i["labels"])
			//props := prop.(neo4j.Node).Props()
			//fmt.Println("type:", reflect.TypeOf(props["type"]))
			//fmt.Println("value:", reflect.TypeOf(props["value"]))
			//fmt.Println("insert_date:", reflect.TypeOf(props["insert_date"]))
			//fmt.Println("update_date:", reflect.TypeOf(props["update_date"]))
			//fmt.Println("is_new:", reflect.TypeOf(props["is_new"]))
			//fmt.Println("overdue_day:", reflect.TypeOf(props["overdue_day"]))
			//fmt.Println("create_date:", reflect.TypeOf(props["create_date"]))
			//fmt.Println("title:", reflect.TypeOf(props["title"]))
			//fmt.Println("loan_times:", reflect.TypeOf(props["loan_times"]))
			//fmt.Println("is_reject:", reflect.TypeOf(props["is_reject"]))
			//fmt.Println("is_fraud:", reflect.TypeOf(props["is_fraud"]))
			//fmt.Println(props["update_date"], props["type"], props["value"], props["insert_date"], props["name"])
		}

		if err = result.Err(); err != nil {
			return nil, err
		}

		return list, nil
	})
	if err != nil {
		return nil, err
	}

	return network.([]string), nil
}

// Neo4j-测试获取结果集
func TestNetworkEdge(t *testing.T) {
	var driver neo4j.Driver
	var err error
	var result []string

	// 创建neo4j驱动
	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "neo4j", ""))
	if err != nil {
		panic(err.Error())
	}
	defer driver.Close()

	// 获取结果
	result, err = GetNetworkEdge(driver)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("------------------------------------------------------------------------------------------",result)
}