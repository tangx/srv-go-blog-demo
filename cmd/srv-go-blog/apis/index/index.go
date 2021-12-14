package index

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jarvis/rum-gonic/pkg/httpx"
)

type Index struct {
	httpx.MethodMulti `path:"/index" methods:"GET,HEAD"`
}

func (Index) Output(c *gin.Context) (interface{}, error) {
	return "powered by go-jarvis", nil
}
