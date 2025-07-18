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

// JWTAuth defines a security scheme that uses JWTs.
var JWTAuth = dsl.JWTSecurity("JWTAuth", func() {
	dsl.Description("Secures endpoint by requiring a valid JWT token")
	dsl.Scope("api:read", "Read only access")
	dsl.Scope("api:write", "Read and Write access")
})

// StaticTokenAuth defines a security scheme for the static bearer token.
var StaticTokenAuth = dsl.APIKeySecurity("StaticTokenAuth", func() {
	dsl.Description("Secures SCIM endpoints with a static Bearer token generated for a service account")
})
