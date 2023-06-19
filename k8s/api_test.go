package k8s

import (
	"context"
	"fmt"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/stretchr/testify/suite"
	"goLibrary/utils"
	"io"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
	"testing"
	"time"
)

type TestK8sSuit struct {
	suite.Suite
	client *kubernetes.Clientset
	ctx    context.Context
}

func (s *TestK8sSuit) SetupTest() {
	s.ctx = context.TODO()
	s.client = NewK8sConn(context.TODO(), nil)
}

func (s *TestK8sSuit) TestConf() {
	ns := s.client.CoreV1().Namespaces()
	nsList, err := ns.List(s.ctx, v1.ListOptions{
		LabelSelector:        "",
		FieldSelector:        "",
		Watch:                false,
		AllowWatchBookmarks:  false,
		ResourceVersion:      "",
		ResourceVersionMatch: "",
		TimeoutSeconds:       nil,
		Limit:                0,
		Continue:             "",
	})
	s.NoError(err)
	for _, v := range nsList.Items {
		fmt.Println(v.Name)
	}
}

func (s *TestK8sSuit) TestGetLog() {
	var (
		byteLimit = int64(50000)
		lineLimit = int64(50)
		writer    = strings.Builder{}
		logs      []LogLine
		//offsetFrom = 200000
		//offsetTo   = 20000
	)
	logSelector := &Selection{
		ReferencePoint: LogLineID{
			LogTimestamp: time.Now().Format(time.RFC3339), // 前端传入
			LineNum:      -1,                              // 前端传入
		},
		OffsetFrom:      10,
		OffsetTo:        100,
		LogFilePosition: "",
	}
	pods, err := s.client.CoreV1().Pods("default").Get(context.TODO(), "logie-7ccc6d4ccb-928dv", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	container := pods.Spec.Containers[0]
	logOptions := &v12.PodLogOptions{
		Container:  container.Name,
		Follow:     false,
		Previous:   false,
		Timestamps: true,
	}

	if logSelector.LogFilePosition == "beginning" {
		logOptions.LimitBytes = &byteLimit
	} else {
		logOptions.TailLines = &lineLimit
	}

	reader, err := s.client.CoreV1().RESTClient().Get().Namespace("default").Name("logie-7ccc6d4ccb-928dv").Resource("pods").
		SubResource("log").VersionedParams(logOptions, scheme.ParameterCodec).Stream(s.ctx)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(&writer, reader)
	utils.NoErr(err)
	for _, line := range strings.Split(writer.String(), "\n") {
		if line == "" {
			continue
		}
		id := strings.Index(line, " ") // 找到第一个空格的位置
		if id > 0 && ('0' <= line[0] && line[0] <= '9') {
			logs = append(logs, LogLine{
				Content:   line[id+1:],
				Timestamp: line[:id],
			})
		}
	}
	logSelector.OffsetFrom = 2000000
	logSelector.OffsetTo = 2000050
	requestedNumItems := logSelector.OffsetTo - logSelector.OffsetFrom // 50
	referenceLineIndex := len(logs) - 1
	switch logSelector.ReferencePoint.LogTimestamp {
	case "newset":
		referenceLineIndex = len(logs) - 1 // 4999
	}
	fromIndex := referenceLineIndex + logSelector.OffsetFrom // 4999 + 2000000 = 2004999
	toIndex := referenceLineIndex + logSelector.OffsetTo     // 4999 + 2000100 = 2005099
	lastPage := false
	if requestedNumItems > len(logs) {

	} else if toIndex > len(logs) {
		fromIndex -= toIndex - len(logs)
		toIndex = len(logs)
		lastPage = false
	}
	newSelection := Selection{
		ReferencePoint: LogLineID{
			LogTimestamp: logs[referenceLineIndex].Timestamp, // todo
			LineNum:      -1,                                 // todo
		},
		OffsetFrom:      fromIndex - len(logs)/2,
		OffsetTo:        toIndex - len(logs)/2,
		LogFilePosition: "",
	}
	var (
		newLogs       []LogLine = logs[fromIndex:toIndex]
		newTimestamp            = logs[fromIndex].Timestamp // todo 起始时间
		new2Timestamp           = logs[toIndex-1].Timestamp // todo 最后时间
	)

	fmt.Println(newLogs, newTimestamp, new2Timestamp, lastPage, newSelection)
}

// 分页
func (s *TestK8sSuit) TestAPI() {
	s.client.CoreV1()

}

func (s *TestK8sSuit) TestHook() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestK8sSuit))
}
