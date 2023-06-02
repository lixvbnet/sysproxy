package sysproxy

import (
	"fmt"
	"strings"
)

type Darwin struct {
	NIC, Service string
}

func New() SysProxy {
	p := &Darwin{}
	p.NIC, p.Service = p.GetActiveNetwork()
	return p
}

func (p *Darwin) GetActiveNetwork() (nic, service string) {
	var cmd string
	// get current NIC and network Service name
	cmd = `scutil --nwi | awk "/Network interfaces/ {print \$3}"`
	nic = RunBashCmd(cmd)
	cmd = fmt.Sprintf(`networksetup -listnetworkserviceorder | grep -B1 "%s" | grep -E "\(\d+\)" | awk '{print $2}'`, nic)
	service = RunBashCmd(cmd)
	return nic, service
}

func (p *Darwin) On(host string, port int) {
	var cmd string
	// set HTTP proxy
	cmd = fmt.Sprintf(`networksetup -setwebproxy "%s" "%s" %d && networksetup -setwebproxystate "%s" on`, p.Service, host, port, p.Service)
	RunBashCmd(cmd)
	// set HTTPS proxy
	cmd = fmt.Sprintf(`networksetup -setsecurewebproxy "%s" "%s" %d && networksetup -setsecurewebproxystate "%s" on`, p.Service, host, port, p.Service)
	RunBashCmd(cmd)
	// set SOCKS proxy
	cmd = fmt.Sprintf(`networksetup -setsocksfirewallproxy "%s" "%s" %d && networksetup -setsocksfirewallproxystate "%s" on`, p.Service, host, port, p.Service)
	RunBashCmd(cmd)
	// set bypass domains
	cmd = fmt.Sprintf(`networksetup -setproxybypassdomains "%s" "192.168.0.0/16" "10.0.0.0/8" "172.16.0.0/12" "127.0.0.1" "localhost" "*.local" "timestamp.apple.com"`, p.Service)
	RunBashCmd(cmd)
}

func (p *Darwin) Off(host string, port int) {
	var cmd string
	// turn off HTTP proxy
	cmd = fmt.Sprintf(`networksetup -setwebproxystate "%s" off`, p.Service)
	RunBashCmd(cmd)
	// turn off HTTPS proxy
	cmd = fmt.Sprintf(`networksetup -setsecurewebproxystate "%s" off`, p.Service)
	RunBashCmd(cmd)
	// turn off SOCKS proxy
	cmd = fmt.Sprintf(`networksetup -setsocksfirewallproxystate "%s" off`, p.Service)
	RunBashCmd(cmd)
}

func (p *Darwin) Show() (output string) {
	sb := strings.Builder{}
	commands := map[string]string{
		"HTTP":           fmt.Sprintf(`networksetup -getwebproxy "%s"`, p.Service),
		"HTTPS":          fmt.Sprintf(`networksetup -getsecurewebproxy "%s"`, p.Service),
		"SOCKS":          fmt.Sprintf(`networksetup -getsocksfirewallproxy "%s"`, p.Service),
		"bypass domains": fmt.Sprintf(`networksetup -getproxybypassdomains "%s"`, p.Service),
	}
	for title, cmd := range commands {
		sb.WriteString(fmt.Sprintf("------------- %s ------------\n", title))
		sb.WriteString(RunBashCmd(cmd))
		sb.WriteByte('\n')
	}
	sb.WriteString("-------------------------")
	output = sb.String()
	fmt.Println(output)
	return output
}
