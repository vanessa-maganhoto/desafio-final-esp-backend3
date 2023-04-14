package main

import (
	"database/sql"

	"github.com/ctd/esp-backend-desafio-2.git/cmd/server/handler"
	"github.com/ctd/esp-backend-desafio-2.git/internal/appointment"
	"github.com/ctd/esp-backend-desafio-2.git/internal/dentist"
	"github.com/ctd/esp-backend-desafio-2.git/internal/patient"
	"github.com/ctd/esp-backend-desafio-2.git/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/my_db")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	storageDentist := store.NewSqlStoreDentist(db)
	repoDentist := dentist.NewRepository(storageDentist)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)

	storagePatient := store.NewSqlStorePatient(db)
	repoPatient := patient.NewRepository(storagePatient)
	servicePatient := patient.NewService(repoPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)

	storageAppointment := store.NewSqlStoreAppointment(db)
	repoAppointment := appointment.NewRepositoryAppointment(storageAppointment)
	serviceAppointment := appointment.NewService(repoAppointment)
	appointmentHandler := handler.NewAppointmentHandler(serviceAppointment)
	
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentists := r.Group("/dentists")

	dentists.GET("", dentistHandler.GetAll())
	dentists.GET(":id", dentistHandler.GetByID())
	dentists.POST("", dentistHandler.Post())
	dentists.DELETE(":id", dentistHandler.Delete())
	dentists.PATCH(":id", dentistHandler.Patch())
	dentists.PUT(":id", dentistHandler.Put())

	patients := r.Group("/patients")
	patients.GET("", patientHandler.GetAll())
	patients.GET(":id", patientHandler.GetByID())
	patients.POST("", patientHandler.Post())
	patients.DELETE(":id", patientHandler.Delete())
	patients.PATCH(":id", patientHandler.Patch())
	patients.PUT(":id", patientHandler.Put())

	appointments := r.Group("/appointments")
	appointments.GET(":id", appointmentHandler.GetByID())
	appointments.POST("", appointmentHandler.Post())
	appointments.DELETE(":id", appointmentHandler.Delete())
	appointments.PATCH(":id", appointmentHandler.Patch())
	appointments.PUT(":id", appointmentHandler.Put()) 
	appointments.GET("patient/:rg", appointmentHandler.GetByRg()) 
	r.Run(":8083")
}
