package k8s

import (
	"context"
	"goLibrary/utils"
	"io"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"time"
)

type LogLineID struct {
	LogTimestamp string `json:"log_timestamp"`
	LineNum      int    `json:"line_num"`
}

type Selection struct {
	ReferencePoint  LogLineID `json:"reference_point"`
	OffsetFrom      int       `json:"offset_from"`
	OffsetTo        int       `json:"offset_to"`
	LogFilePosition string    `json:"log_file_position"`
}

type LogLine struct {
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

func NewLog() {
	var (
		conn       = NewK8sConn(context.TODO(), nil)
		byteLimit  = int64(50000)
		lineLimit  = int64(50)
		writer     = strings.Builder{}
		logs       []LogLine
		offsetFrom = 200000
		offsetTo   = 20000
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
	pods, err := conn.CoreV1().Pods("default").Get(context.TODO(), "logie-7ccc6d4ccb-928dv", v1.GetOptions{})
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
	reader, err := conn.CoreV1().RESTClient().Get().Name("default").Name("logie-7ccc6d4ccb-928dv").Resource("pods").
		SubResource("log").VersionedParams(logOptions, v1.ParameterCodec).Stream(context.TODO())
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
				Timestamp: line[id+1:],
				Content:   line[:id],
			})
		}
	}
	var point int
	var (
		index int //

	)
	index = len(logs) - 1 // 先考虑从最后一行开始

	var (
		referenceLineIndex = index //
		requestedNumItems  = logSelector.OffsetFrom - logSelector.OffsetTo
	)
	if referenceLineIndex == -1 || requestedNumItems <= 0 || len(logs) == 0 {
		return
	}

	fromIndex := referenceLineIndex + logSelector.OffsetFrom
	toIndex := referenceLineIndex + logSelector.OffsetTo
	lastPage := false
	// 要查询的日志行数大于实际日志行数
	if requestedNumItems > len(logs) {
		fromIndex = 0       // 从第一行开始
		toIndex = len(logs) // 到最后一行结束
		lastPage = true     // 最后一页
	} else if toIndex > len(logs) { // 首次查询，toindex 一定会 > len(logs)
		fromIndex -= toIndex - len(logs) // 首次查询，fromIndex 一定会 远远 > len(logs)
		toIndex = len(logs)              // 最大行数
		lastPage = false                 // 不是最后一页
	} else if fromIndex < 0 {
		// todo
	}

	newSelection := Selection{
		ReferencePoint:  LogLineID{},
		OffsetFrom:      0,
		OffsetTo:        0,
		LogFilePosition: "",
	}
}
