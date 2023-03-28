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

func configDeltaStartupLogsRouter(router *httprouter.Router) {
	router.GET("/deltastartuplogs", GetAllDeltaStartupLogs)
	router.GET("/deltastartuplogs/:deltaStartupLogsID", GetDeltaStartupLogs)
}

func configGinDeltaStartupLogsRouter(router gin.IRoutes) {
	router.GET("/deltastartuplogs", ConverHttprouterToGin(GetAllDeltaStartupLogs))
	router.GET("/deltastartuplogs/:deltaStartupLogsID", ConverHttprouterToGin(GetDeltaStartupLogs))
}

// GetAllDeltaStartupLogs is a function to get a slice of record(s) from delta_startup_logs table in the estuary database
// @Summary Get list of DeltaStartupLogs
// @Tags DeltaStartupLogs
// @Description GetAllDeltaStartupLogs is a handler to get a slice of record(s) from delta_startup_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.DeltaStartupLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /deltastartuplogs [get]
// http "http://localhost:8080/deltastartuplogs?page=0&pagesize=20" X-Api-User:user123
func GetAllDeltaStartupLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "delta_startup_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDeltaStartupLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetDeltaStartupLogs is a function to get a single record from the delta_startup_logs table in the estuary database
// @Summary Get record from table DeltaStartupLogs by  deltaStartupLogsID
// @Tags DeltaStartupLogs
// @ID deltaStartupLogsID
// @Description GetDeltaStartupLogs is a function to get a single record from the delta_startup_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  deltaStartupLogsID path int64 true "id"
// @Success 200 {object} model.DeltaStartupLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /deltastartuplogs/{deltaStartupLogsID} [get]
// http "http://localhost:8080/deltastartuplogs/1" X-Api-User:user123
func GetDeltaStartupLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	deltaStartupLogsID, err := parseInt64(ps, "deltaStartupLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_startup_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDeltaStartupLogs(ctx, deltaStartupLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddDeltaStartupLogs add to add a single record to delta_startup_logs table in the estuary database
// @Summary Add an record to delta_startup_logs table
// @Description add to add a single record to delta_startup_logs table in the estuary database
// @Tags DeltaStartupLogs
// @Accept  json
// @Produce  json
// @Param DeltaStartupLogs body model.DeltaStartupLogs true "Add DeltaStartupLogs"
// @Success 200 {object} model.DeltaStartupLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /deltastartuplogs [post]
// echo '{"id": 17,"nodeInfo": "MiifxOjnqsUpOjndrGokBhUuL","osDetails": "fIDOfHOZfJMMKAdBmIxiBvYpC","ipAddress": "VogTZvQFdkLBQruBIclgLKiTa","createdAt": "2299-11-22T20:06:45.836550913-05:00","updatedAt": "2059-05-30T21:42:11.112312668-04:00","deltaNodeUuid": "DSimhMAQGdBANWNucLQNitfKW"}' | http POST "http://localhost:8080/deltastartuplogs" X-Api-User:user123
func AddDeltaStartupLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	deltastartuplogs := &model.DeltaStartupLogs{}

	if err := readJSON(r, deltastartuplogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := deltastartuplogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	deltastartuplogs.Prepare()

	if err := deltastartuplogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_startup_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	deltastartuplogs, _, err = dao.AddDeltaStartupLogs(ctx, deltastartuplogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, deltastartuplogs)
}

// UpdateDeltaStartupLogs Update a single record from delta_startup_logs table in the estuary database
// @Summary Update an record in table delta_startup_logs
// @Description Update a single record from delta_startup_logs table in the estuary database
// @Tags DeltaStartupLogs
// @Accept  json
// @Produce  json
// @Param  deltaStartupLogsID path int64 true "id"
// @Param  DeltaStartupLogs body model.DeltaStartupLogs true "Update DeltaStartupLogs record"
// @Success 200 {object} model.DeltaStartupLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /deltastartuplogs/{deltaStartupLogsID} [put]
// echo '{"id": 17,"nodeInfo": "MiifxOjnqsUpOjndrGokBhUuL","osDetails": "fIDOfHOZfJMMKAdBmIxiBvYpC","ipAddress": "VogTZvQFdkLBQruBIclgLKiTa","createdAt": "2299-11-22T20:06:45.836550913-05:00","updatedAt": "2059-05-30T21:42:11.112312668-04:00","deltaNodeUuid": "DSimhMAQGdBANWNucLQNitfKW"}' | http PUT "http://localhost:8080/deltastartuplogs/1"  X-Api-User:user123
func UpdateDeltaStartupLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	deltaStartupLogsID, err := parseInt64(ps, "deltaStartupLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	deltastartuplogs := &model.DeltaStartupLogs{}
	if err := readJSON(r, deltastartuplogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := deltastartuplogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	deltastartuplogs.Prepare()

	if err := deltastartuplogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_startup_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	deltastartuplogs, _, err = dao.UpdateDeltaStartupLogs(ctx,
		deltaStartupLogsID,
		deltastartuplogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, deltastartuplogs)
}

// DeleteDeltaStartupLogs Delete a single record from delta_startup_logs table in the estuary database
// @Summary Delete a record from delta_startup_logs
// @Description Delete a single record from delta_startup_logs table in the estuary database
// @Tags DeltaStartupLogs
// @Accept  json
// @Produce  json
// @Param  deltaStartupLogsID path int64 true "id"
// @Success 204 {object} model.DeltaStartupLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /deltastartuplogs/{deltaStartupLogsID} [delete]
// http DELETE "http://localhost:8080/deltastartuplogs/1" X-Api-User:user123
func DeleteDeltaStartupLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	deltaStartupLogsID, err := parseInt64(ps, "deltaStartupLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_startup_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDeltaStartupLogs(ctx, deltaStartupLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
