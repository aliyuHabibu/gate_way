// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: lock.sql

package repository

import (
	"context"
)

const advisoryTransactionLock = `-- name: AdvisoryTransactionLock :exec
SELECT pg_advisory_xact_lock($1)
`

func (q *Queries) AdvisoryTransactionLock(ctx context.Context, pgAdvisoryXactLock int64) error {
	_, err := q.db.Exec(ctx, advisoryTransactionLock, pgAdvisoryXactLock)
	return err
}

const advisoryTransactionUnlock = `-- name: AdvisoryTransactionUnlock :exec
SELECT pg_advisory_xact_lock($1)
`

func (q *Queries) AdvisoryTransactionUnlock(ctx context.Context, pgAdvisoryXactLock int64) error {
	_, err := q.db.Exec(ctx, advisoryTransactionUnlock, pgAdvisoryXactLock)
	return err
}