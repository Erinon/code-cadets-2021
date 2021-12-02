package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betStatusValidator BetStatusValidator
	betService         BetService
	betMapper		   BetMapper
}

// NewController creates a new instance of Controller
func NewController(betStatusValidator BetStatusValidator, betService BetService, betMapper BetMapper) *Controller {
	return &Controller{
		betStatusValidator: betStatusValidator,
		betService:         betService,
		betMapper: 		    betMapper,
	}
}

// GetBet returns bet by ID as response.
func (e *Controller) GetBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		betId := ctx.Param("id")

		bet, err := e.betService.GetBet(ctx, betId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		if bet == nil {
			ctx.String(http.StatusNotFound, "bet with specified id does not exist.")
			return
		}

		betDto := e.betMapper.MapDomainBetToBetDto(*bet)

		ctx.JSON(http.StatusOK, betDto)
	}
}

// GetUserBets returns bets by user ID as response.
func (e *Controller) GetUserBets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")

		bets, err := e.betService.GetUserBets(ctx, userId)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		var betDtos []models.BetResponseDto

		for _, bet := range bets {
			betDtos = append(betDtos, e.betMapper.MapDomainBetToBetDto(bet))
		}

		ctx.JSON(http.StatusOK, betDtos)
	}
}

// GetBetsWithStatus returns bets by status as response.
func (e *Controller) GetBetsWithStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")

		if !e.betStatusValidator.BetStatusIsValid(status) {
			ctx.String(http.StatusBadRequest, "status query is not valid.")
			return
		}

		bets, err := e.betService.GetBetsWithStatus(ctx, status)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		var betDtos []models.BetResponseDto

		for _, bet := range bets {
			betDtos = append(betDtos, e.betMapper.MapDomainBetToBetDto(bet))
		}

		ctx.JSON(http.StatusOK, betDtos)
	}
}
