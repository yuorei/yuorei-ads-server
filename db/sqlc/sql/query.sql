-- name: GetUserById :one
SELECT * FROM users WHERE user_id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (user_id,role) VALUES (?,?);

-- name: DeleteUser :execresult
UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE user_id = ?;

-- name: GetCampaignById :one
SELECT * FROM campaigns WHERE campaign_id = ? LIMIT 1;

-- name: GetCampaignsByUserId :many
SELECT * FROM campaigns WHERE user_id = ?;

-- name: CreateCampaign :execresult
INSERT INTO campaigns (campaign_id, user_id, name, budget, start_date, end_date) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateCampaign :execresult
UPDATE campaigns SET name = ?, budget = ?, start_date = ?, end_date = ?, updated_at = CURRENT_TIMESTAMP WHERE campaign_id = ?;

-- name: DeleteCampaign :execresult
UPDATE campaigns SET deleted_at = CURRENT_TIMESTAMP WHERE campaign_id = ?;

-- name: GetAdGroupById :one
SELECT * FROM ad_groups WHERE ad_group_id = ? LIMIT 1;

-- name: GetAdGroupsByCampaignId :many
SELECT * FROM ad_groups WHERE campaign_id = ?;

-- name: CreateAdGroup :execresult
INSERT INTO ad_groups (ad_group_id, campaign_id, name) VALUES (?, ?, ?);

-- name: UpdateAdGroup :execresult
UPDATE ad_groups SET name = ?, updated_at = CURRENT_TIMESTAMP WHERE ad_group_id = ?;

-- name: DeleteAdGroup :execresult
UPDATE ad_groups SET deleted_at = CURRENT_TIMESTAMP WHERE ad_group_id = ?;

-- name: GetAdById :one
SELECT * FROM ads WHERE ad_id = ? LIMIT 1;

-- name: GetAdsByAdGroupId :many
SELECT * FROM ads WHERE ad_group_id = ?;

-- name: CreateAd :execresult
INSERT INTO ads (ad_id, ad_group_id, type, content) VALUES (?, ?, ?, ?);

-- name: UpdateAd :execresult
UPDATE ads SET type = ?, content = ?, updated_at = CURRENT_TIMESTAMP WHERE ad_id = ?;

-- name: DeleteAd :execresult
UPDATE ads SET deleted_at = CURRENT_TIMESTAMP WHERE ad_id = ?;

-- name: GetImpressionById :one
SELECT * FROM impressions WHERE impression_id = ? LIMIT 1;

-- name: GetImpressionsByAdId :many
SELECT * FROM impressions WHERE ad_id = ?;

-- name: CreateImpression :execresult
INSERT INTO impressions (impression_id, ad_id, date, impressions, clicks) VALUES (?, ?, ?, ?, ?);

-- name: UpdateImpression :execresult
UPDATE impressions SET impressions = ?, clicks = ?, updated_at = CURRENT_TIMESTAMP WHERE impression_id = ?;

-- name: DeleteImpression :execresult
UPDATE impressions SET deleted_at = CURRENT_TIMESTAMP WHERE impression_id = ?;

-- name: GetTargetingById :one
SELECT * FROM targeting WHERE targeting_id = ? LIMIT 1;

-- name: GetTargetingByAdId :many
SELECT * FROM targeting WHERE ad_id = ?;

-- name: CreateTargeting :execresult
INSERT INTO targeting (targeting_id, ad_id, type, value) VALUES (?, ?, ?, ?);

-- name: UpdateTargeting :execresult
UPDATE targeting SET type = ?, value = ?, updated_at = CURRENT_TIMESTAMP WHERE targeting_id = ?;

-- name: DeleteTargeting :execresult
UPDATE targeting SET deleted_at = CURRENT_TIMESTAMP WHERE targeting_id = ?;
