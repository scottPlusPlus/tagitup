package domain

type Tag struct {
	Value string
}

type JoinedTag struct {
	Tag   Tag
	Count int
}

func NewTag(tag string) Tag {
	return Tag{
		Value: tag,
	}
}

func NewJoinedTag(tag string, count int) JoinedTag {
	return JoinedTag{
		Tag:   NewTag(tag),
		Count: count,
	}
}
