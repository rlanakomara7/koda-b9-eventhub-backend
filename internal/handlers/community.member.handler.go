package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/dto"
	"github.com/rlanakomara7/koda-b9-eventhub-backend/internal/services"
)

type CommunityMemberHandler struct {
	Service *services.CommunityMemberService
}

func NewCommunityMemberHandler(
	service *services.CommunityMemberService,
) *CommunityMemberHandler {

	return &CommunityMemberHandler{
		Service: service,
	}
}

// Join Community
//
// @Summary      Join Community
// @Description  Join community for currently logged in user
// @Tags         communities
// @Produce      json
// @Security     BearerToken
// @Param        id path int true "Community ID"
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /api/communities/{id}/join [post]
func (h *CommunityMemberHandler) JoinCommunity(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "invalid community id",
			})
		return
	}

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized,
			dto.ErrorResponse{
				Message: "unautorizhed",
			})
		return
	}

	err = h.Service.JoinCommunity(
		uint(id),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
			})
		return
	}

	c.JSON(
		http.StatusOK,
		dto.Response{
			Message: "join community success",
		},
	)
}

// Leave Community
//
// @Summary      Leave Community
// @Description  Leave community for currently logged in user
// @Tags         communities
// @Produce      json
// @Security     BearerToken
// @Param        id path int true "Community ID"
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.ErrorResponse
// @Failure      401 {object} dto.ErrorResponse
// @Router       /api/communities/{id}/leave [delete]
func (h *CommunityMemberHandler) LeaveCommunity(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: "invalid community id",
			},
		)
		return
	}

	userID := c.GetUint("user_id")

	if userID == 0 {
		c.JSON(http.StatusUnauthorized,
			dto.ErrorResponse{
				Message: "unauthorized",
			})
		return
	}

	err = h.Service.LeaveCommunity(
		uint(id),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			dto.ErrorResponse{
				Message: err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK,
		dto.Response{
			Message: "leave community success",
		})

}
