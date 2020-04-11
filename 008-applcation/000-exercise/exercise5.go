package main

import (
	"fmt"
	"sort"
)

type user5 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user5

func (a ByAge) Len() int               { return len(a) }
func (a ByAge) Swap(i int, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i int, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user5

func (a ByLast) Len() int               { return len(a) }
func (a ByLast) Swap(i int, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i int, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user5{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user5{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user5{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	user5s := []user5{u1, u2, u3}

	fmt.Println(user5s)

	// your code goes here
	sort.Sort(ByAge(user5s))
	fmt.Println(user5s)
	sort.Sort(ByLast(user5s))
	fmt.Println(user5s)

}
