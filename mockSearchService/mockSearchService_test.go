package mockSearchService

import (
	"fmt"
	"testing"

	"github.com/scottPlusPlus/tagItUp-v2/core"
	"github.com/scottPlusPlus/tagItUp-v2/core/domain"
	"github.com/scottPlusPlus/tagItUp-v2/mockGalleryService"
)

const tagA string = "tagA"
const tagB string = "tagB"
const urlA string = "urlA"
const urlB string = "urlB"

type searchTestAggregate struct {
	SearchService  MockSearchService
	GalleryService domain.IGalleryService
}

func testSearchService() searchTestAggregate {
	actorService := core.ActorService{}
	agg := searchTestAggregate{}
	agg.GalleryService = mockGalleryService.NewMockGalleryService(actorService)
	agg.SearchService = MockSearchService{
		GalleryService: agg.GalleryService,
	}
	return agg
}

func tagListContains(list domain.ListJoinedTag, tag domain.JoinedTag) bool {
	for _, t := range list {
		if t.Tag == tag.Tag {
			if t.Count == tag.Count {
				return true
			}
		}
	}
	return false
}

func validateJoinedItems(expected []domain.JoinedItem, have []domain.JoinedItem) error {
	for _, expectedItem := range expected {
		matchedItem, err := findMatchingItem(have, expectedItem)
		if err != nil {
			return err
		}
		for _, expectedTag := range expectedItem.Tags {
			if !tagListContains(matchedItem.Tags, expectedTag) {
				return fmt.Errorf("Expected item %v to contain tag %v. Have: %v", expectedItem.URL, expectedTag, expectedItem)
			}
		}
	}
	return nil
}

func findMatchingItem(items []domain.JoinedItem, item domain.JoinedItem) (domain.JoinedItem, error) {
	for _, itemi := range items {
		if itemi.URL == item.URL {
			return itemi, nil
		}
	}
	return item, fmt.Errorf("Could not find an item that matches %v", item.URL)
}

func newJoinedItem(url string, tags []string, counts []int) domain.JoinedItem {
	if len(tags) != len(counts) {
		fmt.Println("Invalid Joined Item")
	}
	item := domain.JoinedItem{
		URL:  url,
		Tags: make([]domain.JoinedTag, len(tags)),
	}
	for index := 0; index < len(tags); index++ {
		item.Tags[index].Tag = domain.NewTag(tags[index])
		item.Tags[index].Count = counts[index]
	}
	return item
}

//actual tests

func Test_SearchSingleGallery(t *testing.T) {
	agg := testSearchService()
	actor := domain.NewTestActor(0)
	gallery, err := agg.GalleryService.CreateGallery(actor)
	if err != nil {
		t.Errorf("Expected CreateGallery to succed, got err: %v", err.Error())
		return
	}

	item := domain.NewItem(urlA, tagA, tagB)
	err = agg.GalleryService.AddItem(gallery.ID, item, actor)
	if err != nil {
		t.Errorf("Expected adding item to succeed. Got error: %v", err.Error())
		return
	}

	search := domain.Search{
		AllOf:     domain.NewListTag(domain.NewTag(tagA)),
		Galleries: []domain.GalleryID{gallery.ID},
	}
	results, err := agg.SearchService.Search(search)
	if err != nil {
		t.Errorf("Expected search to succeed. Got error: %v", err.Error())
		return
	}

	expected1 := newJoinedItem(urlA, []string{tagA, tagB}, []int{1, 1})
	err = validateJoinedItems([]domain.JoinedItem{expected1}, results)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Test_SearchSingleGallery: PASSED")
}

func Test_SearchTwoGalleries(t *testing.T) {
	agg := testSearchService()
	actor := domain.NewTestActor(0)
	galleryA, _ := agg.GalleryService.CreateGallery(actor)
	galleryB, _ := agg.GalleryService.CreateGallery(actor)

	item := domain.NewItem(urlA, tagA, tagB)
	agg.GalleryService.AddItem(galleryA.ID, item, actor)
	item = domain.NewItem(urlA, tagA)
	agg.GalleryService.AddItem(galleryB.ID, item, actor)

	search := domain.Search{
		AllOf:     domain.NewListTag(domain.NewTag(tagA)),
		Galleries: []domain.GalleryID{galleryA.ID, galleryB.ID},
	}
	results, err := agg.SearchService.Search(search)
	if err != nil {
		t.Errorf("Expected search to succeed. Got error: %v", err.Error())
		return
	}

	expected1 := newJoinedItem(urlA, []string{tagA, tagB}, []int{2, 1})
	err = validateJoinedItems([]domain.JoinedItem{expected1}, results)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("Test_SearchTwoGalleries: PASSED")
}
