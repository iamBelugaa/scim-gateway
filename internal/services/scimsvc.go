package scimsvc

import (
	"context"

	"github.com/iamBelugaa/scim-gateway/gen/scim"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Retrieves service provider's configuration metadata including supported SCIM
// features and authentication schemes.
func (s *Service) ServiceProviderConfig(context.Context, *scim.ServiceProviderRequest) (
	*scim.ServiceProviderConfigResponse, error,
) {
	return &scim.ServiceProviderConfigResponse{
		Schemas:               []string{},
		DocumentationURI:      "",
		AuthenticationSchemes: []*scim.AuthenticationScheme{},
		Patch:                 &scim.Supported{Supported: true},
		Bulk:                  &scim.Supported{Supported: true},
		ChangePassword:        &scim.Supported{Supported: true},
		Sort:                  &scim.Supported{Supported: false},
		Etag:                  &scim.Supported{Supported: false},
		Filter:                &scim.FilterSupported{Supported: true, MaxResults: 50},
	}, nil
}

// Retrieve the supported schemas.
func (s *Service) ListSchemas(context.Context, *scim.ServiceProviderRequest) (*scim.ListSchemaResponse, error) {
	return nil, nil
}

// Retrieve a specific schema by its ID.
func (s *Service) GetSchema(context.Context, *scim.GetSchemaPayload) (*scim.SCIMSchema, error) {
	return nil, nil
}

// Retrieve the supported resource types.
func (s *Service) ResourceTypes(context.Context, *scim.ServiceProviderRequest) (*scim.ListResourceResponse, error) {
	return nil, nil
}
