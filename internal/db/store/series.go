package store

import (
	"github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/series"
)

type Series interface {
	base.BasicCURD[series.Series]
}
