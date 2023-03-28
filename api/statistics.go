package api

import (
	"github.com/application-research/delta-metrics-rest/dao"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func configGinStatisticsRouter(router gin.IRoutes) {
	router.GET("/open/stats/totals/info", ConverHttprouterToGin(GetOpenTotalInfoStats))
}

func GetOpenTotalInfoStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	record, err := dao.GetOpenTotalInfoStats()
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}
