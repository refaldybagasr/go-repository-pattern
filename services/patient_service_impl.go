package services

import (
	"github.com/gin-gonic/gin"
	"github.com/refaldybagasr/go-repository-pattern/model/web"
	"github.com/refaldybagasr/go-repository-pattern/repository"
)

type PatientServiceImpl struct {
	repository repository.PatientRepository
}

func NewPatientService(repository repository.PatientRepository) PatientService {
	return &PatientServiceImpl{repository: repository}
}

func (service *PatientServiceImpl) FindAll(ctx *gin.Context) []web.PatientResponse {
	var patientResponses []web.PatientResponse
	patients := service.repository.FindAll(ctx)
	for _, patient := range patients {
		patientResponse := web.PatientResponse{
			Id:            patient.Id,
			Name:          patient.Name,
			MedicalRecord: patient.MedicalRecord,
			Sex:           patient.Sex,
		}
		patientResponses = append(patientResponses, patientResponse)
	}
	return patientResponses
}

func (service *PatientServiceImpl) Create(ctx *gin.Context, request web.PatientCreateRequest) web.PatientResponse {
	//TODO implement me
	panic("implement me")
}

func (service *PatientServiceImpl) FindById(ctx *gin.Context, patientId int) web.PatientResponse {
	patient := service.repository.FindById(ctx, patientId)
	patientResponse := web.PatientResponse{
		Id:            patient.Id,
		Name:          patient.Name,
		MedicalRecord: patient.MedicalRecord,
		Sex:           patient.Sex,
	}
	return patientResponse
}

func (service *PatientServiceImpl) Update(ctx *gin.Context, request web.PatientUpdateRequest) web.PatientResponse {
	//TODO implement me
	panic("implement me")
}

func (service *PatientServiceImpl) Delete(ctx *gin.Context, id int) {
	//TODO implement me
	panic("implement me")
}
