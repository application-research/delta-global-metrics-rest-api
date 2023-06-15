DROP MATERIALIZED VIEW IF EXISTS mv_content_logs_tbl;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_content_logs_tbl
AS
select * from content_logs;
CREATE UNIQUE INDEX ON mv_content_logs_tbl(id);

DROP MATERIALIZED VIEW IF EXISTS mv_content_deal_logs_tbl;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_content_deal_logs_tbl
AS
select * from content_deal_logs;
CREATE UNIQUE INDEX ON mv_content_deal_logs_tbl(id);

DROP MATERIALIZED VIEW IF EXISTS mv_content_deal_proposal_logs_tbl;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_content_deal_proposal_logs_tbl
AS
select * from content_deal_proposal_logs;
CREATE UNIQUE INDEX ON mv_content_deal_proposal_logs_tbl(id);

DROP MATERIALIZED VIEW IF EXISTS mv_content_miner_logs_tbl;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_content_miner_logs_tbl
AS
select * from content_miner_logs;
CREATE UNIQUE INDEX ON mv_content_miner_logs_tbl(id);