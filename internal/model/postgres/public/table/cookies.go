//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Cookies = newCookiesTable("public", "cookies", "")

type cookiesTable struct {
	postgres.Table

	//Columns
	UserID    postgres.ColumnInteger
	Value     postgres.ColumnString
	ExpiresAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CookiesTable struct {
	cookiesTable

	EXCLUDED cookiesTable
}

// AS creates new CookiesTable with assigned alias
func (a CookiesTable) AS(alias string) *CookiesTable {
	return newCookiesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CookiesTable with assigned schema name
func (a CookiesTable) FromSchema(schemaName string) *CookiesTable {
	return newCookiesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CookiesTable with assigned table prefix
func (a CookiesTable) WithPrefix(prefix string) *CookiesTable {
	return newCookiesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CookiesTable with assigned table suffix
func (a CookiesTable) WithSuffix(suffix string) *CookiesTable {
	return newCookiesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCookiesTable(schemaName, tableName, alias string) *CookiesTable {
	return &CookiesTable{
		cookiesTable: newCookiesTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newCookiesTableImpl("", "excluded", ""),
	}
}

func newCookiesTableImpl(schemaName, tableName, alias string) cookiesTable {
	var (
		UserIDColumn    = postgres.IntegerColumn("user_id")
		ValueColumn     = postgres.StringColumn("value")
		ExpiresAtColumn = postgres.TimestampzColumn("expires_at")
		allColumns      = postgres.ColumnList{UserIDColumn, ValueColumn, ExpiresAtColumn}
		mutableColumns  = postgres.ColumnList{UserIDColumn, ValueColumn, ExpiresAtColumn}
	)

	return cookiesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:    UserIDColumn,
		Value:     ValueColumn,
		ExpiresAt: ExpiresAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
