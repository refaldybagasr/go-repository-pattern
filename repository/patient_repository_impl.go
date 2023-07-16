package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/refaldybagasr/go-repository-pattern/model/domain"
)

type PatientRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewPatientRepository(db *pgxpool.Pool) PatientRepository {
	return &PatientRepositoryImpl{db: db}
}

func (repository *PatientRepositoryImpl) Save(ctx *gin.Context, patient domain.Patient) domain.Patient {
	panic("implement me")
}

func (repository *PatientRepositoryImpl) FindAll(ctx *gin.Context) []domain.Patient {
	SQL := "SELECT id, name, medical_record, sex FROM patients"
	rows, err := repository.db.Query(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var patients []domain.Patient
	for rows.Next() {
		var patient domain.Patient
		err := rows.Scan(&patient.Id, &patient.Name, &patient.MedicalRecord, &patient.Sex)
		if err != nil {
			panic(err)
		}
		patients = append(patients, patient)
	}

	return patients
}

func (repository *PatientRepositoryImpl) FindById(ctx *gin.Context, patientId int) domain.Patient {
	var patient domain.Patient

	SQL := "SELECT id, name, medical_record, sex FROM patients where id = $1"
	err := repository.db.QueryRow(ctx, SQL, patientId).Scan(&patient.Id, &patient.Name, &patient.MedicalRecord, &patient.Sex)
	if err != nil {
		panic(err)
	}

	return patient
}

func (repository *PatientRepositoryImpl) Update(ctx *gin.Context, patient domain.Patient) domain.Patient {
	//TODO implement me
	panic("implement me")
}

func (repository *PatientRepositoryImpl) Delete(ctx *gin.Context, patient domain.Patient) {
	//TODO implement me
	panic("implement me")
}
