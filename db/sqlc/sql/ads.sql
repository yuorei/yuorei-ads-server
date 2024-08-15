-- name: CreateAd :execresult
insert into ads (ad_id, campaign_id, ad_type, is_approval,is_open,ad_link) values (?, ?, ?, ? , ?, ?);

-- name: CreateAdImage :execresult
insert into ad_images (ad_id, title, description, image_url) values (?, ?, ?, ?);

-- name: CreateAdVideo :execresult
insert into ad_videos (ad_id, title, description, thumbnail_url,video_url) values (?, ?, ?, ?,?);

-- name: GetAdById :one
select * from ads where ad_id = ? limit 1;

-- name: UpdateAd :execresult
update ads set is_approval = ?, is_open = ?, updated_at = CURRENT_TIMESTAMP where ad_id = ?;
