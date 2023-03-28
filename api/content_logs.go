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

func configContentLogsRouter(router *httprouter.Router) {
	router.GET("/contentlogs", GetAllContentLogs)
	router.POST("/contentlogs", AddContentLogs)
	router.GET("/contentlogs/:argID", GetContentLogs)
	router.PUT("/contentlogs/:argID", UpdateContentLogs)
	router.DELETE("/contentlogs/:argID", DeleteContentLogs)
}

func configGinContentLogsRouter(router gin.IRoutes) {
	router.GET("/contentlogs", ConverHttprouterToGin(GetAllContentLogs))
	router.POST("/contentlogs", ConverHttprouterToGin(AddContentLogs))
	router.GET("/contentlogs/:argID", ConverHttprouterToGin(GetContentLogs))
	router.PUT("/contentlogs/:argID", ConverHttprouterToGin(UpdateContentLogs))
	router.DELETE("/contentlogs/:argID", ConverHttprouterToGin(DeleteContentLogs))
}

// GetAllContentLogs is a function to get a slice of record(s) from content_logs table in the estuary database
// @Summary Get list of ContentLogs
// @Tags ContentLogs
// @Description GetAllContentLogs is a handler to get a slice of record(s) from content_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ContentLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentlogs [get]
// http "http://localhost:8080/contentlogs?page=0&pagesize=20" X-Api-User:user123
func GetAllContentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "content_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetContentLogs is a function to get a single record from the content_logs table in the estuary database
// @Summary Get record from table ContentLogs by  argID
// @Tags ContentLogs
// @ID argID
// @Description GetContentLogs is a function to get a single record from the content_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ContentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentlogs/{argID} [get]
// http "http://localhost:8080/contentlogs/1" X-Api-User:user123
func GetContentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContentLogs add to add a single record to content_logs table in the estuary database
// @Summary Add an record to content_logs table
// @Description add to add a single record to content_logs table in the estuary database
// @Tags ContentLogs
// @Accept  json
// @Produce  json
// @Param ContentLogs body model.ContentLogs true "Add ContentLogs"
// @Success 200 {object} model.ContentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentlogs [post]
// echo '{"id": 90,"name": "ASnOqmbJwxJRhmRbnOGDJixaK","size": 46,"cid": "fwOkmQhekAyhfPXeyAsceSMfE","requesting_api_key": "wCYsLEJxerQPPSjbfZOpBkqRF","piece_commitment_id": 17,"status": "qFHjvjNejYOLPAmldrTyyCSRP","connection_mode": "yyLqyHFuEHxjlNPfNAvKCkmkY","last_message": "QoxDAOYZpaxSxYBDjRFERtYsJ","node_info": "GmPwLHeTToiHGixmnlFYsGyQW","requester_info": "qNmEOqjASNXSRExXKilJriyCh","system_content_id": 87,"created_at": "2224-09-27T23:57:35.226433687-04:00","updated_at": "2223-10-07T22:03:30.973380902-04:00","delta_node_uuid": "asobGhiCRNLAmxIUxePboDwlT"}' | http POST "http://localhost:8080/contentlogs" X-Api-User:user123
func AddContentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contentlogs := &model.ContentLogs{}

	if err := readJSON(r, contentlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentlogs.Prepare()

	if err := contentlogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contentlogs, _, err = dao.AddContentLogs(ctx, contentlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentlogs)
}

// UpdateContentLogs Update a single record from content_logs table in the estuary database
// @Summary Update an record in table content_logs
// @Description Update a single record from content_logs table in the estuary database
// @Tags ContentLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ContentLogs body model.ContentLogs true "Update ContentLogs record"
// @Success 200 {object} model.ContentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentlogs/{argID} [put]
// echo '{"id": 90,"name": "ASnOqmbJwxJRhmRbnOGDJixaK","size": 46,"cid": "fwOkmQhekAyhfPXeyAsceSMfE","requesting_api_key": "wCYsLEJxerQPPSjbfZOpBkqRF","piece_commitment_id": 17,"status": "qFHjvjNejYOLPAmldrTyyCSRP","connection_mode": "yyLqyHFuEHxjlNPfNAvKCkmkY","last_message": "QoxDAOYZpaxSxYBDjRFERtYsJ","node_info": "GmPwLHeTToiHGixmnlFYsGyQW","requester_info": "qNmEOqjASNXSRExXKilJriyCh","system_content_id": 87,"created_at": "2224-09-27T23:57:35.226433687-04:00","updated_at": "2223-10-07T22:03:30.973380902-04:00","delta_node_uuid": "asobGhiCRNLAmxIUxePboDwlT"}' | http PUT "http://localhost:8080/contentlogs/1"  X-Api-User:user123
func UpdateContentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentlogs := &model.ContentLogs{}
	if err := readJSON(r, contentlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentlogs.Prepare()

	if err := contentlogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentlogs, _, err = dao.UpdateContentLogs(ctx,
		argID,
		contentlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentlogs)
}

// DeleteContentLogs Delete a single record from content_logs table in the estuary database
// @Summary Delete a record from content_logs
// @Description Delete a single record from content_logs table in the estuary database
// @Tags ContentLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ContentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentlogs/{argID} [delete]
// http DELETE "http://localhost:8080/contentlogs/1" X-Api-User:user123
func DeleteContentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
