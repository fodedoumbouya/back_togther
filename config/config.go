package config

import (
	model "back_togther/structure"
)

var Mysql = model.Connection{
	Host:     "localhost:3306",
	Password: "130231322",
	DbName:   "togther",
}
var ConnectionParameters = "root:" + Mysql.Password + "@tcp(" + Mysql.Host + ")/" + Mysql.DbName + ""
