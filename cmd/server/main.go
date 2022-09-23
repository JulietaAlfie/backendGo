package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/JulietaAlfie/backendGo.git/cmd/server/handler"
	"github.com/JulietaAlfie/backendGo.git/docs"
	"github.com/JulietaAlfie/backendGo.git/internal/appointment"
	"github.com/JulietaAlfie/backendGo.git/internal/dentist"
	"github.com/JulietaAlfie/backendGo.git/internal/patient"
	"github.com/JulietaAlfie/backendGo.git/pkg/middleware"
	"github.com/JulietaAlfie/backendGo.git/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Certified Tech Developer - Julieta Alfie
// @version 1.0
// @description Clinica Odontologica.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	dataSource := "root:root@tcp(localhost:3306)/my_db"
	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	storageDentist := store.NewSqlStoreDentist(storageDB)
	repositoryDentist := dentist.NewRepository(storageDentist)
	serviceDentist := dentist.NewService(repositoryDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)

	storagePatient := store.NewSqlStorePatient(storageDB)
	repositoryPatient := patient.NewRepository(storagePatient)
	servicePatient := patient.NewService(repositoryPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)

	storageAppointment := store.NewSqlStoreAppointment(storageDB)
	repositoryAppointment := appointment.NewRepository(storageAppointment)
	serviceAppointment := appointment.NewService(repositoryAppointment)
	appointmentHandler := handler.NewAppointmentHandler(serviceAppointment)

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger(), middleware.AllowAll())
	  
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.GET("", dentistHandler.GetAll())
		dentists.POST("", middleware.Authentication(), dentistHandler.Post())
		dentists.DELETE(":id", middleware.Authentication(), dentistHandler.Delete())
		dentists.PATCH(":id", middleware.Authentication(), dentistHandler.Patch())
		dentists.PUT(":id", middleware.Authentication(), dentistHandler.Put())
	}

	patients := r.Group("/patients")
	{
		patients.GET(":id", patientHandler.GetByID())
		patients.GET("", patientHandler.GetAll())
		patients.POST("", middleware.Authentication(), patientHandler.Post())
		patients.DELETE(":id", middleware.Authentication(), patientHandler.Delete())
		patients.PATCH(":id", middleware.Authentication(), patientHandler.Patch())
		patients.PUT(":id", middleware.Authentication(), patientHandler.Put())
	}

	appointments := r.Group("/appointments")
	{
		appointments.GET("", appointmentHandler.GetAll())
		appointments.GET(":id", appointmentHandler.GetByID())
		appointments.GET("/dni/:dni", appointmentHandler.GetByDni())
		appointments.POST("", middleware.Authentication(), appointmentHandler.Post())
		appointments.POST(":dni/:license", middleware.Authentication(), appointmentHandler.PostByDniAndLicence())
		appointments.DELETE(":id", middleware.Authentication(), appointmentHandler.Delete())
		appointments.PATCH(":id", middleware.Authentication(), appointmentHandler.Patch())
		appointments.PUT(":id", middleware.Authentication(), appointmentHandler.Put())
	}

	if err = r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}


