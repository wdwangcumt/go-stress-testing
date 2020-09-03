package api

import (
	"net/http"
)

type StressTask struct {
	TaskId       int          `json:"taskId"`
	HttpConfig   HttpConfig   `json:"httpConfig"`
	StressConfig StressConfig `json:"stressConfig"`
}

type HttpConfig struct {
	Url      string            `json:"url"`
	Method   string            `json:"method"`
	Headers  map[string]string `json:"headers"`
	BodyData string            `json:"bodyData"`
	Timeout  int               `json:"timeout"`
}

type StressConfig struct {
	ConcurrencyQuantity         int `json:"concurrencyQuantity"`         // 并发数
	TotalRequestsPerConcurrency int `json:"totalRequestsPerConcurrency"` // 单个并发的请求数
}

func doStressTask(w http.ResponseWriter, r *http.Request) {

}

func fetchStressTaskStatus(w http.ResponseWriter, r *http.Request) {

}
