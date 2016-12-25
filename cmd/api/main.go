package main

import (
	"flag"
	"github.com/facebookgo/grace/gracehttp"
	"net/http"
	"github.com/distillate/zeppelin/lib/api"
)

var (
	listen string
	apiUrl string
)

func init() {
	flag.StringVar(&listen, "listen", ":8080", "listen host and port")
	flag.StringVar(&apiUrl, "apiUrl", "https://localhost:8443", "kubernetes API url")
	flag.Parse()
}

func main() {
	gracehttp.Serve([]*http.Server{
		{Addr: listen, Handler: api.Mux()},
	}...)
}
