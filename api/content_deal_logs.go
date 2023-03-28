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
	router.GET("/contentdeallogs/:contentDealLogsID", GetContentDealLogs)
}

func configGinContentDealLogsRouter(router gin.IRoutes) {
	router.GET("/contentdeallogs", ConverHttprouterToGin(GetAllContentDealLogs))
	router.GET("/contentdeallogs/:contentDealLogsID", ConverHttprouterToGin(GetContentDealLogs))
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
// @Summary Get record from table ContentDealLogs by  contentDealLogsID
// @Tags ContentDealLogs
// @ID contentDealLogsID
// @Description GetContentDealLogs is a function to get a single record from the content_deal_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  contentDealLogsID path int64 true "id"
// @Success 200 {object} model.ContentDealLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentdeallogs/{contentDealLogsID} [get]
// http "http://localhost:8080/contentdeallogs/1" X-Api-User:user123
func GetContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contentDealLogsID, err := parseInt64(ps, "contentDealLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentDealLogs(ctx, contentDealLogsID)
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
// echo '{"id": 99,"content": 55,"propCid": "xknsQTytaZWgYojHSDkHVfPHa","dealUuid": "GxKYowuykjCRKgEIOgeQqqsen","miner": "OInOCDbSbQxplmMthtVnPierV","dealId": 85,"failed": true,"verified": false,"slashed": false,"failedAt": "2104-07-28T14:23:06.964124494-04:00","dtChan": "AcIPDCIUYkjdwjTvidljlPmEx","transferStarted": "2255-11-25T21:46:23.259616019-05:00","transferFinished": "2185-12-19T08:50:07.455413034-05:00","onChainAt": "2029-02-05T01:46:10.623356801-05:00","sealedAt": "2028-02-10T19:11:30.660121509-05:00","lastMessage": "vNFrivKDIRlSgyxmVrQZtcNPp","dealProtocolVersion": "SrvIxXavVsTojSXZWVUyHfVsG","minerVersion": "dZasDhnmXPWxsTJTkkjqsXekw","nodeInfo": "eaPQgRleeUStfJjFeLaqbMGAl","requesterInfo": "IPnKALNplIsAFbfNNEJjxEEbV","requestingApiKey": "GfxZduGGhytByYTANrfWicIOx","systemContentDealId": 68,"createdAt": "2236-06-18T16:44:22.565027633-04:00","updatedAt": "2108-07-24T20:06:59.490583119-04:00","deltaNodeUuid": "DBoqKOqeVRvCTTBXExXyHoKfB"}' | http POST "http://localhost:8080/contentdeallogs" X-Api-User:user123
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
// @Param  contentDealLogsID path int64 true "id"
// @Param  ContentDealLogs body model.ContentDealLogs true "Update ContentDealLogs record"
// @Success 200 {object} model.ContentDealLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdeallogs/{contentDealLogsID} [put]
// echo '{"id": 99,"content": 55,"propCid": "xknsQTytaZWgYojHSDkHVfPHa","dealUuid": "GxKYowuykjCRKgEIOgeQqqsen","miner": "OInOCDbSbQxplmMthtVnPierV","dealId": 85,"failed": true,"verified": false,"slashed": false,"failedAt": "2104-07-28T14:23:06.964124494-04:00","dtChan": "AcIPDCIUYkjdwjTvidljlPmEx","transferStarted": "2255-11-25T21:46:23.259616019-05:00","transferFinished": "2185-12-19T08:50:07.455413034-05:00","onChainAt": "2029-02-05T01:46:10.623356801-05:00","sealedAt": "2028-02-10T19:11:30.660121509-05:00","lastMessage": "vNFrivKDIRlSgyxmVrQZtcNPp","dealProtocolVersion": "SrvIxXavVsTojSXZWVUyHfVsG","minerVersion": "dZasDhnmXPWxsTJTkkjqsXekw","nodeInfo": "eaPQgRleeUStfJjFeLaqbMGAl","requesterInfo": "IPnKALNplIsAFbfNNEJjxEEbV","requestingApiKey": "GfxZduGGhytByYTANrfWicIOx","systemContentDealId": 68,"createdAt": "2236-06-18T16:44:22.565027633-04:00","updatedAt": "2108-07-24T20:06:59.490583119-04:00","deltaNodeUuid": "DBoqKOqeVRvCTTBXExXyHoKfB"}' | http PUT "http://localhost:8080/contentdeallogs/1"  X-Api-User:user123
func UpdateContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contentDealLogsID, err := parseInt64(ps, "contentDealLogsID")
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
		contentDealLogsID,
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
// @Param  contentDealLogsID path int64 true "id"
// @Success 204 {object} model.ContentDealLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentdeallogs/{contentDealLogsID} [delete]
// http DELETE "http://localhost:8080/contentdeallogs/1" X-Api-User:user123
func DeleteContentDealLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	contentDealLogsID, err := parseInt64(ps, "contentDealLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deal_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentDealLogs(ctx, contentDealLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
