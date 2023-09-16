package service

import (
	"context"
	"go-blog/internal/db/model/site"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type SiteSrv interface {
	Create(ctx context.Context, example *site.Site, option *meta.CreateOption) error
	Get(ctx context.Context, example *site.Site, option *meta.GetOption) (*site.Site, error)
	Update(ctx context.Context, example *site.Site, option *meta.UpdateOption) error
	Del(ctx context.Context, example *site.Site, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*site.Site, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *site.Site, option *meta.ListOption) ([]*site.Site, int64, error)
	ListByWhereNode(ctx context.Context, example *site.Site, node *meta.WhereNode, option *meta.ListOption) ([]*site.Site, int64, error)

	DeleteSiteAllTags(ctx context.Context, siteID uint) error
}

type siteSrv struct {
	store store.Factory
}

func (ss *siteSrv) Create(ctx context.Context, example *site.Site, option *meta.CreateOption) error {
	return ss.store.Sites().Create(ctx, example, option)
}

func (ss *siteSrv) Get(ctx context.Context, example *site.Site, option *meta.GetOption) (*site.Site, error) {
	return ss.store.Sites().Get(ctx, example, option)
}

func (ss *siteSrv) Update(ctx context.Context, example *site.Site, option *meta.UpdateOption) error {
	return ss.store.Sites().Update(ctx, example, option)
}

func (ss *siteSrv) Del(ctx context.Context, example *site.Site, option *meta.DeleteOption) error {
	return ss.store.Sites().Delete(ctx, example, option)
}

func (ss *siteSrv) DeleteCollection(ctx context.Context, examples []*site.Site, option *meta.DeleteCollectionOption) []error {
	return ss.store.Sites().DeleteCollection(ctx, examples, option)
}

func (ss *siteSrv) List(ctx context.Context, example *site.Site, option *meta.ListOption) ([]*site.Site, int64, error) {
	total, err := ss.store.Sites().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}

	list, err := ss.store.Sites().List(ctx, example, option)
	return list, total, err
}

func (ss *siteSrv) ListByWhereNode(ctx context.Context, example *site.Site, node *meta.WhereNode, option *meta.ListOption) ([]*site.Site, int64, error) {
	total, err := ss.store.Sites().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}

	list, err := ss.store.Sites().ListComplex(ctx, example, node, option)
	return list, total, err
}

func (ss *siteSrv) DeleteSiteAllTags(ctx context.Context, siteID uint) error {
	return ss.store.SiteTags().Delete(ctx, &site.SiteTag{SiteID: siteID}, nil)
}

func newSiteSrv(store store.Factory) SiteSrv {
	return &siteSrv{store: store}
}
