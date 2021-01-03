package filters

import (
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

type RequestTracer struct {
	log *logrus.Logger
}

func NewRequestTracer(log *logrus.Logger) *RequestTracer {
	return &RequestTracer{log: log}
}

func (rcv *RequestTracer) LogRequest(
	request *restful.Request,
	response *restful.Response,
	chain *restful.FilterChain,
) {
	rcv.log.Tracef(" %8s | %s", request.Request.Method, request.SelectedRoutePath())
	chain.ProcessFilter(request, response)
}
