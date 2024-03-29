package json_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type JsonSuite struct {
	suite.Suite
}

func (s *JsonSuite) SetupTest() {
}

// TestMarshal :
func (s *JsonSuite) TestUmarshal() {

	type NodeEntity struct {
		ID       string        `json:"_id"`
		Name     string        `json:"name"`
		Content  string        `json:"content"`
		Depend   string        `json:"depend"`
		Father   string        `json:"father"`
		FatherID string        `json:"father_id"`
		Done     bool          `json:"done"`
		Status   string        `json:"status"`
		Note     string        `json:"note"`
		Tags     []string      `json:"tags"`
		Children []string      `json:"children"`
		Nodes    []*NodeEntity `json:"nodes"`
	}
	var x = NodeEntity{}
	res, err := json.Marshal(x)
	s.NoError(err)
	fmt.Println(string(res))
}

// TestMarshal :
func (s *JsonSuite) TestHelloWorld() {
	data := []byte("{\"key\":\"word\",\"context\": {\"k1\":\"v1\"}}")
	var a = struct {
		Key     string      `json:"key"`
		Context interface{} `json:"context"`
	}{}
	err := json.Unmarshal(data, &a)
	s.NoError(err)
	fmt.Println(a.Context)
}

// 比较两个json数据
func (s *JsonSuite) TestDeepEqual() {
	var (
		map1 = map[string]interface{}{}
		map2 = map[string]interface{}{}
	)

	data := []byte("{\n  \"log_level\": 0,\n  \"resource_type\": \"ReplicaSet\",\n  \"resource_id\": \"esid\",\n  \"time\": \"1680599874\",\n  \"detail\": {\n    \"cluster_name\": \"devk8s\",\n    \"event\": {\n      \"apiVersion\": \"v1\",\n      \"firstTimestamp\": \"2023-03-19T14:57:30Z\",\n      \"involvedObject\": {\n        \"apiVersion\": \"v1\",\n        \"fieldPath\": \"spec.containers{filebeat}\",\n        \"kind\": \"Pod\",\n        \"name\": \"filebeat-7cb69b857d-pgjhb\",\n        \"namespace\": \"default\",\n        \"resourceVersion\": \"8986862\",\n        \"uid\": \"770c2417-6091-4cfc-abf1-66c74098c5e8\"\n      },\n      \"source\": {\n        \"component\": \"kubelet\",\n        \"host\": \"minikube\"\n      },\n      \"type\": \"Normal\"\n    }\n  }\n}")
	data2 := []byte("{\n  \"resource_type\": \"ReplicaSet\",\n  \"resource_id\": \"esid\",\n  \"log_level\": 0,\n  \"time\": \"1680599874\",\n  \"detail\": {\n    \"cluster_name\": \"devk8s\",\n    \"event\": {\n      \"apiVersion\": \"v1\",\n      \"firstTimestamp\": \"2023-03-19T14:57:30Z\",\n      \"source\": {\n        \"component\": \"kubelet\",\n        \"host\": \"minikube\"\n      },\n      \"involvedObject\": {\n        \"apiVersion\": \"v1\",\n        \"resourceVersion\": \"8986862\",\n        \"fieldPath\": \"spec.containers{filebeat}\",\n        \"kind\": \"Pod\",\n        \"name\": \"filebeat-7cb69b857d-pgjhb\",\n        \"namespace\": \"default\",\n        \"uid\": \"770c2417-6091-4cfc-abf1-66c74098c5e8\"\n      },\n      \"type\": \"Normal\"\n    }\n  }\n}")
	s.NoError(json.Unmarshal(data, &map1))
	s.NoError(json.Unmarshal(data2, &map2))
	s.Equal(true, reflect.DeepEqual(map2, map1))
}

func (s *JsonSuite) TestJsonReformat() {
	jsonStr := `{
     "k":"v",
		"key1": {
		"k":1}
		}`
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	s.NoError(err)
	jsonByte, err := json.Marshal(data)
	fmt.Println(string(jsonByte))

}

// 直接赋值
// 需求 我有段json {"key":"value"} 要直接给到另一个json的某个字段下
func (s *JsonSuite) TestStyax() {
	var (
		map1 = map[string]interface{}{}
	)
	data := map[string]interface{}{}
	jsonStr := `{"key":"word","context": {"k1":"v1"}}`
	json.Unmarshal([]byte(jsonStr), &data)
	map1["keyx"] = data
	sx, _ := json.Marshal(map1)
	fmt.Println(string(sx))
}

func TestJSONConfiguration(t *testing.T) {
	suite.Run(t, new(JsonSuite))
}
