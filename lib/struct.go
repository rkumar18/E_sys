package lib


var Postgres_config Postgres
var Services_config Service

type PradhanData struct{
	Username string
	Password string
}

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db_Name  string `yaml:"dbname"`

}

type Service struct{
	Port string `yaml:"port"`
}