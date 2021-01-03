package app

import "github.com/emicklei/go-restful"

type MicroService interface {
	Run() error
}

type RESTAPIVersionedService interface {
	ProvideWebService() *restful.WebService
}
