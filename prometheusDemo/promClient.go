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
