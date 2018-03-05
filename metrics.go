package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Metric struct {
	Duration      time.Duration
	StatusCode    int
	BytesRecieved int
	Error         *string
}

type Metrics struct {
	StartTime      time.Time
	AvgReqTime     time.Duration
	AvgBytesPerReq float64
	TotalReqsTime  time.Duration
	NumberReqs     int64
	TimeSinceStart time.Duration
	TotalBytes     int64
	Throughput     float64
	StatusCounts   map[int]int
	ErrorCounts    map[string]int
	MetricChan     chan Metric
}

func (m *Metrics) Update(metric Metric) {
	m.NumberReqs++
	m.TimeSinceStart = time.Now().Sub(m.StartTime)

	if metric.Error != nil {
		m.ErrorCounts[*metric.Error]++
		return
	}

	m.TotalBytes += int64(metric.BytesRecieved)
	m.Throughput = float64(m.TotalBytes) / float64(m.TimeSinceStart.Seconds())
	//m.Throughput = float64(m.TotalBytes) / float64(m.TotalReqsTime.Seconds())
	m.StatusCounts[metric.StatusCode]++
	m.TotalReqsTime += metric.Duration
	m.AvgReqTime = m.TotalReqsTime / time.Duration(m.NumberReqs)
	m.AvgBytesPerReq = float64(m.TotalBytes) / float64(m.NumberReqs)
	if m.NumberReqs%100 == 0 {
		m.Print()
	}
}
func (m *Metrics) FormatErrors() string {
	sb := strings.Builder{}
	for v, k := range m.ErrorCounts {
		sb.WriteString(strconv.Itoa(k) + ":" + v + "\n")
	}
	return sb.String()
}

func (m *Metrics) Print() {
	fmt.Printf("#requests=%d, average time per call: %v, total time: %v, average response size: %fKB throughput: %dKBps \nstatuses: %v\n", m.NumberReqs, m.AvgReqTime, m.TimeSinceStart, m.AvgBytesPerReq/1000, int(m.Throughput/1000), m.StatusCounts)
	if len(m.ErrorCounts) > 0 {
		fmt.Printf("%v", m.FormatErrors())
	}
}

func NewMetrics() *Metrics {
	m := &Metrics{}
	m.MetricChan = make(chan Metric)
	m.StatusCounts = make(map[int]int)
	m.ErrorCounts = make(map[string]int)
	m.StartTime = time.Now()
	go func() {
		for {
			metric := <-m.MetricChan
			m.Update(metric)
		}
	}()

	return m
}
