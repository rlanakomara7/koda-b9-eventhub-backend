package services

import (
	"errors"

	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/repositories"
)

type CommunityMemberService struct {
	CommunityMemberRepo *repositories.CommunityMemberRepository
}

func NewCommunityMemberService(
	communityMemberRepo *repositories.CommunityMemberRepository,
) *CommunityMemberService {

	return &CommunityMemberService{
		CommunityMemberRepo: communityMemberRepo,
	}
}

func (s *CommunityMemberService) JoinCommunity(communityID uint, userID uint) error {

	isMember, err := s.CommunityMemberRepo.IsMember(
		communityID,
		userID,
	)

	if err != nil {
		return err
	}

	if isMember {
		return errors.New("already joined community")
	}

	return s.CommunityMemberRepo.JoinCommunity(
		communityID,
		userID,
	)
}

func (s *CommunityMemberService) LeaveCommunity(
	communityID uint,
	userID uint,
) error {

	isMember, err := s.CommunityMemberRepo.IsMember(
		communityID,
		userID,
	)

	if err != nil {
		return err
	}

	if !isMember {
		return errors.New("not joined community")
	}

	return s.CommunityMemberRepo.LeaveCommunity(
		communityID,
		userID,
	)
}
