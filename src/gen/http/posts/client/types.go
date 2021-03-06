// Code generated by goa v3.1.2, DO NOT EDIT.
//
// posts HTTP client types
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package client

import (
	posts "github.com/k-yomo/elastic_blog_search/src/gen/posts"
	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "posts" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	Posts []*PostParamsRequestBody `form:"posts" json:"posts" xml:"posts"`
}

// SearchResponseBody is the type of the "posts" service "search" endpoint HTTP
// response body.
type SearchResponseBody struct {
	Posts     PostOutputCollectionResponseBody `form:"posts,omitempty" json:"posts,omitempty" xml:"posts,omitempty"`
	Page      *uint                            `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
	TotalPage *uint                            `form:"totalPage,omitempty" json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

// RelatedPostsResponseBody is the type of the "posts" service "relatedPosts"
// endpoint HTTP response body.
type RelatedPostsResponseBody struct {
	Posts PostOutputCollectionResponseBody `form:"posts,omitempty" json:"posts,omitempty" xml:"posts,omitempty"`
	Count *uint                            `form:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
}

// PostParamsRequestBody is used to define fields on request body types.
type PostParamsRequestBody struct {
	// Post's id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Post's title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Post's description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Post's screen image url
	ScreenImageURL *string `form:"screenImageUrl,omitempty" json:"screenImageUrl,omitempty" xml:"screenImageUrl,omitempty"`
	// Post's body
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
}

// PostOutputCollectionResponseBody is used to define fields on response body
// types.
type PostOutputCollectionResponseBody []*PostOutputResponseBody

// PostOutputResponseBody is used to define fields on response body types.
type PostOutputResponseBody struct {
	// Post's id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Post's title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Post's description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Post's screen image url
	ScreenImageURL *string `form:"screenImageUrl,omitempty" json:"screenImageUrl,omitempty" xml:"screenImageUrl,omitempty"`
}

// NewRegisterRequestBody builds the HTTP request body from the payload of the
// "register" endpoint of the "posts" service.
func NewRegisterRequestBody(p *posts.RegisterPayload) *RegisterRequestBody {
	body := &RegisterRequestBody{}
	if p.Posts != nil {
		body.Posts = make([]*PostParamsRequestBody, len(p.Posts))
		for i, val := range p.Posts {
			body.Posts[i] = marshalPostsPostParamsToPostParamsRequestBody(val)
		}
	}
	return body
}

// NewSearchResultOK builds a "posts" service "search" endpoint result from a
// HTTP "OK" response.
func NewSearchResultOK(body *SearchResponseBody) *posts.SearchResult {
	v := &posts.SearchResult{
		Page:      *body.Page,
		TotalPage: *body.TotalPage,
	}
	v.Posts = make([]*posts.PostOutput, len(body.Posts))
	for i, val := range body.Posts {
		v.Posts[i] = unmarshalPostOutputResponseBodyToPostsPostOutput(val)
	}

	return v
}

// NewRelatedPostsResultOK builds a "posts" service "relatedPosts" endpoint
// result from a HTTP "OK" response.
func NewRelatedPostsResultOK(body *RelatedPostsResponseBody) *posts.RelatedPostsResult {
	v := &posts.RelatedPostsResult{
		Count: *body.Count,
	}
	v.Posts = make([]*posts.PostOutput, len(body.Posts))
	for i, val := range body.Posts {
		v.Posts[i] = unmarshalPostOutputResponseBodyToPostsPostOutput(val)
	}

	return v
}

// ValidateSearchResponseBody runs the validations defined on SearchResponseBody
func ValidateSearchResponseBody(body *SearchResponseBody) (err error) {
	if body.Posts == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("posts", "body"))
	}
	if body.Page == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page", "body"))
	}
	if body.TotalPage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("totalPage", "body"))
	}
	if err2 := ValidatePostOutputCollectionResponseBody(body.Posts); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// ValidateRelatedPostsResponseBody runs the validations defined on
// RelatedPostsResponseBody
func ValidateRelatedPostsResponseBody(body *RelatedPostsResponseBody) (err error) {
	if body.Posts == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("posts", "body"))
	}
	if body.Count == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("count", "body"))
	}
	if err2 := ValidatePostOutputCollectionResponseBody(body.Posts); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// ValidatePostOutputCollectionResponseBody runs the validations defined on
// PostOutputCollectionResponseBody
func ValidatePostOutputCollectionResponseBody(body PostOutputCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidatePostOutputResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidatePostOutputResponseBody runs the validations defined on
// PostOutputResponseBody
func ValidatePostOutputResponseBody(body *PostOutputResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.ScreenImageURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("screenImageUrl", "body"))
	}
	return
}
