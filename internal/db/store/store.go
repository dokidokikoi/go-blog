package store

import "fmt"

type Factory interface {
	Transaction() Transaction
	Article() Article
	ArticleBody() ArticleBody
	Category() Category
	Tag() Tag
	ArticleSeries() Series
	ArticleTag() ArticleTag

	Users() Users
	Roles() Roles
}

var storePointer Factory

func GetStoreFactory() (Factory, error) {
	if storePointer == nil {
		return nil, fmt.Errorf("数据层未初始化")
	}
	return storePointer, nil
}

func SetStoreFactory(factory Factory) {
	storePointer = factory
}
