package domain

type Patient struct {
	Id            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	Lastname      string `json:"lastname" binding:"required"`
	Residence     string `json:"residence" binding:"required"`
	DNI           int    `json:"dni" binding:"required"`
	DischargeDate string `json:"discharge_date" binding:"required"`
}
