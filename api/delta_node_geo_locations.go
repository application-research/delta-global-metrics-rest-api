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

func configDeltaNodeGeoLocationsRouter(router *httprouter.Router) {
	router.GET("/deltanodegeolocations", GetAllDeltaNodeGeoLocations)
	router.GET("/deltanodegeolocations/:deltaNodeGeoLocationsID", GetDeltaNodeGeoLocations)
}

func configGinDeltaNodeGeoLocationsRouter(router gin.IRoutes) {
	router.GET("/deltanodegeolocations", ConverHttprouterToGin(GetAllDeltaNodeGeoLocations))
	router.GET("/deltanodegeolocations/:deltaNodeGeoLocationsID", ConverHttprouterToGin(GetDeltaNodeGeoLocations))
}

// GetAllDeltaNodeGeoLocations is a function to get a slice of record(s) from delta_node_geo_locations table in the estuary database
// @Summary Get list of DeltaNodeGeoLocations
// @Tags DeltaNodeGeoLocations
// @Description GetAllDeltaNodeGeoLocations is a handler to get a slice of record(s) from delta_node_geo_locations table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.DeltaNodeGeoLocations}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /deltanodegeolocations [get]
// http "http://localhost:8080/deltanodegeolocations?page=0&pagesize=20" X-Api-User:user123
func GetAllDeltaNodeGeoLocations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "delta_node_geo_locations", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDeltaNodeGeoLocations(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetDeltaNodeGeoLocations is a function to get a single record from the delta_node_geo_locations table in the estuary database
// @Summary Get record from table DeltaNodeGeoLocations by  deltaNodeGeoLocationsID
// @Tags DeltaNodeGeoLocations
// @ID deltaNodeGeoLocationsID
// @Description GetDeltaNodeGeoLocations is a function to get a single record from the delta_node_geo_locations table in the estuary database
// @Accept  json
// @Produce  json
// @Param  deltaNodeGeoLocationsID path int64 true "id"
// @Success 200 {object} model.DeltaNodeGeoLocations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /deltanodegeolocations/{deltaNodeGeoLocationsID} [get]
// http "http://localhost:8080/deltanodegeolocations/1" X-Api-User:user123
func GetDeltaNodeGeoLocations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	deltaNodeGeoLocationsID, err := parseInt64(ps, "deltaNodeGeoLocationsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_node_geo_locations", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDeltaNodeGeoLocations(ctx, deltaNodeGeoLocationsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddDeltaNodeGeoLocations add to add a single record to delta_node_geo_locations table in the estuary database
// @Summary Add an record to delta_node_geo_locations table
// @Description add to add a single record to delta_node_geo_locations table in the estuary database
// @Tags DeltaNodeGeoLocations
// @Accept  json
// @Produce  json
// @Param DeltaNodeGeoLocations body model.DeltaNodeGeoLocations true "Add DeltaNodeGeoLocations"
// @Success 200 {object} model.DeltaNodeGeoLocations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /deltanodegeolocations [post]
// echo '{"id": 89,"ip": "uwpBtXBntBWgkCKraXSpkctDe","country": "GHPtnoTwpFebPPVvcspCrMqZE","city": "kNBasfOCYAWmojkCMuiBaimiw","region": "vbJjsoGjvIDUSYIYsbvtGKygD","zip": "VRnvGNQBsiJRbbUgODZEJDkcB","lat": 0.2523771892476505,"lon": 0.8991950175159545,"createdAt": "2205-06-16T18:43:47.606371433-04:00","updatedAt": "2166-09-03T02:07:18.062428741-04:00"}' | http POST "http://localhost:8080/deltanodegeolocations" X-Api-User:user123
func AddDeltaNodeGeoLocations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	deltanodegeolocations := &model.DeltaNodeGeoLocations{}

	if err := readJSON(r, deltanodegeolocations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := deltanodegeolocations.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	deltanodegeolocations.Prepare()

	if err := deltanodegeolocations.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_node_geo_locations", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	deltanodegeolocations, _, err = dao.AddDeltaNodeGeoLocations(ctx, deltanodegeolocations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, deltanodegeolocations)
}

// UpdateDeltaNodeGeoLocations Update a single record from delta_node_geo_locations table in the estuary database
// @Summary Update an record in table delta_node_geo_locations
// @Description Update a single record from delta_node_geo_locations table in the estuary database
// @Tags DeltaNodeGeoLocations
// @Accept  json
// @Produce  json
// @Param  deltaNodeGeoLocationsID path int64 true "id"
// @Param  DeltaNodeGeoLocations body model.DeltaNodeGeoLocations true "Update DeltaNodeGeoLocations record"
// @Success 200 {object} model.DeltaNodeGeoLocations
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /deltanodegeolocations/{deltaNodeGeoLocationsID} [put]
// echo '{"id": 89,"ip": "uwpBtXBntBWgkCKraXSpkctDe","country": "GHPtnoTwpFebPPVvcspCrMqZE","city": "kNBasfOCYAWmojkCMuiBaimiw","region": "vbJjsoGjvIDUSYIYsbvtGKygD","zip": "VRnvGNQBsiJRbbUgODZEJDkcB","lat": 0.2523771892476505,"lon": 0.8991950175159545,"createdAt": "2205-06-16T18:43:47.606371433-04:00","updatedAt": "2166-09-03T02:07:18.062428741-04:00"}' | http PUT "http://localhost:8080/deltanodegeolocations/1"  X-Api-User:user123
func UpdateDeltaNodeGeoLocations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	deltaNodeGeoLocationsID, err := parseInt64(ps, "deltaNodeGeoLocationsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	deltanodegeolocations := &model.DeltaNodeGeoLocations{}
	if err := readJSON(r, deltanodegeolocations); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := deltanodegeolocations.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	deltanodegeolocations.Prepare()

	if err := deltanodegeolocations.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_node_geo_locations", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	deltanodegeolocations, _, err = dao.UpdateDeltaNodeGeoLocations(ctx,
		deltaNodeGeoLocationsID,
		deltanodegeolocations)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, deltanodegeolocations)
}

// DeleteDeltaNodeGeoLocations Delete a single record from delta_node_geo_locations table in the estuary database
// @Summary Delete a record from delta_node_geo_locations
// @Description Delete a single record from delta_node_geo_locations table in the estuary database
// @Tags DeltaNodeGeoLocations
// @Accept  json
// @Produce  json
// @Param  deltaNodeGeoLocationsID path int64 true "id"
// @Success 204 {object} model.DeltaNodeGeoLocations
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /deltanodegeolocations/{deltaNodeGeoLocationsID} [delete]
// http DELETE "http://localhost:8080/deltanodegeolocations/1" X-Api-User:user123
func DeleteDeltaNodeGeoLocations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	deltaNodeGeoLocationsID, err := parseInt64(ps, "deltaNodeGeoLocationsID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "delta_node_geo_locations", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDeltaNodeGeoLocations(ctx, deltaNodeGeoLocationsID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
