package prometheusDemo

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net/http"
	"os"
	"runtime"
	"time"
)

func HelloWorld() {
	client, err := api.NewClient(api.Config{
		Address: "localhost:9091",
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

// NewPromMetrics 上报metric
func NewPromMetrics() {
	cpu := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "ian_test_cpu",
		Help: "hei",
	})
	failures := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "errors_total",
		Help: "errors",
	}, []string{"device"})

	count := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "count",
		Help: "x",
	})

	goroutines := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   "",
		Subsystem:   "",
		Name:        "goroutine_nums",
		Help:        "",
		ConstLabels: nil,
	})

	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpu)
	prometheus.MustRegister(failures, count, goroutines)

	cpu.Set(65.3)
	failures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	count.Inc()
	goroutines.Set(cast.ToFloat64(runtime.NumGoroutine()))
	http.Handle("/metrics", promhttp.Handler())
	logrus.Fatal(http.ListenAndServe(":8080", nil))

}

func init() {

}
