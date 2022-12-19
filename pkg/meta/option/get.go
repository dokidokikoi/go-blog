package meta

const (
	INNER_JOIN  = "INNER JOIN"
	LEFT_JOIN   = "LEFT JOIN"
	RIGHRT_JOIN = "RIGHT JOIN"
)

type GetOption struct {
	Include []string
	Preload []string
	Select  []string
	Join    []*Join
}

type Join struct {
	Method             string
	Table              string
	JoinTable          string
	TableField         string
	JoinTableField     string
	JoinTableCondition []Condition
}
