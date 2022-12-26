package store

type Transaction interface {
	TransactionBegin() Factory
	TransactionRollback()
	TransactionCommit()
}
