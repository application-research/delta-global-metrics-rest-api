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

func configContentDealLogsRouter(router *httprouter.Router) {
	router.GET("/contentdeallogs", GetAllContentDealLogs)
	router.POST("/contentdeallogs", AddContentDealLogs)
	router.GET("/contentdeallogs/:argID", GetContentDealLogs)
	router.PUT("/contentdeallogs/:argID", UpdateContentDealLogs)
	router.DELETE("/contentdeallogs/:argID", DeleteContentDealLogs)
}

func configGinContentDealLogsRouter(router gin.IRoutes) {
	router.GET("/contentdeallogs", ConverHttprouterToGin(GetAllContentDealLogs))
	router.POST("/contentdeallogs", ConverHttprouterToGin(AddContentDealLogs))
	router.GET("/contentdeallogs/:argID", ConverHttprouterToGin(GetContentDealLogs))
	router.PUT("/contentdeallogs/:argID", ConverHttprouterToGin(UpdateContentDealLogs))
	router.DELETE("/contentdeallogs/:argID", ConverHttprouterToGin(DeleteContentDealLogs))
}

// GetAllContentDealLogs is a function to get a slice of record(s) from content_deal_logs table in the estuary database
// @Summary Get list of ContentDealLogs
// @Tags ContentDealLogs
// @Description GetAllContentDealLogs is a handler to get a slice of record(s) from content_deal_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ContentDealLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdeallogs [get]
// http "http://localhost:8080/contentdeallogs?page=0&pagesize=20" X-Api-User:user123
func GetAllContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "content_deal_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentDealLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetContentDealLogs is a function to get a single record from the content_deal_logs table in the estuary database
// @Summary Get record from table ContentDealLogs by  argID
// @Tags ContentDealLogs
// @ID argID
// @Description GetContentDealLogs is a function to get a single record from the content_deal_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ContentDealLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentdeallogs/{argID} [get]
// http "http://localhost:8080/contentdeallogs/1" X-Api-User:user123
func GetContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentDealLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContentDealLogs add to add a single record to content_deal_logs table in the estuary database
// @Summary Add an record to content_deal_logs table
// @Description add to add a single record to content_deal_logs table in the estuary database
// @Tags ContentDealLogs
// @Accept  json
// @Produce  json
// @Param ContentDealLogs body model.ContentDealLogs true "Add ContentDealLogs"
// @Success 200 {object} model.ContentDealLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdeallogs [post]
// echo '{"id": 78,"content": 50,"prop_cid": "wPitHmApMJNqFkbCVCEXLPOgO","deal_uuid": "KfpfjKjQsUEdUlCKdpDmZPHSm","miner": "oNPSjFMgtvtqevtwSaCronoHs","deal_id": 69,"failed": false,"verified": false,"slashed": false,"failed_at": "2073-03-13T07:58:47.787105952-04:00","dt_chan": "gbajoJITYJHctKypowgyQwouy","transfer_started": "2288-03-25T04:38:08.461808637-04:00","transfer_finished": "2054-06-03T09:38:09.659274379-04:00","on_chain_at": "2242-09-05T02:11:35.872053883-04:00","sealed_at": "2029-08-07T07:31:54.707466325-04:00","last_message": "fJmUPhsRNdLSGkcfdTaGGgnAX","deal_protocol_version": "etTRJJaJfuZJNqcywDOanNvWe","miner_version": "XjJCulvZcXNXKflGaedJVMgFF","node_info": "onYUJQpmPnpMgOHnWwjaWDeUd","requester_info": "NGGHKoXWAibSANLnomZsFWOQB","requesting_api_key": "njTDXKscmeIOOHGiipNNRcYyH","system_content_deal_id": 1,"created_at": "2226-07-15T00:10:30.581336342-04:00","updated_at": "2098-07-14T10:42:36.953905498-04:00","delta_node_uuid": "HbTqKrUqmSCnaYUYmRoPtipKX"}' | http POST "http://localhost:8080/contentdeallogs" X-Api-User:user123
func AddContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contentdeallogs := &model.ContentDealLogs{}

	if err := readJSON(r, contentdeallogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdeallogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdeallogs.Prepare()

	if err := contentdeallogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contentdeallogs, _, err = dao.AddContentDealLogs(ctx, contentdeallogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdeallogs)
}

// UpdateContentDealLogs Update a single record from content_deal_logs table in the estuary database
// @Summary Update an record in table content_deal_logs
// @Description Update a single record from content_deal_logs table in the estuary database
// @Tags ContentDealLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ContentDealLogs body model.ContentDealLogs true "Update ContentDealLogs record"
// @Success 200 {object} model.ContentDealLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdeallogs/{argID} [put]
// echo '{"id": 78,"content": 50,"prop_cid": "wPitHmApMJNqFkbCVCEXLPOgO","deal_uuid": "KfpfjKjQsUEdUlCKdpDmZPHSm","miner": "oNPSjFMgtvtqevtwSaCronoHs","deal_id": 69,"failed": false,"verified": false,"slashed": false,"failed_at": "2073-03-13T07:58:47.787105952-04:00","dt_chan": "gbajoJITYJHctKypowgyQwouy","transfer_started": "2288-03-25T04:38:08.461808637-04:00","transfer_finished": "2054-06-03T09:38:09.659274379-04:00","on_chain_at": "2242-09-05T02:11:35.872053883-04:00","sealed_at": "2029-08-07T07:31:54.707466325-04:00","last_message": "fJmUPhsRNdLSGkcfdTaGGgnAX","deal_protocol_version": "etTRJJaJfuZJNqcywDOanNvWe","miner_version": "XjJCulvZcXNXKflGaedJVMgFF","node_info": "onYUJQpmPnpMgOHnWwjaWDeUd","requester_info": "NGGHKoXWAibSANLnomZsFWOQB","requesting_api_key": "njTDXKscmeIOOHGiipNNRcYyH","system_content_deal_id": 1,"created_at": "2226-07-15T00:10:30.581336342-04:00","updated_at": "2098-07-14T10:42:36.953905498-04:00","delta_node_uuid": "HbTqKrUqmSCnaYUYmRoPtipKX"}' | http PUT "http://localhost:8080/contentdeallogs/1"  X-Api-User:user123
func UpdateContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdeallogs := &model.ContentDealLogs{}
	if err := readJSON(r, contentdeallogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdeallogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdeallogs.Prepare()

	if err := contentdeallogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdeallogs, _, err = dao.UpdateContentDealLogs(ctx,
		argID,
		contentdeallogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdeallogs)
}

// DeleteContentDealLogs Delete a single record from content_deal_logs table in the estuary database
// @Summary Delete a record from content_deal_logs
// @Description Delete a single record from content_deal_logs table in the estuary database
// @Tags ContentDealLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ContentDealLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentdeallogs/{argID} [delete]
// http DELETE "http://localhost:8080/contentdeallogs/1" X-Api-User:user123
func DeleteContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentDealLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
