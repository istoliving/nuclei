// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"

	"github.com/jackc/pgtype"
)

type Issue struct {
	Matchedat     sql.NullString
	Title         sql.NullString
	Severity      sql.NullString
	Createdat     sql.NullTime
	Updatedat     sql.NullTime
	Scansource    sql.NullString
	Issuestate    sql.NullString
	Description   sql.NullString
	Author        sql.NullString
	Cvss          []int32
	Cwe           []int32
	Labels        []string
	Issuedata     sql.NullString
	Issuetemplate sql.NullString
	Remediation   sql.NullString
	Debug         sql.NullString
	ID            int64
	Scanid        sql.NullInt64
}

type Scan struct {
	Name       sql.NullString
	Status     sql.NullString
	Scantime   sql.NullTime
	Hosts      sql.NullInt64
	Scansource sql.NullString
	Progress   sql.NullFloat64
	Templates  []string
	Targets    []string
	Debug      sql.NullBool
	ID         int64
}

type Setting struct {
	Alerting pgtype.JSON
	Config   pgtype.JSON
}

type Target struct {
	ID        int64
	Name      sql.NullString
	Filepath  string
	Total     sql.NullInt64
	Createdat sql.NullTime
	Updatedat sql.NullTime
}

type Template struct {
	ID        int64
	Name      sql.NullString
	Folder    sql.NullString
	Path      string
	Contents  string
	Createdat sql.NullTime
	Updatedat sql.NullTime
}
