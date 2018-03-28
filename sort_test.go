package slice

import (
	"testing"
	"fmt"
)


func TestSort(t *testing.T) {
	type Student struct{
		Name string
		Age int
	}

	students := []*Student{
		{
			Name:"lily",
			Age:10,
		},
		{
			Name:"lucy",
			Age:20,
		},
	}

	Sort(students[:], func(i, j int) bool {
		return students[i].Age > students[j].Age
	})

	for _, s := range students{
		fmt.Printf("Student = %+v", s)
	}
}

