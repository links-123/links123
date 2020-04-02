package api

import "github.com/emicklei/go-restful"

type RESTAPIVersionedService interface {
	ProvideWebService() *restful.WebService
}
