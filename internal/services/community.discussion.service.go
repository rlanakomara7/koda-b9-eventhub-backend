package services

import (
	"context"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
)

func (s *CommunityMemberService) GetCommunityDiscussions(
	ctx context.Context,
	communityID uint,
) ([]dto.CommunityDiscussionResponse, error) {
	return s.CommunityMemberRepo.GetCommunityDiscussions(ctx, communityID)
}

func (s *CommunityMemberService) CreateCommunityDiscussion(
	ctx context.Context,
	communityID uint,
	userID uint,
	message string,
) (dto.CommunityDiscussionResponse, error) {
	return s.CommunityMemberRepo.CreateCommunityDiscussion(
		ctx,
		communityID,
		userID,
		message,
	)
}
