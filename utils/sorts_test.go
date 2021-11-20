package utils

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"sort"

	"testing"
)

type SortSuite struct {
	suite.Suite
}

func (s *SortSuite) SetupTest() {
}

// TestMarshal :
func (s *SortSuite) TestHelloWorld() {
	users := persons{
		{"a", 1},
		{"w", 3},
		{"z", 4},
		{"x", 2},
	}

	sort.Sort(users)
	for _, u := range users {
		fmt.Println(u.Name)
	}
}

func (s *SortSuite) TestAnySort() {
	users := []SortDemo{
		{&Rank{index: 2}, "a"},
		{&Rank{index: 1}, "w"},
		{&Rank{index: 4}, "z"},
		{&Rank{index: 12}, "x"},
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].index < users[j].index
	})
	for _, u := range users {
		fmt.Println(u.Name)
	}
}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(SortSuite))
}
