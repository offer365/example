package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var (
	daemon *bool
	addr   *string
	path   *string
)

var usage = `Usage:
%s  -a  address -p dir
Example:
%s  -a  :8888   -p /home/
`

func init() {
	daemon = flag.Bool("d", false, "child process")
	addr = flag.String("a", "", "specify service address,example :8080")
	path = flag.String("p", "", "specify path,example /home/")
	flag.Parse()
}

func child(path, addr string) int {
	args := []string{"-d", "-p", path, "-a", addr}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return 0
	}
	return cmd.Process.Pid
}

func main() {
	if addr == nil || *addr == "" || path == nil || *path == "" {
		fmt.Printf(usage, os.Args[0], os.Args[0])
		return
	}
	if daemon != nil && *daemon == true {
		http.ListenAndServe(*addr, http.FileServer(http.Dir(*path)))
		return
	}
	pid := child(*path, *addr)
	fmt.Println("Child process pid:", pid)
	// usage: ./wfs  -a :8888 -p /root/
}
