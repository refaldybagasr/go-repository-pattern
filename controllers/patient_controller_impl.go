package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/refaldybagasr/go-repository-pattern/model/web"
	"github.com/refaldybagasr/go-repository-pattern/services"
	"net/http"
	"strconv"
)

type PatientControllerImpl struct {
	service services.PatientService
}

func NewPatientController(service services.PatientService) PatientController {
	return &PatientControllerImpl{service: service}
}

func (controller *PatientControllerImpl) FindAll(ctx *gin.Context) {
	patients := controller.service.FindAll(ctx)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   patients,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PatientControllerImpl) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *PatientControllerImpl) FindById(ctx *gin.Context) {
	patientId := ctx.Param("patientId")
	id, err := strconv.Atoi(patientId)
	if err != nil {
		panic(err)
	}
	patient := controller.service.FindById(ctx, id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   patient,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PatientControllerImpl) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *PatientControllerImpl) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
