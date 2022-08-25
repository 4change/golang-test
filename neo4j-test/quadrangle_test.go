package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"testing"
)

func TestGetQuadrangle(t *testing.T) {

	// 创建neo4j驱动, 切记关闭驱动
	driver, driverErr := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "neo4j", ""))
	if driverErr != nil {
		panic(driverErr.Error())
	}
	defer driver.Close()

	// 获取neo4j session, 切记关闭session
	session, sessionErr := driver.Session(neo4j.AccessModeRead)
	if sessionErr != nil {
		panic(sessionErr.Error())
	}
	defer session.Close()

	// 执行neo4j的读事务
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		CYPHER := "match p=(a1)-[]->(x1)<-[]-(a2)-[]->(x2)<-[]-(a1) " +
			" where a2.value<>a1.value " +
			"return p "
		result, resErr := tx.Run(CYPHER, nil)

		recordNum := 0
		var nodeMap = make(map[int]neo4j.Node)
		var relMap = make(map[int]neo4j.Relationship)
		for result.Next() {
			// 记录
			recordNum++
			recordValues := result.Record().Values()[0]

			// 节点
			nodes := recordValues.(neo4j.Path).Nodes()
			for _,v := range nodes {
				nodeMap[int(v.Id())] = v
			}

			// 关系
			relationships := recordValues.(neo4j.Path).Relationships()
			for _,v := range relationships {
				relMap[int(v.Id())] = v
			}
		}
		fmt.Println("nodeMap len:", len(nodeMap))
		fmt.Println("relMap len:", len(relMap))
		fmt.Println("recordNum:", recordNum)

		for i, v := range nodeMap {
			labels := v.Labels()
			fmt.Println(i, v.Id(), labels)
		}

		for i, v := range relMap {
			fmt.Println(i, v.Id(), v.Type(), v.StartId(), v.EndId(), v.Props())
		}

		if resErr != nil {
			panic(resErr.Error())
		}

		return result, resErr
	})

	if err != nil {
		panic(err.Error())
	}
}
