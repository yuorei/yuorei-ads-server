-- name: CheckOrganization :one
SELECT * FROM organizations LEFT JOIN organizations_users ON organizations.organization_id = organizations_users.organization_id WHERE organizations.organization_id = ? AND organizations_users.user_id = ?;

-- name: ListCampaignByOrganizationID :many
SELECT *
FROM campaigns AS c
LEFT JOIN organizations_users AS ou ON c.user_id = ou.user_id
WHERE ou.organization_id = ?
LIMIT ? OFFSET ?;

-- name: GetCampaign :one
SELECT * FROM campaigns WHERE campaign_id = ? LIMIT 1;

-- name: ListAds :many
SELECT * FROM ads LIMIT ? OFFSET ?;

-- name: ListAdsByCampaignID :many
SELECT * FROM ads WHERE campaign_id = ? LIMIT ? OFFSET ?;

-- name: GetAd :one
SELECT * FROM ads WHERE ad_id = ? LIMIT 1;

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

-- name: GetAdVideos :many
SELECT
    v.ad_id,
    v.title,
    v.description,
    v.thumbnail_url,
    v.video_url,
    a.ad_link
FROM
    ad_videos AS v
LEFT JOIN
    ads AS a
ON
    v.ad_id = a.ad_id
WHERE
    a.is_approval = TRUE
    AND a.is_open = TRUE
    AND v.deleted_at IS NULL
    AND a.deleted_at IS NULL;
