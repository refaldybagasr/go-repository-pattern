package main

import (
	"github.com/gin-gonic/gin"
	"github.com/refaldybagasr/go-repository-pattern/app"
	"github.com/refaldybagasr/go-repository-pattern/controllers"
	"github.com/refaldybagasr/go-repository-pattern/repository"
	"github.com/refaldybagasr/go-repository-pattern/services"
	"net/http"
)

func main() {
	db := app.NewDB()
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	patientRepository := repository.NewPatientRepository(db)
	patientService := services.NewPatientService(patientRepository)
	patientController := controllers.NewPatientController(patientService)
	patients := r.Group("/patients")
	{
		patients.GET("", patientController.FindAll)
		patients.GET("/:patientId", patientController.FindById)
		patients.PATCH("/:patientId", patientController.Update)
		patients.POST("", patientController.Create)
		patients.DELETE("", patientController.Delete)
	}
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
