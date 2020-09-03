package api


import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunServer(){
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/mobius-stress/api/v1", func(r chi.Router){
		r.Post("/tasks",doStressTask)
		r.Get("/tasks/{taskId}", fetchStressTaskStatus)
	})

	host:="0.0.0.0"
	port:="6666"
	logrus.WithFields(
		logrus.Fields{
			"host": host,
			"port": port,
		},
	).Infoln("starting the http server")
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r)
}

