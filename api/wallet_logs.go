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

func configWalletLogsRouter(router *httprouter.Router) {
	router.GET("/walletlogs", GetAllWalletLogs)
	router.GET("/walletlogs/:walletLogsID", GetWalletLogs)
}

func configGinWalletLogsRouter(router gin.IRoutes) {
	router.GET("/walletlogs", ConverHttprouterToGin(GetAllWalletLogs))
	router.GET("/walletlogs/:walletLogsID", ConverHttprouterToGin(GetWalletLogs))
}

// GetAllWalletLogs is a function to get a slice of record(s) from wallet_logs table in the estuary database
// @Summary Get list of WalletLogs
// @Tags WalletLogs
// @Description GetAllWalletLogs is a handler to get a slice of record(s) from wallet_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.WalletLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /walletlogs [get]
// http "http://localhost:8080/walletlogs?page=0&pagesize=20" X-Api-User:user123
func GetAllWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "wallet_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllWalletLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetWalletLogs is a function to get a single record from the wallet_logs table in the estuary database
// @Summary Get record from table WalletLogs by  walletLogsID
// @Tags WalletLogs
// @ID walletLogsID
// @Description GetWalletLogs is a function to get a single record from the wallet_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  walletLogsID path int64 true "id"
// @Success 200 {object} model.WalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /walletlogs/{walletLogsID} [get]
// http "http://localhost:8080/walletlogs/1" X-Api-User:user123
func GetWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	walletLogsID, err := parseInt64(ps, "walletLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "wallet_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetWalletLogs(ctx, walletLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddWalletLogs add to add a single record to wallet_logs table in the estuary database
// @Summary Add an record to wallet_logs table
// @Description add to add a single record to wallet_logs table in the estuary database
// @Tags WalletLogs
// @Accept  json
// @Produce  json
// @Param WalletLogs body model.WalletLogs true "Add WalletLogs"
// @Success 200 {object} model.WalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /walletlogs [post]
// echo '{"id": 80,"uuId": "peFbMgHivRACqmNCvSorweBqP","addr": "DZHnFIoBEmueGkCKsbhBavqtY","owner": "EiInDJEyANGkwlltXGTjhuYGx","keyType": "vDZKOPdjBbTrnnYuaGJHQaYtJ","privateKey": "wUkalhcPceXJobZGwJmepOywB","nodeInfo": "vgQknMVpqfrrAcQbfEKeknEyr","requesterInfo": "oByfGPwsEqniJLssjQCfDFPAQ","requestingApiKey": "MRemQfcdJkewUiFHrosfOFoKp","systemWalletId": 70,"createdAt": "2067-08-04T12:59:08.474556194-04:00","updatedAt": "2094-09-21T03:52:50.420485709-04:00","deltaNodeUuid": "BJLhHabuFNncXQTqTrxbJQkLE"}' | http POST "http://localhost:8080/walletlogs" X-Api-User:user123
func AddWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	walletlogs := &model.WalletLogs{}

	if err := readJSON(r, walletlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := walletlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	walletlogs.Prepare()

	if err := walletlogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "wallet_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	walletlogs, _, err = dao.AddWalletLogs(ctx, walletlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, walletlogs)
}

// UpdateWalletLogs Update a single record from wallet_logs table in the estuary database
// @Summary Update an record in table wallet_logs
// @Description Update a single record from wallet_logs table in the estuary database
// @Tags WalletLogs
// @Accept  json
// @Produce  json
// @Param  walletLogsID path int64 true "id"
// @Param  WalletLogs body model.WalletLogs true "Update WalletLogs record"
// @Success 200 {object} model.WalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /walletlogs/{walletLogsID} [put]
// echo '{"id": 80,"uuId": "peFbMgHivRACqmNCvSorweBqP","addr": "DZHnFIoBEmueGkCKsbhBavqtY","owner": "EiInDJEyANGkwlltXGTjhuYGx","keyType": "vDZKOPdjBbTrnnYuaGJHQaYtJ","privateKey": "wUkalhcPceXJobZGwJmepOywB","nodeInfo": "vgQknMVpqfrrAcQbfEKeknEyr","requesterInfo": "oByfGPwsEqniJLssjQCfDFPAQ","requestingApiKey": "MRemQfcdJkewUiFHrosfOFoKp","systemWalletId": 70,"createdAt": "2067-08-04T12:59:08.474556194-04:00","updatedAt": "2094-09-21T03:52:50.420485709-04:00","deltaNodeUuid": "BJLhHabuFNncXQTqTrxbJQkLE"}' | http PUT "http://localhost:8080/walletlogs/1"  X-Api-User:user123
func UpdateWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	walletLogsID, err := parseInt64(ps, "walletLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	walletlogs := &model.WalletLogs{}
	if err := readJSON(r, walletlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := walletlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	walletlogs.Prepare()

	if err := walletlogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "wallet_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	walletlogs, _, err = dao.UpdateWalletLogs(ctx,
		walletLogsID,
		walletlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, walletlogs)
}

// DeleteWalletLogs Delete a single record from wallet_logs table in the estuary database
// @Summary Delete a record from wallet_logs
// @Description Delete a single record from wallet_logs table in the estuary database
// @Tags WalletLogs
// @Accept  json
// @Produce  json
// @Param  walletLogsID path int64 true "id"
// @Success 204 {object} model.WalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /walletlogs/{walletLogsID} [delete]
// http DELETE "http://localhost:8080/walletlogs/1" X-Api-User:user123
func DeleteWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	walletLogsID, err := parseInt64(ps, "walletLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "wallet_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteWalletLogs(ctx, walletLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
