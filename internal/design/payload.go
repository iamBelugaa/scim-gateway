package design

import (
	"goa.design/goa/v3/dsl"
)

// Supported defines a boolean flag indicating whether a specific feature is supported
// by the service provider.
var Supported = dsl.Type("Supported", func() {
	dsl.Description("Indicates whether a specific capability or feature is supported by the service provider.")
	dsl.Attribute("supported", dsl.Boolean, "True if the feature is supported.")
	dsl.Example(map[string]bool{"supported": true})
	dsl.Required("supported")
})

// FilterSupported defines the support for filtering functionality with an optional
// limit on the number of results.
var FilterSupported = dsl.Type("FilterSupported", func() {
	dsl.Description("Specifies whether filtering is supported and optionally the maximum number of results that can be returned.")
	dsl.Attribute("supported", dsl.Boolean, "True if filtering is supported.")
	dsl.Attribute("maxResults", dsl.UInt, "Maximum number of results that can be returned.")

	dsl.Example(map[string]any{"supported": true, "maxResults": 100})
	dsl.Required("supported", "maxResults")
})

// AuthenticationScheme describes a method used for client authentication.
var AuthenticationScheme = dsl.Type("AuthenticationScheme", func() {
	dsl.Description("Defines the authentication mechanism supported by the service provider.")
	dsl.Attribute("type", dsl.String, "Type of the authentication scheme, e.g., 'oauth2', 'basic'.")
	dsl.Attribute("name", dsl.String, "Human-readable name of the authentication scheme.")
	dsl.Attribute("description", dsl.String, "A detailed description of the authentication scheme.")
	dsl.Attribute("specUri", dsl.String, "URI pointing to the authentication scheme specification.")
	dsl.Attribute("documentationUri", dsl.String, "URI pointing to the documentation of the authentication usage.")
	dsl.Attribute("primary", dsl.Boolean, "Indicates if this is the primary authentication scheme.")

	dsl.Example(map[string]any{
		"type":             "oauth2",
		"name":             "OAuth2",
		"description":      "OAuth2 based authentication",
		"specUri":          "https://example.com/specs/oauth2",
		"documentationUri": "https://example.com/docs/auth",
		"primary":          true,
	})
	dsl.Required("type", "name", "description", "specUri", "documentationUri", "primary")
})

// ServiceProviderConfigResponse represents the configuration settings of the
// service provider, including supported features and authentication mechanisms.
var ServiceProviderConfigResponse = dsl.Type("ServiceProviderConfigResponse", func() {
	dsl.Description("Holds metadata and supported capabilities of the service provider, including authentication mechanisms and SCIM features.")
	dsl.Attribute("schemas", dsl.ArrayOf(dsl.String), "List of schema URIs used to define the structure of this resource.")
	dsl.Attribute("documentationUri", dsl.String, "URI pointing to service provider help or documentation.")
	dsl.Attribute("authenticationSchemes", dsl.ArrayOf(AuthenticationScheme), "List of supported authentication schemes.")
	dsl.Attribute("patch", Supported, "Indicates if PATCH operation is supported.")
	dsl.Attribute("bulk", Supported, "Indicates if bulk operations are supported.")
	dsl.Attribute("filter", FilterSupported, "Indicates if filtering is supported and the maximum number of results.")
	dsl.Attribute("changePassword", Supported, "Indicates if password change operation is supported.")
	dsl.Attribute("sort", Supported, "Indicates if sorting is supported.")
	dsl.Attribute("etag", Supported, "Indicates if ETag-based versioning is supported.")

	dsl.Example(map[string]any{
		"schemas":          []string{"urn:ietf:params:scim:schemas:core:2.0:ServiceProviderConfig"},
		"documentationUri": "https://example.com/help",
		"authenticationSchemes": []any{
			map[string]any{
				"type":             "oauthbearertoken",
				"name":             "OAuth Bearer Token",
				"description":      "Authentication scheme using the OAuth Bearer Token Standard",
				"specUri":          "https://www.rfc-editor.org/info/rfc6750",
				"documentationUri": "https://example.com/help",
				"primary":          true,
			},
		},
		"patch":          map[string]bool{"supported": true},
		"bulk":           map[string]bool{"supported": false},
		"changePassword": map[string]bool{"supported": false},
		"sort":           map[string]bool{"supported": false},
		"etag":           map[string]bool{"supported": false},
		"filter":         map[string]any{"supported": true, "maxResults": 100},
	})
	dsl.Required(
		"schemas", "documentationUri", "authenticationSchemes",
		"patch", "bulk", "filter", "changePassword", "sort", "etag",
	)
})

// ServiceProviderRequest is used to authenticate requests to the service
// provider using a static API key.
var ServiceProviderRequest = dsl.Type("ServiceProviderRequest", func() {
	dsl.Description("Describes the request format for service provider operations that require API key-based authentication.")
	dsl.APIKey("StaticTokenAuth", "apiKey", dsl.String, func() {
		dsl.Description("API Key for authentication. Pass this in the 'Authorization' header as: StaticTokenAuth apiKey=<key>")
	})
	dsl.Example("X-API-KEY: Bearer <your-api-key>")
	dsl.Required("apiKey")
})
