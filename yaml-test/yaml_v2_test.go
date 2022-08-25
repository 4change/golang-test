package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

//Nginx nginx  配置
type Nginx struct {
	Port int `yaml:"Port"`
	LogPath string `yaml:"LogPath"`
	Path string `yaml:"Path"`
}
//Config   系统配置配置
type Config struct{
	Name string `yaml:"SiteName"`
	Addr string `yaml:"SiteAddr"`
	HTTPS bool `yaml:"Https"`
	SiteNginx  Nginx `yaml:"Nginx"`
}

// 测试读取yaml文件中的对应属性到一个结构体
func TestReadAttributes(t *testing.T) {
	var setting Config
	config, err := ioutil.ReadFile("./first.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config,&setting)

	fmt.Println(setting.Name)
	fmt.Println(setting.Addr)
	fmt.Println(setting.HTTPS)
	fmt.Println(setting.SiteNginx.Port)
	fmt.Println(setting.SiteNginx.LogPath)
	fmt.Println(setting.SiteNginx.Path)
}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


//定义conf类型
//类型里的属性，全是配置文件里的属性
type conf struct {
	Host   string `yaml: "host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

// 测试读取yaml文件到一个简单的结构体
func TestReadSimpleToStruct(t *testing.T) {
	var c conf
	//读取yaml配置文件
	conf := c.getConf()
	fmt.Println("conf: ", conf)

	//将对象，转换成json格式
	data, err := json.Marshal(conf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(data))
}

//读取Yaml配置文件,
//并转换成conf对象
func (c *conf) getConf() *conf {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("/home/cx4gxf/GoEnv/src/go-test/yaml-test/one.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type KafkaCluster struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml: "kind"`
	Metadata   Metadata `yaml: "metadata"`
	Spec       Spec     `yaml: "spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
	//map类型
	Labels map[string]*NodeServer `yaml:"labels"`
}

type NodeServer struct {
	Address string `yaml: "address"`
	Id      string `yaml: "id"`
	Name    string `yaml: "name"`
	//注意，属性里，如果有大写的话，tag里不能存在空格
	//如yaml: "nodeName" 格式是错误的，中间多了一个空格，不能识别的
	NodeName string `yaml:"nodeName"`
	Role     string `yaml: "role"`
}

type Spec struct {
	Replicas int    `yaml: "replicas"`
	Name     string `yaml: "name"`
	Image    string `yaml: "iamge"`
	Ports    int    `yaml: "ports"`
	//slice类型
	Conditions []Conditions `yaml: "conditions"`
}

type Conditions struct {
	ContainerPort string   `yaml:"containerPort"`
	Requests      Requests `yaml: "requests"`
	Limits        Limits   `yaml: "limits"`
}

type Requests struct {
	CPU    string `yaml: "cpu-profile"`
	MEMORY string `yaml: "memory"`
}

type Limits struct {
	CPU    string `yaml: "cpu-profile"`
	MEMORY string `yaml: "memory"`
}

// 测试读取yaml文件到一个复杂的结构体
func TestReadToComplexStruct(t *testing.T) {
	var c KafkaCluster
	//读取yaml配置文件, 将yaml配置文件，转换struct类型
	conf := c.getConf()

	//将对象，转换成json格式
	data, err := json.Marshal(conf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(data))
}

//读取Yaml配置文件,
//并转换成conf对象  struct结构
func (kafkaCluster *KafkaCluster) getConf() *KafkaCluster {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("two.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	//err = yaml.Unmarshal(yamlFile, kafkaCluster)
	err = yaml.UnmarshalStrict(yamlFile, kafkaCluster)

	if err != nil {
		fmt.Println(err.Error())
	}

	return kafkaCluster
}
