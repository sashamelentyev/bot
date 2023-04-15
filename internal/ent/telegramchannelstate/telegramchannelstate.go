// Code generated by ent, DO NOT EDIT.

package telegramchannelstate

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the telegramchannelstate type in the database.
	Label = "telegram_channel_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChannelID holds the string denoting the channel_id field in the database.
	FieldChannelID = "channel_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPts holds the string denoting the pts field in the database.
	FieldPts = "pts"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the telegramchannelstate in the database.
	Table = "telegram_channel_states"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "telegram_channel_states"
	// UserInverseTable is the table name for the TelegramUserState entity.
	// It exists in this package in order to avoid circular dependency with the "telegramuserstate" package.
	UserInverseTable = "telegram_user_states"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for telegramchannelstate fields.
var Columns = []string{
	FieldID,
	FieldChannelID,
	FieldUserID,
	FieldPts,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPts holds the default value on creation for the "pts" field.
	DefaultPts int
)

// Order defines the ordering method for the TelegramChannelState queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChannelID orders the results by the channel_id field.
func ByChannelID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldChannelID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByPts orders the results by the pts field.
func ByPts(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldPts, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
