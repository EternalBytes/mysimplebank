package api

import (
	db "github.com/eternalbytes/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves http requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a http server and set routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	// add routes to server
	server.router = router
	return server
}

func errorMessage(err error) gin.H {
	return gin.H{"error": err.Error()}
}
