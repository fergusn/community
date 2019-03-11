package commento

import (
	"net/http"

	"github.com/documize/community/core/env"
	"github.com/documize/community/domain/section/provider"
	"github.com/documize/community/domain/store"
)

// Provider represents Draw.io
type Provider struct {
	Runtime *env.Runtime
	Store   *store.Store
}

// Meta describes us
func (*Provider) Meta() provider.TypeMeta {
	section := provider.TypeMeta{}

	section.ID = "934ad-4ff2-4aa4-98ab-8d54c13c2dc7"
	section.Title = "Commento Discussion"
	section.Description = "Commento powered discussion"
	section.ContentType = "commento"
	section.PageType = "section"
	section.Order = 9992

	return section
}

// Command stub.
func (p *Provider) Command(ctx *provider.Context, w http.ResponseWriter, r *http.Request) {
}

// Render returns data as-is (HTML).
func (p *Provider) Render(ctx *provider.Context, config, data string) string {
	url, _ := p.Store.Setting.Get("commento", "url")

	return url
	//return data
}

// Refresh just sends back data as-is.
func (*Provider) Refresh(ctx *provider.Context, config, data string) string {
	return data
}
