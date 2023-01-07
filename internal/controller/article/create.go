package article

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/model/category"
	myErrors "go-blog/internal/errors"
	"go-blog/pkg/log/zap"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Create(ctx *gin.Context) {
	var a ArticleCreate
	if err := ctx.ShouldBindJSON(&a); err != nil {
		zap.Suger().Errorf("%s err: %v", myErrors.ErrValidation.Error(), err)
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	var err error
	targetArticleBody := &article.ArticleBody{
		Content: a.ArticleBody.Content,
	}
	if err = c.srv.Article().CreateBody(ctx, targetArticleBody); err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	tartgetCategory := &category.Category{
		ID:           a.Category.ID,
		CategoryName: a.Category.CategoryName,
		Summary:      a.Summary,
		Type:         category.ARTICLE,
	}
	ok, err := c.srv.Category().IsExist(ctx, tartgetCategory)
	if err != nil || !ok {
		c.srv.Category().Create(ctx, tartgetCategory)
	}

	targetSeries := &article.Series{
		ID:         a.Series.ID,
		SeriesName: a.Series.SeriesName,
		Summary:    a.Series.Summary,
	}
	ok, err = c.srv.ArticleSeries().IsExist(ctx, targetSeries)
	if err != nil || !ok {
		c.srv.ArticleSeries().Create(ctx, targetSeries)
	}

	targetArtcle := &article.Article{
		Title:         a.Title,
		Summary:       a.Summary,
		Cover:         a.Cover,
		CategoryID:    tartgetCategory.ID,
		SeriesID:      targetSeries.ID,
		ArticleBodyID: targetArticleBody.ID,
		AuthorID:      2, // TODO 暂时写死
	}
	if err = c.srv.Article().CreateArticle(ctx, targetArtcle); err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	var targetArticleTag []*article.ArticleTag
	for _, tag := range a.Tags {
		targetTag := &article.Tag{
			ID:      tag.ID,
			TagName: tag.TagName,
		}
		ok, err = c.srv.ArticleTag().IsExist(ctx, targetTag)
		if err != nil || !ok {
			c.srv.ArticleTag().Create(ctx, targetTag)
		}
		targetArticleTag = append(targetArticleTag, &article.ArticleTag{ArticleID: targetArtcle.ID, TagId: targetTag.ID})
	}

	errs := c.srv.ArticleTag().CreateArticleTagCollection(ctx, targetArticleTag)
	for _, err = range errs {
		if err != nil {
			zap.Suger().Errorf("create article_tag err: %v", err)
		}
	}

	core.WriteResponse(ctx, myErrors.Success("创建成功"), nil)
}
