CREATE MATERIALIZED VIEW IF NOT EXISTS mv_deals_attempted
AS
select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c group by system_content_id) subquery;

CREATE MATERIALIZED VIEW IF NOT EXISTS mv_deals_attempted_size
AS
select sum(size) as total_size_sum from (select c.size as size,system_content_id from content_logs c where (system_content_id is null or system_content_id is not null) and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '') group by c.size,system_content_id) subquery;

--mv_e2e_deals_attempted
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_e2e_deals_attempted
AS
select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' group by system_content_id) subquery;

--mv_e2e_deals_attempted_size
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_e2e_deals_attempted_size
AS
select sum(size) as total_size_sum from (select c.size as size,system_content_id from content_logs c where c.connection_mode = 'e2e' and (system_content_id is null or system_content_id is not null) group by c.size,system_content_id) subquery;

--mv_import_deals_attempted
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_import_deals_attempted
AS
select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' group by system_content_id) subquery;

--mv_import_deals_attempted_size
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_import_deals_attempted_size
AS
select sum(size) as total_size_sum from (select c.size as size,system_content_id from content_logs c where c.connection_mode = 'import' and (system_content_id is null or system_content_id is not null) group by c.size,system_content_id) subquery;

--mv_deals_succeeded
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_deals_succeeded
AS
select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where status in ('deal-proposal-sent','transfer-started','transfer-finished') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery;

--mv_deals_succeeded_size
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_deals_succeeded_size
AS
select sum(size) as total_size_sum from (select p.padded_piece_size as size,system_content_id from content_logs c, piece_commitment_logs p where c.piece_commitment_id = p.system_content_piece_commitment_id and c.status in ('deal-proposal-sent','transfer-started','transfer-finished') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id, p.padded_piece_size) subquery;

--mv_e2e_deals_succeeded
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_e2e_deals_succeeded
AS
select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'e2e' and status in ('transfer-started','transfer-finished') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery;

--mv_e2e_deals_succeeded_size
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_e2e_deals_succeeded_size
AS
select sum(size) as total_size_sum from (select p.padded_piece_size as size,system_content_id from content_logs c, piece_commitment_logs p where c.piece_commitment_id = p.system_content_piece_commitment_id and c.status in ('transfer-started','transfer-finished') and c.connection_mode = 'e2e' and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id, p.padded_piece_size) subquery;

--mv_import_deals_succeeded
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_import_deals_succeeded
AS
select sum(cnt) as total_rows from (select count(*) as cnt from content_logs c where c.connection_mode = 'import' and status in ('deal-proposal-sent') and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id) subquery;

--mv_import_deals_succeeded_size
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_import_deals_succeeded_size
AS
select sum(size) as total_size_sum from (select p.padded_piece_size as size,system_content_id from content_logs c, piece_commitment_logs p where c.piece_commitment_id = p.system_content_piece_commitment_id and c.status in ('deal-proposal-sent') and c.connection_mode = 'import' and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')  group by system_content_id, p.padded_piece_size) subquery;

--mv_commp_compute_succeeded
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_commp_compute_succeeded
AS
select sum(cnt) as total_rows from (select count(p.piece) as cnt from piece_commitment_logs p where p.status = 'committed' group by p.piece) subquery;

--mv_commp_compute_succeeded_size
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_commp_compute_succeeded_size
AS
select sum(size) as total_size_sum from (select p.size as size from piece_commitment_logs p where p.status = 'committed' group by p.size,p.piece) subquery;

--mv_commp_compute_attempted
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_commp_compute_attempted
AS
select sum(cnt) as total_rows from (select count(p.piece) as cnt from piece_commitment_logs p group by p.piece) subquery;

--mv_commp_compute_attempted_size
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_commp_compute_attempted_size
AS
select sum(size) as total_size_sum from (select p.size as size from piece_commitment_logs p group by p.size,p.piece) subquery;

--mv_number_of_sps_work_with
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_number_of_sps_work_with
AS
select count(miners) as total_rows from (select distinct(miner) as miners from content_miner_logs group by miner) subquery;

--mv_number_of_unique_delta_nodes
CREATE MATERIALIZED VIEW  IF NOT EXISTS mv_number_of_unique_delta_nodes
AS
select count(delta_node) as total_rows from (select distinct(delta_node_uuid) as delta_node from delta_startup_logs group by delta_node_uuid) subquery;