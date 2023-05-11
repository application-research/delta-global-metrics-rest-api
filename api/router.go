package api

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/application-research/delta-metrics-rest/dao"
	"github.com/application-research/delta-metrics-rest/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

var (
	_             = time.Second // import time.Second for unknown usage in api
	crudEndpoints map[string]*CrudAPI
)

// CrudAPI describes requests available for tables in the database
type CrudAPI struct {
	Name            string           `json:"name"`
	CreateURL       string           `json:"createUrl"`
	RetrieveOneURL  string           `json:"retrieveOneUrl"`
	RetrieveManyURL string           `json:"retrieveManyUrl"`
	UpdateURL       string           `json:"updateUrl"`
	DeleteURL       string           `json:"deleteUrl"`
	FetchDDLURL     string           `json:"fetchDdlUrl"`
	TableInfo       *model.TableInfo `json:"tableInfo"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"pageSize"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"totalRecords"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ConfigGinRouter configure gin router
func ConfigGinRouter(router gin.IRoutes) {
	configGinStatisticsRouter(router)
	configGinStatisticsTimeSeriesRouter(router)
	router.GET("/ddl/:argID", ConverHttprouterToGin(GetDdl))
	router.GET("/ddl", ConverHttprouterToGin(GetDdlEndpoints))
	return
}

// ConverHttprouterToGin wrap httprouter.Handle to gin.HandlerFunc
func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func initializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var RequestValidator RequestValidatorFunc

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var ContextInitializer ContextInitializerFunc

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func writeRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func returnError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJSON(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func parseUint8(ps httprouter.Params, key string) (uint8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}
func parseUint16(ps httprouter.Params, key string) (uint16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}
func parseUint32(ps httprouter.Params, key string) (uint32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}
func parseUint64(ps httprouter.Params, key string) (uint64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}
func parseInt(ps httprouter.Params, key string) (int, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}
func parseInt8(ps httprouter.Params, key string) (int8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}
func parseInt16(ps httprouter.Params, key string) (int16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}
func parseInt32(ps httprouter.Params, key string) (int32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}
func parseInt64(ps httprouter.Params, key string) (int64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err
}
func parseString(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
func parseUUID(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}

// GetDdl is a function to get table info for a table in the estuary database
// @Summary Get table info for a table in the estuary database by argID
// @Tags TableInfo
// @ID argID
// @Description GetDdl is a function to get table info for a table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} api.CrudAPI
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /ddl/{argID} [get]
// http "http://localhost:8080/ddl/xyz" X-Api-User:user123
func GetDdl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID := ps.ByName("argID")

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, ok := crudEndpoints[argID]
	if !ok {
		returnError(ctx, w, r, fmt.Errorf("unable to find table: %s", argID))
		return
	}

	writeJSON(ctx, w, record)
}

// GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the estuary database
// @Summary Gets a list of ddl endpoints available for tables in the estuary database
// @Tags TableInfo
// @Description GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the estuary database
// @Accept  json
// @Produce  json
// @Success 200 {object} api.CrudAPI
// @Router /ddl [get]
// http "http://localhost:8080/ddl" X-Api-User:user123
func GetDdlEndpoints(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, crudEndpoints)
}

func init() {
	crudEndpoints = make(map[string]*CrudAPI)

	var tmp *CrudAPI

	tmp = &CrudAPI{
		Name:            "content_deal_logs",
		CreateURL:       "/contentdeallogs",
		RetrieveOneURL:  "/contentdeallogs",
		RetrieveManyURL: "/contentdeallogs",
		UpdateURL:       "/contentdeallogs",
		DeleteURL:       "/contentdeallogs",
		FetchDDLURL:     "/ddl/content_deal_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("content_deal_logs")
	crudEndpoints["content_deal_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "content_deal_proposal_logs",
		CreateURL:       "/contentdealproposallogs",
		RetrieveOneURL:  "/contentdealproposallogs",
		RetrieveManyURL: "/contentdealproposallogs",
		UpdateURL:       "/contentdealproposallogs",
		DeleteURL:       "/contentdealproposallogs",
		FetchDDLURL:     "/ddl/content_deal_proposal_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("content_deal_proposal_logs")
	crudEndpoints["content_deal_proposal_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "content_deal_proposal_parameters_logs",
		CreateURL:       "/contentdealproposalparameterslogs",
		RetrieveOneURL:  "/contentdealproposalparameterslogs",
		RetrieveManyURL: "/contentdealproposalparameterslogs",
		UpdateURL:       "/contentdealproposalparameterslogs",
		DeleteURL:       "/contentdealproposalparameterslogs",
		FetchDDLURL:     "/ddl/content_deal_proposal_parameters_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("content_deal_proposal_parameters_logs")
	crudEndpoints["content_deal_proposal_parameters_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "content_logs",
		CreateURL:       "/contentlogs",
		RetrieveOneURL:  "/contentlogs",
		RetrieveManyURL: "/contentlogs",
		UpdateURL:       "/contentlogs",
		DeleteURL:       "/contentlogs",
		FetchDDLURL:     "/ddl/content_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("content_logs")
	crudEndpoints["content_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "content_miner_logs",
		CreateURL:       "/contentminerlogs",
		RetrieveOneURL:  "/contentminerlogs",
		RetrieveManyURL: "/contentminerlogs",
		UpdateURL:       "/contentminerlogs",
		DeleteURL:       "/contentminerlogs",
		FetchDDLURL:     "/ddl/content_miner_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("content_miner_logs")
	crudEndpoints["content_miner_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "content_wallet_logs",
		CreateURL:       "/contentwalletlogs",
		RetrieveOneURL:  "/contentwalletlogs",
		RetrieveManyURL: "/contentwalletlogs",
		UpdateURL:       "/contentwalletlogs",
		DeleteURL:       "/contentwalletlogs",
		FetchDDLURL:     "/ddl/content_wallet_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("content_wallet_logs")
	crudEndpoints["content_wallet_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "delta_node_geo_locations",
		CreateURL:       "/deltanodegeolocations",
		RetrieveOneURL:  "/deltanodegeolocations",
		RetrieveManyURL: "/deltanodegeolocations",
		UpdateURL:       "/deltanodegeolocations",
		DeleteURL:       "/deltanodegeolocations",
		FetchDDLURL:     "/ddl/delta_node_geo_locations",
	}

	tmp.TableInfo, _ = model.GetTableInfo("delta_node_geo_locations")
	crudEndpoints["delta_node_geo_locations"] = tmp

	tmp = &CrudAPI{
		Name:            "delta_startup_logs",
		CreateURL:       "/deltastartuplogs",
		RetrieveOneURL:  "/deltastartuplogs",
		RetrieveManyURL: "/deltastartuplogs",
		UpdateURL:       "/deltastartuplogs",
		DeleteURL:       "/deltastartuplogs",
		FetchDDLURL:     "/ddl/delta_startup_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("delta_startup_logs")
	crudEndpoints["delta_startup_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "instance_meta_logs",
		CreateURL:       "/instancemetalogs",
		RetrieveOneURL:  "/instancemetalogs",
		RetrieveManyURL: "/instancemetalogs",
		UpdateURL:       "/instancemetalogs",
		DeleteURL:       "/instancemetalogs",
		FetchDDLURL:     "/ddl/instance_meta_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("instance_meta_logs")
	crudEndpoints["instance_meta_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "log_events",
		CreateURL:       "/logevents",
		RetrieveOneURL:  "/logevents",
		RetrieveManyURL: "/logevents",
		UpdateURL:       "/logevents",
		DeleteURL:       "/logevents",
		FetchDDLURL:     "/ddl/log_events",
	}

	tmp.TableInfo, _ = model.GetTableInfo("log_events")
	crudEndpoints["log_events"] = tmp

	tmp = &CrudAPI{
		Name:            "piece_commitment_logs",
		CreateURL:       "/piececommitmentlogs",
		RetrieveOneURL:  "/piececommitmentlogs",
		RetrieveManyURL: "/piececommitmentlogs",
		UpdateURL:       "/piececommitmentlogs",
		DeleteURL:       "/piececommitmentlogs",
		FetchDDLURL:     "/ddl/piece_commitment_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("piece_commitment_logs")
	crudEndpoints["piece_commitment_logs"] = tmp

	tmp = &CrudAPI{
		Name:            "wallet_logs",
		CreateURL:       "/walletlogs",
		RetrieveOneURL:  "/walletlogs",
		RetrieveManyURL: "/walletlogs",
		UpdateURL:       "/walletlogs",
		DeleteURL:       "/walletlogs",
		FetchDDLURL:     "/ddl/wallet_logs",
	}

	tmp.TableInfo, _ = model.GetTableInfo("wallet_logs")
	crudEndpoints["wallet_logs"] = tmp

}
