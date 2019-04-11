//go:generate genny -in=$GOFILE -out=./gen-$GOFILE.go gen "TypeValue=TagWithCount ListType=ListTagWithCount"

package generics

import (
	"github.com/cheekybits/genny/generic"
)

type TypeValue generic.Type

type ListType []TypeValue

func (list ListType) areSame(item1 TypeValue, item2 TypeValue) bool {
	return item1 == item2
}

func NewListType(items ...TypeValue) ListType {
	slice := make([]TypeValue, 0)
	for _, item := range items {
		slice = append(slice, item)
	}
	return slice
}

//TODO - handle comparors
func (list ListType) AppendIfUnique(item TypeValue) ListType {
	if list.Contains(item) == false {
		return append(list, item)
	}
	return list
}

func (list ListType) Contains(item TypeValue) bool {
	for _, listItem := range list {
		if list.areSame(listItem, item) {
			return true
		}
	}
	return false
}

func (list ListType) Remove(items ...TypeValue) ListType {
	newList := NewListType()
	toRemove := NewListType(items...)
	for _, existingItem := range list {
		if !toRemove.Contains(existingItem) {
			newList = append(newList, existingItem)
		}
	}
	return newList
}

func (list ListType) IndexOf(item TypeValue) int {
	for index, existingItem := range list {
		if list.areSame(existingItem, item) {
			return index
		}
	}
	return -1
}
