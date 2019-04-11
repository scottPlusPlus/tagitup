package domain

import "testing"

func TestCreateNewItem(t *testing.T) {

	item := NewItem("myUrl", "tag1", "tag2")

	if item.URL != "myUrl" {
		t.Errorf("Expected item url == 'myUrl', got %v", item.URL)
		return
	}

	tag1 := Tag{Value: "tag1"}
	tagx := Tag{Value: "x"}

	var contains bool
	contains = item.Tags.Contains(tag1)
	if !contains {
		t.Errorf("Expected item tags to contain 'tag1'")
		return
	}
	contains = item.Tags.Contains(tagx)
	if contains {
		t.Errorf("Expected item tags to not contain 'x'")
		return
	}

}
