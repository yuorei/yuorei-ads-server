// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: organization.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createOrganization = `-- name: CreateOrganization :execresult
insert into organizations (organization_id, organization_name, representative_name, representative_email, purpose, category) values (?, ?, ?, ?, ?, ?)
`

type CreateOrganizationParams struct {
	OrganizationID      string
	OrganizationName    string
	RepresentativeName  string
	RepresentativeEmail string
	Purpose             string
	Category            string
}

func (q *Queries) CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createOrganization,
		arg.OrganizationID,
		arg.OrganizationName,
		arg.RepresentativeName,
		arg.RepresentativeEmail,
		arg.Purpose,
		arg.Category,
	)
}
