package utils

import (
	"github.com/cheekybits/genny/generic"
)

// golang 泛型

type KeyType generic.Type

type Arraylist struct {
	items []KeyType
}

func NewArraylist() *Arraylist {
	return &Arraylist{
		items: make([]KeyType, 0),
	}
}
func (a *Arraylist) Push(key KeyType) {
	a.items = append(a.items, key)
}

func (a *Arraylist) Pop() KeyType {
	res := a.items[0]
	a.items = a.items[1:]
	return res
}
