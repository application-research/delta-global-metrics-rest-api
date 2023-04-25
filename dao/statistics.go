package dao

import (
	"fmt"
	"github.com/application-research/delta-metrics-rest/model"
)

// function to get all totals info
func GetOpenTotalInfoStats() (interface{}, error) {

	var totalContentConsumed int64
	//SELECT SUM(cnt) as total_rows
	//FROM (
	//         SELECT COUNT(*) as cnt
	//         FROM content_logs
	//         GROUP BY system_content_id, delta_node_uuid
	//     ) subquery;
	DB.Raw("SELECT SUM(cnt) as total_rows FROM (SELECT COUNT(*) as cnt FROM content_logs GROUP BY system_content_id, delta_node_uuid) subquery").Find(&totalContentConsumed)
	//DB.Raw("select count(*) from content_logs").Count(&totalContentConsumed)

	var totalTransferStarted int64
	//SELECT SUM(cnt) as total_rows
	//FROM (
	//         SELECT COUNT(*) as cnt
	//         FROM content_logs
	//         WHERE status = 'transfer-started'
	//         GROUP BY system_content_id, delta_node_uuid
	//     ) subquery;
	err := DB.Model(model.ContentLogs{}).Where("status = ?", "transfer-started").Count(&totalTransferStarted).Error
	if err != nil {
		fmt.Println("Error in GetOpenTotalInfoStats", err)
		return nil, err
	}

	var totalTransferFinished int64
	//SELECT SUM(cnt) as total_rows
	//FROM (
	//         SELECT COUNT(*) as cnt
	//         FROM content_logs
	//         WHERE status = 'transfer-started'
	//         GROUP BY system_content_id, delta_node_uuid
	//     ) subquery;
	err = DB.Model(model.ContentLogs{}).Where("status = ?", "transfer-finished").Count(&totalTransferFinished).Error
	if err != nil {
		fmt.Println("Error in GetOpenTotalInfoStats", err)
		return nil, err
	}
	var totalProposalMade int64
	//SELECT SUM(cnt) as total_rows
	//FROM (
	//	SELECT COUNT(*) as cnt
	//FROM content_deal_proposal_logs
	//GROUP BY system_content_deal_proposal_id, delta_node_uuid
	//) subquery;
	DB.Raw("select count(*) from content_deal_proposal_logs").Find(&totalProposalMade)

	var totalCommitmentPiece int64
	//SELECT SUM(cnt) as total_rows
	//FROM (
	//	SELECT COUNT(*) as cnt
	//FROM piece_commitment_logs
	//GROUP BY system_content_piece_commitment_id, delta_node_uuid
	//) subquery;
	DB.Raw("select count(*) from piece_commitment_logs").Find(&totalCommitmentPiece)

	var totalPieceCommitted int64
	//SELECT SUM(cnt) as total_rows
	//FROM (
	//	SELECT COUNT(*) as cnt
	//FROM piece_commitment_logs
	//WHERE status = 'committed'
	//GROUP BY system_content_piece_commitment_id, delta_node_uuid
	//) subquery;
	DB.Raw("select count(*) from piece_commitment_logs where status = 'committed'").Find(&totalPieceCommitted)

	var totalMiners int64
	//SELECT count(cnt) as total_rows
	//FROM (
	//	SELECT miner as cnt
	//FROM content_miner_logs
	//GROUP BY miner
	//) subquery;
	rows, err := DB.Raw("select distinct(miner) from content_miner_logs").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var miner string
		rows.Scan(&miner)
		totalMiners++
	}

	var totalStorageAllocated int64
	DB.Raw("select sum(size) from content_logs").Count(&totalStorageAllocated)

	var totalProposalSent int64
	DB.Raw("select count(*) from content_logs c where status = 'deal-proposal-sent' and c.system_content_id in (select c.system_content_id from content_logs where status = 'deal-proposal-sent')").Count(&totalProposalSent)

	var totalSealedDealInBytes int64
	DB.Raw("select sum(size) from content_logs where status in ('transfer-started','transfer-finished','deal-proposal-sent')").Scan(&totalSealedDealInBytes)

	var totalImportDeals int64
	DB.Raw("select count(*) from content_logs where connection_mode = 'import'").Count(&totalImportDeals)

	var totalE2EDeals int64
	DB.Raw("select count(*) from content_logs where connection_mode = 'e2e'").Count(&totalE2EDeals)

	var totalE2EDealsInBytes int64
	DB.Raw("select sum(size) from content_logs where connection_mode = 'e2e'").Count(&totalE2EDealsInBytes)

	var totalImportDealsInBytes int64
	DB.Raw("select sum(size) from content_logs where connection_mode = 'import'").Count(&totalImportDealsInBytes)

	return map[string]interface{}{
		"total_content_consumed":      totalContentConsumed,
		"total_transfer_started":      totalTransferStarted,
		"total_transfer_finished":     totalTransferFinished,
		"total_piece_commitment_made": totalCommitmentPiece,
		"total_piece_committed":       totalPieceCommitted,
		"total_miners":                totalMiners,
		"total_storage_allocated":     totalStorageAllocated,
		"total_proposal_made":         totalProposalMade,
		"total_proposal_sent":         totalProposalSent,
		"total_sealed_deal_in_bytes":  totalSealedDealInBytes,
		"total_import_deals":          totalImportDeals,
		"total_e2e_deals":             totalE2EDeals,
		"total_e2e_deals_in_bytes":    totalE2EDealsInBytes,
		"total_import_deals_in_bytes": totalImportDealsInBytes,
	}, nil
}
