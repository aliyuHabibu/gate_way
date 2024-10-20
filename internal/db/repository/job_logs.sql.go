// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: job_logs.sql

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgtype"
)

const createJobLog = `-- name: CreateJobLog :exec
insert into job_logs (created_at, level, job_id, message, metadata)
values ($1, $2, $3, $4, $5)
`

type CreateJobLogParams struct {
	CreatedAt time.Time
	Level     int16
	JobID     sql.NullString
	Message   string
	Metadata  pgtype.JSONB
}

func (q *Queries) CreateJobLog(ctx context.Context, arg CreateJobLogParams) error {
	_, err := q.db.Exec(ctx, createJobLog,
		arg.CreatedAt,
		arg.Level,
		arg.JobID,
		arg.Message,
		arg.Metadata,
	)
	return err
}

const listJobLogsByID = `-- name: ListJobLogsByID :many
select id, created_at, level, job_id, message, metadata from job_logs where job_id = $1 order by id limit $2
`

type ListJobLogsByIDParams struct {
	JobID sql.NullString
	Limit int32
}

func (q *Queries) ListJobLogsByID(ctx context.Context, arg ListJobLogsByIDParams) ([]JobLog, error) {
	rows, err := q.db.Query(ctx, listJobLogsByID, arg.JobID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []JobLog
	for rows.Next() {
		var i JobLog
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Level,
			&i.JobID,
			&i.Message,
			&i.Metadata,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
