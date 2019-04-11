package mockSearchService

import (
	"github.com/scottPlusPlus/tagItUp-v2/core/domain"
)

type MockSearchService struct {
	GalleryService domain.IGalleryService
}

func (srv MockSearchService) Search(search domain.Search) ([]domain.JoinedItem, error) {
	nilReturn := make([]domain.JoinedItem, 0)
	galleries, err := srv.GalleryService.Get(search.Galleries)
	if err != nil {
		return nilReturn, err
	}

	items := make(map[string]domain.TagCloud)

	addItemToResults := func(item domain.Item) {
		cloud, ok := items[item.URL]
		if !ok {
			cloud = domain.NewTagCloud()
		}
		items[item.URL] = cloud.AddList(item.Tags, 1)
	}

	for _, gallery := range galleries {
		for _, item := range gallery.Items {
			added := false
			for _, tag := range search.AllOf {
				if item.Tags.Contains(tag) {
					added = true
					break
				}
			}
			if added {
				addItemToResults(item)
			}
		}
	}

	//convert results to list
	results := make([]domain.JoinedItem, len(items))
	i := 0
	for url, cloud := range items {
		list := domain.NewListJoinedTag()
		for tag, count := range cloud {
			joinedTag := domain.JoinedTag{
				Tag:   tag,
				Count: count,
			}
			list = list.AppendIfUnique(joinedTag)
		}
		results[i] = domain.JoinedItem{
			URL:  url,
			Tags: list,
		}
		i++
	}
	return results, nil
}
