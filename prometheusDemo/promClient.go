package prometheusDemo

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/model"
	"os"
	"time"
)

func HelloWorld() {
	client, err := api.NewClient(api.Config{
		Address: "localhost:7111",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx := context.Background()
	var s = time.Time{}
	v, err := v1api.Query(ctx, "ianleto{ip='172.27.25.135',diff='\"\"',inspect='udisk-vm-all-patrol'}[1w]", s)
	x, _ := v.(model.Matrix)
	for _, v := range x {
		fmt.Println(v.Metric.String())
		fmt.Println(v.Values)
		for _, i2 := range v.Values {
			fmt.Println(i2.Timestamp.String())
			fmt.Println(i2.Value.String())
		}
	}

	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
}

func ValueModel() {
	//var res = model.SampleStream{
	//	Metric: map[model.LabelName]model.LabelValue{
	//		"labelName1": "labelValue1",
	//		"labelName2": "labelValue2",
	//		"labelName3": "labelValue3",
	//		"diff":       "testDiff",
	//	},
	//	Values: []model.SamplePair{
	//		{Timestamp: 1621378684104,
	//			Value: 11},
	//	},
	//}
	//sz := model.Matrix{&res}
	//for _, stream := range sz {
	//	for lablename, lablevalue := range stream.Metric {
	//
	//		fmt.Println(string(map[string]string{}))
	//		//fmt.Println(conv.String(lablevalue))
	//
	//	}
	//
	//}
}

// 处理返回值
func FormatModelValue(value model.Value) {

}

func Pusher() {
	reg := prometheus.NewRegistry()
	records := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "Ian",
		Help: "",
	}, []string{"ip"})
	reg.MustRegister(records)
	records.WithLabelValues("199.01").Set(float64(10))
	pusher := push.New("http://172.20.178.124:9091", "inspect").Gatherer(reg)
	fmt.Println("ready to push")
	if err := pusher.Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}

}
