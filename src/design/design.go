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

var PostsParams = ArrayOf(Type("Post", func() {
	Attribute("id", String, "Post's id")
	Attribute("title", String, "Post's title")
	Attribute("description", String, "Post's description")
	Attribute("body", String, "Post's body")
	Required("id", "title", "description", "body")
}))

var _ = Service("register", func() {
	Description("register service registers blog posts to be searched")
	Method("register", func() {
		Payload(PostsParams, func() {
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

var _ = Service("openapi", func() {
	Files("/swagger.json", "../gen/http/openapi.json")
})
