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

func configLogEventsRouter(router *httprouter.Router) {
	router.GET("/logevents", GetAllLogEvents)
	router.GET("/logevents/:logEventsID", GetLogEvents)
}

func configGinLogEventsRouter(router gin.IRoutes) {
	router.GET("/logevents", ConverHttprouterToGin(GetAllLogEvents))
	router.GET("/logevents/:logEventsID", ConverHttprouterToGin(GetLogEvents))
}

// GetAllLogEvents is a function to get a slice of record(s) from log_events table in the estuary database
// @Summary Get list of LogEvents
// @Tags LogEvents
// @Description GetAllLogEvents is a handler to get a slice of record(s) from log_events table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.LogEvents}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /logevents [get]
// http "http://localhost:8080/logevents?page=0&pagesize=20" X-Api-User:user123
func GetAllLogEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "log_events", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllLogEvents(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetLogEvents is a function to get a single record from the log_events table in the estuary database
// @Summary Get record from table LogEvents by  logEventsID
// @Tags LogEvents
// @ID logEventsID
// @Description GetLogEvents is a function to get a single record from the log_events table in the estuary database
// @Accept  json
// @Produce  json
// @Param  logEventsID path int64 true "id"
// @Success 200 {object} model.LogEvents
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /logevents/{logEventsID} [get]
// http "http://localhost:8080/logevents/1" X-Api-User:user123
func GetLogEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	logEventsID, err := parseInt64(ps, "logEventsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "log_events", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetLogEvents(ctx, logEventsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddLogEvents add to add a single record to log_events table in the estuary database
// @Summary Add an record to log_events table
// @Description add to add a single record to log_events table in the estuary database
// @Tags LogEvents
// @Accept  json
// @Produce  json
// @Param LogEvents body model.LogEvents true "Add LogEvents"
// @Success 200 {object} model.LogEvents
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /logevents [post]
// echo '{"id": 40,"logEventType": "utLCvdSTXAVpXmZBtJIFHfJTD","logEventObject": "WnAJsllMZuniavJewVbJDhPaM","logEventId": 11,"logEvent": "VoWiowlmuxqWgHNwOMrwCMnWt","createdAt": "2079-09-24T02:09:13.876912282-04:00","updatedAt": "2148-01-31T17:30:24.030190491-05:00","sourceHost": "rgljTEOgYPOGyYOxsLgOFLfQs","sourceIp": "rYhpAoCBhenbrcXTiOGyqTQMH","deltaUuid": "cyFGOCqwcETcuaAJpREQqAylk"}' | http POST "http://localhost:8080/logevents" X-Api-User:user123
func AddLogEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	logevents := &model.LogEvents{}

	if err := readJSON(r, logevents); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := logevents.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	logevents.Prepare()

	if err := logevents.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "log_events", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	logevents, _, err = dao.AddLogEvents(ctx, logevents)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, logevents)
}

// UpdateLogEvents Update a single record from log_events table in the estuary database
// @Summary Update an record in table log_events
// @Description Update a single record from log_events table in the estuary database
// @Tags LogEvents
// @Accept  json
// @Produce  json
// @Param  logEventsID path int64 true "id"
// @Param  LogEvents body model.LogEvents true "Update LogEvents record"
// @Success 200 {object} model.LogEvents
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /logevents/{logEventsID} [put]
// echo '{"id": 40,"logEventType": "utLCvdSTXAVpXmZBtJIFHfJTD","logEventObject": "WnAJsllMZuniavJewVbJDhPaM","logEventId": 11,"logEvent": "VoWiowlmuxqWgHNwOMrwCMnWt","createdAt": "2079-09-24T02:09:13.876912282-04:00","updatedAt": "2148-01-31T17:30:24.030190491-05:00","sourceHost": "rgljTEOgYPOGyYOxsLgOFLfQs","sourceIp": "rYhpAoCBhenbrcXTiOGyqTQMH","deltaUuid": "cyFGOCqwcETcuaAJpREQqAylk"}' | http PUT "http://localhost:8080/logevents/1"  X-Api-User:user123
func UpdateLogEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	logEventsID, err := parseInt64(ps, "logEventsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	logevents := &model.LogEvents{}
	if err := readJSON(r, logevents); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := logevents.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	logevents.Prepare()

	if err := logevents.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "log_events", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	logevents, _, err = dao.UpdateLogEvents(ctx,
		logEventsID,
		logevents)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, logevents)
}

// DeleteLogEvents Delete a single record from log_events table in the estuary database
// @Summary Delete a record from log_events
// @Description Delete a single record from log_events table in the estuary database
// @Tags LogEvents
// @Accept  json
// @Produce  json
// @Param  logEventsID path int64 true "id"
// @Success 204 {object} model.LogEvents
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /logevents/{logEventsID} [delete]
// http DELETE "http://localhost:8080/logevents/1" X-Api-User:user123
func DeleteLogEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	logEventsID, err := parseInt64(ps, "logEventsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "log_events", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteLogEvents(ctx, logEventsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
