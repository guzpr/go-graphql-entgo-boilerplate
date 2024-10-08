// Code generated by ent, DO NOT EDIT.

package googleauth

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sekalahita/epirus/internal/ent/gen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldDeletedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldUserID, v))
}

// GoogleID applies equality check predicate on the "google_id" field. It's identical to GoogleIDEQ.
func GoogleID(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldGoogleID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNotNull(FieldDeletedAt))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldContainsFold(FieldUserID, v))
}

// GoogleIDEQ applies the EQ predicate on the "google_id" field.
func GoogleIDEQ(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEQ(FieldGoogleID, v))
}

// GoogleIDNEQ applies the NEQ predicate on the "google_id" field.
func GoogleIDNEQ(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNEQ(FieldGoogleID, v))
}

// GoogleIDIn applies the In predicate on the "google_id" field.
func GoogleIDIn(vs ...string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldIn(FieldGoogleID, vs...))
}

// GoogleIDNotIn applies the NotIn predicate on the "google_id" field.
func GoogleIDNotIn(vs ...string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldNotIn(FieldGoogleID, vs...))
}

// GoogleIDGT applies the GT predicate on the "google_id" field.
func GoogleIDGT(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGT(FieldGoogleID, v))
}

// GoogleIDGTE applies the GTE predicate on the "google_id" field.
func GoogleIDGTE(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldGTE(FieldGoogleID, v))
}

// GoogleIDLT applies the LT predicate on the "google_id" field.
func GoogleIDLT(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLT(FieldGoogleID, v))
}

// GoogleIDLTE applies the LTE predicate on the "google_id" field.
func GoogleIDLTE(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldLTE(FieldGoogleID, v))
}

// GoogleIDContains applies the Contains predicate on the "google_id" field.
func GoogleIDContains(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldContains(FieldGoogleID, v))
}

// GoogleIDHasPrefix applies the HasPrefix predicate on the "google_id" field.
func GoogleIDHasPrefix(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldHasPrefix(FieldGoogleID, v))
}

// GoogleIDHasSuffix applies the HasSuffix predicate on the "google_id" field.
func GoogleIDHasSuffix(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldHasSuffix(FieldGoogleID, v))
}

// GoogleIDEqualFold applies the EqualFold predicate on the "google_id" field.
func GoogleIDEqualFold(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldEqualFold(FieldGoogleID, v))
}

// GoogleIDContainsFold applies the ContainsFold predicate on the "google_id" field.
func GoogleIDContainsFold(v string) predicate.GoogleAuth {
	return predicate.GoogleAuth(sql.FieldContainsFold(FieldGoogleID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.GoogleAuth {
	return predicate.GoogleAuth(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.GoogleAuth {
	return predicate.GoogleAuth(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoogleAuth) predicate.GoogleAuth {
	return predicate.GoogleAuth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoogleAuth) predicate.GoogleAuth {
	return predicate.GoogleAuth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GoogleAuth) predicate.GoogleAuth {
	return predicate.GoogleAuth(func(s *sql.Selector) {
		p(s.Not())
	})
}
