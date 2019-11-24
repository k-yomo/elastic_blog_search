package design

import (
	. "goa.design/goa/v3/dsl"
)

var PostPayload = Type("PostParams", func() {
	Attribute("id", String, "Post's id")
	Attribute("title", String, "Post's title")
	Attribute("description", String, "Post's description")
	Attribute("body", String, "Post's body")
})


var SearchResult = ResultType("application/vnd.posts", func() {
	TypeName("PostOutput")
	ContentType("application/json")
	Attribute("id", String, "Post's id")
	Attribute("title", String, "Post's title")
	Attribute("description", String, "Post's description")
	Required("id", "title", "description")
})

var _ = Service("posts", func() {
	Security(APIKeyAuth)

	Method("register", func() {
		Description("registers blog posts to be searched")
		Payload(func() {
			APIKey("api_key", "key", String, func() {
				Description("API key used to perform authorization")
			})
			Attribute("posts", ArrayOf(PostPayload), func() {
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

	Method("search", func() {
		Description("search blog posts")
		NoSecurity()
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
	})

	Error("badRequest")
	Error("unauthenticated")
	Error("internal", func() {
		Fault()
	})
})
