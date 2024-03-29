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
	ListTagSite(ctx context.Context, tagID uint, option *meta.ListOption) ([]*site.Site, int64, error)
	CreateSiteTagCollection(ctx context.Context, sts []*site.SiteTag, option *meta.CreateCollectionOption) []error
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

func (as *siteSrv) ListTagSite(ctx context.Context, tagID uint, option *meta.ListOption) ([]*site.Site, int64, error) {
	siteTags, err := as.store.SiteTags().List(ctx, &site.SiteTag{TagID: tagID}, &meta.ListOption{PageSize: 1000})
	if err != nil {
		return nil, 0, err
	}
	siteIDs := []uint{}
	for _, st := range siteTags {
		siteIDs = append(siteIDs, st.SiteID)
	}
	node := &meta.WhereNode{
		Conditions: []*meta.Condition{
			{
				Field:    "id",
				Operator: meta.IN,
				Value:    siteIDs,
			},
		},
	}

	sites, err := as.store.Sites().ListComplex(ctx, nil, node, option)
	return sites, int64(len(sites)), err
}

func (as *siteSrv) CreateSiteTagCollection(ctx context.Context, sts []*site.SiteTag, option *meta.CreateCollectionOption) []error {
	return as.store.SiteTags().CreateCollection(ctx, sts, option)
}

func newSiteSrv(store store.Factory) SiteSrv {
	return &siteSrv{store: store}
}
