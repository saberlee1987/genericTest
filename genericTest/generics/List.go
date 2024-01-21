package generics

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
)

type List[T any] struct {
	Items []T
}

func (List *List[T]) Add(item T) {
	List.Items = append(List.Items, item)
}
func (List *List[T]) InsertAt(index int, item T) {
	fmt.Printf("item %s insert at index %d\n", item, index)
	List.Add(item)
	copy(List.Items[index+1:], List.Items[index:])
	List.Items[index] = item
}
func (List *List[T]) RemoveAt(index int) {
	fmt.Printf("remove item %s at index : %d \n", List.Items[index], index)
	List.Items = append(List.Items[:index], List.Items[index+1:]...)
}
func (List *List[T]) Remove(item T) {
	index := List.Find(item)
	if index != -1 {
		List.RemoveAt(index)
	}
}
func (List *List[T]) Get(index int) T {
	return List.Items[index]
}
func (List *List[T]) Find(item T) int {
	for index, itemValue := range List.Items {
		if cmp.Equal(itemValue, item) {
			return index
		}
	}
	return -1
}
