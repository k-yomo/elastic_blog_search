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

var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("secret api key for authentication")
})

var Post = Type("Post", func() {
	Attribute("id", String, "Post's id")
	Attribute("title", String, "Post's title")
	Attribute("description", String, "Post's description")
	Attribute("body", String, "Post's body")
})

var _ = Service("register", func() {
	Description("register service registers blog posts to be searched")
	Security(APIKeyAuth)
	Method("register", func() {
		Payload(func() {
			APIKey("api_key", "key", String, func() {
				Description("API key used to perform authorization")
			})
			Attribute("posts", ArrayOf(Post), func() {
				MinLength(1)
			})
			Required("key", "posts")
		})
		Result(Int)
		HTTP(func() {
			POST("/posts/bulk")
			Header("key:Authorization")
			Response(StatusCreated)
		})
	})
	Error("badRequest")
	Error("unauthenticated")
	Error("internal", func() {
		Fault()
	})
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
		Payload(func() {
			Description("search params")
			Attribute("query", String, "search query")
			Attribute("page", UInt, func() {
				Description("page")
				Default(1)
			})
			Attribute("pageSize", UInt, func() {
				Description("results per page")
				Default(50)
			})
			Required("query")
		})
		Result(func() {
			Attribute("posts", CollectionOf(SearchResult))
			Attribute("page", UInt)
			Attribute("totalPage", UInt)
			Required("posts", "page", "totalPage")
		})
		HTTP(func() {
			GET("/posts/search")
			Params(func() {
				Param("query", String)
				Param("page", UInt)
				Param("pageSize", UInt)
				Required("query")
			})
			Response(StatusOK)
		})
		Error("BadRequest")
		Error("internal", func() {
			Fault()
		})
	})
})

var _ = Service("openapi", func() {
	Files("/swagger.json", "../gen/http/openapi.json")
})
