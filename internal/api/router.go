package api

import (
	"sikabiz/user-importer/internal/api/handler/user"
)

// SetupAPIRoutes
// @title						user Importer Service
// @version         			1.0.0
// @description     			This APIs create server for importing users and fetch them
// @Host 						localhost:8080
// @BasePath  					/
// @Schemes 					https
func (s *Server) SetupAPIRoutes(
	userHandler *user.UserHandler,
) {
	r := s.engine

	v1 := r.Group("v1")
	{
		v1.GET("/users/:id", userHandler.GetUser)
	}
}
