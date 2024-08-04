-- name: CreateOrganization :execresult
insert into organizations (organization_id, organization_name, representative_name, representative_email, purpose, category) values (?, ?, ?, ?, ?, ?);