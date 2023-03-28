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

func configInstanceMetaLogsRouter(router *httprouter.Router) {
	router.GET("/instancemetalogs", GetAllInstanceMetaLogs)
	router.POST("/instancemetalogs", AddInstanceMetaLogs)
	router.GET("/instancemetalogs/:argID", GetInstanceMetaLogs)
	router.PUT("/instancemetalogs/:argID", UpdateInstanceMetaLogs)
	router.DELETE("/instancemetalogs/:argID", DeleteInstanceMetaLogs)
}

func configGinInstanceMetaLogsRouter(router gin.IRoutes) {
	router.GET("/instancemetalogs", ConverHttprouterToGin(GetAllInstanceMetaLogs))
	router.POST("/instancemetalogs", ConverHttprouterToGin(AddInstanceMetaLogs))
	router.GET("/instancemetalogs/:argID", ConverHttprouterToGin(GetInstanceMetaLogs))
	router.PUT("/instancemetalogs/:argID", ConverHttprouterToGin(UpdateInstanceMetaLogs))
	router.DELETE("/instancemetalogs/:argID", ConverHttprouterToGin(DeleteInstanceMetaLogs))
}

// GetAllInstanceMetaLogs is a function to get a slice of record(s) from instance_meta_logs table in the estuary database
// @Summary Get list of InstanceMetaLogs
// @Tags InstanceMetaLogs
// @Description GetAllInstanceMetaLogs is a handler to get a slice of record(s) from instance_meta_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.InstanceMetaLogs}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /instancemetalogs [get]
// http "http://localhost:8080/instancemetalogs?page=0&pagesize=20" X-Api-User:user123
func GetAllInstanceMetaLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "instance_meta_logs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllInstanceMetaLogs(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetInstanceMetaLogs is a function to get a single record from the instance_meta_logs table in the estuary database
// @Summary Get record from table InstanceMetaLogs by  argID
// @Tags InstanceMetaLogs
// @ID argID
// @Description GetInstanceMetaLogs is a function to get a single record from the instance_meta_logs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.InstanceMetaLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /instancemetalogs/{argID} [get]
// http "http://localhost:8080/instancemetalogs/1" X-Api-User:user123
func GetInstanceMetaLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "instance_meta_logs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetInstanceMetaLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddInstanceMetaLogs add to add a single record to instance_meta_logs table in the estuary database
// @Summary Add an record to instance_meta_logs table
// @Description add to add a single record to instance_meta_logs table in the estuary database
// @Tags InstanceMetaLogs
// @Accept  json
// @Produce  json
// @Param InstanceMetaLogs body model.InstanceMetaLogs true "Add InstanceMetaLogs"
// @Success 200 {object} model.InstanceMetaLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /instancemetalogs [post]
// echo '{"id": 13,"instance_uuid": "QZqJQFhMxvKdcoPYRhoYIGKXY","instance_host_name": "NkObRXNTrSgVTmrnWfaqQtSMG","instance_node_name": "uUtfSFLYSNkXgRVjSHnQbEUPM","os_details": "IOJXCpfFNFyWWXWGtwIamAglk","public_ip": "GJlJfHlSXyKHEaurudKHiDnHf","memory_limit": 16,"cpu_limit": 6,"storage_limit": 58,"disable_request": true,"disable_commitment_piece_generation": true,"disable_storage_deal": false,"disable_online_deals": true,"disable_offline_deals": true,"number_of_cpus": 60,"storage_in_bytes": 55,"system_memory": 96,"heap_memory": 50,"heap_in_use": 76,"stack_in_use": 52,"instance_start": "2097-04-10T11:10:18.072518734-04:00","bytes_per_cpu": 62,"node_info": "wOyeKPZPqxkMJQmhilThcIMoZ","requester_info": "njINpEXWJnnqYmgtpxmHybHPq","delta_node_uuid": "yhiYTereSvRGaSamamuAPWUPV","system_instance_meta_id": 5,"created_at": "2147-06-07T15:53:28.278550966-04:00","updated_at": "2070-01-08T01:13:47.373733577-05:00"}' | http POST "http://localhost:8080/instancemetalogs" X-Api-User:user123
func AddInstanceMetaLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	instancemetalogs := &model.InstanceMetaLogs{}

	if err := readJSON(r, instancemetalogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := instancemetalogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	instancemetalogs.Prepare()

	if err := instancemetalogs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "instance_meta_logs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	instancemetalogs, _, err = dao.AddInstanceMetaLogs(ctx, instancemetalogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, instancemetalogs)
}

// UpdateInstanceMetaLogs Update a single record from instance_meta_logs table in the estuary database
// @Summary Update an record in table instance_meta_logs
// @Description Update a single record from instance_meta_logs table in the estuary database
// @Tags InstanceMetaLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  InstanceMetaLogs body model.InstanceMetaLogs true "Update InstanceMetaLogs record"
// @Success 200 {object} model.InstanceMetaLogs
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /instancemetalogs/{argID} [put]
// echo '{"id": 13,"instance_uuid": "QZqJQFhMxvKdcoPYRhoYIGKXY","instance_host_name": "NkObRXNTrSgVTmrnWfaqQtSMG","instance_node_name": "uUtfSFLYSNkXgRVjSHnQbEUPM","os_details": "IOJXCpfFNFyWWXWGtwIamAglk","public_ip": "GJlJfHlSXyKHEaurudKHiDnHf","memory_limit": 16,"cpu_limit": 6,"storage_limit": 58,"disable_request": true,"disable_commitment_piece_generation": true,"disable_storage_deal": false,"disable_online_deals": true,"disable_offline_deals": true,"number_of_cpus": 60,"storage_in_bytes": 55,"system_memory": 96,"heap_memory": 50,"heap_in_use": 76,"stack_in_use": 52,"instance_start": "2097-04-10T11:10:18.072518734-04:00","bytes_per_cpu": 62,"node_info": "wOyeKPZPqxkMJQmhilThcIMoZ","requester_info": "njINpEXWJnnqYmgtpxmHybHPq","delta_node_uuid": "yhiYTereSvRGaSamamuAPWUPV","system_instance_meta_id": 5,"created_at": "2147-06-07T15:53:28.278550966-04:00","updated_at": "2070-01-08T01:13:47.373733577-05:00"}' | http PUT "http://localhost:8080/instancemetalogs/1"  X-Api-User:user123
func UpdateInstanceMetaLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	instancemetalogs := &model.InstanceMetaLogs{}
	if err := readJSON(r, instancemetalogs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := instancemetalogs.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	instancemetalogs.Prepare()

	if err := instancemetalogs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "instance_meta_logs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	instancemetalogs, _, err = dao.UpdateInstanceMetaLogs(ctx,
		argID,
		instancemetalogs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, instancemetalogs)
}

// DeleteInstanceMetaLogs Delete a single record from instance_meta_logs table in the estuary database
// @Summary Delete a record from instance_meta_logs
// @Description Delete a single record from instance_meta_logs table in the estuary database
// @Tags InstanceMetaLogs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.InstanceMetaLogs
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /instancemetalogs/{argID} [delete]
// http DELETE "http://localhost:8080/instancemetalogs/1" X-Api-User:user123
func DeleteInstanceMetaLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "instance_meta_logs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteInstanceMetaLogs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
