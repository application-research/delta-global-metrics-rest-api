package dao

// function to get all totals info
func GetOpenTotalInfoStats() (interface{}, error) {

	var totalContentConsumed int64
	DB.Raw("select count(*) from content_logs").Count(&totalContentConsumed)

	var totalTransferStarted int64
	DB.Raw("select count(*) from content_logs where status = 'transfer-started'").Count(&totalTransferStarted)

	var totalTransferFinished int64
	DB.Raw("select count(*) from content_logs where status = 'transfer-finished'").Count(&totalTransferFinished)

	var totalProposalMade int64
	DB.Raw("select count(*) from content_deal_proposal_logs").Count(&totalProposalMade)

	var totalCommitmentPiece int64
	DB.Raw("select count(*) from piece_commitment_logs").Count(&totalCommitmentPiece)

	var totalPieceCommitted int64
	DB.Raw("select count(*) from piece_commitment_logs where status = 'committed'").Count(&totalPieceCommitted)

	var totalMiners int64
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
	DB.Raw("select count(*) from content_logs where status = 'deal-proposal-sent'").Count(&totalProposalSent)

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
