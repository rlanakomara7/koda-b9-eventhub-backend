package services

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)


func (s *CommunityService) GetPublicCommunities(
	ctx context.Context,
) ([]dto.CommunityResponse, error) {
	return s.CommunityRepo.GetPublicCommunities(ctx)
}
