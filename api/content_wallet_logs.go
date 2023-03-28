package api

import (
	"net/http"

	"github.com/application-research/delta-metrics-rest/dao"
	"github.com/application-research/delta-metrics-rest/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configContentWalletLogsRouter(router *httprouter.Router) {
	router.GET("/contentwalletlogs", GetAllContentWalletLogs)
	router.GET("/contentwalletlogs/:contenWalletLogsID", GetContentWalletLogs)
}

func configGinContentWalletLogsRouter(router gin.IRoutes) {
	router.GET("/contentwalletlogs", ConverHttprouterToGin(GetAllContentWalletLogs))
	router.GET("/contentwalletlogs/:contenWalletLogsID", ConverHttprouterToGin(GetContentWalletLogs))
}

// GetAllContentWalletLogs is a function to get a slice of record(s) from content_wallet_logs table in the estuary database
// @Summary Get list of ContentWalletLogs
// @Tags ContentWalletLogs
// @Description GetAllContentWalletLogs is a handler to get a slice of record(s) from content_wallet_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ContentWalletLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentwalletlogs [get]
// http "http://localhost:8080/contentwalletlogs?page=0&pagesize=20" X-Api-User:user123
func GetAllContentWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "content_wallet_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentWalletLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetContentWalletLogs is a function to get a single record from the content_wallet_logs table in the estuary database
// @Summary Get record from table ContentWalletLogs by  contenWalletLogsID
// @Tags ContentWalletLogs
// @ID contenWalletLogsID
// @Description GetContentWalletLogs is a function to get a single record from the content_wallet_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  contenWalletLogsID path int64 true "id"
// @Success 200 {object} model.ContentWalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentwalletlogs/{contenWalletLogsID} [get]
// http "http://localhost:8080/contentwalletlogs/1" X-Api-User:user123
func GetContentWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contenWalletLogsID, err := parseInt64(ps, "contenWalletLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_wallet_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentWalletLogs(ctx, contenWalletLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContentWalletLogs add to add a single record to content_wallet_logs table in the estuary database
// @Summary Add an record to content_wallet_logs table
// @Description add to add a single record to content_wallet_logs table in the estuary database
// @Tags ContentWalletLogs
// @Accept  json
// @Produce  json
// @Param ContentWalletLogs body model.ContentWalletLogs true "Add ContentWalletLogs"
// @Success 200 {object} model.ContentWalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentwalletlogs [post]
// echo '{"id": 32,"content": 30,"wallet": "IPWZJQDFqOtrZAiOqxXsTcxGA","nodeInfo": "bqfDOZUYMqlEMJMpMNWWCCAFU","requesterInfo": "RGnuoAwqrHeiSIbCNaAtinoEn","requestingApiKey": "oKedRqCcjjdtkhmRKWWoKYqkN","systemContentWalletId": 40,"createdAt": "2253-07-27T03:41:17.410938197-04:00","updatedAt": "2272-09-27T09:42:15.555605978-04:00","deltaNodeUuid": "MmyGprUgsHYiZKBoEcxYVUCrN","walletId": 32}' | http POST "http://localhost:8080/contentwalletlogs" X-Api-User:user123
func AddContentWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contentwalletlogs := &model.ContentWalletLogs{}

	if err := readJSON(r, contentwalletlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentwalletlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentwalletlogs.Prepare()

	if err := contentwalletlogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_wallet_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contentwalletlogs, _, err = dao.AddContentWalletLogs(ctx, contentwalletlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentwalletlogs)
}

// UpdateContentWalletLogs Update a single record from content_wallet_logs table in the estuary database
// @Summary Update an record in table content_wallet_logs
// @Description Update a single record from content_wallet_logs table in the estuary database
// @Tags ContentWalletLogs
// @Accept  json
// @Produce  json
// @Param  contenWalletLogsID path int64 true "id"
// @Param  ContentWalletLogs body model.ContentWalletLogs true "Update ContentWalletLogs record"
// @Success 200 {object} model.ContentWalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentwalletlogs/{contenWalletLogsID} [put]
// echo '{"id": 32,"content": 30,"wallet": "IPWZJQDFqOtrZAiOqxXsTcxGA","nodeInfo": "bqfDOZUYMqlEMJMpMNWWCCAFU","requesterInfo": "RGnuoAwqrHeiSIbCNaAtinoEn","requestingApiKey": "oKedRqCcjjdtkhmRKWWoKYqkN","systemContentWalletId": 40,"createdAt": "2253-07-27T03:41:17.410938197-04:00","updatedAt": "2272-09-27T09:42:15.555605978-04:00","deltaNodeUuid": "MmyGprUgsHYiZKBoEcxYVUCrN","walletId": 32}' | http PUT "http://localhost:8080/contentwalletlogs/1"  X-Api-User:user123
func UpdateContentWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contenWalletLogsID, err := parseInt64(ps, "contenWalletLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentwalletlogs := &model.ContentWalletLogs{}
	if err := readJSON(r, contentwalletlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentwalletlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentwalletlogs.Prepare()

	if err := contentwalletlogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_wallet_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentwalletlogs, _, err = dao.UpdateContentWalletLogs(ctx,
		contenWalletLogsID,
		contentwalletlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentwalletlogs)
}

// DeleteContentWalletLogs Delete a single record from content_wallet_logs table in the estuary database
// @Summary Delete a record from content_wallet_logs
// @Description Delete a single record from content_wallet_logs table in the estuary database
// @Tags ContentWalletLogs
// @Accept  json
// @Produce  json
// @Param  contenWalletLogsID path int64 true "id"
// @Success 204 {object} model.ContentWalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentwalletlogs/{contenWalletLogsID} [delete]
// http DELETE "http://localhost:8080/contentwalletlogs/1" X-Api-User:user123
func DeleteContentWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contenWalletLogsID, err := parseInt64(ps, "contenWalletLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_wallet_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentWalletLogs(ctx, contenWalletLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
