package lib

import(
	"fmt"
	"io/ioutil"
	"github.com/go-yaml/yaml"

)
func init(){
	GetConfiguration()
}




func GetConfiguration(){
	
	yamlFile, err := ioutil.ReadFile("config/postgres.yaml")
	if err != nil {
		fmt.Println("Error! ReadFile[GetPostgresConfigration]:", err)
		
	}
	err = yaml.Unmarshal(yamlFile, &Postgres_config)
	if err != nil {
		fmt.Println("Error! Unmarshal[GetTestConfigration]:", err)
		
	}
	yamlFile1, err := ioutil.ReadFile("config/services.yaml")
	if err != nil {
		fmt.Println("Error! ReadFile[GetPostgresConfigration]:", err)
		
	}
	err = yaml.Unmarshal(yamlFile1, &Services_config)
	if err != nil {
		fmt.Println("Error! Unmarshal[GetTestConfigration]:", err)
		
		
	}
	
}