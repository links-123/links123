package helpers

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"

	"github.com/links-123/links123/shared/gateway/representation"
)

func logError(logger *logrus.Logger, err error) {
	var message = "error is empty"

	if err != nil {
		message = err.Error()
	}

	logger.Errorf("error during request processing, %s", message)
}

func RespondWithBadRequest(response *restful.Response, log *logrus.Logger, err error, message string) {
	logError(log, err)

	if err := response.WriteHeaderAndEntity(http.StatusBadRequest, &representation.ErrorResponse{
		Message: message,
	}); err != nil {
		logError(log, err)
	}
}

func RespondWithInternalError(response *restful.Response, log *logrus.Logger, err error) {
	logError(log, err)

	if err := response.WriteHeaderAndEntity(http.StatusInternalServerError, &representation.ErrorResponse{
		Message: http.StatusText(http.StatusInternalServerError),
	}); err != nil {
		logError(log, err)
	}
}

func RespondWithOK(response *restful.Response, log *logrus.Logger, data interface{}) {
	if err := response.WriteHeaderAndEntity(http.StatusOK, data); err != nil {
		logError(log, err)
	}
}

func RespondWithNoContent(response *restful.Response, log *logrus.Logger) {
	response.WriteHeader(http.StatusNoContent)
}
