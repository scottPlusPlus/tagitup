package domain

type Item struct {
	URL  string
	Tags ListTag
}

func NewItem(url string, tags ...string) Item {
	item := Item{
		URL: url,
	}
	item.Tags = make([]Tag, len(tags))
	for index, tag := range tags {
		t := Tag{
			Value: tag,
		}
		item.Tags[index] = t
	}
	return item
}
