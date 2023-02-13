package users_db

import (
	"database/sql"
	"fmt"
	"github.com/Davide202/golang-users/utils/logger"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)
const mysqlUsersHost = "mysql_users_host"
/*
const (

	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"

)
*/
var (
	Client *sql.DB
	
	/*
		username = os.Getenv(mysqlUsersUsername)
		password = os.Getenv(mysqlUsersPassword)
		host     = os.Getenv(mysqlUsersHost) "172.17.0.1"(container)
		schema   = os.Getenv(mysqlUsersSchema)
	*/
)

func init() {

	host, maybe := os.LookupEnv(mysqlUsersHost)
	var dataSourceName string 
	if maybe {
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "admin", host, "db_users")
	}else{
		dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "admin", "localhost", "db_users")
	}

	

	//fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8","root", "admin", "localhost", "db_users",)
	log.Println("database connection: " + dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		//panic(err)
		log.Println("database connection: " + err.Error())
	}
	if err = Client.Ping(); err != nil {
		//panic(err)
		log.Println("database connection: " + err.Error())
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
