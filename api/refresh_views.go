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

// write a temporary storage to track view refreshes

func configGinRefreshViewsRouter(router gin.IRoutes) {
	router.GET("/admin/views/refresh/:view_name", ConverHttprouterToGin(RefreshView))
	router.GET("/admin/views/refresh/status/:view_name", ConverHttprouterToGin(GetRefreshViewStatus))
}

type RefreshViewResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func GetRefreshViewStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// use the tracker to return the result
	ctx := initializeContext(r)
	viewName := ps.ByName("view_name")
	if viewName == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(ctx, w, "view_name required")
		return
	}

	// check if the view is being refreshed
	if dao.ViewRefreshes[viewName] {
		w.WriteHeader(http.StatusAccepted)
		writeJSON(ctx, w,
			struct {
				Message string `json:"message"`
			}{
				Message: "refresh in progress for view " + viewName,
			})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		writeJSON(ctx, w, "no refresh in progress for view "+viewName)
		return
	}

}

func RefreshView(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fmt.Println("refreshing view")
	ctx := initializeContext(r)
	authParts := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(ctx, w, struct {
			Message string `json:"message"`
		}{
			Message: "invalid authorization header",
		})
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
		writeJSON(ctx, w,
			struct {
				Message string `json:"message"`
			}{
				Message: "invalid api key",
			})
		return
	}

	// get query params
	viewName := ps.ByName("view_name")

	if viewName == viewTypeRefreshGlobalStats {
		go func() {
			_, err := dao.RefreshGlobalStatsView(viewName)
			if err != nil {
				returnError(ctx, w, r, err)
				return
			}
		}()
		// add tracking
		dao.ViewRefreshes[viewName] = true
	}

	if viewName == viewTypeRefreshAllTableViews {
		go func() {
			_, err := dao.RefreshGlobalAllTableView(viewName)
			if err != nil {
				returnError(ctx, w, r, err)
				return
			}
		}()
		dao.ViewRefreshes[viewName] = true
	}

	w.WriteHeader(http.StatusOK)
	writeJSON(ctx, w,
		struct {
			Message string `json:"message"`
		}{
			Message: "Refresh view request received. Please wait for a few minutes for the view to be refreshed.",
		},
	)
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
