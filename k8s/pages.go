package k8s

import (
	"context"
	"io"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type LogLineID struct {
	LogTimestamp string `json:"log_timestamp"`
	LineNum      int    `json:"line_num"`
}

type Selection struct {
	ReferencePoint  *LogLineID `json:"reference_point"`
	OffsetFrom      int        `json:"offset_from"`
	OffsetTo        int        `json:"offset_to"`
	LogFilePosition string     `json:"log_file_position"`
}

type LogLine struct {
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

func Pages(offsetFrom, offsetTo int, times string, referenceLineNum int, logFilePosition string) (Selection, []LogLine, string, string, bool) {

	var (
		conn      = NewK8sConn(context.TODO(), nil)
		byteLimit = int64(50000)
		lineLimit = int64(100)
		writer    = strings.Builder{}
		logs      []LogLine
	)
	logSelector := &Selection{
		ReferencePoint: &LogLineID{
			LogTimestamp: times,
			LineNum:      referenceLineNum,
		},
		OffsetFrom:      offsetFrom,
		OffsetTo:        offsetTo,
		LogFilePosition: logFilePosition,
	}
	pods, _ := conn.CoreV1().Pods("default").Get(context.TODO(), "logie-7ccc6d4ccb-928dv", v1.GetOptions{})
	container := pods.Spec.Containers[0]
	logOptions := &v12.PodLogOptions{
		Container:  container.Name,
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
	requestedNumItems := logSelector.OffsetTo - logSelector.OffsetFrom
	referenceIndex := 0
	if logSelector.ReferencePoint == nil || logSelector.ReferencePoint.LogTimestamp == "newset" {
		referenceIndex = len(logs) - 1
	} else if logSelector.ReferencePoint.LogTimestamp == "oldest" {
		referenceIndex = 0
	}
	logTimestamp := logSelector.ReferencePoint.LogTimestamp
	lineMatched := 0
	matchingStartAt := 0
	for idx := range logs {
		if logs[idx].Timestamp == logTimestamp {
			if lineMatched == 0 {
				matchingStartAt = idx
			}
			lineMatched += 1
		} else if lineMatched > 0 {
			break
		}
	}
	var logLineId = logSelector.ReferencePoint
	var offset int
	if logLineId.LineNum < 0 {
		offset = lineMatched + logLineId.LineNum
	} else {
		offset = logLineId.LineNum - 1
	}

	if offset >= 0 && offset < lineMatched {
		referenceIndex = matchingStartAt + offset
	} else {
		referenceIndex = -1
	}
	fromIndex := referenceIndex + logSelector.OffsetFrom
	toIndex := referenceIndex + logSelector.OffsetTo
	lastPage := false
	if requestedNumItems > len(logs) {
		fromIndex = 0
		toIndex = len(logs)
		lastPage = true
	} else if toIndex > len(logs) {
		fromIndex -= toIndex - len(logs)
		toIndex = len(logs)
		lastPage = logSelector.LogFilePosition == "beginning"
	} else if fromIndex < 0 {
		toIndex += -fromIndex
		fromIndex = 0
		lastPage = logSelector.LogFilePosition == "end"
	}
	logTimestamp2 := logs[len(logs)/2].Timestamp
	var step int
	if logs[len(logs)-1].Timestamp == logTimestamp2 {
		step = 1
	} else {
		step = -1
	}
	offset2 := step
	for ; len(logs)/2 >= 0 && len(logs)/2-offset2 < len(logs); offset2 += step {
		if !(logs[len(logs)-offset2].Timestamp == logTimestamp2) {
			break
		}
	}
	newSelection := Selection{
		ReferencePoint: &LogLineID{
			LogTimestamp: logTimestamp2,
			LineNum:      offset2,
		},
		OffsetFrom:      fromIndex - len(logs)/2,
		OffsetTo:        toIndex - len(logs)/2,
		LogFilePosition: logSelector.LogFilePosition,
	}

	var (
		logLines3     = logs[fromIndex:toIndex]
		fromData3     = logs[fromIndex].Timestamp
		toDate3       = logs[toIndex-1].Timestamp
		logSelection3 = newSelection
		lastPage3     = lastPage
	)

	return logSelection3, logLines3, fromData3, toDate3, lastPage3

}
