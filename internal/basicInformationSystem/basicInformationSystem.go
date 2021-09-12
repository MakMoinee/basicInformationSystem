package basicinformationsystem

import (
	"github.com/MakMoinee/basicInformationSystem/internal/repository/mysql"
	log "github.com/sirupsen/logrus"
)

type service struct {
	MySqlRepo mysql.IMySql
}

type IBasic interface {
	GetStudentInfoByTeacherID(id string)
}

func NewService() IBasic {
	svc := service{}
	svc.MySqlRepo = mysql.NewService()
	return &svc
}

func (svc *service) GetStudentInfoByTeacherID(id string) {
	log.Info("Inside basicService,GetStudentInfoByTeacherID")
}
