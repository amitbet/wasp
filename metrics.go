package main

import (
	"fmt"
	"time"
)

type Metric struct {
	Duration      time.Duration
	StatusCode    int
	BytesRecieved int
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
	MetricChan     chan Metric
}

func (m *Metrics) Update(metric Metric) {
	m.TotalBytes += int64(metric.BytesRecieved)
	m.TimeSinceStart = time.Now().Sub(m.StartTime)
	m.Throughput = float64(m.TotalBytes) / float64(m.TimeSinceStart.Seconds())
	//m.Throughput = float64(m.TotalBytes) / float64(m.TotalReqsTime.Seconds())
	m.StatusCounts[metric.StatusCode]++

	m.TotalReqsTime += metric.Duration
	m.NumberReqs++
	m.AvgReqTime = m.TotalReqsTime / time.Duration(m.NumberReqs)
	m.AvgBytesPerReq = float64(m.TotalBytes) / float64(m.NumberReqs)
	if m.NumberReqs%100 == 0 {
		m.Print()
	}
}

func (m *Metrics) Print() {
	fmt.Printf("#requests=%d, average time per call: %v, total time: %v, average response size: %fKB throughput: %dKBps \nstatuses: %v\n", m.NumberReqs, m.AvgReqTime, m.TimeSinceStart, m.AvgBytesPerReq/1000, int(m.Throughput/1000), m.StatusCounts)

}

func NewMetrics() *Metrics {
	m := &Metrics{}
	m.MetricChan = make(chan Metric)
	m.StatusCounts = make(map[int]int)
	m.StartTime = time.Now()
	go func() {
		for {
			metric := <-m.MetricChan
			m.Update(metric)
		}
	}()

	return m
}
