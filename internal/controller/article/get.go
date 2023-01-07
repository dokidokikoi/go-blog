package article

import (
	"go-blog/internal/code"
	"go-blog/internal/core"
	"go-blog/internal/db/model/category"
	myErrors "go-blog/internal/errors"
	meta "go-blog/pkg/meta/option"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	option := meta.GetOption{}
	option.Preload = append(
		option.Preload,
		[]interface{}{"Category", "type = ?", category.ARTICLE},
		[]interface{}{"Tags"},
		[]interface{}{"Series"},
		[]interface{}{"Author"},
		[]interface{}{"ArticleBody"})
	res, err := c.srv.Article().GetArticleByID(ctx, uint(id), &option)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ClientFailed("文章未找到", code.ErrArticleNotFound), nil)
		return
	}

	c.srv.Article().UpdateViewCntByID(ctx, uint(id))

	core.WriteResponse(ctx, nil, res)
}
