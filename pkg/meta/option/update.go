package meta

type UpdateCollectionOption struct {
	UpdateOption
}

type UpdateOption struct {
	// 选取哪几列更新
	Select []string
}
