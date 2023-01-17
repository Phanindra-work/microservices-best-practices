package main

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/microservice/mysql"
	"github.com/microservice/tracing"

	"github.com/microservice/config"
	"github.com/microservice/retryablehttp"

	"github.com/go-openapi/loads"
	"github.com/microservice/restapi"

	apierrors "github.com/go-openapi/errors"
	"github.com/microservice/log"

	studenthandler "github.com/microservice/restapi/handlers/student"
	teacherhandler "github.com/microservice/restapi/handlers/teacher"
	"github.com/microservice/restapi/operations"
	studentservice "github.com/microservice/services/students"
	teacherservice "github.com/microservice/services/teacher"
)

func main() {
	config.Init()
	switch config.Vars.Environment {
	case config.Local:
		log.Local()
	case config.Development:
		log.Development()
	case config.Production:
		log.Development()
	}
	l := log.Sugar()
	tracing.InitTracer(
		config.Vars.ZipkinEndpoint, fmt.Sprintf(":%d", config.Vars.Port), false, config.Vars.ZipkinServiceName,
	)
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		panic(err)
	}

	db, err := db.NewManager(config.Vars.DatabaseURI, config.Vars.DatabaseName)

	if err != nil {
		l.Panicf("mongo connection error ", err)
	}
	studentservice := studentservice.New(db)
	studenthandler.Init(studentservice)
	teacherservice := teacherservice.New(db)
	teacherhandler.Init(teacherservice)

	apierrors.DefaultHTTPCode = http.StatusBadRequest
	api := operations.NewMicroserviceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = 8888
	server.EnabledListeners = []string{"http"}
	server.CleanupTimeout = 10 * time.Second
	if err := server.MaxHeaderSize.Set("10KB"); err != nil {
		panic(err)
	}
	server.KeepAlive = 3 * time.Minute
	server.ReadTimeout = 30 * time.Second
	server.WriteTimeout = 60 * time.Second

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
