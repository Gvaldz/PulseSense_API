package server

import (
	signsRouters "pulse_sense/src/internal/sensores/signos/infrastructure"
	loginRouters "pulse_sense/src/internal/services/auth/infrastructure"
	websocketRouters "pulse_sense/src/internal/services/websocket/infrastructure"
	patientRouters "pulse_sense/src/internal/sensores/patients/infrastructure"
	usersRouters "pulse_sense/src/internal/users/infrastructure"
	hospitalsRouters "pulse_sense/src/internal/hospitals/infrastructure"
	workersRouters "pulse_sense/src/internal/workers/infrastructure"
	caregiversRouters "pulse_sense/src/internal/caregivers/infrastructure"
	shiftsRouters "pulse_sense/src/internal/shifts/infrastructure"
	userpatientRouters "pulse_sense/src/internal/userpatient/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine             *gin.Engine
	signsRouters	   *signsRouters.SignsRoutes
	usersRouters       *usersRouters.UserRoutes
	patientRouters     *patientRouters.PatientRoutes
	hospitalsRouters   *hospitalsRouters.HospitalRoutes
	workersRouters     *workersRouters.WorkerRoutes
	caregiversRouters  *caregiversRouters.CaregiverRoutes
	shiftsRouters      *shiftsRouters.ShiftRoutes
	userpatientRouters *userpatientRouters.UserPatientRoutes
	loginRouters       *loginRouters.AuthRoutes
	websocketRouters   *websocketRouters.WebSocketRoutes
	authMiddleware     gin.HandlerFunc
}

func NewServer(
	signRoutes *signsRouters.SignsRoutes,
	userRoutes *usersRouters.UserRoutes,
	patientRoutes *patientRouters.PatientRoutes,
	loginRoutes *loginRouters.AuthRoutes,
	hospitalsRoutes *hospitalsRouters.HospitalRoutes,
	workerRoutes *workersRouters.WorkerRoutes,
	caregiversRoutes *caregiversRouters.CaregiverRoutes,
	userpatientRoutes *userpatientRouters.UserPatientRoutes,
	shiftsRoutes *shiftsRouters.ShiftRoutes,
	wsRoutes *websocketRouters.WebSocketRoutes,
	authMiddleware gin.HandlerFunc,
) *Server {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	return &Server{
		engine:             r,
		signsRouters: signRoutes,
		usersRouters:       userRoutes,
		patientRouters:     patientRoutes,
		hospitalsRouters:   hospitalsRoutes,
		workersRouters:     workerRoutes,
		caregiversRouters:  caregiversRoutes,
		shiftsRouters:      shiftsRoutes,
		userpatientRouters: userpatientRoutes,
		loginRouters:       loginRoutes,
		websocketRouters:   wsRoutes,
		authMiddleware:     authMiddleware,
	}
}

func (s *Server) Run() error {
	s.signsRouters.AttachRoutes(s.engine)
	s.usersRouters.AttachRoutes(s.engine)
	s.patientRouters.AttachRoutes(s.engine)
	s.hospitalsRouters.AttachRoutes(s.engine)
	s.workersRouters.AttachRoutes(s.engine)
	s.caregiversRouters.AttachRoutes(s.engine)
	s.shiftsRouters.AttachRoutes(s.engine)
	s.userpatientRouters.AttachRoutes(s.engine)
	s.loginRouters.AttachRoutes(s.engine)
	s.websocketRouters.AttachRoutes(s.engine)
	return s.engine.Run(":8080")
}
