package comment

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/model/comment"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateComment
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	if input.PID != 0 {
		_, err := c.srv.Comment().Get(ctx, &comment.Comment{ID: input.PID}, nil)
		if err != nil {
			zaplog.L().Error("创建评论失败", zap.Error(err))
			core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
			return
		}
	}
	if input.ArticleID != 0 {
		_, err := c.srv.Article().Get(ctx, &article.Article{ID: input.ArticleID}, nil)
		if err != nil {
			zaplog.L().Error("创建评论失败", zap.Error(err))
			core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
			return
		}
	}

	err := c.srv.Comment().Create(
		ctx,
		&comment.Comment{
			PID:        input.PID,
			ArticleID:  input.ArticleID,
			Content:    input.Content,
			Nickname:   input.Nickname,
			ToNickname: input.ToNickname,
		},
		nil,
	)
	if err != nil {
		zaplog.L().Error("创建评论失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	go func() {
		cnt := 0
		for cnt < 10 {
			s, err := c.srv.Article().Get(ctx, &article.Article{ID: input.ArticleID}, nil)
			if err != nil {
				return
			}
			node := &meta.WhereNode{
				Conditions: []*meta.Condition{
					{
						Field:    "comment_counts",
						Operator: meta.EQUAL,
						Value:    s.CommentCounts,
					},
				},
			}
			err = c.srv.Article().UpdateByWhereNode(ctx, &article.Article{ID: s.ID, CommentCounts: s.CommentCounts + 1}, node, nil)
			if err == nil {
				return
			}
			cnt++
		}
	}()
	core.WriteResponse(ctx, nil, nil)
}
