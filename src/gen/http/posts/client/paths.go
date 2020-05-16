// Code generated by goa v3.1.2, DO NOT EDIT.
//
// HTTP request path constructors for the posts service.
//
// Command:
// $ goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/

package client

// RegisterPostsPath returns the URL path to the posts service register HTTP endpoint.
func RegisterPostsPath() string {
	return "/posts/bulk"
}

// SearchPostsPath returns the URL path to the posts service search HTTP endpoint.
func SearchPostsPath() string {
	return "/posts/search"
}
