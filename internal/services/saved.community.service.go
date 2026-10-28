package services

import "github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"

func (s *CommunityMemberService) SaveCommunity(communityID, userID uint) error {
	return s.CommunityMemberRepo.SaveCommunity(communityID, userID)
}

func (s *CommunityMemberService) UnsaveCommunity(communityID, userID uint) error {
	return s.CommunityMemberRepo.UnsaveCommunity(communityID, userID)
}

func (s *CommunityMemberService) GetSavedCommunities(userID uint) ([]dto.CommunityResponse, error) {
	return s.CommunityMemberRepo.GetSavedCommunities(userID)
}
