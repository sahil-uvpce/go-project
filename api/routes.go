package api

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Root    *mux.Router // ''
	APIRoot *mux.Router // 'api/v1'
	User    *mux.Router // 'api/v1/user'
	Setting *mux.Router // 'api/v1/setting'
	Metrics *mux.Router // 'api/v1/metrics'
}

func (a *API) setupRoutes() {
	a.Router.Root = a.MainRouter
	a.Router.APIRoot = a.MainRouter.PathPrefix("/api/v1").Subrouter()
	a.Router.User = a.Router.APIRoot.PathPrefix("/user").Subrouter()
	a.Router.Setting = a.Router.APIRoot.PathPrefix("/setting").Subrouter()
	a.Router.Metrics = a.Router.APIRoot.PathPrefix("/metrics").Subrouter()

	a.InitHealthCheck()
	a.InitUser()
	a.InitUserDetail()
	a.InitUserLogin()
	a.InitSetting()

	a.Metrics.SetupHttpHandler(a.Router.Metrics)
}
