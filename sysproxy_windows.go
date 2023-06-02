package sysproxy

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// This program is downloaded from https://github.com/getlantern/sysproxy/raw/main/binaries/windows/sysproxy_amd64.exe
//
//go:embed bin/windows/sysproxy_helper.exe
var bin []byte

const defaultPerm = 0744

var dir, program = os.TempDir() + "/.sysproxy", "sysproxy_helper.exe"
var programPath = filepath.Join(dir, program)

type Windows struct{}

func New() SysProxy {
	err := os.MkdirAll(dir, defaultPerm)
	if err != nil {
		log.Fatal(err)
	}
	SaveToFile(programPath, bin, defaultPerm)
	p := &Windows{}
	return p
}

func (p *Windows) On(host string, port int) {
	RunCmd(programPath, "on", host, strconv.Itoa(port))
}

func (p *Windows) Off(host string, port int) {
	RunCmd(programPath, "off", host, strconv.Itoa(port))
}

func (p *Windows) Show() (output string) {
	output = RunCmd(programPath, "show")
	fmt.Println(output)
	return output
}
