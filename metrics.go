package main

import (
	"fmt"
	"time"
)

type Metrics struct {
	AvgReqTime    time.Duration
	TotalReqsTime time.Duration
	NumberReqs    time.Duration
	MetricChan    chan time.Duration
}

func (m *Metrics) Update(input time.Duration) {
	m.TotalReqsTime += input
	m.NumberReqs++
	m.AvgReqTime = m.TotalReqsTime / m.NumberReqs
	if m.NumberReqs%100 == 0 {
		fmt.Printf("#requests=%d, average time per call: %v, total time: %v\n", m.NumberReqs, m.AvgReqTime, m.TotalReqsTime)
	}
}

func NewMetrics() *Metrics {
	m := &Metrics{}
	m.MetricChan = make(chan time.Duration)
	go func() {
		for {
			metric := <-m.MetricChan
			m.Update(metric)
		}
	}()

	return m
}
