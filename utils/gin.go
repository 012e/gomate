package utils

import (
	"net/http"

	"github.com/012e/gomate/utils/json"
	"github.com/gin-gonic/gin"
)

// return json with status internal server error, with json message: "something went wrong"
func UnknownFail(g *gin.Context) {
	g.JSON(http.StatusInternalServerError, json.Fail("something went wrong"))
	return
}
