package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/refaldybagasr/go-repository-pattern/model/domain"
)

type PatientRepository interface {
	Save(ctx *gin.Context, patient domain.Patient) domain.Patient
	FindAll(ctx *gin.Context) []domain.Patient
	FindById(ctx *gin.Context, patientId int) domain.Patient
	Update(ctx *gin.Context, patient domain.Patient) domain.Patient
	Delete(ctx *gin.Context, patient domain.Patient)
}
