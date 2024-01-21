package main

import (
	"Test7/genericTest/generics"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	genericInt()
	fmt.Println("======================================================================================")
	genericString()
	fmt.Println("======================================================================================")
	genericPerson()
}
func genericInt() {
	list := generics.List[int]{Items: []int{}}
	list.Add(12)
	list.Add(45)
	list.Add(23)
	list.Add(29)
	fmt.Printf("list int ===> %+v\n", list.Items)
	list.InsertAt(2, 18)
	fmt.Printf("list int ===> %+v\n", list.Items)
	list.RemoveAt(3)
	fmt.Printf("list int ===> %+v\n", list.Items)
	findIndex := list.Find(29)
	fmt.Printf("find index item %d ==> %d\n", 29, findIndex)
	list.Remove(45)
	fmt.Printf("list int ===> %+v\n", list.Items)
}
func genericString() {
	list := generics.List[string]{Items: []string{}}
	list.Add("Saber")
	list.Add("Bruce")
	list.Add("Ali")
	list.Add("Jackie")
	fmt.Printf("list string ===> %+v\n", list.Items)
	list.InsertAt(2, "Jet")
	fmt.Printf("list string ===> %+v\n", list.Items)
	list.RemoveAt(3)
	fmt.Printf("list string ===> %+v\n", list.Items)
	findIndex := list.Find("Jet")
	fmt.Printf("find index item %s ==> %d\n", "Jet", findIndex)
	list.Remove("Bruce")
	fmt.Printf("list string ===> %+v\n", list.Items)
}
func genericPerson() {
	list := generics.List[Person]{Items: []Person{}}
	list.Add(Person{FirstName: "Saber", LastName: "Azizi", Age: 36})
	list.Add(Person{FirstName: "Bruce", LastName: "Lee", Age: 33})
	list.Add(Person{FirstName: "Ali", LastName: "Omidian", Age: 27})
	list.Add(Person{FirstName: "Jackie", LastName: "Chan", Age: 70})
	fmt.Printf("list Person ===> %+v\n", list.Items)
	list.InsertAt(2, Person{FirstName: "Jet", LastName: "Li", Age: 57})
	fmt.Printf("list Person ===> %+v\n", list.Items)
	list.RemoveAt(3)
	fmt.Printf("list Person ===> %+v\n", list.Items)
	findIndex := list.Find(Person{FirstName: "Jet", LastName: "Li", Age: 57})
	fmt.Printf("find index item %s ==> %d\n", "Ali", findIndex)
	list.Remove(Person{FirstName: "Bruce", LastName: "Lee", Age: 33})
	fmt.Printf("list Person ===> %+v\n", list.Items)
}
