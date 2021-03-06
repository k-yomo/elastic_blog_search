// Code generated by goa v3.1.2, DO NOT EDIT.
//
// posts service
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package posts

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Service is the posts service interface.
type Service interface {
	// registers blog posts to be searched
	Register(context.Context, *RegisterPayload) (res int, err error)
	// search blog posts
	Search(context.Context, *SearchPayload) (res *SearchResult, err error)
	// get related blog posts
	RelatedPosts(context.Context, *RelatedPostsPayload) (res *RelatedPostsResult, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// APIKeyAuth implements the authorization logic for the APIKey security scheme.
	APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "posts"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"register", "search", "relatedPosts"}

// RegisterPayload is the payload type of the posts service register method.
type RegisterPayload struct {
	// API key used to perform authorization
	Key   string
	Posts []*PostParams
}

// search params
type SearchPayload struct {
	// search query
	Query string
	// page
	Page uint
	// results per page
	PageSize uint
}

// SearchResult is the result type of the posts service search method.
type SearchResult struct {
	Posts     PostOutputCollection
	Page      uint
	TotalPage uint
}

// params
type RelatedPostsPayload struct {
	// post's url
	URL string
	// count
	Count uint
}

// RelatedPostsResult is the result type of the posts service relatedPosts
// method.
type RelatedPostsResult struct {
	Posts PostOutputCollection
	Count uint
}

type PostParams struct {
	// Post's id
	ID *string
	// Post's title
	Title *string
	// Post's description
	Description *string
	// Post's screen image url
	ScreenImageURL *string
	// Post's body
	Body *string
}

type PostOutputCollection []*PostOutput

type PostOutput struct {
	// Post's id
	ID string
	// Post's title
	Title string
	// Post's description
	Description string
	// Post's screen image url
	ScreenImageURL string
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "badRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnauthenticated builds a goa.ServiceError from an error.
func MakeUnauthenticated(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthenticated",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeInternal builds a goa.ServiceError from an error.
func MakeInternal(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
		Fault:   true,
	}
}
