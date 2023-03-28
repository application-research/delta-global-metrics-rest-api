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

func configContentMinerLogsRouter(router *httprouter.Router) {
	router.GET("/contentminerlogs", GetAllContentMinerLogs)
	router.GET("/contentminerlogs/:contenMinerLogsID", GetContentMinerLogs)
}

func configGinContentMinerLogsRouter(router gin.IRoutes) {
	router.GET("/contentminerlogs", ConverHttprouterToGin(GetAllContentMinerLogs))
	router.GET("/contentminerlogs/:contenMinerLogsID", ConverHttprouterToGin(GetContentMinerLogs))
}

// GetAllContentMinerLogs is a function to get a slice of record(s) from content_miner_logs table in the estuary database
// @Summary Get list of ContentMinerLogs
// @Tags ContentMinerLogs
// @Description GetAllContentMinerLogs is a handler to get a slice of record(s) from content_miner_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ContentMinerLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentminerlogs [get]
// http "http://localhost:8080/contentminerlogs?page=0&pagesize=20" X-Api-User:user123
func GetAllContentMinerLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "content_miner_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentMinerLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetContentMinerLogs is a function to get a single record from the content_miner_logs table in the estuary database
// @Summary Get record from table ContentMinerLogs by  contenMinerLogsID
// @Tags ContentMinerLogs
// @ID contenMinerLogsID
// @Description GetContentMinerLogs is a function to get a single record from the content_miner_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  contenMinerLogsID path int64 true "id"
// @Success 200 {object} model.ContentMinerLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentminerlogs/{contenMinerLogsID} [get]
// http "http://localhost:8080/contentminerlogs/1" X-Api-User:user123
func GetContentMinerLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contenMinerLogsID, err := parseInt64(ps, "contenMinerLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_miner_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentMinerLogs(ctx, contenMinerLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContentMinerLogs add to add a single record to content_miner_logs table in the estuary database
// @Summary Add an record to content_miner_logs table
// @Description add to add a single record to content_miner_logs table in the estuary database
// @Tags ContentMinerLogs
// @Accept  json
// @Produce  json
// @Param ContentMinerLogs body model.ContentMinerLogs true "Add ContentMinerLogs"
// @Success 200 {object} model.ContentMinerLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentminerlogs [post]
// echo '{"id": 69,"content": 0,"miner": "NXGJnUGafqmmNBPShVEUEfmmp","nodeInfo": "WDxXEDCrIJYyvsBAVKmNeiPEZ","requesterInfo": "DFRSNyUcGDJhibePMKhtctiaZ","requestingApiKey": "xNwUUtBHNrehAgaLbqFLwmxgh","systemContentMinerId": 7,"createdAt": "2257-06-15T10:59:42.700134707-04:00","updatedAt": "2207-11-18T23:49:37.467225423-05:00","deltaNodeUuid": "jIkpxjIHoCSjlPeyacQcPuKtJ"}' | http POST "http://localhost:8080/contentminerlogs" X-Api-User:user123
func AddContentMinerLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contentminerlogs := &model.ContentMinerLogs{}

	if err := readJSON(r, contentminerlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentminerlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentminerlogs.Prepare()

	if err := contentminerlogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_miner_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contentminerlogs, _, err = dao.AddContentMinerLogs(ctx, contentminerlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentminerlogs)
}

// UpdateContentMinerLogs Update a single record from content_miner_logs table in the estuary database
// @Summary Update an record in table content_miner_logs
// @Description Update a single record from content_miner_logs table in the estuary database
// @Tags ContentMinerLogs
// @Accept  json
// @Produce  json
// @Param  contenMinerLogsID path int64 true "id"
// @Param  ContentMinerLogs body model.ContentMinerLogs true "Update ContentMinerLogs record"
// @Success 200 {object} model.ContentMinerLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentminerlogs/{contenMinerLogsID} [put]
// echo '{"id": 69,"content": 0,"miner": "NXGJnUGafqmmNBPShVEUEfmmp","nodeInfo": "WDxXEDCrIJYyvsBAVKmNeiPEZ","requesterInfo": "DFRSNyUcGDJhibePMKhtctiaZ","requestingApiKey": "xNwUUtBHNrehAgaLbqFLwmxgh","systemContentMinerId": 7,"createdAt": "2257-06-15T10:59:42.700134707-04:00","updatedAt": "2207-11-18T23:49:37.467225423-05:00","deltaNodeUuid": "jIkpxjIHoCSjlPeyacQcPuKtJ"}' | http PUT "http://localhost:8080/contentminerlogs/1"  X-Api-User:user123
func UpdateContentMinerLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contenMinerLogsID, err := parseInt64(ps, "contenMinerLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentminerlogs := &model.ContentMinerLogs{}
	if err := readJSON(r, contentminerlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentminerlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentminerlogs.Prepare()

	if err := contentminerlogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_miner_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentminerlogs, _, err = dao.UpdateContentMinerLogs(ctx,
		contenMinerLogsID,
		contentminerlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentminerlogs)
}

// DeleteContentMinerLogs Delete a single record from content_miner_logs table in the estuary database
// @Summary Delete a record from content_miner_logs
// @Description Delete a single record from content_miner_logs table in the estuary database
// @Tags ContentMinerLogs
// @Accept  json
// @Produce  json
// @Param  contenMinerLogsID path int64 true "id"
// @Success 204 {object} model.ContentMinerLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentminerlogs/{contenMinerLogsID} [delete]
// http DELETE "http://localhost:8080/contentminerlogs/1" X-Api-User:user123
func DeleteContentMinerLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contenMinerLogsID, err := parseInt64(ps, "contenMinerLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_miner_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentMinerLogs(ctx, contenMinerLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
