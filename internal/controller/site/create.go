package site

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/site"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateSite
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	s := &site.Site{
		SiteName:   input.SiteName,
		Summary:    input.Summary,
		Logo:       input.Logo,
		Url:        input.Url,
		CategoryID: input.CategoryID,
	}
	err := c.srv.Site().Create(
		ctx,
		s,
		nil,
	)
	if err != nil {
		zaplog.L().Error("创建网站失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	var tags []*site.SiteTag
	for _, t := range input.Tags {
		tags = append(tags, &site.SiteTag{
			SiteID: s.ID,
			TagID:  t,
		})
	}
	errs := c.srv.Site().CreateSiteTagCollection(ctx, tags, nil)
	if errs != nil {
		zaplog.L().Error("添加网站标签失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}

// func (c *Controller) Create(ctx *gin.Context) {
// 	var input CreateSite
// 	if ctx.ShouldBindJSON(&input) != nil {
// 		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
// 		return
// 	}

// 	var tags []tag.Tag
// 	for _, t := range input.Tags {
// 		tags = append(tags, tag.Tag{
// 			ID:      t.ID,
// 			TagName: t.TagName,
// 			Type:    tag.TAG_TYPE_SITE,
// 		})
// 	}

// 	err := c.srv.Site().Create(
// 		ctx,
// 		&site.Site{
// 			SiteName: input.SiteName,
// 			Summary:  input.Summary,
// 			Logo:     input.Logo,
// 			Url:      input.Url,
// 			Category: category.Category{
// 				ID:           input.Category.ID,
// 				CategoryName: input.Category.CategoryName,
// 				Summary:      input.Category.Summary,
// 				Type:         category.CATE_TYPE_SITE,
// 			},
// 			Tags: tags,
// 		},
// 		nil,
// 	)
// 	if err != nil {
// 		zaplog.L().Error("创建网站失败", zap.Error(err))
// 		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
// 		return
// 	}
// 	core.WriteResponse(ctx, nil, nil)
// }
