package main

import (
	"flag"
	"fmt"
	"github.com/lixvbnet/sysproxy"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const Version = "0.1"

var Name = filepath.Base(os.Args[0])
var GitHash string

var (
	V = flag.Bool("v", false, "show version")
	H = flag.Bool("h", false, "show help and exit")
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] <on|off|show> [<host> <port>]\n", Name)
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "options\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if !flag.Parsed() {
		flag.Usage()
		return
	}

	if *H {
		flag.Usage()
		return
	}

	if *V {
		fmt.Printf("%s version %s %s\n", Name, Version, GitHash)
		return
	}

	if len(flag.Args()) < 1 {
		flag.Usage()
		return
	}

	op := flag.Arg(0)
	host, port := "", -1
	var err error
	if op == "on" || op == "off" {
		if len(flag.Args()) < 3 {
			log.Fatalf("Error: insufficient number of arguments for operation %s\n", op)
		}
		host = flag.Arg(1)
		port, err = strconv.Atoi(flag.Arg(2))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(host, port)
	}

	//fmt.Println(runtime.GOOS, runtime.GOARCH)
	// OS specific "New()" method, determined by filename suffix, such as "_darwin.go", "_windows.go".
	proxy := sysproxy.New()
	if op == "on" {
		proxy.On(host, port)
	} else if op == "off" {
		proxy.Off(host, port)
	} else if op == "show" {
		proxy.Show()
	} else {
		log.Fatalf("Unsupported operation: %s\n", op)
	}


	//// Convert to concrete type if needed.
	//if pDarwin, ok := proxy.(*sysproxy.Darwin); ok {
	//	fmt.Println("proxy is of type *Darwin!!")
	//	fmt.Println(pDarwin.NIC)
	//	fmt.Println(pDarwin.GetSystemProxy())
	//} else {
	//	fmt.Println("Not of type Darwin")
	//	fmt.Printf("%T\n", proxy)
	//}
}
