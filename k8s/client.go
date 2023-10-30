package k8s

import (
	"context"
	"flag"
	"goLibrary/config"
	_ "k8s.io/api/core/v1"                     // k8s 核心组件的包,也就是传统组件
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"   // k8s 容器管理,在线商店之类的包
	_ "k8s.io/apimachinery/pkg/runtime/schema" // k8s
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func NewK8sConn(ctx context.Context, conf *config.Config) *kubernetes.Clientset {
	k8sconfig := flag.String("k8sconfig", "/Users/ian/.kube/config", "kubernetes config file path")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *k8sconfig)
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
