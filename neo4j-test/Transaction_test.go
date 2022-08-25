package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"testing"
)

func addPersonInSession(driver neo4j.Driver, name string) error {
	var err error
	var session neo4j.Session
	var result neo4j.Result

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer session.Close()

	if result, err = session.Run("CREATE (a:Person {name: $name})", map[string]interface{}{"name": name}); err != nil {
		return err
	}

	if _, err = result.Consume(); err != nil {
		return err
	}

	return nil
}

func TestAddPersonInSession(t *testing.T) {
	var driver neo4j.Driver
	var err error

	// 创建neo4j驱动
	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "123456", ""))
	if err != nil {
		fmt.Println("Get Driver Failed: " + err.Error())
	}
	defer driver.Close()

	err = addPersonInSession(driver, "test")
	if err != nil {
		fmt.Println("Get Driver Failed: " + err.Error())
	}
}
