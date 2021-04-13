package prometheusDemo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"io/ioutil"
	"net/http"
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
	//for i := 0; i < 60; i++ {
	//	wg.Add(1)
	//	go func() {
	//
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		wg.Done()
	//	}()
	//
	//	//fmt.Printf("Result:\n%v\n", result)
	//}
	//result, err := v1api.Query(ctx, "up", time.Now())
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
		panic(err)
	}

	//fmt.Println(v)

	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}

}

func PromQueryRequest(url string, requestParam string) ([]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("INIT HTTP ERR %s", err.Error())
	}
	client := &http.Client{}
	q := req.URL.Query()
	q.Add("query", requestParam)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("CLIENT DO ERROR %s", err.Error())
	}
	defer resp.Body.Close()
	promResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("BODY READ ERR %s", err.Error())
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(promResponse, &dat); err != nil {
		return nil, fmt.Errorf("UNMARSHAL ERR %s", err.Error())
	}
	if _, ok := dat["status"]; !ok {
		return nil, fmt.Errorf("NOT FOUNT STATUS ERR.")
	}
	status := dat["status"].(string)
	if status != "success" {
		return nil, fmt.Errorf("PROME REQUEST ERR %s", dat["errorType"].(string))
	}
	fmt.Println(1111)

	//if _, ok := dat["data"]; !ok {
	//	return nil, fmt.Errorf("NOT FOUNT DATA ERR.")
	//}
	//err = ParamIllegalMapCheck(dat["data"])
	//if err != nil {
	//	return nil, err
	//}
	//data := dat["data"].(map[string]interface{})
	//if _, ok := data["result"]; !ok {
	//	return nil, fmt.Errorf("NOT FOUNT RESULT ERR.")
	//}
	//err = ParamIllegalListInterfaceCheck(data["result"])
	//if err != nil {
	//	return nil, err
	//}
	//resultList := data["result"].([]interface{})
	return nil, nil
}
