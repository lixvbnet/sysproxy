package sysproxy

type SysProxy interface {
	On(host string, port int)
	Off(host string, port int)
	// Show prints current proxy settings to console, and also returns the output
	Show() (output string)
}
