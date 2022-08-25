package main

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"strings"
	"testing"
)

func TestReadYaml(t *testing.T) {

	dir, _ := os.Getwd()
	replace := strings.Replace(dir, "yaml-test", "", -1)
	fmt.Println(replace)
	file, err := yaml.ReadFile( "test.yaml")

	if err != nil {
		panic(err.Error())
	}

	str := "apiVersion"
	apiVersion, error := file.Get(str)
	if error != nil {
		panic(error.Error())
	}

	fmt.Println("=apiVersion===:\t", apiVersion)

}
