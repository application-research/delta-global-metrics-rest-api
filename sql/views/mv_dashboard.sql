DROP MATERIALIZED VIEW IF EXISTS mv_top_sp_miners;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_top_sp_miners AS
select
    cml.miner,
    ROUND(SUM(cl.size) / (1024*1024*1024),2) AS size_gb,
    ROUND(SUM(cl.size) / (1024*1024*1024),2) / 1000 AS size_tb,
    cml.created_at,
    cml.updated_at
from content_miner_logs cml, content_logs cl where cl.system_content_id = cml.content and cl.delta_node_uuid = cml.delta_node_uuid and cl.status in ('transfer-started','transfer-finished','deal-proposal-sent')
group by cml.miner, cml.created_at, cml.updated_at order by size_gb desc;

DROP MATERIALIZED VIEW IF EXISTS mv_top_delta_nodes;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_top_delta_nodes AS
select
    cml.delta_node_uuid,
    iml.os_details,
    iml.public_ip,
    ROUND(SUM(cl.size) / (1024*1024*1024),2) AS size_gb,
    ROUND(SUM(cl.size) / (1024*1024*1024),2) / 1000 AS size_tb,
    cml.created_at,
    cml.updated_at
from content_miner_logs cml, content_logs cl, instance_meta_logs iml where cl.delta_node_uuid = cml.delta_node_uuid and iml.delta_node_uuid = cml.delta_node_uuid                                                        and cml.delta_node_uuid <> ''
group by cml.miner, cml.delta_node_uuid, iml.os_details, iml.instance_host_name, iml.public_ip, cml.created_at, cml.updated_at order by size_gb desc;