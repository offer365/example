package ssh

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Cli interface {
	Run(cmd string) (string, error)
	Close()
}

type client struct {
	cli     *ssh.Client
	session *ssh.Session
}

func (c *client) Run(cmd string) (string, error) {

	byt, err := c.session.CombinedOutput(cmd)
	c.session.Stdout = c
	return string(byt), err
}

func (c *client) Close() {
	c.cli.Close()
	c.session.Close()
}
func (c *client) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func NewCli(user, pwd, host, port string) Cli {
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(pwd))
	config := &ssh.ClientConfig{
		User: user,
		Auth: auth,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	var err error
	sc := new(client)
	sc.cli, err = ssh.Dial("tcp", host+":"+port, config)
	sc.session, err = sc.cli.NewSession()
	fmt.Println(err)
	return sc
}

func ReadFile(path string) (cmd []string, err error) {
	byt, err := ioutil.ReadFile(path)
	cmd = strings.Split(string(byt), "\n")
	return
}
