package infrastructure

import (
	"pulse_sense/src/core"
	"pulse_sense/src/internal/hospitals/infrastructure/controllers"
	tokenService "pulse_sense/src/internal/services/auth/domain"
	"pulse_sense/src/server/middleware"

	"github.com/gin-gonic/gin"
)

type HospitalRoutes struct {
	CreateHospitalsController    *controllers.CreateHospitalController
	GetAllHospitalsController    *controllers.GetAllHospitalsController
	GetHospitalsController       *controllers.GetHospitalByIDController
	GetHospitalsByUserController *controllers.GetHospitalByUserController
	UpdateHospitalsController    *controllers.UpdateHospitalController
	SearchHospitalCOntroller 	 *controllers.SearchHospitalController
	TokenService                 tokenService.TokenService
	AuthRepo                     *core.AuthRepository
}

func NewHospitalRoutes(
	createHospitalsController	 *controllers.CreateHospitalController,
	getAllHospitalsController	 *controllers.GetAllHospitalsController,
	getHospitalsController 		 *controllers.GetHospitalByIDController,
	getHospitalsByUserController *controllers.GetHospitalByUserController,
	updateHospitalsController 	 *controllers.UpdateHospitalController,
	searchHospitalController 	 *controllers.SearchHospitalController,
	tokenService 				 tokenService.TokenService,
	authRepo					 *core.AuthRepository,

) *HospitalRoutes {
	return &HospitalRoutes{
		CreateHospitalsController:    createHospitalsController,
		GetAllHospitalsController:    getAllHospitalsController,
		GetHospitalsController:       getHospitalsController,
		GetHospitalsByUserController: getHospitalsByUserController,
		UpdateHospitalsController:    updateHospitalsController,
		SearchHospitalCOntroller:     searchHospitalController,
		TokenService:                 tokenService,
		AuthRepo:                     authRepo,
	}
}

func (r *HospitalRoutes) AttachRoutes(router *gin.Engine) {
	
		userGroup := router.Group("/hospitals")
		{
			userGroup.GET("/:id", r.GetHospitalsController.GetHospitalByID)
			userGroup.GET("/user/:id", r.GetHospitalsByUserController.GetByUser)
			userGroup.PUT("/:id", r.UpdateHospitalsController.UpdateUser)
			userGroup.POST("", r.CreateHospitalsController.Create)
			userGroup.GET("", r.GetAllHospitalsController.GetAllHospital)
			userGroup.GET("/search/:name", r.SearchHospitalCOntroller.SearchHospital)
		}
	userAuth := middleware.AuthMiddleware(r.TokenService, r.AuthRepo, "usuario")
	adminAuth := middleware.AuthMiddleware(r.TokenService, r.AuthRepo, "administrador")
	userGroup.Use(userAuth)
	{
	}

	adminGroup := router.Group("/admin/Hospitals")
	adminGroup.Use(adminAuth)
	{
		adminGroup.GET("", r.GetAllHospitalsController.GetAllHospital)
		adminGroup.PUT("/:id", r.UpdateHospitalsController.UpdateUser)
	}
}
