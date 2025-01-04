-- name: CreateOrganization :execresult
insert into organizations (organization_id, organization_name, representative_user_id, purpose, category) values (?, ?, ?, ?, ?);

-- name: CreateOrganizationUser :execresult
insert into organizations_users (organization_id, user_id) values (?, ?);
