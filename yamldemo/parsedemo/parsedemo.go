package main
import (
	"fmt"
	"io/ioutil"
	"log"
	yaml "gopkg.in/yaml.v2"
)
func main() {
	fmt.Println("---start---")

    conf := new(Yaml)
    yamlFile, err := ioutil.ReadFile("test.yaml")
   
 
    log.Println("yamlFile:", string(yamlFile))
    if err != nil {
        log.Printf("yamlFile.Get err %v", err)
	}
	
    err = yaml.Unmarshal(yamlFile, conf)

    if err != nil {
        log.Fatalf("Unmarshal: %v when to struct", err)
    }
	log.Println("conf", conf)
	log.Printf("mysqlHost:%s, port:%d, cacheList:%s", conf.Mysql.Host, conf.Mysql.Port, conf.Cache.List)

		
	configMap := make(map[string]interface{})
    err = yaml.Unmarshal(yamlFile, configMap)

    if err != nil {
        log.Fatalf("Unmarshal: %v when to map", err)
    }

	for k,v := range configMap {
		_, isStr := v.(string)
		_, isInt32 := v.(int32)
		_, isBool := v.(bool)
		if isStr  {
			fmt.Printf("key:%s, value(string):%s \n", v)
		} else if isInt32  {
			fmt.Printf("key:%s, value(int32):%d \n", k,v)
		} else if isBool  {
			fmt.Printf("key:%s, value(bool):%t \n", k,v)
		} else {
			fmt.Printf("key:%s, value:%s \n", k,v)
		} 
	}
	fmt.Printf("---遍历结束---\n")

	// 使用map表示已经提前知道结构是什么样的， 可以将map提取为自己想要的格式
	mysqlValueInterface := configMap["mysql"]
	mysqlValue, ok := mysqlValueInterface.(interface{})
	if ok {
		mysqlMap, ok := mysqlValue.(map[interface{}]interface{})
		if ok {
			log.Printf(" map解析结果。 mysqlHost:%s, port:%d", mysqlMap["host"], mysqlMap["port"])
		} else {
			log.Printf("no mysql info")
		}
    } else {
		log.Printf("no mysql value is not interface")
	}
	log.Printf("---程序结束---\n")
}