package k8s

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func (s *TestK8sSuit) TestAPI() {
	s.client.CoreV1()
	s.client.CoreV1()
}

func (s *TestK8sSuit) TestHook() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestK8sSuit))
}
