package article

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateArticle
	currUser, ok := ctx.Get("current_user")
	u := currUser.(*user.User)
	if !ok || currUser == nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	_, err := c.srv.User().Get(ctx, &user.User{ID: u.ID}, nil)
	if err != nil {
		zaplog.L().Error("用户不存在", zap.Uint("id", u.ID))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	a := &article.Article{
		Title:    input.Title,
		Cover:    input.Cover,
		Summary:  input.Summary,
		Weight:   input.Weight,
		SeriesID: input.SeriesID,
		AuthorID: u.ID,
		ArticleBody: article.ArticleBody{
			Content: input.ArticleBody,
		},
		CategoryID: input.CategoryID,
	}
	err = c.srv.Article().Create(
		ctx,
		a,
		nil,
	)
	if err != nil {
		zaplog.L().Error("创建文章失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	var tags []*article.ArticleTag
	for _, t := range input.Tags {
		tags = append(tags, &article.ArticleTag{
			ArticleID: a.ID,
			TagID:     t,
		})
	}
	errs := c.srv.Article().CreateArticleTagsCollection(ctx, tags, nil)
	if errs != nil {
		zaplog.L().Error("创建文章失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	core.WriteResponse(ctx, nil, gin.H{"id": a.ID})
}

// func (c *Controller) Create(ctx *gin.Context) {
// 	var input CreateArticle
// 	currUser, ok := ctx.Get("current_user")
// 	u := currUser.(*user.User)
// 	if !ok || currUser == nil {
// 		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
// 		return
// 	}
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
// 	_, err := c.srv.User().Get(ctx, &user.User{ID: u.ID}, nil)
// 	if err != nil {
// 		zaplog.L().Error("用户不存在", zap.Uint("id", u.ID))
// 		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
// 		return
// 	}
// 	err = c.srv.Article().Create(
// 		ctx,
// 		&article.Article{
// 			Title:    input.Title,
// 			Cover:    input.Cover,
// 			Summary:  input.Summary,
// 			Weight:   input.Weight,
// 			Series:   input.Series,
// 			AuthorID: u.ID,
// 			ArticleBody: article.ArticleBody{
// 				Content: input.ArticleBody,
// 			},
// 			Category: category.Category{
// 				ID:           input.Category.ID,
// 				CategoryName: input.Category.CategoryName,
// 				Summary:      input.Category.Summary,
// 				Type:         category.CATE_TYPE_ARTICLE,
// 			},
// 			Tags: tags,
// 		},
// 		nil,
// 	)
// 	if err != nil {
// 		zaplog.L().Error("创建文章失败", zap.Error(err))
// 		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
// 		return
// 	}
// 	core.WriteResponse(ctx, nil, nil)
// }
