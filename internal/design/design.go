package design

import (
	"goa.design/goa/v3/dsl"
)

// API defines the global settings and metadata for the SCIM Platform service.
var API = dsl.API("scim", func() {
	dsl.Title("SCIM Gateway Service")
	dsl.Description("A service providing a SCIM v2.0 interface for user and group management.")
	dsl.Version("1.0.0")

	// LICENSE information.
	dsl.License(func() {
		dsl.Name("MIT")
		dsl.URL("https://opensource.org/licenses/MIT")
	})

	// Server block defines where the services are hosted.
	dsl.Server("development", func() {
		dsl.Description("Development server hosting the SCIM service")
		dsl.Host("localhost", func() {
			dsl.Description("Local development environment")
			dsl.URI("http://localhost:8080")
		})
	})
})
