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

func (s *Service) ServiceProviderConfig(
	context.Context, *scim.ServiceProviderRequest,
) (res *scim.ServiceProviderConfigResponse, err error) {
	return &scim.ServiceProviderConfigResponse{}, nil
}
