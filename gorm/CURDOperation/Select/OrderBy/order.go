package order

import (
	. "gorm/Model"

	"gorm.io/gorm/clause"
)
func Order() {
	// DB.Order("FIELD(id, 2, 3, 1)").Find()

  i  := []int{2, 3, 1}
	DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL: "FIELD(id, ?)",
			Vars: []any{i},
			WithoutParentheses: true,
		},
	})
}