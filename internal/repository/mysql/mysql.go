package mysql

import (
	"database/sql"

	"github.com/MakMoinee/basicInformationSystem/internal/common"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

type mysqlService struct {
	DatabaseName     string
	QueryString      string
	DbDriver         string
	ConnectionString string
	Db               *sql.DB
}

type IMySql interface {
	GetStudentInfoByTeacherID(id string)
}

func (svc *mysqlService) Set() {
	svc.DatabaseName = common.DB_NAME
	svc.ConnectionString = common.MYSQL_USERNAME + ":" + common.MYSQL_PASSWORD + "@" + common.CONNECTION_STRING + svc.DatabaseName
	svc.DbDriver = common.DB_DRIVER

	db := svc.openDBConnection()
	testDb(db)
}

func NewService() IMySql {
	svc := mysqlService{}
	svc.Set()
	return &svc
}

func (svc *mysqlService) GetStudentInfoByTeacherID(id string) {
	log.Info("Inside mysql,GetStudentInfoByTeacherID")
}

// openDBConnection creates db connection
func (svc *mysqlService) openDBConnection() *sql.DB {
	db, err := sql.Open(svc.DbDriver, svc.ConnectionString)
	if err != nil {
		log.Error(err.Error())
	}
	return db
}

func testDb(db *sql.DB) {
	_, err := db.Query("select 1")
	defer db.Close()
	if err != nil {
		log.Fatal("Invalid Username or Password, Please Check Your MySql Configurations")
	}
	log.Info("Successfully Connected to DB")
}
