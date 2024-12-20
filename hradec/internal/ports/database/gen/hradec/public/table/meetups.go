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

var Meetups = newMeetupsTable("public", "meetups", "")

type meetupsTable struct {
	postgres.Table

	// Columns
	ID      postgres.ColumnInteger
	PlaceID postgres.ColumnString
	Time    postgres.ColumnTimestamp
	Name    postgres.ColumnString
	UserID  postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type MeetupsTable struct {
	meetupsTable

	EXCLUDED meetupsTable
}

// AS creates new MeetupsTable with assigned alias
func (a MeetupsTable) AS(alias string) *MeetupsTable {
	return newMeetupsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MeetupsTable with assigned schema name
func (a MeetupsTable) FromSchema(schemaName string) *MeetupsTable {
	return newMeetupsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MeetupsTable with assigned table prefix
func (a MeetupsTable) WithPrefix(prefix string) *MeetupsTable {
	return newMeetupsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MeetupsTable with assigned table suffix
func (a MeetupsTable) WithSuffix(suffix string) *MeetupsTable {
	return newMeetupsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMeetupsTable(schemaName, tableName, alias string) *MeetupsTable {
	return &MeetupsTable{
		meetupsTable: newMeetupsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newMeetupsTableImpl("", "excluded", ""),
	}
}

func newMeetupsTableImpl(schemaName, tableName, alias string) meetupsTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		PlaceIDColumn  = postgres.StringColumn("place_id")
		TimeColumn     = postgres.TimestampColumn("time")
		NameColumn     = postgres.StringColumn("name")
		UserIDColumn   = postgres.IntegerColumn("user_id")
		allColumns     = postgres.ColumnList{IDColumn, PlaceIDColumn, TimeColumn, NameColumn, UserIDColumn}
		mutableColumns = postgres.ColumnList{PlaceIDColumn, TimeColumn, NameColumn, UserIDColumn}
	)

	return meetupsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:      IDColumn,
		PlaceID: PlaceIDColumn,
		Time:    TimeColumn,
		Name:    NameColumn,
		UserID:  UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
