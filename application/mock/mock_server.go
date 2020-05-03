package mock

import (
	"github.com/gin-gonic/gin"
)

// ServerMock to mock server gin
func ServerMock(fn func(g *gin.Engine)) {
	g := gin.Default()
	fn(g)
}
