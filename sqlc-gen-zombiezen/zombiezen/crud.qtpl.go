// Code generated by qtc from "crud.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:1
package zombiezen

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:1
func StreamGenerateCRUD(qw422016 *qt422016.Writer, t *GenerateCRUDTable) {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:1
	qw422016.N().S(`
// Code generated by "sqlc-gen-zombiezen". DO NOT EDIT.

package `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:4
	qw422016.E().S(t.PackageName.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:4
	qw422016.N().S(`

import (
    "fmt"
    "zombiezen.com/go/sqlite"
)

type `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:11
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:11
	qw422016.N().S(`Model struct {
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:12
	for _, f := range t.Fields {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:12
		qw422016.N().S(`        `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.E().S(f.Name.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.N().S(` `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.E().S(f.GoType.Original)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.N().S(` `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.N().S(`json:"`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.E().S(f.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.N().S(`"`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:13
		qw422016.N().S(`
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:14
	}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:14
	qw422016.N().S(`}

func Create`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:17
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:17
	qw422016.N().S(`(tx *sqlite.Conn, m *`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:17
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:17
	qw422016.N().S(`Model) error {
    stmt := tx.Prep(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:17
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:17
	qw422016.N().S(`
INSERT INTO `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:19
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:19
	qw422016.N().S(` (
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:20
	for _, f := range t.Fields {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:20
		qw422016.N().S(`        `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:21
		qw422016.E().S(f.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:21
		qw422016.N().S(`,
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:22
	}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:22
	qw422016.N().S(`) VALUES (
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:24
	for range t.Fields {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:24
		qw422016.N().S(`        ?,
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:26
	}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:26
	qw422016.N().S(`)
    `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:26
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:26
	qw422016.N().S(`)
    defer stmt.Reset()

    // Bind parameters
    `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:32
	streambindFields(qw422016, t)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:32
	qw422016.N().S(`

    if _, err := stmt.Step(); err != nil {
        return fmt.Errorf("failed to insert `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:35
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:35
	qw422016.N().S(`: %w", err)
    }

    return nil
}

func Read`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:41
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:41
	qw422016.N().S(`(tx *sqlite.Conn, id int64) (*`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:41
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:41
	qw422016.N().S(`Model, error) {
    stmt := tx.Prep(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:41
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:41
	qw422016.N().S(`
SELECT
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:44
	for _, f := range t.Fields {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:44
		qw422016.N().S(`        `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:45
		qw422016.E().S(f.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:45
		qw422016.N().S(`,
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:46
	}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:46
	qw422016.N().S(`FROM `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:47
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:47
	qw422016.N().S(`
WHERE id = ?
    `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:47
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:47
	qw422016.N().S(`)
    defer stmt.Reset()

    stmt.BindInt64(1, id)

    if hasRow, err := stmt.Step(); err != nil {
        return nil, fmt.Errorf("failed to read `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:55
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:55
	qw422016.N().S(`: %w", err)
    } else if !hasRow {
        return nil, nil
    }

    m := &`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:60
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:60
	qw422016.N().S(`Model{}
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:61
	for i, f := range t.Fields {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:62
		switch f.GoType.Original {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:63
		case "time.Time":
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:63
			qw422016.N().S(`                m.`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:64
			qw422016.E().S(f.Name.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:64
			qw422016.N().S(` = toolbelt.JulianDayToTime(stmt.Column`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:64
			qw422016.E().S(f.SQLType.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:64
			qw422016.N().S(`(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:64
			qw422016.N().D(i)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:64
			qw422016.N().S(`))
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:65
		case "time.Duration":
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:65
			qw422016.N().S(`                m.`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:66
			qw422016.E().S(f.Name.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:66
			qw422016.N().S(` = toolbelt.MillisecondsToDuration(stmt.Column`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:66
			qw422016.E().S(f.SQLType.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:66
			qw422016.N().S(`(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:66
			qw422016.N().D(i)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:66
			qw422016.N().S(`))
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:67
		default:
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:67
			qw422016.N().S(`                m.`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:68
			qw422016.E().S(f.Name.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:68
			qw422016.N().S(` = stmt.Column`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:68
			qw422016.E().S(f.SQLType.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:68
			qw422016.N().S(`(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:68
			qw422016.N().D(i)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:68
			qw422016.N().S(`)
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:69
		}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:70
	}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:70
	qw422016.N().S(`
    return m, nil
}

func Update`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:75
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:75
	qw422016.N().S(`(tx *sqlite.Conn, m *`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:75
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:75
	qw422016.N().S(`Model) error {
    stmt := tx.Prep(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:75
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:75
	qw422016.N().S(`
UPDATE `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:77
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:77
	qw422016.N().S(`
SET
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:79
	for i, f := range t.Fields {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:80
		if i > 0 {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:80
			qw422016.N().S(`        `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:81
			qw422016.E().S(f.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:81
			qw422016.N().S(` = ?`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:81
			qw422016.N().D(i + 1)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:81
			qw422016.N().S(`,
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:82
		}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:83
	}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:83
	qw422016.N().S(`WHERE id = ?1
    `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:83
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:83
	qw422016.N().S(`)
    defer stmt.Reset()

    // Bind parameters
    `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:89
	streambindFields(qw422016, t)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:89
	qw422016.N().S(`

    if _, err := stmt.Step(); err != nil {
        return fmt.Errorf("failed to update `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:92
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:92
	qw422016.N().S(`: %w", err)
    }

    return nil
}

func Delete`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:98
	qw422016.E().S(t.SingleName.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:98
	qw422016.N().S(`(tx *sqlite.Conn, id int64) error {
    stmt := tx.Prep(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:98
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:98
	qw422016.N().S(`
DELETE FROM `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:100
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:100
	qw422016.N().S(`
WHERE id = ?
    `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:100
	qw422016.N().S("`")
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:100
	qw422016.N().S(`)
    defer stmt.Reset()

    stmt.BindInt64(1, id)

    if _, err := stmt.Step(); err != nil {
        return fmt.Errorf("failed to delete `)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:108
	qw422016.E().S(t.Name.Lower)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:108
	qw422016.N().S(`: %w", err)
    }

    return nil
}

`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
}

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
func WriteGenerateCRUD(qq422016 qtio422016.Writer, t *GenerateCRUDTable) {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	qw422016 := qt422016.AcquireWriter(qq422016)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	StreamGenerateCRUD(qw422016, t)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	qt422016.ReleaseWriter(qw422016)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
}

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
func GenerateCRUD(t *GenerateCRUDTable) string {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	qb422016 := qt422016.AcquireByteBuffer()
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	WriteGenerateCRUD(qb422016, t)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	qs422016 := string(qb422016.B)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	qt422016.ReleaseByteBuffer(qb422016)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
	return qs422016
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:114
}

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:116
func streambindFields(qw422016 *qt422016.Writer, tbl *GenerateCRUDTable) {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:117
	for _, f := range tbl.Fields {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:118
		switch f.GoType.Original {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:119
		case "time.Time":
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:119
			qw422016.N().S(`            stmt.Bind`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:120
			qw422016.E().S(f.SQLType.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:120
			qw422016.N().S(`(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:120
			qw422016.N().D(f.Column)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:120
			qw422016.N().S(`, toolbelt.TimeToJulianDay(m.`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:120
			qw422016.E().S(f.Name.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:120
			qw422016.N().S(`))
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:121
		case "time.Duration":
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:121
			qw422016.N().S(`            stmt.Bind`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:122
			qw422016.E().S(f.SQLType.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:122
			qw422016.N().S(`(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:122
			qw422016.N().D(f.Column)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:122
			qw422016.N().S(`, toolbelt.DurationToMilliseconds(m.`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:122
			qw422016.E().S(f.Name.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:122
			qw422016.N().S(`))
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:123
		default:
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:123
			qw422016.N().S(`            stmt.Bind`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:124
			qw422016.E().S(f.SQLType.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:124
			qw422016.N().S(`(`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:124
			qw422016.N().D(f.Column)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:124
			qw422016.N().S(`, m.`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:124
			qw422016.E().S(f.Name.Pascal)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:124
			qw422016.N().S(`)
`)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:125
		}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:126
	}
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
}

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
func writebindFields(qq422016 qtio422016.Writer, tbl *GenerateCRUDTable) {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	qw422016 := qt422016.AcquireWriter(qq422016)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	streambindFields(qw422016, tbl)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	qt422016.ReleaseWriter(qw422016)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
}

//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
func bindFields(tbl *GenerateCRUDTable) string {
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	qb422016 := qt422016.AcquireByteBuffer()
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	writebindFields(qb422016, tbl)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	qs422016 := string(qb422016.B)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	qt422016.ReleaseByteBuffer(qb422016)
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
	return qs422016
//line sqlc-gen-zombiezen/zombiezen/crud.qtpl:127
}