package domain

//TagCloud represents a collection of tags with a count
type TagCloud map[Tag]int

func (cloud TagCloud) Join(cloud2 TagCloud) TagCloud {
	for tag, val := range cloud2 {
		currentVal, ok := cloud[tag]
		if !ok {
			cloud[tag] = val
		} else {
			cloud[tag] = currentVal + val
		}
	}
	return cloud
}

func (cloud TagCloud) Add(tag Tag, val int) TagCloud {
	currentVal, ok := cloud[tag]
	if !ok {
		cloud[tag] = val
	} else {
		cloud[tag] = currentVal + val
	}
	return cloud
}

func (cloud TagCloud) AddList(tags ListTag, val int) TagCloud {
	for _, tag := range tags {
		currentVal, ok := cloud[tag]
		if !ok {
			cloud[tag] = val
		} else {
			cloud[tag] = currentVal + val
		}
	}
	return cloud
}

func NewTagCloud() TagCloud {
	return make(map[Tag]int, 0)
}
