// Code generated by goa v3.0.7, DO NOT EDIT.
//
// posts HTTP server types
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package server

import (
	posts "github.com/k-yomo/elastic_blog_search/src/gen/posts"
	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "posts" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	Posts []*PostParamsRequestBody `form:"posts,omitempty" json:"posts,omitempty" xml:"posts,omitempty"`
}

// SearchResponseBody is the type of the "posts" service "search" endpoint HTTP
// response body.
type SearchResponseBody struct {
	Posts     PostOutputCollectionResponseBody `form:"posts" json:"posts" xml:"posts"`
	Page      uint                             `form:"page" json:"page" xml:"page"`
	TotalPage uint                             `form:"totalPage" json:"totalPage" xml:"totalPage"`
}

// PostOutputCollectionResponseBody is used to define fields on response body
// types.
type PostOutputCollectionResponseBody []*PostOutputResponseBody

// PostOutputResponseBody is used to define fields on response body types.
type PostOutputResponseBody struct {
	// Post's id
	ID string `form:"id" json:"id" xml:"id"`
	// Post's title
	Title string `form:"title" json:"title" xml:"title"`
	// Post's description
	Description string `form:"description" json:"description" xml:"description"`
}

// PostParamsRequestBody is used to define fields on request body types.
type PostParamsRequestBody struct {
	// Post's id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Post's title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Post's description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Post's body
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
}

// NewSearchResponseBody builds the HTTP response body from the result of the
// "search" endpoint of the "posts" service.
func NewSearchResponseBody(res *posts.SearchResult) *SearchResponseBody {
	body := &SearchResponseBody{
		Page:      res.Page,
		TotalPage: res.TotalPage,
	}
	if res.Posts != nil {
		body.Posts = make([]*PostOutputResponseBody, len(res.Posts))
		for i, val := range res.Posts {
			body.Posts[i] = marshalPostsPostOutputToPostOutputResponseBody(val)
		}
	}
	return body
}

// NewRegisterPayload builds a posts service register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody, key string) *posts.RegisterPayload {
	v := &posts.RegisterPayload{}
	v.Posts = make([]*posts.PostParams, len(body.Posts))
	for i, val := range body.Posts {
		v.Posts[i] = unmarshalPostParamsRequestBodyToPostsPostParams(val)
	}
	v.Key = key
	return v
}

// NewSearchPayload builds a posts service search endpoint payload.
func NewSearchPayload(query string, page uint, pageSize uint) *posts.SearchPayload {
	return &posts.SearchPayload{
		Query:    query,
		Page:     page,
		PageSize: pageSize,
	}
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.Posts == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("posts", "body"))
	}
	if len(body.Posts) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.posts", body.Posts, len(body.Posts), 1, true))
	}
	return
}
