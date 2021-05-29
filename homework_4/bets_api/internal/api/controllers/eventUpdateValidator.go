package controllers

import "code-cadets-2021/homework_4/bets_api/internal/api/controllers/models"

// EventUpdateValidator validates event update requests.
type EventUpdateValidator interface {
	EventUpdateIsValid(eventUpdateRequestDto models.EventUpdateRequestDto) bool
}
