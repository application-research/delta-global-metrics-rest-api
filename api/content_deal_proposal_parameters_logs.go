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
	router.POST("/contentdealproposalparameterslogs", AddContentDealProposalParametersLogs)
	router.GET("/contentdealproposalparameterslogs/:argID", GetContentDealProposalParametersLogs)
	router.PUT("/contentdealproposalparameterslogs/:argID", UpdateContentDealProposalParametersLogs)
	router.DELETE("/contentdealproposalparameterslogs/:argID", DeleteContentDealProposalParametersLogs)
}

func configGinContentDealProposalParametersLogsRouter(router gin.IRoutes) {
	router.GET("/contentdealproposalparameterslogs", ConverHttprouterToGin(GetAllContentDealProposalParametersLogs))
	router.POST("/contentdealproposalparameterslogs", ConverHttprouterToGin(AddContentDealProposalParametersLogs))
	router.GET("/contentdealproposalparameterslogs/:argID", ConverHttprouterToGin(GetContentDealProposalParametersLogs))
	router.PUT("/contentdealproposalparameterslogs/:argID", ConverHttprouterToGin(UpdateContentDealProposalParametersLogs))
	router.DELETE("/contentdealproposalparameterslogs/:argID", ConverHttprouterToGin(DeleteContentDealProposalParametersLogs))
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
// @Summary Get record from table ContentDealProposalParametersLogs by  argID
// @Tags ContentDealProposalParametersLogs
// @ID argID
// @Description GetContentDealProposalParametersLogs is a function to get a single record from the content_deal_proposal_parameters_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ContentDealProposalParametersLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentdealproposalparameterslogs/{argID} [get]
// http "http://localhost:8080/contentdealproposalparameterslogs/1" X-Api-User:user123
func GetContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_parameters_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentDealProposalParametersLogs(ctx, argID)
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
// echo '{"id": 79,"content": 27,"label": "TGKBXMsLffqmOVsZYoeTNDive","duration": 73,"start_epoch": 47,"end_epoch": 86,"transfer_params": "WpMPvCtTZoJomXavYlqZbBNxF","remove_unsealed_copy": false,"skip_ip_ni_announce": true,"node_info": "dNcquanpkcUJpmUEERUAatKXA","requester_info": "XMHaKbKpreLvYjeJvnyEJuFOj","requesting_api_key": "VSygsWKMxqBbsmqCXmGPwoldX","system_content_deal_proposal_parameters_id": 4,"created_at": "2200-11-27T04:06:47.214934775-05:00","updated_at": "2284-03-31T09:39:16.663154075-04:00","delta_node_uuid": "ZWNplgCUcugrHxYFSChKHjowb"}' | http POST "http://localhost:8080/contentdealproposalparameterslogs" X-Api-User:user123
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
// @Param  argID path int64 true "id"
// @Param  ContentDealProposalParametersLogs body model.ContentDealProposalParametersLogs true "Update ContentDealProposalParametersLogs record"
// @Success 200 {object} model.ContentDealProposalParametersLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdealproposalparameterslogs/{argID} [put]
// echo '{"id": 79,"content": 27,"label": "TGKBXMsLffqmOVsZYoeTNDive","duration": 73,"start_epoch": 47,"end_epoch": 86,"transfer_params": "WpMPvCtTZoJomXavYlqZbBNxF","remove_unsealed_copy": false,"skip_ip_ni_announce": true,"node_info": "dNcquanpkcUJpmUEERUAatKXA","requester_info": "XMHaKbKpreLvYjeJvnyEJuFOj","requesting_api_key": "VSygsWKMxqBbsmqCXmGPwoldX","system_content_deal_proposal_parameters_id": 4,"created_at": "2200-11-27T04:06:47.214934775-05:00","updated_at": "2284-03-31T09:39:16.663154075-04:00","delta_node_uuid": "ZWNplgCUcugrHxYFSChKHjowb"}' | http PUT "http://localhost:8080/contentdealproposalparameterslogs/1"  X-Api-User:user123
func UpdateContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
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
		argID,
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
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ContentDealProposalParametersLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentdealproposalparameterslogs/{argID} [delete]
// http DELETE "http://localhost:8080/contentdealproposalparameterslogs/1" X-Api-User:user123
func DeleteContentDealProposalParametersLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_proposal_parameters_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentDealProposalParametersLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
