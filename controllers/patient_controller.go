package controllers

import "github.com/gin-gonic/gin"

type PatientController interface {
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
