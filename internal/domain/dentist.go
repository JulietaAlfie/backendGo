package domain

type Dentist struct {
	Id       int    `json:"id"`
	Lastname string `json:"lastname" binding:"required"`
	Name     string `json:"name" binding:"required"`
	License  string `json:"license" binding:"required"`
}
