package lib


var Postgres_config Postgres
var Services_config Service

type Admin struct{
	Username string
	Email    string
	Id       int
	Password string
}

type Users struct{
	Username string
	Email    string
	Phoneno  int
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