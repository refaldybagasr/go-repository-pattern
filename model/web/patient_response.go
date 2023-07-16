package web

type PatientResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	MedicalRecord string `json:"medical_record"`
	Sex           string `json:"sex"`
}
