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

// StaticTokenAuthRequest is used to authenticate requests to the service
// provider using a static API key.
var StaticTokenAuthRequest = dsl.Type("ServiceProviderRequest", func() {
	dsl.Description("Describes the request format for service provider operations that require API key-based authentication.")
	dsl.APIKey("StaticTokenAuth", "apiKey", dsl.String, func() {
		dsl.Description("API Key for authentication. Pass this in the 'X-API-KEY' header as: Bearer <your-api-key>")
	})
	dsl.Example("X-API-KEY: Bearer <your-api-key>")
	dsl.Required("apiKey")
})

// SCIMAttribute defines the metadata for an attribute in a schema.
var SCIMAttribute = dsl.Type("SCIMAttribute", func() {
	dsl.Description("Defines a SCIM attribute or sub attribute, including metadata such as mutability, uniqueness, and whether it's multi valued.")
	dsl.Attribute("name", dsl.String, "Name of the attribute")
	dsl.Attribute("type", dsl.String, "Data type of the attribute")
	dsl.Attribute("multiValued", dsl.Boolean, "Whether the attribute is multi valued")
	dsl.Attribute("description", dsl.String, "Human readable description of the attribute")
	dsl.Attribute("required", dsl.Boolean, "True if the attribute is required")
	dsl.Attribute("caseExact", dsl.Boolean, "True if string comparisons are case sensitive")
	dsl.Attribute("mutability", dsl.String, "Defines whether the attribute is readOnly, readWrite, immutable, or writeOnly")
	dsl.Attribute("returned", dsl.String, "Defines when the attribute is returned in a response")
	dsl.Attribute("uniqueness", dsl.String, "Specifies how the attribute value is unique across the service provider")
	dsl.Attribute("canonicalValues", dsl.ArrayOf(dsl.String), "List of canonical values for the attribute")
	dsl.Attribute("referenceTypes", dsl.ArrayOf(dsl.String), "Valid SCIM resource types if this is a reference")
	// dsl.Attribute("subAttributes", dsl.ArrayOf(SCIMAttribute), "List of sub attributes if this is a complex attribute")

	dsl.Required("name", "type", "multiValued", "description", "required", "mutability", "returned")
})

// SCIMMeta represents the meta information of a schema.
var SCIMMeta = dsl.Type("SCIMMeta", func() {
	dsl.Description("Metadata about the SCIM schema resource including resource type and location URI.")
	dsl.Attribute("resourceType", dsl.String, "Type of SCIM resource")
	dsl.Attribute("location", dsl.String, "URI location of the resource")

	dsl.Example(map[string]any{
		"resourceType": "Schema",
		"location":     "/v2/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
	})

	dsl.Required("resourceType", "location")
})

// SCIMSchema represents a SCIM schema resource.
var SCIMSchema = dsl.Type("SCIMSchema", func() {
	dsl.Description("Represents a SCIM schema definition, including its attributes and metadata.")
	dsl.Attribute("id", dsl.String, "Schema ID (usually a URN)")
	dsl.Attribute("name", dsl.String, "Display name of the schema")
	dsl.Attribute("description", dsl.String, "Description of the schema")
	dsl.Attribute("attributes", dsl.ArrayOf(SCIMAttribute), "List of top-level attributes defined in this schema")
	dsl.Attribute("meta", SCIMMeta, "Metadata associated with the schema")

	dsl.Example(map[string]any{
		"id":          "urn:ietf:params:scim:schemas:core:2.0:User",
		"name":        "User",
		"description": "User Schema",
		"attributes": []any{
			map[string]any{
				"name":        "userName",
				"type":        "string",
				"multiValued": false,
				"description": "Unique identifier for the User",
				"required":    true,
				"caseExact":   false,
				"mutability":  "readWrite",
				"returned":    "default",
				"uniqueness":  "server",
			},
		},
		"meta": map[string]any{
			"resourceType": "Schema",
			"location":     "/v2/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
		},
	})

	dsl.Required("id", "name", "description", "attributes", "meta")
})

// ListSchemaResponse represents the SCIM ListSchemaResponse containing schemas
var ListSchemaResponse = dsl.Type("ListSchemaResponse", func() {
	dsl.Description("SCIM ListResponse structure for returning multiple schema definitions in a paginated format.")
	dsl.Attribute("schemas", dsl.ArrayOf(dsl.String), "The list of schema URIs")
	dsl.Attribute("totalResults", dsl.Int, "Total number of results")
	dsl.Attribute("itemsPerPage", dsl.Int, "Number of items per page")
	dsl.Attribute("startIndex", dsl.Int, "Starting index of returned items")
	dsl.Attribute("Resources", dsl.ArrayOf(SCIMSchema), "Array of SCIM schemas")

	dsl.Example(map[string]any{
		"schemas":      []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"},
		"totalResults": 3,
		"itemsPerPage": 3,
		"startIndex":   1,
		"Resources": []any{
			map[string]any{
				"id":          "urn:ietf:params:scim:schemas:core:2.0:User",
				"name":        "User",
				"description": "User Schema",
				"attributes":  []any{},
				"meta": map[string]any{
					"resourceType": "Schema",
					"location":     "/v2/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
				},
			},
		},
	})

	dsl.Required("schemas", "totalResults", "itemsPerPage", "startIndex", "Resources")
})
