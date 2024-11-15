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

var UserMeetups = newUserMeetupsTable("public", "user_meetups", "")

type userMeetupsTable struct {
	postgres.Table

	// Columns
	UserID   postgres.ColumnInteger
	MeetupID postgres.ColumnInteger
	State    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserMeetupsTable struct {
	userMeetupsTable

	EXCLUDED userMeetupsTable
}

// AS creates new UserMeetupsTable with assigned alias
func (a UserMeetupsTable) AS(alias string) *UserMeetupsTable {
	return newUserMeetupsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserMeetupsTable with assigned schema name
func (a UserMeetupsTable) FromSchema(schemaName string) *UserMeetupsTable {
	return newUserMeetupsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserMeetupsTable with assigned table prefix
func (a UserMeetupsTable) WithPrefix(prefix string) *UserMeetupsTable {
	return newUserMeetupsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserMeetupsTable with assigned table suffix
func (a UserMeetupsTable) WithSuffix(suffix string) *UserMeetupsTable {
	return newUserMeetupsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserMeetupsTable(schemaName, tableName, alias string) *UserMeetupsTable {
	return &UserMeetupsTable{
		userMeetupsTable: newUserMeetupsTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newUserMeetupsTableImpl("", "excluded", ""),
	}
}

func newUserMeetupsTableImpl(schemaName, tableName, alias string) userMeetupsTable {
	var (
		UserIDColumn   = postgres.IntegerColumn("user_id")
		MeetupIDColumn = postgres.IntegerColumn("meetup_id")
		StateColumn    = postgres.StringColumn("state")
		allColumns     = postgres.ColumnList{UserIDColumn, MeetupIDColumn, StateColumn}
		mutableColumns = postgres.ColumnList{StateColumn}
	)

	return userMeetupsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:   UserIDColumn,
		MeetupID: MeetupIDColumn,
		State:    StateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}