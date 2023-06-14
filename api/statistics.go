package api

import (
	"github.com/application-research/delta-metrics-rest/dao"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func configGinStatisticsRouter(router gin.IRoutes) {
	router.GET("/open/stats/totals/info", ConverHttprouterToGin(GetOpenTotalInfoStats))
	router.GET("/open/stats/list/sps", ConverHttprouterToGin(GetAllSps))
	router.GET("/open/stats/list/wallet/addrs", ConverHttprouterToGin(GetWalletsAddrs))
	router.GET("/open/stats/instance/ips", ConverHttprouterToGin(GetDeltaIps))

	// open stats
	//router.GET("/open/stats/onboarded/deals/by-sp/:sp_id", ConverHttprouterToGin(GetDeltaIps))
	// /open/stats/onboarded/deals/by-username/:username
	// /open/stats/onboarded/deals/by-sp/:sp_id
	// /open/stats/onboarded/deals/by-delta-uuid/:delta_uuid
	// /open/stats/onboarded/deals/by-key/ (post)
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

func GetAllSps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	record, err := dao.GetAllSPs()
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

func GetWalletsAddrs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	record, err := dao.GetAllWalletAddrs()
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

func GetDeltaIps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	record, err := dao.GetAllDeltaIps()
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}
