package domain

type Patient struct {
	Id            int    `db:"id"`
	Name          string `db:"name"`
	MedicalRecord string `db:"medical_record"`
	Sex           string `db:"sex"`
}
