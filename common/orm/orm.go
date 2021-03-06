package orm

import (
	"database/sql"

	"main/common/kv"
)

type field struct {
	fieldType  string
	fieldName  string
	fieldValue interface{}
}

var ormDb *sql.DB = nil

func SetDb(db *sql.DB) {
	ormDb = db
}

type Orm interface {
	SetDb(db *sql.DB)
	SetTableName(tbName string)

	Create(item kv.Element) (int64, error)
	First(where, order kv.Element) (kv.Element, error)
	List(where, order kv.Element) (kv.Element, error)
	Page(where, order kv.Element, page, size int) (kv.Element, error)
	Remove(where kv.Element) (int64, error)
	Update(where kv.Element, data kv.Element) (int64, error)
}

func NewOrm(tbName string) Orm {
	return &orm{db: ormDb, tableName: tbName}
}

type orm struct {
	db        *sql.DB
	tableName string
}

func (o *orm) SetDb(db *sql.DB) {
	o.db = db
}
func (o *orm) SetTableName(tbName string) {
	o.tableName = tbName
}

func makeFields(rows *sql.Rows) ([]*field, []interface{}, error) {
	if cts, err := rows.ColumnTypes(); err == nil {
		fields := make([]*field, 0)
		pointers := make([]interface{}, 0)
		for _, ct := range cts {
			fd := &field{fieldType: ct.DatabaseTypeName(), fieldName: ct.Name()}
			switch ct.DatabaseTypeName() {
			case "INT", "BIGINT":
				{
					var v int64
					fd.fieldValue = &v
					pointers = append(pointers, &v)
				}
			case "VARCHAR", "TEXT", "NVARCHAR":
				{
					var v string
					fd.fieldValue = &v
					pointers = append(pointers, &v)
				}
			case "DECIMAL":
				{
					var v float64
					fd.fieldValue = &v
					pointers = append(pointers, &v)
				}
			case "BOOL":
				{
					var v bool
					fd.fieldValue = &v
					pointers = append(pointers, &v)
				}
			}
			fields = append(fields, fd)
		}
		return fields, pointers, nil
	} else {
		return nil, nil, err
	}
}
