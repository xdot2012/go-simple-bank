package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/xdot2012/simple-bank/db/sqlc"
	"github.com/xdot2012/simple-bank/token"
	"github.com/xdot2012/simple-bank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

// setupRouter adds routes to the server
func (server *Server) setupRouter() {
	router := gin.Default()

	// USERS
	router.POST("/user/", server.createUser)
	router.POST("/login/", server.loginUser)
	router.POST("/token/renew/", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// ACCOUNTS
	authRoutes.GET("/accounts/", server.listAccounts)
	authRoutes.POST("/account", server.createAccount)
	authRoutes.GET("/account/:id", server.getAccount)
	authRoutes.PUT("/account/:id", server.updateAccount)
	authRoutes.DELETE("/account/:id", server.deleteAccount)

	// ENTRIES
	authRoutes.GET("/entries/", server.listEntries)
	authRoutes.POST("/entry/", server.createEntry)
	authRoutes.GET("/entry/:id", server.getEntry)
	authRoutes.PUT("/entry/:id", server.updateEntry)
	authRoutes.DELETE("/entry/:id", server.deleteEntry)

	// TRANSFERS
	authRoutes.GET("/transfers/", server.listTransfers)
	authRoutes.POST("/transfer/", server.createTransfer)
	authRoutes.GET("/transfer/:id", server.getTransfer)
	authRoutes.PUT("/transfer/:id", server.updateTransfer)
	authRoutes.DELETE("/transfer/:id", server.deleteTransfer)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
