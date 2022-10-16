package utils_test

import (
	"encoding/json"
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/stretchr/testify/suite"
	"goLibrary/utils"
	"math"
	"strings"

	"testing"
)

type ConvSuite struct {
	suite.Suite
}

func (s *ConvSuite) SetupTest() {
}

// mysql 常用场合
func (s *ConvSuite) TestMySQL() {
	cases := []struct {
		ori    interface{}
		except interface{}
	}{
		{ori: []string{"1", "2", "3"}, except: "1,2,3"},
		{ori: "1,2,3", except: []string{"1", "2", "3"}},
	}
	s.Equal(cases[0].except, utils.ArrToString([]string{"1", "2", "3"}))
	s.Equal(cases[1].except, utils.StringToArr(utils.AnyToString(cases[1].ori.(string))))

}

type Node struct {
	ID    string
	PID   string
	Deep  int
	Child []*Node
}

func (s *ConvSuite) TestConn() {
	cases := []*Node{
		{ID: "2", PID: "", Deep: 0},
		{ID: "1", PID: "", Deep: 0},
		{ID: "1.22", PID: "1", Deep: 0},
		{ID: "1.23", PID: "1", Deep: 0},
		{ID: "1.21", PID: "1", Deep: 0},
		{ID: "1.3.21", PID: "1.23", Deep: 0},
		{ID: "1.2.2.1", PID: "1.3.21", Deep: 0},
		{ID: "1.2.2.1", PID: "1.3.21", Deep: 0},
		{ID: "1.2.2.1.5.5", PID: "1.3.21.1", Deep: 0},
		{ID: "1.2.2.1.4.6", PID: "1.3.21.2", Deep: 0},
		{ID: "12222221", PID: "32221", Deep: 0},
		{ID: "222", PID: "2", Deep: 0},
	}
	res := func(s1, s2 string) bool {
		v1, v2 := strings.Split(s1, "."), strings.Split(s2, ".")
		if len(v1) == len(v2) {
			return false
		}
		//默认v1 小
		if len(v1) > len(v2) {
			v1, v2 = v2, v1
		}
		for i, j := 0, 0; i < len(v1); i, j = i+1, j+1 {
			if v1[i] != v2[i] {
				return false
			}
		}

		return true
	}
	s.Equal(true, res(cases[2].ID, cases[1].ID))
	s.Equal(false, res(cases[3].ID, cases[4].ID))
	s.Equal(true, res(cases[7].ID, cases[8].ID))
	s.Equal(true, res(cases[7].ID, cases[9].ID))
	var deep = map[int][]map[string]Node{}
	for _, node := range cases {
		length := len(strings.Split(node.ID, "."))
		var n = node
		deep[length] = append(deep[length], map[string]Node{node.ID: *n})
	}
	toJs(deep)

	//for i := 1; i < 6; i++ {
	//	for _, m := range deep[i] {
	//		for _, node := range m {
	//
	//		}
	//	}
	//}

}
func (s *ConvSuite) TestConvArr2() {
	memo := make(map[string]*Node)
	cases := []*Node{
		{ID: "2", PID: "", Deep: 0},
		{ID: "1", PID: "", Deep: 0},
		{ID: "1.22", PID: "1", Deep: 0},
		{ID: "1.23", PID: "1", Deep: 0},
		{ID: "1.21", PID: "1", Deep: 0},
		{ID: "1.3.21", PID: "1.23", Deep: 0},
		{ID: "1.2.2.1", PID: "1.3.21", Deep: 0},
		{ID: "1.2.2.1", PID: "1.3.21", Deep: 0},
		{ID: "1.2.2.1.5.5", PID: "1.3.21.1", Deep: 0},
		{ID: "1.2.2.1.4.6", PID: "1.3.21.2", Deep: 0},
		{ID: "12222221", PID: "32221", Deep: 0},
		{ID: "222", PID: "2", Deep: 0},
	}
	for _, v := range cases {
		if _, ok := memo[v.ID]; ok {
			v.Child = memo[v.ID].Child
			memo[v.ID] = v
		} else {
			v.Child = []*Node{}
			memo[v.ID] = v
		}
		if _, ok := memo[v.PID]; ok {
			memo[v.PID].Child = append(memo[v.PID].Child, memo[v.ID])
		} else {
			memo[v.PID] = &Node{
				Child: []*Node{memo[v.ID]},
			}
		}
	}
	mmt(memo)
}

func (s *ConvSuite) TestConvArr() {
	cases := []Node{
		{ID: "1", PID: ""},
		{ID: "321", PID: "23"},
		{ID: "22", PID: "1"},
		{ID: "23", PID: "1"},
		{ID: "21", PID: "1"},
		{ID: "1221", PID: "321"},
		{ID: "2", PID: ""},
		{ID: "222", PID: "2"},
	}
	var trees []*Node
	var hash = map[string]*Node{}
	for _, node := range cases {
		var ephemeral = Node{
			ID:    node.ID,
			PID:   node.PID,
			Child: nil,
		}
		hash[ephemeral.ID] = &ephemeral
	}
	for _, node := range cases {
		var ephemeral = Node{
			ID:    node.ID,
			PID:   node.PID,
			Child: nil,
		}
		if ephemeral.PID == "" {
			trees = append(trees, &ephemeral)
		} else {
			hash[ephemeral.PID].Child = append(hash[ephemeral.PID].Child, &ephemeral)
		}
		hash[ephemeral.ID] = &ephemeral
	}
	res, _ := json.Marshal(trees)
	fmt.Println(string(res))

}

var getTree = func(id string, insert *Node, oldNode []*Node, deep int) {}

func (s *ConvSuite) TestConvArr3() {
	cases := []*Node{
		//
		//{ID: "2", PID: "", Deep: 0},
		//{ID: "1", PID: "", Deep: 0},
		{ID: "1.22", PID: "1", Deep: 0},
		{ID: "1.23", PID: "1", Deep: 0},
		{ID: "1.21", PID: "1", Deep: 0},
		{ID: "1.3.21", PID: "1.23", Deep: 0},
		{ID: "1.2.2.1.5.5", PID: "1.3.21.1", Deep: 0},
		{ID: "1.2.2.1", PID: "1.3.21", Deep: 0},
		{ID: "1.2.2.1.4.6", PID: "1.3.21.2", Deep: 0},
		{ID: "1.2.2.1", PID: "1.3.21", Deep: 0},
		//{ID: "12222221", PID: "32221", Deep: 0},
		//{ID: "222", PID: "2", Deep: 0},
	}
	inde := make([]int, 0)
	for i := range cases {
		cases[i].Deep = len(strings.Split(cases[i].ID, "."))
	}
	// 遍历数组，将当前节点的子节点收集起来
	var getRoot = func(id string, nodes []*Node) []*Node {
		var resR []*Node
		for i := 0; i < len(nodes); i++ {
			if nodes[i].PID == id {
				resR = append(resR, nodes[i])
				x := i
				inde = append(inde, x) // 一对多，收集首尾互连的index
			}
		}

		//for index, node := range nodes {
		//	if node.PID == id {
		//		resR = append(resR, nodes[index])
		//	}
		//}
		return resR
	}
	getTree = func(id string, insert *Node, oldNode []*Node, deep int) {
		// 拿根节点
		epRoot := getRoot(id, oldNode)
		// 如果该根节点是叶子节点 则return
		if epRoot == nil {
			return
		}
		// 否则，继续递归； 遍历当前节点的子节点，并将每一个子节点放到节点的child的中
		for _, node := range epRoot {
			// insert 是当前节点，将收集到的子节点串联
			insert.Child = append(insert.Child, node)
			getTree(node.ID, node, cases, deep)
		}
	}
	roots := getRoot("", cases)
	// 没有头节点
	if len(roots) == 0 {
		// 寻找最小纵深的集合
		var hash = map[int][]string{}
		var min = len(strings.Split(cases[0].ID, "."))
		for _, node := range cases {
			var deep = len(strings.Split(node.ID, "."))
			min = minInt(deep, min)
			hash[deep] = append(hash[deep], node.PID) // 放入父id
		}
		for _, s2 := range RemoveRepeatedElement(hash[min]) {
			roots = append(roots, getRoot(s2, cases)...)
		}
	}

	for _, root := range roots {
		var deep = 1
		getTree(root.ID, root, cases, deep)
	}
	toJs(roots)
	//fmt.Println(inde)
	fmt.Println(utils.Consisten(inde))
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x

}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func mmt(v interface{}) {
	res, _ := json.Marshal(v)
	fmt.Println(string(res))
}

// mysql 常用场合
func (s *ConvSuite) TestConvANny() {
	fmt.Println(math.Min(1.0, 9.2))
	cases := []struct {
		ori    interface{}
		except interface{}
	}{
		{ori: []string{"1", "2", "3"}, except: "1,2,3"},
		{ori: "1,2,3", except: []string{"1", "2", "3"}},
	}
	for _, s2 := range cases {
		res, err := conv.String(s2)
		s.NoError(err)
		fmt.Println(res)
	}

}

func toJs(v interface{}) {
	res, _ := json.Marshal(v)
	fmt.Println(string(res))
}
func BenchmarkHash(b *testing.B) {
	memo := make(map[string]*Node)
	cases := []*Node{
		{ID: "321", PID: "23"},
		{ID: "22", PID: "1"},
		{ID: "23", PID: "1"},
		{ID: "21", PID: "1"},
		{ID: "1221", PID: "321"},
		{ID: "1221", PID: "321"},
		{ID: "12222221", PID: "32221"},
		{ID: "222", PID: "2"},
		{ID: "2", PID: "root"},
		{ID: "1", PID: "root"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range cases {
			if _, ok := memo[v.ID]; ok {
				v.Child = memo[v.ID].Child
				memo[v.ID] = v
			} else {
				v.Child = []*Node{}
				memo[v.ID] = v
			}
			if _, ok := memo[v.PID]; ok {
				memo[v.PID].Child = append(memo[v.PID].Child, memo[v.ID])
			} else {
				memo[v.PID] = &Node{
					Child: []*Node{memo[v.ID]},
				}
			}
		}
	}
	//1 -> 1.1 -> 1.1.1 ->1.1.1.1 -> 1.1.1.2.3.4 -> 1.1.2.3.4.5
	//a-> b -> c -> d ->e -> f
}

//func BenchmarkBack(b *testing.B) {
//	cases := []*Node{
//		{ID: "321", PID: "23"},
//		{ID: "22", PID: "1"},
//		{ID: "23", PID: "1"},
//		{ID: "21", PID: "1"},
//		{ID: "1221", PID: "321"},
//		{ID: "1221", PID: "321"},
//		{ID: "12222221", PID: "32221"},
//		{ID: "222", PID: "2"},
//		{ID: "2", PID: ""},
//		{ID: "1", PID: ""},
//	}
//
//	b.ResetTimer()
//
//	for i := 0; i < b.N; i++ {
//		var getRoot = func(id string, nodes []*Node) []*Node {
//			var resR []*Node
//			for index, node := range nodes {
//				if node.PID == id {
//					resR = append(resR, nodes[index])
//				}
//			}
//			return resR
//		}
//		getTree = func(id string, insert *Node, oldNode []*Node) {
//			// 拿根节点
//			epRoot := getRoot(id, oldNode)
//			// 如果该根节点是叶子节点 则return
//			if epRoot == nil {
//				return
//			}
//			// 否则，继续递归根节点
//			for _, node := range epRoot {
//				insert.Child = append(insert.Child, node)
//				getTree(node.ID, node, oldNode)
//			}
//		}
//		roots := getRoot("", cases)
//		for _, root := range roots {
//			getTree(root.ID, root, cases)
//		}
//	}
//}

func TestConvConfiguration(t *testing.T) {
	suite.Run(t, new(ConvSuite))
}
