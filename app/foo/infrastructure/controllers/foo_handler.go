package controllers

import (
	"github.com/labstack/echo/v4"
	"go-course/demo/app/foo/application/usecases"
	"go-course/demo/app/foo/infrastructure/controllers/web_model"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
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
	webModelFoo := make([]*web_model.FooWebModel,0)
	utils.ConvertEntity(response, &webModelFoo)
	if err != nil {
		log.WithError(err).Info("error executing use case")
		return c.JSON(http.StatusBadRequest,echo.Map{"error": err.Error()})
	}
	log.Info("response resolved successfully")
	return c.JSON(http.StatusOK, webModelFoo)
}
