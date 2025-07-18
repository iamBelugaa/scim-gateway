package design

import (
	"goa.design/goa/v3/dsl"
)

// Service describes the main SCIM (System for Cross-domain Identity Management) service.
// This service implements SCIM 2.0 protocol endpoints.
var _ = dsl.Service("scim", func() {
	dsl.Description("The SCIM service implements the SCIM 2.0 protocol, including discovery and identity management features.")

	// Apply static API key security scheme to all methods in this service.
	dsl.Security(StaticTokenAuth)

	// Base path prefix for all endpoints under the SCIM v2 API.
	dsl.HTTP(func() {
		dsl.Path("/scim/v2/")
	})

	// This method returns the configuration metadata for the SCIM service provider.
	dsl.Method("ServiceProviderConfig", func() {
		dsl.Description("Retrieves service provider's configuration metadata including supported SCIM features and authentication schemes.")

		dsl.Payload(StaticTokenAuthRequest)
		dsl.Result(ServiceProviderConfigResponse)

		dsl.HTTP(func() {
			dsl.GET("/ServiceProviderConfig")
			dsl.Header("apiKey:X-API-KEY")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(ServiceProviderConfigResponse)
			})
		})
	})

	// Method for retrieving supported schemas.
	dsl.Method("ListSchemas", func() {
		dsl.Description("Retrieve the supported schemas.")

		dsl.Payload(StaticTokenAuthRequest)
		dsl.Result(ListSchemaResponse)

		dsl.HTTP(func() {
			dsl.GET("/Schemas")
			dsl.Header("apiKey:X-API-KEY")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(ListSchemaResponse)
			})
		})
	})

	// Method for retrieving a specific schema by ID.
	dsl.Method("GetSchema", func() {
		dsl.Description("Retrieve a specific schema by its ID.")

		dsl.Payload(func() {
			dsl.Extend(StaticTokenAuthRequest)
			dsl.Attribute("id", dsl.String, "Schema ID")
			dsl.Required("id")
		})
		dsl.Result(SCIMSchema)

		dsl.HTTP(func() {
			dsl.GET("/Schemas/{id}")
			dsl.Header("apiKey:X-API-KEY")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(SCIMSchema)
			})
		})
	})

	// Method for retrieving resource types.
	dsl.Method("ResourceTypes", func() {
		dsl.Description("Retrieve the supported resource types.")

		dsl.Payload(StaticTokenAuthRequest)
		dsl.Result(ListResourceResponse)

		dsl.HTTP(func() {
			dsl.GET("/ResourceTypes")
			dsl.Header("apiKey:X-API-KEY")
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(ListResourceResponse)
			})
		})
	})
})
