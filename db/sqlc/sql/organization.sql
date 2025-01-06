-- name: GetOrganization :one
SELECT * FROM organizations WHERE organization_id = ?;

-- name: GetOrganizationByUserID :one
SELECT * FROM organizations LEFT JOIN organizations_users ON organizations.organization_id = organizations_users.organization_id WHERE organizations_users.user_id = ?;

-- name: CreateOrganization :execresult
insert into organizations (organization_id, organization_name, representative_user_id, purpose, category) values (?, ?, ?, ?, ?);

-- name: CreateOrganizationUser :execresult
insert into organizations_users (organization_id, user_id) values (?, ?);
