package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/xdot2012/simple-bank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	// ACCOUNTS
	router.GET("/accounts/", server.listAccounts)
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.PUT("/account/:id", server.updateAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	// ENTRIES
	router.GET("/entries/", server.listEntries)
	router.POST("/entry/", server.createEntry)
	router.GET("/entry/:id", server.getEntry)
	router.PUT("/entry/:id", server.updateEntry)
	router.DELETE("/entry/:id", server.deleteEntry)

	// TRANSFERS
	router.GET("/transfers/", server.listTransfers)
	router.POST("/transfer/", server.createTransfer)
	router.GET("/transfer/:id", server.getTransfer)
	router.PUT("/transfer/:id", server.updateTransfer)
	router.DELETE("/transfer/:id", server.deleteTransfer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
