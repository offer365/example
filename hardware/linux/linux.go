package linux

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

)

const (
	osReleaseFile = "/etc/os-release"

	centOS6Template = `NAME="CentOS Linux"
VERSION="6 %s"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="6"
PRETTY_NAME="CentOS Linux 6 %s"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:6"
HOME_URL="https://www.centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"`

	redhat6Template = `NAME="Red Hat Enterprise Linux Server"
VERSION="%s %s"
ID="rhel"
ID_LIKE="fedora"
VERSION_ID="%s"
PRETTY_NAME="Red Hat Enterprise Linux"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:redhat:enterprise_linux:%s:GA:server"
HOME_URL="https://www.redhat.com/"
BUG_REPORT_URL="https://bugzilla.redhat.com/"`
)

var (
	rePrettyName = regexp.MustCompile(`^PRETTY_NAME=(.*)$`)
	reID         = regexp.MustCompile(`^ID=(.*)$`)
	reVersionID  = regexp.MustCompile(`^VERSION_ID=(.*)$`)
	reUbuntu     = regexp.MustCompile(`[\( ]([\d\.]+)`)
	reCentOS     = regexp.MustCompile(`^CentOS( Linux)? release ([\d\.]+) `)
	reCentOS6    = regexp.MustCompile(`^CentOS release 6\.\d (.*)`)
	reRedhat     = regexp.MustCompile(`[\( ]([\d\.]+)`)
	reRedhat6    = regexp.MustCompile(`^Red Hat Enterprise Linux Server release (.*) (.*)`)
)

func spewFile(path string, data string, perm os.FileMode) {
	_ = ioutil.WriteFile(path, []byte(data+"\n"), perm)
}

func genOSRelease() {
	// CentOS 6.x
	if release := cat("/etc/centos-release"); release != "" {
		if m := reCentOS6.FindStringSubmatch(release); m != nil {
			spewFile(osReleaseFile, fmt.Sprintf(centOS6Template, m[1], m[1]), 0666)
			return
		}
	}

	// RHEL 6.x
	if release := cat("/etc/redhat-release"); release != "" {
		if m := reRedhat6.FindStringSubmatch(release); m != nil {
			version := "6"
			code_name := "()"
			switch l := len(m); l {
			case 3:
				code_name = m[2]
				fallthrough
			case 2:
				version = m[1]
			}
			spewFile(osReleaseFile, fmt.Sprintf(redhat6Template, version, code_name, version, version), 0666)
			return
		}
	}
}



type Linux struct {}

func (l *Linux) CpuInfo() *CPU {
	return &CPU{}
}

func (l *Linux) MemInfo() *Memory {
	panic("implement me")
}

func (l *Linux) NetworkInfo() []NetworkDevice {
	panic("implement me")
}

func (l *Linux) OSInfo() *OS {
	osinfo:=&OS{}
	// This seems to be the best and most portable way to detect OS architecture (NOT kernel!)
	if _, err := os.Stat("/lib64/ld-linux-x86-64.so.2"); err == nil {
		osinfo.Architecture = "amd64"
	} else if _, err := os.Stat("/lib/ld-linux.so.2"); err == nil {
		osinfo.Architecture = "i386"
	}

	if _, err := os.Stat(osReleaseFile); os.IsNotExist(err) {
		genOSRelease()
	}

	f, err := os.Open(osReleaseFile)
	if err != nil {
		return nil
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if m := rePrettyName.FindStringSubmatch(s.Text()); m != nil {
			osinfo.Name = strings.Trim(m[1], `"`)
		} else if m := reID.FindStringSubmatch(s.Text()); m != nil {
			osinfo.Vendor = strings.Trim(m[1], `"`)
		} else if m := reVersionID.FindStringSubmatch(s.Text()); m != nil {
			osinfo.Version = strings.Trim(m[1], `"`)
		}
	}

	switch osinfo.Vendor {
	case "debian":
		osinfo.Release = cat("/etc/debian_version")
	case "ubuntu":
		if m := reUbuntu.FindStringSubmatch(osinfo.Name); m != nil {
			osinfo.Release = m[1]
		}
	case "centos":
		if release := cat("/etc/centos-release"); release != "" {
			if m := reCentOS.FindStringSubmatch(release); m != nil {
				osinfo.Release = m[2]
			}
		}
	case "rhel":
		if release := cat("/etc/redhat-release"); release != "" {
			if m := reRedhat.FindStringSubmatch(release); m != nil {
				osinfo.Release = m[1]
			}
		}
		if len(osinfo.Release) == 0 {
			if m := reRedhat.FindStringSubmatch(osinfo.Name); m != nil {
				osinfo.Release = m[1]
			}
		}
	}
	return osinfo
}

func (l *Linux) BiosInfo() *BIOS {
	return &BIOS{
		Vendor:  cat("/sys/class/dmi/id/bios_vendor"),
		Version: cat("/sys/class/dmi/id/bios_version"),
	}
}

func (l *Linux) BoardInfo() *Board {
	return &Board{
		Name:    cat("/sys/class/dmi/id/board_name"),
		Version: cat("/sys/class/dmi/id/board_version"),
		Vendor:  cat("/sys/class/dmi/id/board_vendor"),
		Serial:  cat("/sys/class/dmi/id/board_serial"),
	}
}

func (l *Linux) ProductInfo() *Product {
	return &Product{
		Name:   cat("/sys/class/dmi/id/product_name"),
		Version: cat("/sys/class/dmi/id/product_version"),
		Vendor:  cat("/sys/class/dmi/id/sys_vendor"),
		Serial:    cat("/sys/class/dmi/id/product_serial"),
	}
}

func cat(path string) string  {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(data))
}

