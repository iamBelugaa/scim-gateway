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

		// Request body must include the API key for authentication.
		dsl.Payload(ServiceProviderRequest)

		// Successful response returns ServiceProviderConfigResponse with supported features.
		dsl.Result(ServiceProviderConfigResponse)

		dsl.HTTP(func() {
			// Define the HTTP method and path for this operation.
			dsl.GET("/ServiceProviderConfig")

			// Map the API key from the request payload to a header named X-API-KEY.
			dsl.Header("apiKey:X-API-KEY")

			// Define the HTTP 200 OK response and bind it to the result type.
			dsl.Response(dsl.StatusOK, func() {
				dsl.Body(ServiceProviderConfigResponse)
			})
		})
	})
})
