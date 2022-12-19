package meta

type DeleteOption struct {
	Select []string
}

type DeleteCollectionOption struct {
	DeleteOption
	Order string
	Limit int64
}
