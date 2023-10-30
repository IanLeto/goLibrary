package k8s

import (
	"context"
	"github.com/stretchr/testify/suite"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"testing"
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

// 分页
func (s *TestK8sSuit) TestAPI() {
	var logOptions = &v1.PodLogOptions{
		Container:    "",    // pod中的容器名称
		Follow:       false, // 是否跟随
		Previous:     false, // 是否获取之前的日志
		SinceSeconds: nil,   // 显示日志的当前时间之前的相对时间,以秒为单位. 当前时间不包括在内.
		// 如果设置了sinceTime,则忽略此字段.如果时间早于pod启动时间,则范湖全部日志.如果设置成了未来时间则啥都不反
		SinceTime:                    nil,
		Timestamps:                   true, // 是否显示时间戳,要加上
		TailLines:                    nil,  // 从日志末尾开始显示,如果设置了sinceSeconds,则忽略此字段
		LimitBytes:                   nil,
		InsecureSkipTLSVerifyBackend: false,
	}
	_, err := s.client.CoreV1().RESTClient().Get().Namespace("default").Resource("pods").
		SubResource("log").VersionedParams(logOptions, scheme.ParameterCodec).Stream(s.ctx)
	s.NoError(err)
}
func (s *TestK8sSuit) TestPodList() {
	_, err := s.client.CoreV1().Pods("default").List(s.ctx, v12.ListOptions{
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
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestK8sSuit))
}
