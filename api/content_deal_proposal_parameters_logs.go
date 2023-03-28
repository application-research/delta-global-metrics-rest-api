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

func configContentDealProposalParametersLogsRouter(router *httprouter.Router) {
	router.GET("/contentdealproposalparameterslogs", GetAllContentDealProposalParametersLogs)
	router.GET("/contentdealproposalparameterslogs/:contentDealProposalParametersLogsID", GetContentDealProposalParametersLogs)
}

func configGinContentDealProposalParametersLogsRouter(router gin.IRoutes) {
	router.GET("/contentdealproposalparameterslogs", ConverHttprouterToGin(GetAllContentDealProposalParametersLogs))
	router.GET("/contentdealproposalparameterslogs/:contentDealProposalParametersLogsID", ConverHttprouterToGin(GetContentDealProposalParametersLogs))
}

// GetAllContentDealProposalParametersLogs is a function to get a slice of record(s) from content_deal_proposal_parameters_logs table in the estuary database
// @Summary Get list of ContentDealProposalParametersLogs
// @Tags ContentDealProposalParametersLogs
// @Description GetAllContentDealProposalParametersLogs is a handler to get a slice of record(s) from content_deal_proposal_parameters_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ContentDealProposalParametersLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdealproposalparameterslogs [get]
// http "http://localhost:8080/contentdealproposalparameterslogs?page=0&pagesize=20" X-Api-User:user123
func GetAllContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "content_deal_proposal_parameters_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentDealProposalParametersLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetContentDealProposalParametersLogs is a function to get a single record from the content_deal_proposal_parameters_logs table in the estuary database
// @Summary Get record from table ContentDealProposalParametersLogs by  contentDealProposalParametersLogsID
// @Tags ContentDealProposalParametersLogs
// @ID contentDealProposalParametersLogsID
// @Description GetContentDealProposalParametersLogs is a function to get a single record from the content_deal_proposal_parameters_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  contentDealProposalParametersLogsID path int64 true "id"
// @Success 200 {object} model.ContentDealProposalParametersLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentdealproposalparameterslogs/{contentDealProposalParametersLogsID} [get]
// http "http://localhost:8080/contentdealproposalparameterslogs/1" X-Api-User:user123
func GetContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contentDealProposalParametersLogsID, err := parseInt64(ps, "contentDealProposalParametersLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_parameters_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentDealProposalParametersLogs(ctx, contentDealProposalParametersLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContentDealProposalParametersLogs add to add a single record to content_deal_proposal_parameters_logs table in the estuary database
// @Summary Add an record to content_deal_proposal_parameters_logs table
// @Description add to add a single record to content_deal_proposal_parameters_logs table in the estuary database
// @Tags ContentDealProposalParametersLogs
// @Accept  json
// @Produce  json
// @Param ContentDealProposalParametersLogs body model.ContentDealProposalParametersLogs true "Add ContentDealProposalParametersLogs"
// @Success 200 {object} model.ContentDealProposalParametersLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdealproposalparameterslogs [post]
// echo '{"id": 11,"content": 28,"label": "gbgXFsMvnaClNRAqhtQdtFhNQ","duration": 43,"startEpoch": 93,"endEpoch": 15,"transferParams": "hDsSoFfCmVRZJFXIPJYBAxGLK","removeUnsealedCopy": false,"skipIpNiAnnounce": true,"nodeInfo": "idaYRkrLBhLFHYQAZbeFIkcXi","requesterInfo": "IwXjAiusvSYuTSiLFiYOJTTCO","requestingApiKey": "TfZrGwdaEivZEPgRZouqkMoVs","systemContentDealProposalParametersId": 54,"createdAt": "2051-07-18T14:58:51.12809116-04:00","updatedAt": "2196-07-15T21:02:52.682828298-04:00","deltaNodeUuid": "vXwPAuNtmYAUodvoDsZMFPAjG"}' | http POST "http://localhost:8080/contentdealproposalparameterslogs" X-Api-User:user123
func AddContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contentdealproposalparameterslogs := &model.ContentDealProposalParametersLogs{}

	if err := readJSON(r, contentdealproposalparameterslogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdealproposalparameterslogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdealproposalparameterslogs.Prepare()

	if err := contentdealproposalparameterslogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_parameters_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contentdealproposalparameterslogs, _, err = dao.AddContentDealProposalParametersLogs(ctx, contentdealproposalparameterslogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdealproposalparameterslogs)
}

// UpdateContentDealProposalParametersLogs Update a single record from content_deal_proposal_parameters_logs table in the estuary database
// @Summary Update an record in table content_deal_proposal_parameters_logs
// @Description Update a single record from content_deal_proposal_parameters_logs table in the estuary database
// @Tags ContentDealProposalParametersLogs
// @Accept  json
// @Produce  json
// @Param  contentDealProposalParametersLogsID path int64 true "id"
// @Param  ContentDealProposalParametersLogs body model.ContentDealProposalParametersLogs true "Update ContentDealProposalParametersLogs record"
// @Success 200 {object} model.ContentDealProposalParametersLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdealproposalparameterslogs/{contentDealProposalParametersLogsID} [put]
// echo '{"id": 11,"content": 28,"label": "gbgXFsMvnaClNRAqhtQdtFhNQ","duration": 43,"startEpoch": 93,"endEpoch": 15,"transferParams": "hDsSoFfCmVRZJFXIPJYBAxGLK","removeUnsealedCopy": false,"skipIpNiAnnounce": true,"nodeInfo": "idaYRkrLBhLFHYQAZbeFIkcXi","requesterInfo": "IwXjAiusvSYuTSiLFiYOJTTCO","requestingApiKey": "TfZrGwdaEivZEPgRZouqkMoVs","systemContentDealProposalParametersId": 54,"createdAt": "2051-07-18T14:58:51.12809116-04:00","updatedAt": "2196-07-15T21:02:52.682828298-04:00","deltaNodeUuid": "vXwPAuNtmYAUodvoDsZMFPAjG"}' | http PUT "http://localhost:8080/contentdealproposalparameterslogs/1"  X-Api-User:user123
func UpdateContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contentDealProposalParametersLogsID, err := parseInt64(ps, "contentDealProposalParametersLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdealproposalparameterslogs := &model.ContentDealProposalParametersLogs{}
	if err := readJSON(r, contentdealproposalparameterslogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdealproposalparameterslogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdealproposalparameterslogs.Prepare()

	if err := contentdealproposalparameterslogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_parameters_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdealproposalparameterslogs, _, err = dao.UpdateContentDealProposalParametersLogs(ctx,
		contentDealProposalParametersLogsID,
		contentdealproposalparameterslogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdealproposalparameterslogs)
}

// DeleteContentDealProposalParametersLogs Delete a single record from content_deal_proposal_parameters_logs table in the estuary database
// @Summary Delete a record from content_deal_proposal_parameters_logs
// @Description Delete a single record from content_deal_proposal_parameters_logs table in the estuary database
// @Tags ContentDealProposalParametersLogs
// @Accept  json
// @Produce  json
// @Param  contentDealProposalParametersLogsID path int64 true "id"
// @Success 204 {object} model.ContentDealProposalParametersLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentdealproposalparameterslogs/{contentDealProposalParametersLogsID} [delete]
// http DELETE "http://localhost:8080/contentdealproposalparameterslogs/1" X-Api-User:user123
func DeleteContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contentDealProposalParametersLogsID, err := parseInt64(ps, "contentDealProposalParametersLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_parameters_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentDealProposalParametersLogs(ctx, contentDealProposalParametersLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
