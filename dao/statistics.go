package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// function to get all totals info
func GetOpenTotalInfoStats() (interface{}, error) {

	// total deals attempted
	var statsTotal struct {
		TotalDealsAttempted                       int `json:"total_deals_attempted,omitempty"`
		TotalE2EDealsAttempted                    int `json:"total_e2e_deals_attempted,omitempty"`
		TotalImportDealsAttempted                 int `json:"total_import_deals_attempted,omitempty"`
		TotalPieceCommitmentsComputeAttempted     int `json:"total_piece_commitments_compute_attempted,omitempty"`
		TotalDealsAttemptedSize                   int `json:"total_deals_attempted_size,omitempty"`
		TotalE2EDealsAttemptedSize                int `json:"total_e2e_deals_attempted_size,omitempty"`
		TotalImportDealsAttemptedSize             int `json:"total_import_deals_attempted_size,omitempty"`
		TotalPieceCommitmentsComputeAttemptedSize int `json:"total_piece_commitments_compute_attempted_size,omitempty"`
		TotalDealsSucceeded                       int `json:"total_deals_succeeded,omitempty"`
		TotalE2ESucceeded                         int `json:"total_e2e_succeeded,omitempty"`
		TotalImportSucceeded                      int `json:"total_import_succeeded,omitempty"`
		TotalPieceCommitmentsComputeSucceeded     int `json:"total_piece_commitments_compute_succeeded,omitempty"`
		TotalDealsSucceededSize                   int `json:"total_deals_succeeded_size,omitempty"`
		TotalE2ESucceededSize                     int `json:"total_e2e_succeeded_size,omitempty"`
		TotalImportSucceededSize                  int `json:"total_import_succeeded_size,omitempty"`
		TotalPieceCommitmentsComputeSucceededSize int `json:"total_piece_commitments_compute_succeeded_size,omitempty"`
		TotalInProgressDeals24h                   int `json:"total_in_progress_deals_24h,omitempty"`
		TotalInProgressE2EDeals24h                int `json:"total_in_progress_e2e_deals_24h,omitempty"`
		TotalInProgressImportDeals24h             int `json:"total_in_progress_import_deals_24h,omitempty"`
		TotalInProgressDealsSize24h               int `json:"total_in_progress_deals_size_24h,omitempty"`
		TotalInProgressE2EDealsSize24h            int `json:"total_in_progress_e2e_deals_size_24h,omitempty"`
		TotalInProgressImportDealsSize24h         int `json:"total_in_progress_import_deals_size_24h,omitempty"`
		TotalNumberOfSpsWorkWith                  int `json:"total_number_of_sps_worked_with,omitempty"`
		TotalNumberOfUniqueDeltaNodes             int `json:"total_number_of_unique_delta_nodes,omitempty"`
	}

	val, ok := Cacher.Get("statsTotal")
	if !ok {

		DB.Transaction(func(tx *gorm.DB) error {

			var totalDealsAttempted int
			row := tx.Raw("select * from mv_deals_attempted").Row()
			err := row.Scan(&totalDealsAttempted)
			if err != nil {
				fmt.Println("Error in getting total deals attempted", err)
				totalDealsAttempted = 0
				//return err
			}
			fmt.Sprintf("totalDealsAttempted: %d", totalDealsAttempted)
			statsTotal.TotalDealsAttempted = totalDealsAttempted

			var totalDealsAttemptedSize int
			row = tx.Raw("select * from mv_deals_attempted_size").Row()
			err = row.Scan(&totalDealsAttemptedSize)
			if err != nil {
				fmt.Println("Error in getting total deals attempted size", err)
				totalDealsAttemptedSize = 0
				//return err
			}
			statsTotal.TotalDealsAttemptedSize = totalDealsAttemptedSize

			var totalE2EDealsAttempted int
			row = tx.Raw("select * from mv_e2e_deals_attempted").Row()
			err = row.Scan(&totalE2EDealsAttempted)
			if err != nil {
				fmt.Println("Error in getting total e2e deals attempted", err)
				totalE2EDealsAttempted = 0
				//return err
			}
			statsTotal.TotalE2EDealsAttempted = totalE2EDealsAttempted

			var totalPieceCommitmentsComputeAttempted int
			row = tx.Raw("select * from mv_commp_compute_attempted").Row()
			err = row.Scan(&totalPieceCommitmentsComputeAttempted)
			if err != nil {
				fmt.Println("Error in getting total piece commitments compute attempted", err)
				totalPieceCommitmentsComputeAttempted = 0
				//return err
			}
			statsTotal.TotalPieceCommitmentsComputeAttempted = totalPieceCommitmentsComputeAttempted

			var totalPieceCommitmentsComputeAttemptedSize int
			row = tx.Raw("select * from mv_commp_compute_attempted_size").Row()
			err = row.Scan(&totalPieceCommitmentsComputeAttemptedSize)
			if err != nil {
				fmt.Println("Error in getting total piece commitments compute attempted size", err)
				totalPieceCommitmentsComputeAttemptedSize = 0
				//return err
			}
			statsTotal.TotalPieceCommitmentsComputeAttemptedSize = totalPieceCommitmentsComputeAttemptedSize

			var totalE2EDealsAttemptedSize int
			row = tx.Raw("select * from mv_e2e_deals_attempted_size").Row()
			err = row.Scan(&totalE2EDealsAttemptedSize)
			if err != nil {
				fmt.Println("Error in getting total e2e deals attempted size", err)
				totalE2EDealsAttemptedSize = 0
				//return err
			}
			statsTotal.TotalE2EDealsAttemptedSize = totalE2EDealsAttemptedSize

			var totalImportDealsAttempted int
			row = tx.Raw("select * from mv_import_deals_attempted").Row()
			err = row.Scan(&totalImportDealsAttempted)
			if err != nil {
				fmt.Println("Error in getting total import deals attempted", err)
				totalImportDealsAttempted = 0
				//return err
			}
			statsTotal.TotalImportDealsAttempted = totalImportDealsAttempted

			var totalImportDealsAttemptedSize int
			row = tx.Raw("select * from mv_import_deals_attempted_size").Row()
			err = row.Scan(&totalImportDealsAttemptedSize)
			if err != nil {
				fmt.Println("Error in getting total import deals attempted size", err)
				totalImportDealsAttemptedSize = 0
				//return err
			}
			statsTotal.TotalImportDealsAttemptedSize = totalImportDealsAttemptedSize

			var totalDealsSucceeded int
			row = tx.Raw("select * from mv_deals_succeeded").Row()
			err = row.Scan(&totalDealsSucceeded)
			if err != nil {
				fmt.Println("Error in getting total deals succeeded", err)
				totalDealsSucceeded = 0
				//return err
			}
			statsTotal.TotalDealsSucceeded = totalDealsSucceeded

			var totalDealsSucceededSize int
			row = tx.Raw("select * from mv_deals_succeeded_size").Row()
			err = row.Scan(&totalDealsSucceededSize)
			if err != nil {
				fmt.Println("Error in getting total deals succeeded size", err)
				totalDealsSucceededSize = 0
				//return err
			}
			statsTotal.TotalDealsSucceededSize = totalDealsSucceededSize

			var totalE2EDealsSucceeded int
			row = tx.Raw("select * from mv_e2e_deals_succeeded").Row()
			err = row.Scan(&totalE2EDealsSucceeded)
			if err != nil {
				fmt.Println("Error in getting total e2e deals succeeded", err)
				totalE2EDealsSucceeded = 0
				//return err
			}
			statsTotal.TotalE2ESucceeded = totalE2EDealsSucceeded

			var totalPieceCommitmentsComputeSucceeded int
			row = tx.Raw("select * from mv_commp_compute_succeeded").Row()
			err = row.Scan(&totalPieceCommitmentsComputeSucceeded)
			if err != nil {
				fmt.Println("Error in getting total piece commitments compute succeeded", err)
				totalPieceCommitmentsComputeSucceeded = 0
				//return err
			}
			statsTotal.TotalPieceCommitmentsComputeSucceeded = totalPieceCommitmentsComputeSucceeded

			var totalPieceCommitmentsComputeSucceededSize int
			row = tx.Raw("select * from mv_commp_compute_succeeded_size").Row()
			err = row.Scan(&totalPieceCommitmentsComputeSucceededSize)
			if err != nil {
				fmt.Println("Error in getting total piece commitments compute succeeded size", err)
				totalPieceCommitmentsComputeSucceededSize = 0
				//return err
			}
			statsTotal.TotalPieceCommitmentsComputeSucceededSize = totalPieceCommitmentsComputeSucceededSize

			var totalE2EDealsSucceededSize int
			row = tx.Raw("select * from mv_e2e_deals_succeeded_size").Row()
			//
			err = row.Scan(&totalE2EDealsSucceededSize)
			if err != nil {
				fmt.Println("Error in getting total e2e deals succeeded size", err)
				totalE2EDealsSucceededSize = 0
				//return err
			}
			statsTotal.TotalE2ESucceededSize = totalE2EDealsSucceededSize

			var totalImportDealsSucceeded int
			row = tx.Raw("select * from mv_import_deals_succeeded").Row()
			err = row.Scan(&totalImportDealsSucceeded)
			if err != nil {
				fmt.Println("Error in getting total import deals succeeded", err)
				totalDealsSucceeded = 0
				//return err
			}
			statsTotal.TotalImportSucceeded = totalImportDealsSucceeded

			var totalImportDealsSucceededSize int
			row = tx.Raw("select * from mv_import_deals_succeeded_size").Row()
			err = row.Scan(&totalImportDealsSucceededSize)
			if err != nil {
				fmt.Println("Error in getting total import deals succeeded size", err)
				totalImportDealsSucceededSize = 0
				//return err
			}
			statsTotal.TotalImportSucceededSize = totalImportDealsSucceededSize

			var totalInProgressDeals int

			row = tx.Raw("select * from mv_total_in_progress_deals_24").Row()
			err = row.Scan(&totalInProgressDeals)
			if err != nil {
				fmt.Println("Error in getting total in progress deals", err)
				totalInProgressDeals = 0
				//return err
			}
			statsTotal.TotalInProgressDeals24h = totalInProgressDeals

			var totalInProgressE2EDeals int
			row = tx.Raw("select * from mv_total_in_progress_e2e_deals_24").Row()
			err = row.Scan(&totalInProgressE2EDeals)
			if err != nil {
				fmt.Println("Error in getting total in progress e2e deals", err)
				totalInProgressE2EDeals = 0
				//return err
			}
			statsTotal.TotalInProgressE2EDeals24h = totalInProgressE2EDeals

			var totalInProgressImportDeals int
			row = tx.Raw("select * from mv_total_in_progress_import_deals_24").Row()
			err = row.Scan(&totalInProgressImportDeals)
			if err != nil {
				fmt.Println("Error in getting total in progress import deals", err)
				totalInProgressImportDeals = 0
				//return err
			}
			statsTotal.TotalInProgressImportDeals24h = totalInProgressImportDeals

			var totalNumberOfSpsWorkWith int
			// select count(miners) as total_rows from (select distinct(miner) as miners from content_miner_logs group by miner) subquery;
			row = tx.Raw("select * from mv_number_of_sps_work_with").Row()
			err = row.Scan(&totalNumberOfSpsWorkWith)
			if err != nil {
				fmt.Println("Error in getting total number of sps work with", err)
				totalNumberOfSpsWorkWith = 0
				//return err
			}
			statsTotal.TotalNumberOfSpsWorkWith = totalNumberOfSpsWorkWith

			var totalNumberOfUniqueDeltaNodes int
			row = tx.Raw("select * from mv_number_of_unique_delta_nodes").Row()
			err = row.Scan(&totalNumberOfUniqueDeltaNodes)
			if err != nil {
				fmt.Println("Error in getting total number of unique delta nodes", err)
				totalNumberOfUniqueDeltaNodes = 0
				//return err
			}
			statsTotal.TotalNumberOfUniqueDeltaNodes = totalNumberOfUniqueDeltaNodes
			return nil
		})
		val = statsTotal
		Cacher.Add("statsTotal", val)
	}
	return val, nil

}

func GetAllWalletAddrs() (interface{}, error) {
	type WalletLog struct {
		Addr string
	}
	// total deals attempted
	var addresses []WalletLog

	val, ok := Cacher.Get("allWalletAddrs")
	if !ok {

		DB.Model(&WalletLog{}).
			Select("addr").
			Group("addr").
			Find(&addresses)
		var allWalletAddrs []string
		for _, addr := range addresses {
			allWalletAddrs = append(allWalletAddrs, addr.Addr)
		}
		val = allWalletAddrs
		Cacher.Add("allWalletAddrs", val)
	}
	return val, nil
}

func GetAllSPs() (interface{}, error) {

	// total deals attempted
	type ContentMinerLog struct {
		Miner string
	}

	var miners []ContentMinerLog
	val, ok := Cacher.Get("allSpsStats")
	if !ok {

		DB.Model(&ContentMinerLog{}).
			Select("miner").
			Group("miner").
			Find(&miners)
		var minersStr []string
		for _, miner := range miners {
			minersStr = append(minersStr, miner.Miner)
		}
		val = minersStr

		Cacher.Add("allSpsStats", minersStr)
	}
	return val, nil
}

func GetAllDeltaIps() (interface{}, error) {

	// total deals attempted
	type DeltaStartupLog struct {
		IPAddress string
	}
	var ipAddresses []DeltaStartupLog
	val, ok := Cacher.Get("allDeltaIps")
	if !ok {

		DB.Model(&DeltaStartupLog{}).
			Select("ip_address").
			Where("ip_address <> ?", "").
			Group("ip_address").
			Find(&ipAddresses)

		var allDeltaIps []string
		for _, ip := range ipAddresses {
			allDeltaIps = append(allDeltaIps, ip.IPAddress)
		}
		val = allDeltaIps
		Cacher.Add("allDeltaIps", val)
	}
	return val, nil
}

//
