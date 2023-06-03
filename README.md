# sysproxy

Go library and CLI tool for setting system proxy. Supports Windows and macOS.



## Import as a library

[example.go](./cmd/example/example.go) 

```go
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
```



## CLI tool

- Install

```shell
go install github.com/lixvbnet/sysproxy/cmd/sysproxy@latest
```

or download pre-compiled binaries from [Releases](https://github.com/lixvbnet/sysproxy/releases) page.

- Usage

```shell
sysproxy show
sysproxy on <host> <port>
sysproxy off <host> <port>
```

Run `sysproxy -h` for more information.

