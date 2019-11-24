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

var _ = Service("openapi", func() {
	Files("/swagger.json", "../gen/http/openapi.json")
})
