package prometheusDemo

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"os"
	"time"
)

func HelloWorld() {
	client, err := api.NewClient(api.Config{
		Address: "http://prometheus.iaas.ucloudadmin.com",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx := context.Background()
	v, err := v1api.Query(ctx, "udisk_hostname{ip='172.27.25.135',diff='\"\"',inspect='udisk-vm-all-patrol'}[1w]", time.Now())
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

// 处理返回值
func FormatModelValue(value model.Value) {

}
