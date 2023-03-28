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

func configPieceCommitmentLogsRouter(router *httprouter.Router) {
	router.GET("/piececommitmentlogs", GetAllPieceCommitmentLogs)
	router.GET("/piececommitmentlogs/:pieceCommitmentLogsID", GetPieceCommitmentLogs)
}

func configGinPieceCommitmentLogsRouter(router gin.IRoutes) {
	router.GET("/piececommitmentlogs", ConverHttprouterToGin(GetAllPieceCommitmentLogs))
	router.GET("/piececommitmentlogs/:pieceCommitmentLogsID", ConverHttprouterToGin(GetPieceCommitmentLogs))
}

// GetAllPieceCommitmentLogs is a function to get a slice of record(s) from piece_commitment_logs table in the estuary database
// @Summary Get list of PieceCommitmentLogs
// @Tags PieceCommitmentLogs
// @Description GetAllPieceCommitmentLogs is a handler to get a slice of record(s) from piece_commitment_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PieceCommitmentLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /piececommitmentlogs [get]
// http "http://localhost:8080/piececommitmentlogs?page=0&pagesize=20" X-Api-User:user123
func GetAllPieceCommitmentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "piece_commitment_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPieceCommitmentLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPieceCommitmentLogs is a function to get a single record from the piece_commitment_logs table in the estuary database
// @Summary Get record from table PieceCommitmentLogs by  pieceCommitmentLogsID
// @Tags PieceCommitmentLogs
// @ID pieceCommitmentLogsID
// @Description GetPieceCommitmentLogs is a function to get a single record from the piece_commitment_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  pieceCommitmentLogsID path int64 true "id"
// @Success 200 {object} model.PieceCommitmentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /piececommitmentlogs/{pieceCommitmentLogsID} [get]
// http "http://localhost:8080/piececommitmentlogs/1" X-Api-User:user123
func GetPieceCommitmentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	pieceCommitmentLogsID, err := parseInt64(ps, "pieceCommitmentLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_commitment_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPieceCommitmentLogs(ctx, pieceCommitmentLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPieceCommitmentLogs add to add a single record to piece_commitment_logs table in the estuary database
// @Summary Add an record to piece_commitment_logs table
// @Description add to add a single record to piece_commitment_logs table in the estuary database
// @Tags PieceCommitmentLogs
// @Accept  json
// @Produce  json
// @Param PieceCommitmentLogs body model.PieceCommitmentLogs true "Add PieceCommitmentLogs"
// @Success 200 {object} model.PieceCommitmentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /piececommitmentlogs [post]
// echo '{"id": 96,"cid": "lSpleRdPmtqSgPIPaupwBTuyl","piece": "rdrggiHeNdpKHBfFylyCggASX","size": 54,"paddedPieceSize": 34,"unPaddedPieceSize": 67,"status": "rZcNQmBiCIqntDBaEJOEBDeUU","lastMessage": "nMdhoVjfwPSQpWiVwprTHtQDv","nodeInfo": "hgMKEJrQUrKFCsxmSNhmKrfns","requesterInfo": "gjlGNNmhGjgIZlGkkRLInnaIl","requestingApiKey": "OdFJvYFetTPCvOnNtqEEefDtu","systemContentPieceCommitmentId": 41,"createdAt": "2263-03-20T02:46:34.0114278-04:00","updatedAt": "2233-07-09T11:09:29.842592668-04:00","deltaNodeUuid": "BHSEdeJWRcPaiTAwASdFCrYor"}' | http POST "http://localhost:8080/piececommitmentlogs" X-Api-User:user123
func AddPieceCommitmentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	piececommitmentlogs := &model.PieceCommitmentLogs{}

	if err := readJSON(r, piececommitmentlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := piececommitmentlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	piececommitmentlogs.Prepare()

	if err := piececommitmentlogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_commitment_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	piececommitmentlogs, _, err = dao.AddPieceCommitmentLogs(ctx, piececommitmentlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, piececommitmentlogs)
}

// UpdatePieceCommitmentLogs Update a single record from piece_commitment_logs table in the estuary database
// @Summary Update an record in table piece_commitment_logs
// @Description Update a single record from piece_commitment_logs table in the estuary database
// @Tags PieceCommitmentLogs
// @Accept  json
// @Produce  json
// @Param  pieceCommitmentLogsID path int64 true "id"
// @Param  PieceCommitmentLogs body model.PieceCommitmentLogs true "Update PieceCommitmentLogs record"
// @Success 200 {object} model.PieceCommitmentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /piececommitmentlogs/{pieceCommitmentLogsID} [put]
// echo '{"id": 96,"cid": "lSpleRdPmtqSgPIPaupwBTuyl","piece": "rdrggiHeNdpKHBfFylyCggASX","size": 54,"paddedPieceSize": 34,"unPaddedPieceSize": 67,"status": "rZcNQmBiCIqntDBaEJOEBDeUU","lastMessage": "nMdhoVjfwPSQpWiVwprTHtQDv","nodeInfo": "hgMKEJrQUrKFCsxmSNhmKrfns","requesterInfo": "gjlGNNmhGjgIZlGkkRLInnaIl","requestingApiKey": "OdFJvYFetTPCvOnNtqEEefDtu","systemContentPieceCommitmentId": 41,"createdAt": "2263-03-20T02:46:34.0114278-04:00","updatedAt": "2233-07-09T11:09:29.842592668-04:00","deltaNodeUuid": "BHSEdeJWRcPaiTAwASdFCrYor"}' | http PUT "http://localhost:8080/piececommitmentlogs/1"  X-Api-User:user123
func UpdatePieceCommitmentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	pieceCommitmentLogsID, err := parseInt64(ps, "pieceCommitmentLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	piececommitmentlogs := &model.PieceCommitmentLogs{}
	if err := readJSON(r, piececommitmentlogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := piececommitmentlogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	piececommitmentlogs.Prepare()

	if err := piececommitmentlogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_commitment_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	piececommitmentlogs, _, err = dao.UpdatePieceCommitmentLogs(ctx,
		pieceCommitmentLogsID,
		piececommitmentlogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, piececommitmentlogs)
}

// DeletePieceCommitmentLogs Delete a single record from piece_commitment_logs table in the estuary database
// @Summary Delete a record from piece_commitment_logs
// @Description Delete a single record from piece_commitment_logs table in the estuary database
// @Tags PieceCommitmentLogs
// @Accept  json
// @Produce  json
// @Param  pieceCommitmentLogsID path int64 true "id"
// @Success 204 {object} model.PieceCommitmentLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /piececommitmentlogs/{pieceCommitmentLogsID} [delete]
// http DELETE "http://localhost:8080/piececommitmentlogs/1" X-Api-User:user123
func DeletePieceCommitmentLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	pieceCommitmentLogsID, err := parseInt64(ps, "pieceCommitmentLogsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_commitment_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePieceCommitmentLogs(ctx, pieceCommitmentLogsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
