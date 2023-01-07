package article

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/model/category"
	myErrors "go-blog/internal/errors"
	"go-blog/pkg/log/zap"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Update(ctx *gin.Context) {
	var input ArticleUpdate
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	var err error
	targetArticleBody := &article.ArticleBody{
		ID:      input.ArticleBody.ID,
		Content: input.ArticleBody.Content,
	}
	if err = c.srv.Article().UpdateBody(ctx, targetArticleBody); err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	input.Category.Type = category.ARTICLE
	ok, err := c.srv.Category().IsExist(ctx, &input.Category)
	if err != nil || !ok {
		c.srv.Category().Create(ctx, &input.Category)
	}

	ok, err = c.srv.ArticleSeries().IsExist(ctx, &input.Series)
	if err != nil || !ok {
		c.srv.ArticleSeries().Create(ctx, &input.Series)
	}

	targetArtcle := &article.Article{
		Title:         input.Title,
		Summary:       input.Summary,
		Cover:         input.Cover,
		CategoryID:    input.Category.ID,
		SeriesID:      input.Series.ID,
		ArticleBodyID: targetArticleBody.ID,
	}
	targetArtcle.ID = input.ID
	if err = c.srv.Article().UpdateArticle(ctx, targetArtcle); err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	errs := c.srv.ArticleTag().UpdateArticleTag(ctx, targetArtcle)
	for _, err := range errs {
		if err != nil {
			zap.Suger().Errorf("create article_tag err: %v", err)
		}
	}

	core.WriteResponse(ctx, myErrors.Success("修改成功"), nil)
}
