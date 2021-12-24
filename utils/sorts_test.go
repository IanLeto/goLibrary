package utils

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"sync"

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

func (s *SortSuite) TestMkc() {
	var mkc = func(params string) {
		url := "http://0.0.0.0:8000/configapi/mkcfe"

		payload := strings.NewReader(params)

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("postman-token", "b420b7be-9656-edd9-6e58-f188f0287df1")

		res, _ := http.DefaultClient.Do(req)

		defer func() { _ = res.Body.Close() }()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))

	}
	var wg = &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rand.Intn(10) < 5 {
				mkc("product=eip&servername=unetfe&cfg-version=unetfe_stable_v1.1.43&svc-version=1.3.41&diff=region:1000001,az:4001")
				fmt.Println("stable:")

			} else {
				mkc("product=eip&servername=unetfe&cfg-version=unetfe_internal_v1.1.43&svc-version=1.3.41&diff=region:1000001,az:4001")
				fmt.Println("interval:")
			}
		}()
	}
	wg.Wait()
}

func TestViperConfiguration(t *testing.T) {
	suite.Run(t, new(SortSuite))
}
