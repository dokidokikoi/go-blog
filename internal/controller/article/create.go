package article

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	myErrors "go-blog/internal/errors"
	"go-blog/pkg/log/zap"

	"github.com/gin-gonic/gin"
)

type ArticleCreate struct {
	Title       string           `json:"title" binding:"required,min=2"`
	Summary     string           `json:"summary"`
	Cover       string           `json:"cover" binding:"required"`
	Category    article.Category `json:"category" binding:"required"`
	Tags        []article.Tag    `json:"tags"`
	Series      article.Series   `json:"series"`
	ArticleBody ArticleBody      `json:"article_body"`
}

type ArticleBody struct {
	Content string `json:"content" binding:"required"`
}

// type ArticleTag struct {
// 	Id      uint   `json:"id"`
// 	TagName string `json:"tag_name"`
// }

// type ArticleCategoty struct {
// 	Id           uint   `json:"id"`
// 	CategoryName string `json:"category_name"`
// 	Summary      string `json:"summary"`
// }

// type ArticleSeries struct {
// 	Id         uint   `json:"id"`
// 	SeriesName string `json:"series_name"`
// 	Summary    string `json:"summary"`
// }

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

	tartgetCategory := &article.Category{
		ID:           a.Category.ID,
		CategoryName: a.Category.CategoryName,
		Summary:      a.Summary,
	}
	ok, err := c.srv.ArticleCategory().IsExist(ctx, tartgetCategory)
	if err != nil || !ok {
		c.srv.ArticleCategory().Create(ctx, tartgetCategory)
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
		targetArticleTag = append(targetArticleTag, &article.ArticleTag{ArticleId: targetArtcle.ID, TagId: targetTag.ID})
	}

	errs := c.srv.ArticleTag().CreateArticleTagCollection(ctx, targetArticleTag)
	for _, err = range errs {
		if err != nil {
			zap.Suger().Errorf("create article_tag err: %v", err)
		}
	}

	core.WriteResponse(ctx, myErrors.Success("创建成功"), nil)
}
