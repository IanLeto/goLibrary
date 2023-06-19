package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type LogResponse struct {
	Logs        string    `json:"logs"`
	EarliestLog time.Time `json:"earliest_log"`
}

func main() {
	var config *rest.Config
	var err error

	k8sconfig := flag.String("k8sconfig", "/Users/ian/.kube/config", "kubernetes config file path")
	flag.Parse()
	config, err = clientcmd.BuildConfigFromFlags("", *k8sconfig)

	if err != nil {
		fmt.Println("Error building Kubernetes config:", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error creating Kubernetes client:", err)
		os.Exit(1)
	}

	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		namespace := r.URL.Query().Get("namespace")
		podName := r.URL.Query().Get("pod")
		containerName := r.URL.Query().Get("container")
		tailLinesStr := r.URL.Query().Get("tailLines")
		earliestLogStr := r.URL.Query().Get("earliestLog")

		tailLines := int64(5)
		if tailLinesStr != "" {
			tailLinesInt, err := strconv.Atoi(tailLinesStr)
			if err != nil {
				http.Error(w, "Invalid tailLines value", http.StatusBadRequest)
				return
			}
			tailLines = int64(tailLinesInt)
		}

		var earliestLog time.Time
		if earliestLogStr != "" {
			earliestLog, err = time.Parse(time.RFC3339, earliestLogStr)
			if err != nil {
				http.Error(w, "Invalid earliestLog value", http.StatusBadRequest)
				return
			}
		}

		logs, earliestLog, err := getPodLogs(clientset, namespace, podName, containerName, tailLines, earliestLog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := LogResponse{
			Logs:        logs,
			EarliestLog: earliestLog,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	port := "8080"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func getPodLogs(clientset *kubernetes.Clientset, namespace, podName, containerName string, tailLines int64, earliestLog time.Time) (string, time.Time, error) {
	podLogOpts := corev1.PodLogOptions{
		Container:  containerName,
		Follow:     false,
		Timestamps: true,
		TailLines:  &tailLines,
	}

	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error in opening log stream: %v", err)
	}
	defer podLogs.Close()

	buf := make([]byte, 2000)
	numBytes, err := podLogs.Read(buf)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("error reading log stream: %v", err)
	}

	// Filter logs based on earliestLog
	logs := string(buf[:numBytes])
	lines := strings.Split(logs, "\n")
	filteredLines := []string{}

	for i := len(lines) - 1; i >= 0; i-- {
		if len(lines[i]) > 0 {
			parts := strings.Split(lines[i], " ")
			timestamp, err := time.Parse(time.RFC3339Nano, parts[0])
			if err != nil {
				return "", time.Time{}, fmt.Errorf("error parsing log timestamp: %v", err)
			}

			if !earliestLog.IsZero() && timestamp.Before(earliestLog) {
				filteredLines = append([]string{lines[i]}, filteredLines...)
			}

			earliestLog = timestamp
		}
	}

	filteredLogs := strings.Join(filteredLines, "\n")

	return filteredLogs, earliestLog, nil
}
