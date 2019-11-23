// Code generated by goa v3.0.7, DO NOT EDIT.
//
// search HTTP client types
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package client

import (
	searchviews "github.com/k-yomo/elastic_blog_search/src/gen/search/views"
	goa "goa.design/goa/v3/pkg"
)

// SearchResponseBody is the type of the "search" service "search" endpoint
// HTTP response body.
type SearchResponseBody []*PostResponse

// PostResponse is used to define fields on response body types.
type PostResponse struct {
	// Post's id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Post's title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Post's description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// NewSearchPostCollectionOK builds a "search" service "search" endpoint result
// from a HTTP "OK" response.
func NewSearchPostCollectionOK(body SearchResponseBody) searchviews.PostCollectionView {
	v := make([]*searchviews.PostView, len(body))
	for i, val := range body {
		v[i] = &searchviews.PostView{
			ID:          val.ID,
			Title:       val.Title,
			Description: val.Description,
		}
	}
	return v
}

// ValidatePostResponse runs the validations defined on PostResponse
func ValidatePostResponse(body *PostResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	return
}
