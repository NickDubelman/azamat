package azamat

import (
	sq "github.com/Masterminds/squirrel"
)

type Table struct {
	Name    string
	Columns []string
}

// Builders

// Select returns a buildable Select query that is bound to a specific table name. If
// no columns are provided, it gets all columns specified by the table
func (t Table) Select(columns ...string) sq.SelectBuilder {
	actualColumns := columns
	if len(columns) == 0 {
		actualColumns = t.Columns
	}

	return sq.Select(actualColumns...).From(t.Name)
}

// Insert returns a buildable Insert statement that is bound to a specific table name
func (t Table) Insert() sq.InsertBuilder {
	return sq.Insert(t.Name)
}

// Update returns a buildable Update statement that is bound to a specific table name
func (t Table) Update() sq.UpdateBuilder {
	return sq.Update(t.Name)
}

// Delete returns a buildable Delete statement that is bound to a specific table name
func (t Table) Delete() sq.DeleteBuilder {
	return sq.Delete(t.Name)
}
