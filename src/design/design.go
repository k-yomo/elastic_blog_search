package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("search", func() {
	Title("Blog Posts Search Service")
	Description("HTTP service for blog posts search")
	Server("server", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var Post = Type("Post", func() {
	Attribute("id", String, "Post's id")
	Attribute("title", String, "Post's title")
	Attribute("description", String, "Post's description")
	Attribute("body", String, "Post's body")
})

var _ = Service("register", func() {
	Description("register service registers blog posts to be searched")
	Method("register", func() {
		Payload(ArrayOf(Post), func() {
			MinLength(1)
		})
		Result(Int)
		HTTP(func() {
			POST("/posts/bulk")
			Response(StatusCreated)
		})
	})
	Error("Unauthorized")
	Error("BadRequest")
})

var SearchResult = ResultType("application/vnd.posts", func() {
	TypeName("Post")
	ContentType("application/json")
	Attribute("id", String, "Post's id")
	Attribute("title", String, "Post's title")
	Attribute("description", String, "Post's description")
	Required("id", "title", "description")
})

var _ = Service("search", func() {
	Description("search service searches blog posts with given params")
	Method("search", func() {
		Result(func() {
			Attribute("posts", CollectionOf(SearchResult))
			Attribute("page", UInt)
			Attribute("totalPage", UInt)
			Required("posts", "page", "totalPage")
		})
		Payload(func() {
			Description("search params")
			Attribute("query", String, "search query")
			Attribute("page", UInt, "page")
			Attribute("pageSize", UInt, "results per page")
			Required("query")
		})
		HTTP(func() {
			GET("/posts/search")
			Params(func() {
				Param("query", String, "search query")
				Param("page", UInt, "page")
				Param("pageSize", UInt, "results per page")
				Required("query")
			})
			Response(StatusOK)
		})
		Error("BadRequest")
	})
})

var _ = Service("openapi", func() {
	Files("/swagger.json", "../gen/http/openapi.json")
})
