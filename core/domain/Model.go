package domain

type ItemID struct {
	Value int
}

type GalleryID struct {
	Value int
}

type UserID struct {
	Value int
}

//TODO - figure out which of these we're using... who needs what?
type JoinedItem struct {
	URL  string
	Tags ListJoinedTag
}

type Gallery struct {
	ID           GalleryID
	Summary      string
	Items        ListItem
	Owner        UserID
	Contributors ListUserID
}

type Search struct {
	AllOf     ListTag
	SomeOf    ListTag
	Galleries ListGalleryID
}
