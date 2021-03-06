package routes

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/MakMoinee/basicInformationSystem/cmd/webapp/response"
	basicinformationsystem "github.com/MakMoinee/basicInformationSystem/internal/basicInformationSystem"
	"github.com/MakMoinee/basicInformationSystem/internal/common"
	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/go-chi/cors"
)

type routeService struct {
	BasicService basicinformationsystem.IBasic
}

type IRoutes interface {
	GetStudentInfoByTeacherID(w http.ResponseWriter, r *http.Request)
}

func Set(httpService *goserve.Service) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	routeHandler := newRoutes()
	httpService.Router.Use(cors.Handler)
	initiateRoute(httpService, routeHandler)
}

func newRoutes() IRoutes {
	svc := routeService{}
	svc.BasicService = basicinformationsystem.NewService()
	return &svc
}

func initiateRoute(httpService *goserve.Service, routeHandler IRoutes) {
	httpService.Router.Get(common.GET_STUDENT_INFO_BY_TEACHER_ID_R, routeHandler.GetStudentInfoByTeacherID)
}

func (svc *routeService) GetStudentInfoByTeacherID(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes,GetStudentInfoByTeacherID")

	// id := r.URL.Query().Get("id")
	str := strings.Split(r.URL.Path, "/")
	id := str[3]

	svc.BasicService.GetStudentInfoByTeacherID(id)

	response.Success(id, w)
}
