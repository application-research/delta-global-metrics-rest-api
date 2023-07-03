package api

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("refreshing view")
	ctx := initializeContext(r)
	authParts := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(ctx, w, "invalid authorization header")
		return
	}

	authSvcApi := viper.Get("AUTH_SVC_API").(string)
	if authSvcApi == "" {
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(ctx, w, "AUTH_SVC_API not set")
		return
	}
	
	response, err := http.Post(
		authSvcApi+"/check-api-key",
		"application/json",
		strings.NewReader(fmt.Sprintf(`{"token": "%s"}`, authParts[1])),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(ctx, w, err.Error())
		return
	}

	authResp, err := GetAuthResponse(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJSON(ctx, w, err.Error())
		return
	}

	if authResp.Result.Validated == false {
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(ctx, w, "invalid token")
		return
	}

	// get query params
	viewName := ps.ByName("view_name")
	go func() {
		if viewName == viewTypeRefreshGlobalStats {
			_, err := dao.RefreshGlobalStatsView()
			if err != nil {
				returnError(ctx, w, r, err)

				return
			}
		}
	}()

	w.WriteHeader(http.StatusOK)
	writeJSON(ctx, w, "Refresh view request received. Please wait for a few minutes for the view to be refreshed.")
	return
}

type AuthResponse struct {
	User struct {
		Username string `json:"username"`
		Perm     int    `json:"perm"`
		Flags    int    `json:"flags"`
	} `json:"user"`
	Result struct {
		Validated bool   `json:"validated"`
		Details   string `json:"details"`
	} `json:"result"`
}

func GetAuthResponse(resp *http.Response) (AuthResponse, error) {
	jsonBody := AuthResponse{}
	err := json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		return AuthResponse{
			User: struct {
				Username string `json:"username"`
				Perm     int    `json:"perm"`
				Flags    int    `json:"flags"`
			}{
				Username: "",
				Perm:     0,
				Flags:    0,
			},
			Result: struct {
				Validated bool   `json:"validated"`
				Details   string `json:"details"`
			}{
				Validated: false,
				Details:   "empty json body",
			},
		}, nil
	}
	return jsonBody, nil
}
