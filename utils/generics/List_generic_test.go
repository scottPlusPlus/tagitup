//go:generate genny -in=$GOFILE -out=./gen-$GOFILE.go gen "TypeValue=TagWithCount ListType=TagWithCount"

package generics

import (
	"testing"
)

func newTypeValueFromInt(i int) TypeValue {
	//TODO - currently each generated type needs to fill this in...
	return i
}

func TestAppendIfUnique_NewValue_Appends(t *testing.T) {

	item0 := newTypeValueFromInt(0)
	item1 := newTypeValueFromInt(1)
	list := NewListType(item0)
	appended := list.AppendIfUnique(item1)

	if len(appended) != 2 {
		t.Errorf("Expected len == 2, found %v", len(appended))
		return
	}
}

func TestAppendIfUnique_ExistingValue_Rejected(t *testing.T) {

	item0 := newTypeValueFromInt(0)
	list := NewListType(item0)
	appended := list.AppendIfUnique(item0)

	if len(appended) != 1 {
		t.Errorf("Expected len == 1, found %v", len(appended))
		return
	}
}

func TestContains_ExistingValue_ReturnsTrue(t *testing.T) {

	item0 := newTypeValueFromInt(0)
	list := NewListType(item0)
	contains := list.Contains(item0)

	if contains == false {
		t.Errorf("Expected contains == true, found %v", contains)
		return
	}
}

func TestContains_NonExistingValue_ReturnsFalse(t *testing.T) {

	item0 := newTypeValueFromInt(0)
	item1 := newTypeValueFromInt(1)
	list := NewListType(item0)
	contains := list.Contains(item1)

	if contains == true {
		t.Errorf("Expected contains == false, found %v", contains)
		return
	}
}

func TestRemove_RemoveSingle(t *testing.T) {
	item0 := newTypeValueFromInt(0)
	item1 := newTypeValueFromInt(1)
	item2 := newTypeValueFromInt(2)
	list := NewListType(item0, item1, item2)

	list = list.Remove(item1)

	if len(list) != 2 {
		t.Errorf("Expected len == 2, found %v", len(list))
		return
	}

	contains := list.Contains(item1)
	if contains == true {
		t.Errorf("Expected contains == false, found %v", contains)
		return
	}
}
