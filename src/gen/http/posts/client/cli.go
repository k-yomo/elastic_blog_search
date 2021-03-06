// Code generated by goa v3.1.2, DO NOT EDIT.
//
// posts HTTP client CLI support package
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	posts "github.com/k-yomo/elastic_blog_search/src/gen/posts"
	goa "goa.design/goa/v3/pkg"
)

// BuildRegisterPayload builds the payload for the posts register endpoint from
// CLI flags.
func BuildRegisterPayload(postsRegisterBody string, postsRegisterKey string) (*posts.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(postsRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"posts\": [\n         {\n            \"body\": \"Harum molestiae nemo quo.\",\n            \"description\": \"Tenetur autem libero sint voluptatem.\",\n            \"id\": \"Est ipsa laboriosam assumenda veritatis sapiente ullam.\",\n            \"screenImageUrl\": \"Possimus illo voluptatibus corrupti.\",\n            \"title\": \"Debitis asperiores quasi.\"\n         },\n         {\n            \"body\": \"Harum molestiae nemo quo.\",\n            \"description\": \"Tenetur autem libero sint voluptatem.\",\n            \"id\": \"Est ipsa laboriosam assumenda veritatis sapiente ullam.\",\n            \"screenImageUrl\": \"Possimus illo voluptatibus corrupti.\",\n            \"title\": \"Debitis asperiores quasi.\"\n         }\n      ]\n   }'")
		}
		if body.Posts == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("posts", "body"))
		}
		if len(body.Posts) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.posts", body.Posts, len(body.Posts), 1, true))
		}
		if err != nil {
			return nil, err
		}
	}
	var key string
	{
		key = postsRegisterKey
	}
	v := &posts.RegisterPayload{}
	if body.Posts != nil {
		v.Posts = make([]*posts.PostParams, len(body.Posts))
		for i, val := range body.Posts {
			v.Posts[i] = marshalPostParamsRequestBodyToPostsPostParams(val)
		}
	}
	v.Key = key

	return v, nil
}

// BuildSearchPayload builds the payload for the posts search endpoint from CLI
// flags.
func BuildSearchPayload(postsSearchQuery string, postsSearchPage string, postsSearchPageSize string) (*posts.SearchPayload, error) {
	var err error
	var query string
	{
		query = postsSearchQuery
	}
	var page uint
	{
		if postsSearchPage != "" {
			var v uint64
			v, err = strconv.ParseUint(postsSearchPage, 10, 64)
			page = uint(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for page, must be UINT")
			}
		}
	}
	var pageSize uint
	{
		if postsSearchPageSize != "" {
			var v uint64
			v, err = strconv.ParseUint(postsSearchPageSize, 10, 64)
			pageSize = uint(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for pageSize, must be UINT")
			}
		}
	}
	v := &posts.SearchPayload{}
	v.Query = query
	v.Page = page
	v.PageSize = pageSize

	return v, nil
}

// BuildRelatedPostsPayload builds the payload for the posts relatedPosts
// endpoint from CLI flags.
func BuildRelatedPostsPayload(postsRelatedPostsURL string, postsRelatedPostsCount string) (*posts.RelatedPostsPayload, error) {
	var err error
	var url_ string
	{
		url_ = postsRelatedPostsURL
	}
	var count uint
	{
		if postsRelatedPostsCount != "" {
			var v uint64
			v, err = strconv.ParseUint(postsRelatedPostsCount, 10, 64)
			count = uint(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for count, must be UINT")
			}
		}
	}
	v := &posts.RelatedPostsPayload{}
	v.URL = url_
	v.Count = count

	return v, nil
}
