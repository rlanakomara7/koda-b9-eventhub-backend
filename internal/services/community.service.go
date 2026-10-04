package services

import (
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

type CommunityService struct {
	CommunityRepo *repositories.CommunityRepository
}

func NewCommunityService(communityRepo *repositories.CommunityRepository) *CommunityService {
	return &CommunityService{
		CommunityRepo: communityRepo,
	}
}

func (s *CommunityService) GetCommunities(search string, category string, userID uint) ([]dto.CommunityResponse, error) {
	return s.CommunityRepo.GetCommunities(
		search,
		category,
		userID,
	)
}
