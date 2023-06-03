package main

import "github.com/lixvbnet/sysproxy"

func main() {
	proxy := sysproxy.New()
	proxy.Show()

	host, port := "127.0.0.1", 8080
	proxy.On(host, port)
	//proxy.Off(host, port)

	proxy.Show()
}
