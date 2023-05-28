package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/theodod3s/simplebank/db/sqlc"
)

// Server serves HTTP requests for the banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// creates a new HTTP server and routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
