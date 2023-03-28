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
	router.POST("/walletlogs", AddWalletLogs)
	router.GET("/walletlogs/:argID", GetWalletLogs)
	router.PUT("/walletlogs/:argID", UpdateWalletLogs)
	router.DELETE("/walletlogs/:argID", DeleteWalletLogs)
}

func configGinWalletLogsRouter(router gin.IRoutes) {
	router.GET("/walletlogs", ConverHttprouterToGin(GetAllWalletLogs))
	router.POST("/walletlogs", ConverHttprouterToGin(AddWalletLogs))
	router.GET("/walletlogs/:argID", ConverHttprouterToGin(GetWalletLogs))
	router.PUT("/walletlogs/:argID", ConverHttprouterToGin(UpdateWalletLogs))
	router.DELETE("/walletlogs/:argID", ConverHttprouterToGin(DeleteWalletLogs))
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
// @Summary Get record from table WalletLogs by  argID
// @Tags WalletLogs
// @ID argID
// @Description GetWalletLogs is a function to get a single record from the wallet_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.WalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /walletlogs/{argID} [get]
// http "http://localhost:8080/walletlogs/1" X-Api-User:user123
func GetWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "wallet_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetWalletLogs(ctx, argID)
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
// echo '{"id": 91,"uu_id": "cSPqDInnUCcpbDAmBOZaVvAOk","addr": "uAKwIKMigAcDoaniswuHEpLyk","owner": "MFjujRvaHPJvnYVCqsZikVXBD","key_type": "GsAKUbVMEAubZxEDDDrAqQGtG","private_key": "fXKujNchwKqimxMFvhflvVylr","node_info": "iaoWVnvacvZASYueskgQlULsQ","requester_info": "CRhHWgtTwaXQtbuAyFjsnmIwf","requesting_api_key": "uhDdMUkIhxkNvySVicLZHDWnh","system_wallet_id": 72,"created_at": "2043-05-31T02:35:05.956558122-04:00","updated_at": "2225-04-02T04:00:20.955268059-04:00","delta_node_uuid": "gHCjAZsyEKkMrELNIUwoccdsy"}' | http POST "http://localhost:8080/walletlogs" X-Api-User:user123
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
// @Param  argID path int64 true "id"
// @Param  WalletLogs body model.WalletLogs true "Update WalletLogs record"
// @Success 200 {object} model.WalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /walletlogs/{argID} [put]
// echo '{"id": 91,"uu_id": "cSPqDInnUCcpbDAmBOZaVvAOk","addr": "uAKwIKMigAcDoaniswuHEpLyk","owner": "MFjujRvaHPJvnYVCqsZikVXBD","key_type": "GsAKUbVMEAubZxEDDDrAqQGtG","private_key": "fXKujNchwKqimxMFvhflvVylr","node_info": "iaoWVnvacvZASYueskgQlULsQ","requester_info": "CRhHWgtTwaXQtbuAyFjsnmIwf","requesting_api_key": "uhDdMUkIhxkNvySVicLZHDWnh","system_wallet_id": 72,"created_at": "2043-05-31T02:35:05.956558122-04:00","updated_at": "2225-04-02T04:00:20.955268059-04:00","delta_node_uuid": "gHCjAZsyEKkMrELNIUwoccdsy"}' | http PUT "http://localhost:8080/walletlogs/1"  X-Api-User:user123
func UpdateWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
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
		argID,
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
// @Param  argID path int64 true "id"
// @Success 204 {object} model.WalletLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /walletlogs/{argID} [delete]
// http DELETE "http://localhost:8080/walletlogs/1" X-Api-User:user123
func DeleteWalletLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "wallet_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteWalletLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
