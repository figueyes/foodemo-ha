package controllers

import (
	"github.com/labstack/echo/v4"
	"go-course/demo/app/foo/application/usecases"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/log"
	"net/http"
	"os"
)

type fooHandler struct {
	useCase usecases.FooListAllUseCase
}

func NewFooHandler(e *echo.Echo, useCase usecases.FooListAllUseCase) *fooHandler {
	fooHandler := &fooHandler{
		useCase: useCase,
	}
	fooGroup := e.Group(os.Getenv(constants.BASE_PATH))
	{
		fooGroup.GET("", fooHandler.listFoo)
		fooGroup.POST("/", nil)
		fooGroup.PUT("/", nil)
		fooGroup.PATCH("/", nil)
	}
	return fooHandler
}

func (f *fooHandler) listFoo(c echo.Context) error {
	response, err := f.useCase.List()
	if err != nil {
		log.WithError(err).Info("error executing use case")
		return c.JSON(http.StatusBadRequest,echo.Map{"error": err.Error()})
	}
	log.Info("response resolved successfully")
	return c.JSON(http.StatusOK, response)
}
