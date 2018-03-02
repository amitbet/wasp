# wasp
Http Wasp - A simple load generator for http / rest calls

Usage of ./wasp:
  -n int
    	number of calls to execute (0 = unlimited)
  -p string
    	the proxy address to use
  -t string
    	the target url to attack (default "http://localhost:9090")
  -w int
    	time between calls in milliseconds (default 10)
      
Example:
./wasp -n 2000 -t http://localhost:9090
#requests=100, average time per call: 695.055µs, total time: 69.505578ms
#requests=200, average time per call: 636.66µs, total time: 127.33212ms
#requests=300, average time per call: 622.105µs, total time: 186.631797ms
#requests=400, average time per call: 613.535µs, total time: 245.414285ms
#requests=500, average time per call: 611.086µs, total time: 305.543323ms
#requests=600, average time per call: 610.331µs, total time: 366.198698ms
#requests=700, average time per call: 608.036µs, total time: 425.625851ms
#requests=800, average time per call: 606.35µs, total time: 485.080287ms
#requests=900, average time per call: 606.385µs, total time: 545.747236ms
#requests=1000, average time per call: 604.453µs, total time: 604.453581ms
#requests=1100, average time per call: 604.88µs, total time: 665.368298ms
#requests=1200, average time per call: 604.091µs, total time: 724.909937ms
#requests=1300, average time per call: 602.708µs, total time: 783.521435ms
#requests=1400, average time per call: 601.8µs, total time: 842.520621ms
#requests=1500, average time per call: 599.47µs, total time: 899.206401ms
#requests=1600, average time per call: 601.148µs, total time: 961.836933ms
#requests=1700, average time per call: 601.119µs, total time: 1.0219028s
#requests=1800, average time per call: 601.71µs, total time: 1.083078481s
#requests=1900, average time per call: 600.765µs, total time: 1.141454017s
#requests=2000, average time per call: 599.087µs, total time: 1.198174371s
