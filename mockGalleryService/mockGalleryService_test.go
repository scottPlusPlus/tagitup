package mockGalleryService

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/scottPlusPlus/tagItUp-v2/core"
	"github.com/scottPlusPlus/tagItUp-v2/core/domain"
)

func testGalleryService() MockGalleryService {
	actorService := core.ActorService{}
	srv := NewMockGalleryService(actorService)
	return srv
}

func singleGalleryFromService(id domain.GalleryID, service domain.IGalleryService) (domain.Gallery, error) {
	results, err := service.Get([]domain.GalleryID{id})
	if err != nil {
		return domain.Gallery{}, fmt.Errorf("Expected get gallery %v to succeed. Got error: %v", id, err.Error())
	}
	for _, gallery := range results {
		if gallery.ID == id {
			return gallery, nil
		}
	}
	return domain.Gallery{}, fmt.Errorf("Expected get galleries (len %v) to contain Gallery %v", len(results), id)
}

func TestCreateGallery(t *testing.T) {
	galleryService := testGalleryService()
	actor := domain.NewTestActor(0)
	gallery, err := galleryService.CreateGallery(actor)
	if err != nil {
		t.Errorf("Expected CreateGallery to succed, got err: %v", err.Error())
		return
	}

	if gallery.Owner != actor.UserID() {
		t.Errorf("Expected new gallery owner to be actor who created it")
		return
	}

	if len(gallery.Contributors) != 1 {
		t.Errorf("Expected new gallery to have 1 contributor, got %v", len(gallery.Contributors))
		return
	}

	if gallery.Contributors[0] != actor.UserID() {
		t.Errorf("Expected new gallery contributor to be actor %v, got %v", actor.UserID(), gallery.Contributors[0])
		return
	}

	if len(gallery.Items) != 0 {
		t.Errorf("Expected new gallery to have no items, got %v", len(gallery.Items))
		return
	}
}

func Test_GetGallery_ReturnsSameID(t *testing.T) {
	galleryService := testGalleryService()
	actor := domain.NewTestActor(0)
	galleryC, err := galleryService.CreateGallery(actor)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	galleries, err := galleryService.Get([]domain.GalleryID{galleryC.ID})
	if err != nil {
		t.Errorf("Expected getting galleries to succeed. Got error: %v", err.Error())
		return
	}

	if len(galleries) != 1 {
		t.Errorf("Expected to get 1 gallery. Got %v", len(galleries))
		return
	}

	if galleries[0].ID != galleryC.ID {
		t.Errorf("Expected returned gallery ID == %v. Got %v", galleryC.ID, galleries[0].ID)
		return
	}
}

func Test_AddThenGet_ReturnsItem(t *testing.T) {
	galleryService := testGalleryService()
	actor := domain.NewTestActor(0)
	galleryC, err := galleryService.CreateGallery(actor)

	item := domain.Item{URL: "test"}
	err = galleryService.AddItem(galleryC.ID, item, actor)
	if err != nil {
		t.Errorf("Expected adding item to succeed. Got error: %v", err.Error())
		return
	}

	galleryG, err := singleGalleryFromService(galleryC.ID, galleryService)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(galleryG.Items) != 1 {
		t.Errorf("Expected returned gallery to contain my item. Got %v items", len(galleryG.Items))
		return
	}

	item2 := domain.Item{URL: "test2"}
	err = galleryService.AddItem(galleryC.ID, item2, actor)
	if err != nil {
		t.Errorf("Expected adding item to succeed. Got error: %v", err.Error())
		return
	}

	galleryG, err = singleGalleryFromService(galleryC.ID, galleryService)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !galleryG.Items.Contains(item2) {
		t.Errorf("Expected gallery to contain second added item")
		json, _ := json.Marshal(galleryG)
		fmt.Println(string(json))
		return
	}
}

func Test_AddItemFromNonContributor_Rejects(t *testing.T) {
	galleryService := testGalleryService()
	actor := domain.NewTestActor(0)
	galleryC, err := galleryService.CreateGallery(actor)

	newActor := domain.NewTestActor(1)

	item := domain.Item{URL: "test"}
	err = galleryService.AddItem(galleryC.ID, item, newActor)
	if err == nil {
		t.Errorf("Expected service to reject AddItem from non-contributor")
		return
	}
}

func Test_AddContributor_AllowsAddItem(t *testing.T) {
	galleryService := testGalleryService()
	owner := domain.NewTestActor(0)
	galleryC, err := galleryService.CreateGallery(owner)

	contributor := domain.NewTestActor(1)
	err = galleryService.AddContributor(galleryC.ID, contributor.UserID(), owner)
	if err != nil {
		t.Errorf("Expected AddContributor to succeed. Got error: %v", err.Error())
		return
	}

	item := domain.Item{URL: "test"}
	err = galleryService.AddItem(galleryC.ID, item, contributor)
	if err != nil {
		t.Errorf("Expected AddContributor to succeed. Got error: %v", err.Error())
		return
	}

	galleryG, err := singleGalleryFromService(galleryC.ID, galleryService)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !galleryG.Items.Contains(item) {
		t.Errorf("Expected gallery to contain contributor's item")
		return
	}
}
