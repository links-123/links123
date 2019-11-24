package api_version

import "github.com/emicklei/go-restful"

type APIVersionService interface {
	ProvideWebService() *restful.WebService
}
