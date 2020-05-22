package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

type redirectHostStruct struct {
	FromHost string `json:"fromHost"`
	ToHost   string `json:"toHost"`
}

type settingStruct struct {
	Host            string               `json:"host"`
	Port            int                  `json:"port"`
	RedirectHostArr []redirectHostStruct `json:"redirectHostArr"`
}

var settingJSONPath string
var logger = NewColorLogger()

func main() {
	if len(os.Args) != 2 {
		settingJSONPath = "setting.json"
		fmt.Printf("Use setting.json -> You can change file with command " +
			"( ./go-virtual-host [*.json] )\n\n")
	} else {
		settingJSONPath = os.Args[1]
	}

	setting, err := readSettingJSON()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()

	for _, redirctHost := range setting.RedirectHostArr {
		urlString, err := url.Parse("http://" + redirctHost.ToHost)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(urlString)
		logger.info("Reverse Proxy %s -> %s", redirctHost.FromHost, redirctHost.ToHost)
		mux.HandleFunc(redirctHost.FromHost+"/", proxyHandler(proxy))
	}

	// Server setup
	address := fmt.Sprintf("%s:%d", setting.Host, setting.Port)
	server := &http.Server{
		Addr:           address,
		Handler:        mux,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.info("Listening on %s", address)

	panic(server.ListenAndServe())
}

func proxyHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accessIP := strings.Split(r.RemoteAddr, ":")[0]  // "1.2.3.4:5678" -> "1.2.3.4"
		decodedURI, _ := url.QueryUnescape(r.RequestURI) // URI decode
		logger.debug("%s - %s %s %s \n %s", accessIP, r.Host, r.Method, r.Proto, decodedURI)
		p.ServeHTTP(w, r)
	}
}
