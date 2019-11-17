package vip

import (
	"github.com/pkg/errors"
	"github.com/vishvananda/netlink"
)

type NetWorkConf interface {
	AddIP() error
	DelIP() error
	IsSet() (bool, error)
	IPAddr() string
	Interface() string
}

type NetLinkNetWorkConf struct {
	addr *netlink.Addr
	link netlink.Link
}

// NewNetlinkNetworkConfig
func NewNetLinkNWC(addr, ifcace string) (nnc NetLinkNetWorkConf, err error) {
	nnc = NetLinkNetWorkConf{}
	// 解析ip地址
	nnc.addr, err = netlink.ParseAddr(addr + "/32")
	if err != nil {
		err = errors.Wrapf(err, "Unable to resolve address %s", addr)
		return
	}
	// 解析网卡
	nnc.link, err = netlink.LinkByName(ifcace)
	if err != nil {
		err = errors.Wrapf(err, "Unable to get the network interface %s", ifcace)
		return
	}
	return
}

// 添加vip
func (conf *NetLinkNetWorkConf) AddIP() (err error) {
	set, err := conf.IsSet()
	if err != nil {
		err = errors.Wrap(err, "+")
		return
	}
	// 已经设置
	if set {
		return nil
	}
	if err = netlink.AddrAdd(conf.link, conf.addr); err != nil {
		err = errors.Wrap(err, "Cannot add this IP")
		return
	}

	return nil
}

// 删除vip
func (conf *NetLinkNetWorkConf) DelIP() (err error) {
	set, err := conf.IsSet()
	if err != nil {
		err = errors.Wrap(err, "Failed to delete IP")
		return
	}
	if !set {
		return nil
	}
	if err = netlink.AddrDel(conf.link, conf.addr); err != nil {
		return errors.Wrap(err, "could not delete ip")
	}

	return nil

}

// 是否设置vip
func (conf *NetLinkNetWorkConf) IsSet() (set bool, err error) {

	AddrList, err := netlink.AddrList(conf.link, 0)
	if err != nil {
		err = errors.Wrap(err, "could not list addresses")
		return
	}
	for _, address := range AddrList {
		if address.Equal(*conf.addr) {
			return true, nil
		}
	}
	return
}

func (conf *NetLinkNetWorkConf) IPAddr() string {
	return conf.addr.IP.String()
}

func (conf *NetLinkNetWorkConf) Interface() string {
	return conf.link.Attrs().Name
}
