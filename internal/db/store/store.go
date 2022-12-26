package store

import "fmt"

type Factory interface {
	Transaction() Transaction
	Article() Article
	ArticleBody() ArticleBody
	ArticleCategory() ArticleCategory
	ArticleTag() ArticleTag
	ArticleSeries() ArticleSeries
	ArticleArticleTag() ArticleArticleTag
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
