package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-course/demo/app/foo/application/usecases"
	"go-course/demo/app/foo/infrastructure/controllers/web_model"
	"go-course/demo/app/shared/domain/constants"
	"go-course/demo/app/shared/log"
	"go-course/demo/app/shared/utils"
	"net/http"
	"os"
	"strconv"
)

type fooHandler struct {
	listAllUseCase         usecases.FooListAllUseCase
	pageableListAllUseCase usecases.FooPageableListAllUseCase
}

func NewFooHandler(e *echo.Echo,
	listAllUseCase usecases.FooListAllUseCase,
	pageableListAllUseCase usecases.FooPageableListAllUseCase) *fooHandler {
	fooHandler := &fooHandler{
		listAllUseCase:         listAllUseCase,
		pageableListAllUseCase: pageableListAllUseCase,
	}
	fooGroup := e.Group(os.Getenv(constants.BASE_PATH))
	{
		fooGroup.GET("", fooHandler.listFooPageable)
		fooGroup.POST("/", nil)
		fooGroup.PUT("/", nil)
		fooGroup.PATCH("/", nil)
	}
	return fooHandler
}

func (f *fooHandler) listFoo(c echo.Context) error {
	response, err := f.listAllUseCase.List()
	webModelFoo := make([]*web_model.FooWebModel, 0)
	utils.ConvertEntity(response, &webModelFoo)
	if err != nil {
		log.WithError(err).Info("error executing use case")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	log.Info("response resolved successfully")
	return c.JSON(http.StatusOK, webModelFoo)
}

func (f *fooHandler) listFooPageable(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	if limit == "" || page == "" {
		return c.JSON(http.StatusBadRequest,
			echo.Map{"error": errors.
				New("both limit and skip cannot be empty").
				Error()})
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.WithError(err).Info("error trying to parse strings to int")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.WithError(err).Info("error trying to parse strings to int")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	response, err := f.pageableListAllUseCase.ListPageable(int64(limitInt), int64(pageInt), "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	data := make([]web_model.DataFooWebModel,0)
	utils.ConvertEntity(response.Data, &data)
	webModelFoo := &web_model.FooWebModel{
		Total: response.Total,
		Page:  response.Page,
		Limit: response.Limit,
		Data: data,
	}
	utils.ConvertEntity(response, &webModelFoo)
	log.Info("response resolved successfully")
	return c.JSON(http.StatusOK, webModelFoo)
}
