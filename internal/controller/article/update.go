package article

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Update(ctx *gin.Context) {
	var input UpdateArticle
	if err := ctx.ShouldBindJSON(&input); err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	var tags []*article.ArticleTag
	for _, t := range input.Tags {
		tags = append(tags, &article.ArticleTag{
			ArticleID: input.ID,
			TagID:     t,
		})
	}

	_, err := c.srv.Article().Get(ctx, &article.Article{ID: input.ID}, nil)
	if err != nil {
		zaplog.L().Error("文章不存在", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}

	err = c.srv.Article().DeleteArticleAllTags(ctx, input.ID)
	if err != nil {
		zaplog.L().Error("删除文章标签失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabasDel, nil)
		return
	}

	errs := c.srv.Article().CreateArticleTagsCollection(ctx, tags, nil)
	if errs != nil {
		zaplog.L().Error("创建文章标签失败", zap.Errors("errs", errs))
	}

	err = c.srv.Article().Update(ctx, &article.Article{
		ID:       input.ID,
		Title:    input.Title,
		Cover:    input.Cover,
		Summary:  input.Summary,
		Weight:   input.Weight,
		SeriesID: input.SeriesID,
		ArticleBody: article.ArticleBody{
			ID:      input.ArticleBodyID,
			Content: input.ArticleBody,
		},
		CategoryID: input.CategoryID,
	}, nil)
	if err != nil {
		zaplog.L().Error("更新文章消息失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiNoUpdateRows, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}

// func (c *Controller) Update(ctx *gin.Context) {
// 	var input UpdateArticle
// 	if ctx.ShouldBindJSON(&input) != nil {
// 		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
// 		return
// 	}

// 	var tags []tag.Tag
// 	for _, t := range input.Tags {
// 		tags = append(tags, tag.Tag{
// 			ID:      t.ID,
// 			TagName: t.TagName,
// 			Type:    tag.TAG_TYPE_ARTICLE,
// 		})
// 	}

// 	err := c.srv.Article().DeleteArticleAllTags(ctx, input.ID)
// 	if err != nil {
// 		zaplog.L().Error("删除文章标签失败", zap.Error(err))
// 		core.WriteResponse(ctx, myErrors.ApiErrDatabasDel, nil)
// 		return
// 	}

// 	a, err := c.srv.Article().Get(ctx, &article.Article{ID: input.ID}, nil)
// 	if err != nil {
// 		zaplog.L().Error("文章不存在", zap.Error(err))
// 		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
// 		return
// 	}

// 	err = c.srv.Article().Update(ctx, &article.Article{
// 		ID:      input.ID,
// 		Title:   input.Title,
// 		Cover:   input.Cover,
// 		Summary: input.Summary,
// 		Weight:  input.Weight,
// 		Series:  input.Series,
// 		ArticleBody: article.ArticleBody{
// 			ID:      a.ArticleBodyID,
// 			Content: input.ArticleBody,
// 		},
// 		Category: category.Category{
// 			ID:           input.Category.ID,
// 			CategoryName: input.Category.CategoryName,
// 			Summary:      input.Category.Summary,
// 			Type:         category.CATE_TYPE_ARTICLE,
// 		},
// 		Tags: tags,
// 	}, nil)
// 	if err != nil {
// 		zaplog.L().Error("更新文章消息失败", zap.Error(err))
// 		core.WriteResponse(ctx, myErrors.ApiNoUpdateRows, nil)
// 		return
// 	}
// 	core.WriteResponse(ctx, nil, nil)
// }
