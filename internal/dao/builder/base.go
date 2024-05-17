package builder

import (
	"gorm.io/gorm/clause"
)

type BuilderT interface {
	Unscoped() bool
	Build() []clause.Expression
}

type baseBuilder struct {
	exprs    []clause.Expression
	unscoped bool
}

func newBaseBuilder() baseBuilder {
	return baseBuilder{
		exprs: []clause.Expression{},
	}
}

func (r *baseBuilder) withEq(column string, value interface{}) {
	r.exprs = append(r.exprs, clause.Eq{Column: column, Value: value})
}

func (r *baseBuilder) withNeq(column string, value interface{}) {
	r.exprs = append(r.exprs, clause.Neq{Column: column, Value: value})
}

func (r *baseBuilder) withGt(column string, value interface{}) {
	r.exprs = append(r.exprs, clause.Gt{Column: column, Value: value})
}

func (r *baseBuilder) withLt(column string, value interface{}) {
	r.exprs = append(r.exprs, clause.Lt{Column: column, Value: value})
}

func (r *baseBuilder) withLike(column string, value interface{}) {
	r.exprs = append(r.exprs, clause.Like{Column: column, Value: value})
}

func (r *baseBuilder) withStringIn(column string, value []string) {
	r.exprs = append(r.exprs, clause.IN{Column: column, Values: r.convertStringSlice(value)})
}

func (r *baseBuilder) withIntIn(column string, value []int) {
	r.exprs = append(r.exprs, clause.IN{Column: column, Values: r.convertIntSlice(value)})
}

func (r *baseBuilder) withInt64In(column string, value []int64) {
	r.exprs = append(r.exprs, clause.IN{Column: column, Values: r.convertInt64Slice(value)})
}

func (r *baseBuilder) withSQL(sql string, vars ...interface{}) {
	i := make([]interface{}, 0, len(vars))
	for _, v := range vars {
		i = append(i, v)
	}
	r.exprs = append(r.exprs, clause.Expr{SQL: sql, Vars: i})
}

func (r *baseBuilder) withOrder(order string) {
	expr := clause.Expr{SQL: order}
	r.exprs = append(r.exprs, clause.OrderBy{Expression: expr})
}

func (r *baseBuilder) withSelect(fields string) {
	expr := clause.Expr{SQL: fields}
	r.exprs = append(r.exprs, clause.Select{Expression: expr})
}

func (r *baseBuilder) Build() []clause.Expression {
	return r.exprs
}

func (r *baseBuilder) Unscoped() bool {
	return r.unscoped
}

func (r *baseBuilder) convertIntSlice(s []int) []interface{} {
	res := make([]interface{}, 0)
	m := map[int]bool{}
	for _, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		res = append(res, v)
		m[v] = true
	}
	return res
}

func (r *baseBuilder) convertInt64Slice(s []int64) []interface{} {
	res := make([]interface{}, 0)
	m := map[int64]bool{}
	for _, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		res = append(res, v)
		m[v] = true
	}
	return res
}

func (r *baseBuilder) convertStringSlice(s []string) []interface{} {
	res := make([]interface{}, 0)
	m := map[string]bool{}
	for _, v := range s {
		if _, ok := m[v]; ok {
			continue
		}
		res = append(res, v)
		m[v] = true
	}
	return res
}
