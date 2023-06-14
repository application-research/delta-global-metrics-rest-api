DROP MATERIALIZED VIEW IF EXISTS mv_onboarded_deals_by_sp_uuid_key;
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_onboarded_deals_by_sp_uuid_key
AS
select p.padded_piece_size as size,
       system_content_id, c.delta_node_uuid,c.requesting_api_key, cd.miner, cd.deal_uuid, cd.deal_id
from content_logs c, piece_commitment_logs p, content_deal_logs cd
where c.piece_commitment_id = p.system_content_piece_commitment_id
  and c.system_content_id = cd.content
  and c.status in ('deal-proposal-sent','transfer-started','transfer-finished')
group by system_content_id, p.padded_piece_size,c.delta_node_uuid,c.requesting_api_key,cd.miner,cd.deal_uuid, cd.deal_id;
