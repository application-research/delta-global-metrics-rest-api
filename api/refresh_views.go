package api

import (
	"github.com/application-research/delta-metrics-rest/dao"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

const viewTypeRefreshGlobalStats = "global_stats"
const viewTypeRefreshAllTableViews = "all_table_views"

func configGinRefreshViewsRouter(router gin.IRoutes) {
	router.GET("/admin/views/refresh/:view_name", ConverHttprouterToGin(RefreshView))
}

func RefreshView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	authKey := viper.Get("REFRESH_VIEW")
	if authKey == nil {
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(ctx, w, "REFRESH_VIEW not set")
		return
	}

	// check auth header
	authorizationString := r.Header.Get("Authorization")
	authParts := strings.Split(authorizationString, " ")
	if len(authParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(ctx, w, "invalid authorization header")
		return
	}
	if authParts[1] != authKey {
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(ctx, w, "invalid authorization header")
		return
	}

	// get query params
	viewName := ps.ByName("view_name")
	if viewName == viewTypeRefreshGlobalStats {
		record, err := dao.RefreshGlobalStatsView()
		if err != nil {
			returnError(ctx, w, r, err)

			return
		}
		writeJSON(ctx, w, record)
	}

	return
}
