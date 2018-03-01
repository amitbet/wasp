package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func createClient(proxyUrlStr string) *http.Client {
	proxyUrl, err := url.Parse(proxyUrlStr)
	if err != nil {
		log.Fatal("problem while creating http client: ", err)
	}
	if proxyUrlStr == "" {
		return &http.Client{}
	}

	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	return myClient
}

func getAndCalcTimes(wg *sync.WaitGroup, client *http.Client, metrics *Metrics, url string) {
	start := time.Now()
	res, err := client.Get(url)
	if err != nil {
		fmt.Println("Error while connecting to server: ", err)
		os.Exit(-1)
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response body: ", err)
		return
	}
	end := time.Now()
	delta := end.Sub(start)
	metrics.MetricChan <- delta
	//fmt.Println("req done, time for call: ", delta)
	if wg != nil {
		wg.Done()
	}

}

// func fmtDuration(d time.Duration) string {
// 	d = d.Round(time.Millisecond)
// 	h := d / time.Hour
// 	d -= h * time.Hour
// 	m := d / time.Minute
// 	s := d / time.Second
// 	ms := d / time.Millisecond
// 	return fmt.Sprintf("%02d:%02d:%02d.%d", h, m, s, ms)
// }

func main() {
	numberOfCalls := flag.Int("n", 0, "number of calls to execute (0 = unlimited)")
	timeBetweenCalls := flag.Int("w", 10, "time between calls in milliseconds")
	proxyAddress := flag.String("p", "", "the proxy address to use")
	targetAddress := flag.String("t", "http://localhost:9090", "the target url to attack")
	flag.Parse()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	client := createClient(*proxyAddress)
	metrics := NewMetrics()
	wg := &sync.WaitGroup{}

	if *numberOfCalls == 0 {
		for {
			go getAndCalcTimes(nil, client, metrics, *targetAddress) //http://www.google.com/robots.txt
			time.Sleep(time.Duration(*timeBetweenCalls) * time.Millisecond)

			select {
			case signal := <-sigc:
				if signal != nil {
					time.Sleep(2 * time.Second)
					fmt.Printf("#requests=%d, average time per call: %v, total time: %v\n", metrics.NumberReqs, metrics.AvgReqTime, metrics.TotalReqsTime)
					os.Exit(1)
				}
			default:
			}
		}
	} else {
		wg.Add(*numberOfCalls)
		for i := 0; i < *numberOfCalls; i++ {
			go getAndCalcTimes(wg, client, metrics, *targetAddress) //http://www.google.com/robots.txt
			time.Sleep(time.Duration(*timeBetweenCalls) * time.Millisecond)
		}
		wg.Wait()
	}

	fmt.Printf("average time per call: %v, total time: %v\n", metrics.AvgReqTime, metrics.TotalReqsTime)
}
