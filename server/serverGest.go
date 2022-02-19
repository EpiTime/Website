package serverGest

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	E *gin.Engine
}

func NewServer() (serv Server) {
	serv.E = gin.Default()
	return
}
