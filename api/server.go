package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() (*Server, error) {
	router := gin.Default()
	
	server := &Server{
		router: router,
	}

	server.registerRoutes()	
	
	return server, nil
}

func (s *Server) registerRoutes() {
	// Basic HTTP method endpoints
	httpMethod := NewHttpMethod()
	s.router.GET("/get", httpMethod.HandleGet)
	s.router.POST("/post", httpMethod.HandlePost)
	s.router.PUT("/put", httpMethod.HandlePut)
	s.router.DELETE("/delete", httpMethod.HandleDelete)
	s.router.PATCH("/patch", httpMethod.HandlePatch)


	// Status endpoint
	httpStatus := NewHttpStatus()
	s.router.Any("/status/:code", httpStatus.HandleStatus)

	// Format endpoint
	httpFormat := NewHttpFormat()
	s.router.GET("/brotli", httpFormat.HandleBrotli)
	s.router.GET("/deflate", httpFormat.HandleDeflate)
	s.router.GET("/deny", httpFormat.HandleDeny)
	s.router.GET("/gzip", httpFormat.HandleGzip)
	s.router.GET("/html", httpFormat.HandleHtml)
	s.router.GET("/encoding/utf8", httpFormat.handleUTF8)
	
	// Register middleware
	s.router.Use(ContentLengthMiddleware())

}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}



