CREATE MATERIALIZED VIEW IF NOT EXISTS mv_onboarded_deals_by_sp_uuid_key
AS
select p.padded_piece_size as size,
       system_content_id, c.delta_node_uuid,c.requesting_api_key from content_logs c,
    piece_commitment_logs p
    where c.piece_commitment_id = p.system_content_piece_commitment_id
    and c.system_content_id in (
        select cd.content from content_deal_logs cd
        where cd.failed = false
        and cd.content in (
            select c.system_content_id
            from content_logs c, piece_commitment_logs p
            where c.piece_commitment_id = p.system_content_piece_commitment_id
            and c.status in ('deal-proposal-sent','transfer-started','transfer-finished')
            and (c.delta_node_uuid is not null or c.delta_node_uuid is null or c.delta_node_uuid = '')
    group by system_content_id, p.padded_piece_size))
group by system_content_id, p.padded_piece_size,c.delta_node_uuid,c.requesting_api_key;