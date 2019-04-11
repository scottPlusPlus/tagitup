package mockGalleryService

import (
	"errors"
	"fmt"

	"github.com/scottPlusPlus/tagItUp-v2/core"
	"github.com/scottPlusPlus/tagItUp-v2/core/domain"
)

type MockGalleryService struct {
	Data         map[domain.GalleryID]domain.Gallery
	ActorService core.IActorService
}

func NewMockGalleryService(actorService core.IActorService) MockGalleryService {
	srv := MockGalleryService{
		ActorService: actorService,
	}
	srv.Data = make(map[domain.GalleryID]domain.Gallery, 0)
	return srv
}

func (srv MockGalleryService) CreateGallery(actor domain.Actor) (domain.Gallery, error) {
	_, err := srv.ActorService.IsValidActor(actor)
	if err != nil {
		return domain.Gallery{}, err
	}

	//TODO... don't want to fill the system up with empty galleries
	//maybe a limit on galleries per user?

	gallery := domain.Gallery{}
	gallery.ID = domain.GalleryID{
		Value: len(srv.Data),
	}
	gallery.Items = make([]domain.Item, 0)
	gallery.Owner = actor.UserID()
	gallery.Contributors = make([]domain.UserID, 1)
	gallery.Contributors[0] = actor.UserID()
	srv.saveGallery(gallery)
	return gallery, nil
}

func (srv MockGalleryService) AddItem(galleryID domain.GalleryID, item domain.Item, actor domain.Actor) error {
	_, err := srv.ActorService.IsValidActor(actor)
	if err != nil {
		return err
	}

	gallery, err := srv.galleryByID(galleryID)
	if err != nil {
		return err
	}

	err = srv.assertActorIsContributor(actor, gallery)
	if err != nil {
		return err
	}

	gallery.Items = gallery.Items.AppendIfUnique(item)
	srv.saveGallery(gallery)
	return nil
}

func (srv MockGalleryService) UpdateItem(galleryID domain.GalleryID, item domain.Item, actor domain.Actor) error {
	_, err := srv.ActorService.IsValidActor(actor)
	if err != nil {
		return err
	}

	gallery, err := srv.galleryByID(galleryID)
	if err != nil {
		return err
	}

	err = srv.assertActorIsContributor(actor, gallery)
	if err != nil {
		return err
	}

	index := gallery.Items.IndexOf(item)
	if index < 0 {
		return fmt.Errorf("Gallery %v does not contain an item %v", galleryID, item.URL)
	}

	gallery.Items[index] = item
	srv.saveGallery(gallery)
	return nil
}

func (srv MockGalleryService) Get(ids []domain.GalleryID) ([]domain.Gallery, error) {
	galleries := make([]domain.Gallery, 0)
	for _, id := range ids {
		gallery, err := srv.galleryByID(id)
		if err != nil {
			return galleries, err
		}
		galleries = append(galleries, gallery)
	}

	return galleries, nil
}

func (srv MockGalleryService) AddContributor(galleryID domain.GalleryID, contributor domain.UserID, actor domain.Actor) error {
	_, err := srv.ActorService.IsValidActor(actor)
	if err != nil {
		return err
	}

	gallery, err := srv.galleryByID(galleryID)
	if err != nil {
		return err
	}

	if actor.UserID() != gallery.Owner {
		return errors.New("Only a gallery owner can add contributors")
	}

	gallery.Contributors = append(gallery.Contributors, contributor)
	srv.saveGallery(gallery)
	return nil
}

func (srv MockGalleryService) RemoveContributor(galleryID domain.GalleryID, contributor domain.UserID, actor domain.Actor) error {
	_, err := srv.ActorService.IsValidActor(actor)
	if err != nil {
		return err
	}

	gallery, err := srv.galleryByID(galleryID)
	if err != nil {
		return err
	}

	err = srv.assertActorIsOwner(actor, gallery)
	if err != nil {
		return err
	}

	gallery.Contributors = append(gallery.Contributors, contributor)
	srv.saveGallery(gallery)
	return nil
}

func (srv MockGalleryService) ChangeOwner(galleryID domain.GalleryID, newOwner domain.UserID, actor domain.Actor) error {
	_, err := srv.ActorService.IsValidActor(actor)
	if err != nil {
		return err
	}

	gallery, err := srv.galleryByID(galleryID)
	if err != nil {
		return err
	}

	err = srv.assertActorIsOwner(actor, gallery)
	if err != nil {
		return err
	}

	gallery.Owner = newOwner
	gallery.Contributors = gallery.Contributors.AppendIfUnique(newOwner)
	srv.saveGallery(gallery)
	return nil
}

func (srv MockGalleryService) DestroyGallery(galleryID domain.GalleryID, actor domain.Actor) error {
	_, err := srv.ActorService.IsValidActor(actor)
	if err != nil {
		return err
	}

	gallery, err := srv.galleryByID(galleryID)
	if err != nil {
		return err
	}

	err = srv.assertActorIsOwner(actor, gallery)
	if err != nil {
		return err
	}

	delete(srv.Data, galleryID)
	return nil
}

func (srv MockGalleryService) galleryByID(id domain.GalleryID) (domain.Gallery, error) {
	gallery, ok := srv.Data[id]
	if !ok {
		return domain.Gallery{}, fmt.Errorf("Dont't have any data for gallery %v", id)
	}
	return gallery, nil
}

func (srv MockGalleryService) saveGallery(gallery domain.Gallery) error {
	srv.Data[gallery.ID] = gallery
	return nil
}

func (srv MockGalleryService) assertActorIsOwner(actor domain.Actor, gallery domain.Gallery) error {
	if actor.UserID() != gallery.Owner {
		return fmt.Errorf("Actor %v does not have those permissions for gallery %v", actor.UserID(), gallery.ID)
	}
	return nil
}

func (srv MockGalleryService) assertActorIsContributor(actor domain.Actor, gallery domain.Gallery) error {
	if !gallery.Contributors.Contains(actor.UserID()) {
		return fmt.Errorf("Actor %v does not have those permissions for gallery %v", actor.UserID(), gallery.ID)
	}
	return nil
}

// func (srv MockGalleryService) assertActorIsValidAndOwner(actor domain.Actor, gallery domain.Gallery) error {
// 	_, err := srv.ActorService.IsValidActor(actor)
// 	if err != nil {
// 		return err
// 	}

// 	if actor.UserID() != gallery.Owner {
// 		return fmt.Errorf("Actor is not the owner of gallery %v", gallery.ID)
// 	}
// 	return nil
// }
