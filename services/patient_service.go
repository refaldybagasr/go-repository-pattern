package services

import (
	"github.com/gin-gonic/gin"
	"github.com/refaldybagasr/go-repository-pattern/model/web"
)

type PatientService interface {
	FindAll(ctx *gin.Context) []web.PatientResponse
	Create(ctx *gin.Context, request web.PatientCreateRequest) web.PatientResponse
	FindById(ctx *gin.Context, patientId int) web.PatientResponse
	Update(ctx *gin.Context, request web.PatientUpdateRequest) web.PatientResponse
	Delete(ctx *gin.Context, id int)
}
