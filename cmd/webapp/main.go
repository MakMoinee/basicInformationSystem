package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/MakMoinee/basicInformationSystem/cmd/webapp/config"
	"github.com/MakMoinee/basicInformationSystem/cmd/webapp/routes"
	"github.com/MakMoinee/basicInformationSystem/internal/common"
	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func main() {
	config.Set()
	config.InitializeConfig()
	log.Info("Starting the service at port ", common.PORT)
	service := startService(common.PORT)
	service.EnableProfiling(common.TLS_ENABLED)
	routes.Set(service)
	err := service.Start()
	if err != nil {
		panic(err)
	}
}

// startService custom method for service initialization of http server
func startService(port string) *goserve.Service {

	tlsEnabled := common.TLS_ENABLED
	mInfo := make(map[string]interface{})
	mInfo["SERVICE_VERSION"] = common.SERVICE_VERSION
	cert := viper.GetString("cert_file")
	key := viper.GetString("key_file")
	r := chi.NewRouter()

	customService := goserve.GetService(r, port, cert, key, tlsEnabled, mInfo)
	return customService
}
