package dao

import (
	"log"
	"os"
)

type RefreshGlobalStatsViewResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

func RefreshGlobalStatsView() (interface{}, error) {

	refreshViews, err := os.ReadFile("sql/views/refresh_mv_stats.sql")
	refreshViewsStr := string(refreshViews)
	if err != nil {
		log.Fatalf("Got error when reading refresh_mv_stats.sql, the error is '%v'", err)
	}
	if err := DB.Exec(refreshViewsStr); err != nil {
		panic(err)
	}

	refreshGlobalStatsViewResponse := RefreshGlobalStatsViewResponse{
		Message:     "Success",
		Description: "Refreshed Global Stats View",
	}

	return refreshGlobalStatsViewResponse, nil

}
