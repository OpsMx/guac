// Code generated by ent, DO NOT EDIT.

package scorecard

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the scorecard type in the database.
	Label = "scorecard"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChecks holds the string denoting the checks field in the database.
	FieldChecks = "checks"
	// FieldAggregateScore holds the string denoting the aggregate_score field in the database.
	FieldAggregateScore = "aggregate_score"
	// FieldTimeScanned holds the string denoting the time_scanned field in the database.
	FieldTimeScanned = "time_scanned"
	// FieldScorecardVersion holds the string denoting the scorecard_version field in the database.
	FieldScorecardVersion = "scorecard_version"
	// FieldScorecardCommit holds the string denoting the scorecard_commit field in the database.
	FieldScorecardCommit = "scorecard_commit"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldCollector holds the string denoting the collector field in the database.
	FieldCollector = "collector"
	// EdgeCertifications holds the string denoting the certifications edge name in mutations.
	EdgeCertifications = "certifications"
	// Table holds the table name of the scorecard in the database.
	Table = "scorecards"
	// CertificationsTable is the table that holds the certifications relation/edge.
	CertificationsTable = "certify_scorecards"
	// CertificationsInverseTable is the table name for the CertifyScorecard entity.
	// It exists in this package in order to avoid circular dependency with the "certifyscorecard" package.
	CertificationsInverseTable = "certify_scorecards"
	// CertificationsColumn is the table column denoting the certifications relation/edge.
	CertificationsColumn = "scorecard_id"
)

// Columns holds all SQL columns for scorecard fields.
var Columns = []string{
	FieldID,
	FieldChecks,
	FieldAggregateScore,
	FieldTimeScanned,
	FieldScorecardVersion,
	FieldScorecardCommit,
	FieldOrigin,
	FieldCollector,
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
	// DefaultAggregateScore holds the default value on creation for the "aggregate_score" field.
	DefaultAggregateScore float64
	// DefaultTimeScanned holds the default value on creation for the "time_scanned" field.
	DefaultTimeScanned func() time.Time
)

// OrderOption defines the ordering options for the Scorecard queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAggregateScore orders the results by the aggregate_score field.
func ByAggregateScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAggregateScore, opts...).ToFunc()
}

// ByTimeScanned orders the results by the time_scanned field.
func ByTimeScanned(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimeScanned, opts...).ToFunc()
}

// ByScorecardVersion orders the results by the scorecard_version field.
func ByScorecardVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScorecardVersion, opts...).ToFunc()
}

// ByScorecardCommit orders the results by the scorecard_commit field.
func ByScorecardCommit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScorecardCommit, opts...).ToFunc()
}

// ByOrigin orders the results by the origin field.
func ByOrigin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrigin, opts...).ToFunc()
}

// ByCollector orders the results by the collector field.
func ByCollector(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollector, opts...).ToFunc()
}

// ByCertificationsCount orders the results by certifications count.
func ByCertificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCertificationsStep(), opts...)
	}
}

// ByCertifications orders the results by certifications terms.
func ByCertifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCertificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCertificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CertificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CertificationsTable, CertificationsColumn),
	)
}
