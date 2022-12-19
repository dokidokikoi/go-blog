package meta

const (
	EQUAL    = "="
	NOTEQUAL = "<>"
	GTE      = ">="
	GT       = ">"
	LTE      = "<="
	LT       = "<"
	IN       = "IN"
	LIKE     = "LIKE"
)

type Condition struct {
	Field    string
	Operator string
	Value    any
}

// WhereNode 查询条件组织
// @param Conditions 单个条件 内部是'或'关系
// @param Nodes 条件组 内部是'或'关系
// @param Next       为'与'关系
type WhereNode struct {
	Conditions []*Condition
	Nodes      []*WhereNode
	Next       *WhereNode
}

func (c *Condition) Validate() error {
	return nil
}
