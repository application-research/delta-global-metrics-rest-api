package dao

//-- deals attempted and size (with created_at filter)
//select sum(cnt) as total_rows from (select count(dt_chan) as cnt from content_deal_logs group by dt_chan) subquery;
//select sum(cnt) as total_rows from (select count(dt_chan) as cnt from content_deal_logs where (created_at between '2023-01-01' and '2023-12-30') group by dt_chan) subquery;
//select sum(size) as total_size_sum from (select c.size as size,cd.dt_chan from content_deal_logs cd, content_logs c where cd.content = c.system_content_id and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '') group by c.size,cd.dt_chan) subquery;
//select sum(size) as total_size_sum from (select c.size as size,cd.dt_chan from content_deal_logs cd, content_logs c where cd.content = c.system_content_id and (created_at between '2023-01-01' and '2023-12-30') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '') group by c.size,cd.dt_chan) subquery;
//
//-- e2e deals attempted and size (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and (created_at between '2023-01-01' and '2023-12-30') group by system_content_id) subquery;
//select sum(size) as total_size_sum from (select c.size as size,system_content_id from content_logs c where c.connection_mode = 'e2e' and (system_content_id is null or system_content_id is not null) and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '') group by c.size,system_content_id) subquery;
//select sum(size) as total_size_sum from (select c.size as size,system_content_id from content_logs c where c.connection_mode = 'e2e' and (created_at between '2023-01-01' and '2023-12-30') and (system_content_id is null or system_content_id is not null) and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '') group by c.size,system_content_id) subquery;
//
//-- import deals attempted and size (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and (created_at between '2023-01-01' and '2023-12-30') group by system_content_id) subquery;
//select sum(size) as total_size_sum from (select c.size as size,system_content_id from content_logs c where c.connection_mode = 'import' and (system_content_id is null or system_content_id is not null) and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '') group by c.size,system_content_id) subquery;
//select sum(size) as total_size_sum from (select c.size as size,system_content_id from content_logs c where c.connection_mode = 'import' and (created_at between '2023-01-01' and '2023-12-30') and (system_content_id is null or system_content_id is not null) and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '') group by c.size,system_content_id) subquery;
//
//-- e2e succeeded (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started','transfer-finished') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started','transfer-finished') and (created_at between '2023-01-01' and '2023-12-30') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery;
//
//-- import succeeded (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status in ('deal-proposal-sent') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status in ('deal-proposal-sent') and (created_at between '2023-01-01' and '2023-12-30') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery;
//
//-- e2e failed (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status not in ('transfer-started','transfer-finished') group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status not in ('transfer-started','transfer-finished') and (created_at between '2023-01-01' and '2023-12-30') group by system_content_id) subquery;
//
//-- import failed (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status not in ('deal-proposal-sent') group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status not in ('deal-proposal-sent') and (created_at between '2023-01-01' and '2023-12-30') group by system_content_id) subquery;
//
//-- import failed (with created_at filter)
//select sum(cnt) as total_rows from (select count(system_content_deal_id) as cnt from content_deal_logs cd, content_logs c where c.system_content_id = cd.content and c.connection_mode = 'import' and c.status not in ('deal-proposal-sent') group by system_content_deal_id) subquery;
//select sum(cnt) as total_rows from (select count(system_content_deal_id) as cnt from content_deal_logs cd, content_logs c where c.system_content_id = cd.content and c.connection_mode = 'import' and c.status not in ('deal-proposal-sent') and (created_at between '2023-01-01' and '2023-12-30') group by system_content_deal_id) subquery;
//
//-- total deals active (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started') group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started') and (created_at between '2023-01-01' and '2023-12-30') group by system_content_id) subquery;
//
//-- total e2e deals active (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started') group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started') and (created_at between '2023-01-01' and '2023-12-30') group by system_content_id) subquery;
//
//-- total import deals active (with created_at filter)
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status in ('making-deal-proposal') group by system_content_id) subquery;
//select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status in ('making-deal-proposal') and (created_at between '2023-01-01' and '2023-12-30') group by system_content_id) subquery;
//
//-- total number of SPs (with created_at filter)
//select count(miners) as total_rows from (select distinct(miner) as miners from content_miner_logs group by miner) subquery;
//
//-- total number of unique delta nodes (with created_at filter)
//select count(delta_node) as total_rows from (select distinct(delta_node_uuid) as delta_node from delta_startup_logs group by delta_node_uuid) subquery;
//
//select distinct(delta_node_uuid) as delta_node from delta_startup_logs group by delta_node_uuid;
//
//select * from content_deal_logs cd where (cd.delta_node_uuid is null or cd.delta_node_uuid = '') and requester_info = '139.178.81.109' order by created_at desc;
//
//select count(*) from content_deal_logs cd where (cd.delta_node_uuid is null or cd.delta_node_uuid = '') and node_info = 'shuttle-4';
//select * from content_deal_logs cd where cd.delta_node_uuid is null or cd.delta_node_uuid = '' order by created_at desc;
//select * from content_deal_logs cd where cd.delta_node_uuid is null order by created_at desc;

// function to get all totals info
func GetOpenTotalInfoStats() (interface{}, error) {

	// total deals attempted
	var totalDealsAttempted int64
	DB.Raw("select sum(cnt) as total_deals_attempted from (select count(dt_chan) as cnt from content_deal_logs group by dt_chan) subquery").Scan(&totalDealsAttempted)

	// total e2e deals attempted
	var totalE2EDealsAttempted int
	DB.Raw("select sum(cnt) as total_rows from (select count(dt_chan) as cnt from content_deal_logs where connection_mode = 'e2e' group by dt_chan) subquery").Find(&totalE2EDealsAttempted)

	// total import deals attempted
	var totalImportDealsAttempted int
	DB.Raw("select sum(cnt) as total_rows from (select count(dt_chan) as cnt from content_deal_logs where connection_mode = 'import' group by dt_chan) subquery").Find(&totalImportDealsAttempted)

	// total e2e deals succeeded
	var totalE2EDealsSucceeded int
	DB.Raw("select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started','transfer-finished') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery").Find(&totalE2EDealsSucceeded)

	// total import deals succeeded
	var totalImportDealsSucceeded int
	DB.Raw("select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status in ('making-deal-proposal','deal-proposal-sent','deal-proposal-accepted','deal-proposal-rejected','deal-activated','deal-terminated','deal-completed') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery").Find(&totalImportDealsSucceeded)

	// total number of sps
	var totalSPs int
	DB.Raw("select count(miners) as total_rows from (select distinct(miner) as miners from content_miner_logs group by miner) subquery").Find(&totalSPs)

	var totalDeltaNodes int
	DB.Raw("select count(delta_node) as total_rows from (select distinct(delta_node_uuid) as delta_node from delta_startup_logs group by delta_node_uuid) subquery").Find(&totalDeltaNodes)

	return map[string]interface{}{
		"total_deals_attempted":        totalDealsAttempted,
		"total_e2e_deals_attempted":    totalE2EDealsAttempted,
		"total_import_deals_attempted": totalImportDealsAttempted,
		"total_e2e_deals_succeeded":    totalE2EDealsSucceeded,
		"total_import_deals_succeeded": totalImportDealsSucceeded,
		"total_number_of_sps":          totalSPs,
		"total_number_of_delta_nodes":  totalDeltaNodes,
	}, nil

}
