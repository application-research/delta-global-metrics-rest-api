package api

import (
	"github.com/application-research/delta-metrics-rest/dao"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// TotalDealsAttempted                       int `json:"total_deals_attempted,omitempty"`
// TotalE2EDealsAttempted                    int `json:"total_e2e_deals_attempted,omitempty"`
// TotalImportDealsAttempted                 int `json:"total_import_deals_attempted,omitempty"`
// TotalPieceCommitmentsComputeAttempted     int `json:"total_piece_commitments_compute_attempted,omitempty"`
// TotalDealsAttemptedSize                   int `json:"total_deals_attempted_size,omitempty"`
// TotalE2EDealsAttemptedSize                int `json:"total_e2e_deals_attempted_size,omitempty"`
// TotalImportDealsAttemptedSize             int `json:"total_import_deals_attempted_size,omitempty"`
// TotalPieceCommitmentsComputeAttemptedSize int `json:"total_piece_commitments_compute_attempted_size,omitempty"`
// TotalDealsSucceeded                       int `json:"total_deals_succeeded,omitempty"`
// TotalE2ESucceeded                         int `json:"total_e2e_succeeded,omitempty"`
// TotalImportSucceeded                      int `json:"total_import_succeeded,omitempty"`
// TotalPieceCommitmentsComputeSucceeded     int `json:"total_piece_commitments_compute_succeeded,omitempty"`
// TotalDealsSucceededSize                   int `json:"total_deals_succeeded_size,omitempty"`
// TotalE2ESucceededSize                     int `json:"total_e2e_succeeded_size,omitempty"`
// TotalImportSucceededSize                  int `json:"total_import_succeeded_size,omitempty"`
// TotalPieceCommitmentsComputeSucceededSize int `json:"total_piece_commitments_compute_succeeded_size,omitempty"`
// TotalInProgressDeals24h                   int `json:"total_in_progress_deals_24h,omitempty"`
// TotalInProgressE2EDeals24h                int `json:"total_in_progress_e2e_deals_24h,omitempty"`
// TotalInProgressImportDeals24h             int `json:"total_in_progress_import_deals_24h,omitempty"`
// TotalInProgressDealsSize24h               int `json:"total_in_progress_deals_size_24h,omitempty"`
// TotalInProgressE2EDealsSize24h            int `json:"total_in_progress_e2e_deals_size_24h,omitempty"`
// TotalInProgressImportDealsSize24h         int `json:"total_in_progress_import_deals_size_24h,omitempty"`
// TotalNumberOfSpsWorkWith                  int `json:"total_number_of_sps_worked_with,omitempty"`
// TotalNumberOfUniqueDeltaNodes             int `json:"total_number_of_unique_delta_nodes,omitempty"`
func configGinStatisticsTimeSeriesRouter(router gin.IRoutes) {
	router.GET("/stats/deals-attempted", ConverHttprouterToGin(GetRangeOfDealsAttempted))
}

func GetRangeOfDealsAttempted(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	record, err := dao.GetDealsAttemptedInRange(from, to)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}
