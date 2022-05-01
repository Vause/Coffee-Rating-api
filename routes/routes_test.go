package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetUpRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	routeHandler := SetUpRouter()

	if routeHandler != nil {
		t.Logf("Route initialization complete")
	} else {
		t.Fail()
	}
}
