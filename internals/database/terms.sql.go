// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: terms.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllTerms = `-- name: GetAllTerms :many
SELECT terms.id, terms.created_at, terms.updated_at, session_id, term_id, vacation_date, next_resumption_date, number_school_opened, lock, sessions.id, sessions.created_at, sessions.updated_at, name, sessions.name || ' - ' || term_id AS display_name
FROM terms 
LEFT JOIN sessions on terms.session_id = sessions.id
ORDER BY terms.created_at DESC
`

type GetAllTermsRow struct {
	ID                 pgtype.UUID
	CreatedAt          pgtype.Timestamp
	UpdatedAt          pgtype.Timestamp
	SessionID          pgtype.UUID
	TermID             TermEnum
	VacationDate       pgtype.Timestamp
	NextResumptionDate pgtype.Timestamp
	NumberSchoolOpened pgtype.Int8
	Lock               bool
	ID_2               pgtype.UUID
	CreatedAt_2        pgtype.Timestamp
	UpdatedAt_2        pgtype.Timestamp
	Name               pgtype.Text
	DisplayName        interface{}
}

func (q *Queries) GetAllTerms(ctx context.Context) ([]GetAllTermsRow, error) {
	rows, err := q.db.Query(ctx, getAllTerms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllTermsRow
	for rows.Next() {
		var i GetAllTermsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SessionID,
			&i.TermID,
			&i.VacationDate,
			&i.NextResumptionDate,
			&i.NumberSchoolOpened,
			&i.Lock,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Name,
			&i.DisplayName,
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

const getAllUnlockedTerms = `-- name: GetAllUnlockedTerms :many
SELECT terms.id, terms.created_at, terms.updated_at, session_id, term_id, vacation_date, next_resumption_date, number_school_opened, lock, sessions.id, sessions.created_at, sessions.updated_at, name, sessions.name || ' - ' || term_id AS display_name
FROM terms 
LEFT JOIN sessions on terms.session_id = sessions.id
WHERE lock = false
ORDER BY terms.created_at DESC
`

type GetAllUnlockedTermsRow struct {
	ID                 pgtype.UUID
	CreatedAt          pgtype.Timestamp
	UpdatedAt          pgtype.Timestamp
	SessionID          pgtype.UUID
	TermID             TermEnum
	VacationDate       pgtype.Timestamp
	NextResumptionDate pgtype.Timestamp
	NumberSchoolOpened pgtype.Int8
	Lock               bool
	ID_2               pgtype.UUID
	CreatedAt_2        pgtype.Timestamp
	UpdatedAt_2        pgtype.Timestamp
	Name               pgtype.Text
	DisplayName        interface{}
}

func (q *Queries) GetAllUnlockedTerms(ctx context.Context) ([]GetAllUnlockedTermsRow, error) {
	rows, err := q.db.Query(ctx, getAllUnlockedTerms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUnlockedTermsRow
	for rows.Next() {
		var i GetAllUnlockedTermsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SessionID,
			&i.TermID,
			&i.VacationDate,
			&i.NextResumptionDate,
			&i.NumberSchoolOpened,
			&i.Lock,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Name,
			&i.DisplayName,
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
