![hey](http://i.imgur.com/szzD9q0.png)

[![Build Status](https://travis-ci.org/rakyll/hey.svg?branch=master)](https://travis-ci.org/rakyll/hey)

------

TODO:
+ fasthttp支持time trace
+ fasthttp支持redirect

**An Enhanced Version of [rakyll/hey](https://github.com/rakyll/hey).**

Adds an option `-fast` to use [valyala/fasthttp](https://github.com/valyala/fasthttp) instead of [golang/net/http](https://github.com/golang/net).

Provides shorter return time, higher concurrency, and lower memory footprint.

Install with `go get -u github.com/Pingze-github/hey`.

Benchmark:
``` 
goos: windows
goarch: amd64
pkg: github.com/Pingze-github/hey
BenchmarkRunNetHTTP-4                      20000             86766 ns/op            5811 B/op         79 allocs/op
BenchmarkRunFastHTTP-4                     30000             54656 ns/op            1419 B/op         14 allocs/op
BenchmarkRunParallelNetHTTP-4              50000             25200 ns/op            5804 B/op         78 allocs/op
BenchmarkRunParallelFastHTTP-4            100000             18641 ns/op            1412 B/op         14 allocs/op
BenchmarkRunParallelNetHTTP128-4           30000             52020 ns/op            8259 B/op         92 allocs/op
BenchmarkRunParallelFastHTTP128-4         100000             17802 ns/op            1538 B/op         14 allocs/op
```

------

hey is a tiny program that sends some load to a web application.

hey was originally called boom and was influenced from Tarek Ziade's
tool at [tarekziade/boom](https://github.com/tarekziade/boom). Using the same name was a mistake as it resulted in cases
where binary name conflicts created confusion.
To preserve the name for its original owner, we renamed this project to hey.

## Installation

    go get -u github.com/rakyll/hey

## Usage

hey runs provided number of requests in the provided concurrency level and prints stats.

It also supports HTTP2 endpoints.

```
Usage: hey [options...] <url>

Options:
  -n  Number of requests to run. Default is 200.
  -c  Number of requests to run concurrently. Total number of requests cannot
      be smaller than the concurrency level. Default is 50.
  -q  Rate limit, in queries per second (QPS). Default is no rate limit.
  -z  Duration of application to send requests. When duration is reached,
      application stops and exits. If duration is specified, n is ignored.
      Examples: -z 10s -z 3m.
  -o  Output type. If none provided, a summary is printed.
      "csv" is the only supported alternative. Dumps the response
      metrics in comma-separated values format.

  -m  HTTP method, one of GET, POST, PUT, DELETE, HEAD, OPTIONS.
  -H  Custom HTTP header. You can specify as many as needed by repeating the flag.
      For example, -H "Accept: text/html" -H "Content-Type: application/xml" .
  -t  Timeout for each request in seconds. Default is 20, use 0 for infinite.
  -A  HTTP Accept header.
  -d  HTTP request body.
  -D  HTTP request body from file. For example, /home/user/file.txt or ./file.txt.
  -T  Content-type, defaults to "text/html".
  -a  Basic authentication, username:password.
  -x  HTTP Proxy address as host:port.
  -h2 Enable HTTP/2.

  -host	HTTP Host header.

  -disable-compression  Disable compression.
  -disable-keepalive    Disable keep-alive, prevents re-use of TCP
                        connections between different HTTP requests.
  -disable-redirects    Disable following of HTTP redirects
  -cpus                 Number of used cpu cores.
                        (default for current machine is 8 cores)
```

Previously known as [github.com/rakyll/boom](https://github.com/rakyll/boom).

Note: Requires go 1.7 or greater.
