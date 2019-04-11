package core

import (
	"github.com/scottPlusPlus/tagItUp-v2/core/domain"
)

type IActorService interface {
	IsValidActor(actor domain.Actor) (bool, error)
}

type ActorService struct {
}

func (s ActorService) IsValidActor(actor domain.Actor) (bool, error) {
	return true, nil
}
