package mockUserService

import (
	"fmt"

	"github.com/scottPlusPlus/tagItUp-v2/core/domain"
)

type MockUserService struct {
	Data map[domain.UserID]User
}

func (srv MockUserService) DebugSignIn(id domain.UserID) (domain.Actor, error) {
	actor := domain.Actor{
		ID_: id,
	}
	return actor, nil
}

func (srv MockUserService) GetGalleries(actor domain.Actor) ([]domain.GalleryID, error) {
	emptyList := make([]domain.GalleryID, 0)
	user, err := srv.getUser(actor)
	if err != nil {
		return emptyList, err
	}
	return user.Galleries, nil
}

func (srv MockUserService) FollowGallery(gallery domain.GalleryID, actor domain.Actor) error {
	user, err := srv.getUser(actor)
	if err != nil {
		return err
	}
	user.Galleries = user.Galleries.AppendIfUnique(gallery) //TODO, check if unique
	srv.setUser(user)
	return nil
}

func (srv MockUserService) UnFollowGallery(gallery domain.GalleryID, actor domain.Actor) error {
	user, err := srv.getUser(actor)
	if err != nil {
		return err
	}
	user.Galleries = user.Galleries.Remove(gallery)
	srv.setUser(user)
	return nil
}

func (srv MockUserService) getUser(actor domain.Actor) (User, error) {
	user, ok := srv.Data[actor.UserID()]
	if !ok {
		return User{}, fmt.Errorf("Dont't have any data for user %v", actor.UserID())
	}
	return user, nil
}

func (srv MockUserService) setUser(user User) error {
	srv.Data[user.ID] = user
	return nil
}

type User struct {
	ID        domain.UserID
	Galleries domain.ListGalleryID
}
