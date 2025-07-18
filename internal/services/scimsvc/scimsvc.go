package scimsvc

import (
	"context"

	"github.com/iamBelugaa/scim-gateway/gen/scim"
	"github.com/iamBelugaa/scim-gateway/pkg/logger"
	"goa.design/goa/v3/security"
)

type Service struct {
	log *logger.Logger
}

func NewService(log *logger.Logger) *Service {
	return &Service{log: log}
}

// Retrieves service provider's configuration metadata including supported SCIM
// features and authentication schemes.
func (s *Service) ServiceProviderConfig(context.Context, *scim.ServiceProviderRequest) (
	*scim.ServiceProviderConfigResponse, error,
) {
	return &scim.ServiceProviderConfigResponse{}, nil
}

// Retrieve the supported schemas.
func (s *Service) ListSchemas(context.Context, *scim.ServiceProviderRequest) (*scim.ListSchemaResponse, error) {
	return &scim.ListSchemaResponse{}, nil
}

// Retrieve a specific schema by its ID.
func (s *Service) GetSchema(context.Context, *scim.GetSchemaPayload) (*scim.SCIMSchema, error) {
	return &scim.SCIMSchema{}, nil
}

// Retrieve the supported resource types.
func (s *Service) ResourceTypes(context.Context, *scim.ServiceProviderRequest) (*scim.ListResourceResponse, error) {
	return &scim.ListResourceResponse{}, nil
}

// APIKeyAuth implements the authorization logic for the APIKey security scheme.
func (s *Service) APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error) {
	return ctx, nil
}
