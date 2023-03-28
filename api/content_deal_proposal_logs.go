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

func configContentDealProposalLogsRouter(router *httprouter.Router) {
	router.GET("/contentdealproposallogs", GetAllContentDealProposalLogs)
	router.POST("/contentdealproposallogs", AddContentDealProposalLogs)
	router.GET("/contentdealproposallogs/:argID", GetContentDealProposalLogs)
	router.PUT("/contentdealproposallogs/:argID", UpdateContentDealProposalLogs)
	router.DELETE("/contentdealproposallogs/:argID", DeleteContentDealProposalLogs)
}

func configGinContentDealProposalLogsRouter(router gin.IRoutes) {
	router.GET("/contentdealproposallogs", ConverHttprouterToGin(GetAllContentDealProposalLogs))
	router.POST("/contentdealproposallogs", ConverHttprouterToGin(AddContentDealProposalLogs))
	router.GET("/contentdealproposallogs/:argID", ConverHttprouterToGin(GetContentDealProposalLogs))
	router.PUT("/contentdealproposallogs/:argID", ConverHttprouterToGin(UpdateContentDealProposalLogs))
	router.DELETE("/contentdealproposallogs/:argID", ConverHttprouterToGin(DeleteContentDealProposalLogs))
}

// GetAllContentDealProposalLogs is a function to get a slice of record(s) from content_deal_proposal_logs table in the estuary database
// @Summary Get list of ContentDealProposalLogs
// @Tags ContentDealProposalLogs
// @Description GetAllContentDealProposalLogs is a handler to get a slice of record(s) from content_deal_proposal_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ContentDealProposalLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdealproposallogs [get]
// http "http://localhost:8080/contentdealproposallogs?page=0&pagesize=20" X-Api-User:user123
func GetAllContentDealProposalLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "content_deal_proposal_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentDealProposalLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetContentDealProposalLogs is a function to get a single record from the content_deal_proposal_logs table in the estuary database
// @Summary Get record from table ContentDealProposalLogs by  argID
// @Tags ContentDealProposalLogs
// @ID argID
// @Description GetContentDealProposalLogs is a function to get a single record from the content_deal_proposal_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ContentDealProposalLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentdealproposallogs/{argID} [get]
// http "http://localhost:8080/contentdealproposallogs/1" X-Api-User:user123
func GetContentDealProposalLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentDealProposalLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContentDealProposalLogs add to add a single record to content_deal_proposal_logs table in the estuary database
// @Summary Add an record to content_deal_proposal_logs table
// @Description add to add a single record to content_deal_proposal_logs table in the estuary database
// @Tags ContentDealProposalLogs
// @Accept  json
// @Produce  json
// @Param ContentDealProposalLogs body model.ContentDealProposalLogs true "Add ContentDealProposalLogs"
// @Success 200 {object} model.ContentDealProposalLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdealproposallogs [post]
// echo '{"id": 95,"content": 54,"unsigned": "BdPlZSqEDRmQVZWRcAStrfDVo","signed": "TFZfPHSlKLHIbkpGDLiTMOFBG","meta": "agqScMxqrltdqEfAJEnKPZNXY","node_info": "UGOjVwkfRJGHIvoIEhmErJIAJ","requester_info": "BcLEWYWZcXWfMctbJSsEjJEwE","requesting_api_key": "EtByipcglOXRuXtLBXFvEXUxv","system_content_deal_proposal_id": 89,"created_at": "2155-03-27T07:30:00.872085132-04:00","updated_at": "2158-08-27T22:03:34.596286539-04:00","delta_node_uuid": "xFrsHASTncOBrVcYtmntQWTIC"}' | http POST "http://localhost:8080/contentdealproposallogs" X-Api-User:user123
func AddContentDealProposalLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contentdealproposallogs := &model.ContentDealProposalLogs{}

	if err := readJSON(r, contentdealproposallogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdealproposallogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdealproposallogs.Prepare()

	if err := contentdealproposallogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contentdealproposallogs, _, err = dao.AddContentDealProposalLogs(ctx, contentdealproposallogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdealproposallogs)
}

// UpdateContentDealProposalLogs Update a single record from content_deal_proposal_logs table in the estuary database
// @Summary Update an record in table content_deal_proposal_logs
// @Description Update a single record from content_deal_proposal_logs table in the estuary database
// @Tags ContentDealProposalLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ContentDealProposalLogs body model.ContentDealProposalLogs true "Update ContentDealProposalLogs record"
// @Success 200 {object} model.ContentDealProposalLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdealproposallogs/{argID} [put]
// echo '{"id": 95,"content": 54,"unsigned": "BdPlZSqEDRmQVZWRcAStrfDVo","signed": "TFZfPHSlKLHIbkpGDLiTMOFBG","meta": "agqScMxqrltdqEfAJEnKPZNXY","node_info": "UGOjVwkfRJGHIvoIEhmErJIAJ","requester_info": "BcLEWYWZcXWfMctbJSsEjJEwE","requesting_api_key": "EtByipcglOXRuXtLBXFvEXUxv","system_content_deal_proposal_id": 89,"created_at": "2155-03-27T07:30:00.872085132-04:00","updated_at": "2158-08-27T22:03:34.596286539-04:00","delta_node_uuid": "xFrsHASTncOBrVcYtmntQWTIC"}' | http PUT "http://localhost:8080/contentdealproposallogs/1"  X-Api-User:user123
func UpdateContentDealProposalLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdealproposallogs := &model.ContentDealProposalLogs{}
	if err := readJSON(r, contentdealproposallogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdealproposallogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdealproposallogs.Prepare()

	if err := contentdealproposallogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdealproposallogs, _, err = dao.UpdateContentDealProposalLogs(ctx,
		argID,
		contentdealproposallogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdealproposallogs)
}

// DeleteContentDealProposalLogs Delete a single record from content_deal_proposal_logs table in the estuary database
// @Summary Delete a record from content_deal_proposal_logs
// @Description Delete a single record from content_deal_proposal_logs table in the estuary database
// @Tags ContentDealProposalLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ContentDealProposalLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentdealproposallogs/{argID} [delete]
// http DELETE "http://localhost:8080/contentdealproposallogs/1" X-Api-User:user123
func DeleteContentDealProposalLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentDealProposalLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
