package domain

type Appointment struct {
	Id          int     `json:"id"`
	Patient     Patient `json:"patient" binding:"required"`
	Dentist     Dentist `json:"dentist" binding:"required"`
	Date        string  `json:"date" binding:"required"`
	Time        string  `json:"time" binding:"required"`
	Description string  `json:"description" binding:"required"`
}


