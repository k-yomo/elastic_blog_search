// Code generated by goa v3.1.2, DO NOT EDIT.
//
// posts client
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package posts

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "posts" service client.
type Client struct {
	RegisterEndpoint goa.Endpoint
	SearchEndpoint   goa.Endpoint
}

// NewClient initializes a "posts" service client given the endpoints.
func NewClient(register, search goa.Endpoint) *Client {
	return &Client{
		RegisterEndpoint: register,
		SearchEndpoint:   search,
	}
}

// Register calls the "register" endpoint of the "posts" service.
func (c *Client) Register(ctx context.Context, p *RegisterPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.RegisterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Search calls the "search" endpoint of the "posts" service.
func (c *Client) Search(ctx context.Context, p *SearchPayload) (res *SearchResult, err error) {
	var ires interface{}
	ires, err = c.SearchEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SearchResult), nil
}
