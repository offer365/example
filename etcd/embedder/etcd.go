package embedder

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.etcd.io/etcd/auth/authpb"
	"go.etcd.io/etcd/embed"
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/pkg/types"
)

const (
	Username = "root"
	Password = "613f#8d164df4ACPF49@93a510df49!66f98b*d6"
)

type etcdEmbed struct {
	options *Options
	conf    *embed.Config
	ee      *embed.Etcd
}

func (e *etcdEmbed) Init(ctx context.Context, opts ...Option) (err error) {
	e.options = DefaultOpts()
	for _, opt := range opts {
		opt(e.options)
	}
	e.conf = embed.NewConfig()
	e.conf.Name = e.options.name
	e.conf.Dir = e.options.dir
	e.conf.InitialClusterToken = "odin-token"
	e.conf.ClusterState = e.options.clusterState // "new" or "existing"
	e.conf.EnablePprof = false
	e.conf.TickMs = 200
	e.conf.ElectionMs = 2000
	e.conf.EnableV2 = false
	// 压缩数据
	e.conf.AutoCompactionMode = "periodic"
	e.conf.AutoCompactionRetention = "1h"
	e.conf.QuotaBackendBytes = 8 * 1024 * 1024 * 1024

	e.conf.HostWhitelist = e.hostWhitelist(e.options.cluster)
	e.conf.CORS = e.hostWhitelist(e.options.cluster)

	// metrics 监控
	if e.options.metricsUrl != "" {
		e.conf.Metrics = e.options.metrics //  "extensive" or "base"
		if e.conf.ListenMetricsUrls, err = types.NewURLs([]string{e.options.metricsUrl}); err != nil {
			return
		}
	}

	e.conf.Logger = "zap"    // Logger is logger options: "zap", "capnslog".
	e.conf.LogLevel = "warn" // "debug" "info" "warn" "error"

	if e.conf.LCUrls, err = types.NewURLs([]string{"http://" + e.options.clientAddr}); err != nil {
		return
	}

	if e.conf.ACUrls, err = types.NewURLs([]string{"http://" + e.options.clientAddr}); err != nil {
		return
	}

	if e.conf.LPUrls, err = types.NewURLs([]string{"http://" + e.options.peerAddr}); err != nil {
		return
	}
	if e.conf.APUrls, err = types.NewURLs([]string{"http://" + e.options.peerAddr}); err != nil {
		return
	}
	e.conf.InitialCluster = e.initialCluster()
	return
}

func (e *etcdEmbed) Run(ready chan struct{}) (err error) {
	e.ee, err = embed.StartEtcd(e.conf)
	if err != nil {
		Sugar.Fatal("embed start failed. error: ", err)
	}

	defer e.ee.Close()

	select {
	case <-e.ee.Server.ReadyNotify():
		ready <- struct{}{}
		Sugar.Info("embed server is Ready!")
	case <-time.After(3600 * time.Second):
		e.ee.Server.Stop() // trigger a shutdown
		Sugar.Error("embed server took too long to start!")
	}
	Sugar.Fatal(<-e.ee.Err())
	return
}

func (e *etcdEmbed) SetAuth(username, password string) (err error) {
	var (
		ul *etcdserverpb.AuthUserListResponse
		rl *etcdserverpb.AuthRoleListResponse
	)
	ee := e.ee
	// 添加用户
	ul, err = ee.Server.AuthStore().UserList(&etcdserverpb.AuthUserListRequest{})
	if ul.Users == nil || len(ul.Users) == 0 || ul.Users[0] != username {
		user := &etcdserverpb.AuthUserAddRequest{
			Name:     username,
			Password: password,
			Options: &authpb.UserAddOptions{
				NoPassword: false,
			},
		}
		_, err = ee.Server.AuthStore().UserAdd(user)
		if err != nil {
			Sugar.Error("embed set auth UserAdd failed. error: ", err)
			return
		}
	}

	// 添加角色
	rl, err = ee.Server.AuthStore().RoleList(&etcdserverpb.AuthRoleListRequest{})
	if rl.Roles == nil || len(rl.Roles) == 0 || rl.Roles[0] != username {
		_, err = ee.Server.AuthStore().RoleAdd(&etcdserverpb.AuthRoleAddRequest{Name: username})
		if err != nil {
			Sugar.Error("embed set auth RoleAdd failed. error: ", err)
			return
		}
		perm := &etcdserverpb.AuthRoleGrantPermissionRequest{
			Name: username,
			Perm: &authpb.Permission{
				PermType: 2,
				Key:      []byte("/*"),
				RangeEnd: []byte("/*"),
			},
		}
		_, err = ee.Server.AuthStore().RoleGrantPermission(perm)

		if err != nil {
			Sugar.Error("embed set auth RoleGrantPermission failed. error: ", err)
			return
		}
	}

	// 关联角色用户
	_, err = ee.Server.AuthStore().UserGrantRole(&etcdserverpb.AuthUserGrantRoleRequest{User: username, Role: username})
	if err != nil {
		Sugar.Error("embed set auth UserGrantRole failed. error: ", err)
		return
	}

	// 开启认证
	if !ee.Server.AuthStore().IsAuthEnabled() {
		err = ee.Server.AuthStore().AuthEnable()
		if err != nil {
			Sugar.Error("embed set auth AuthEnable failed. error: ", err)
			return
		}
	}
	return
}

func (e *etcdEmbed) IsLeader() bool {
	return e.ee.Server.Leader().String() == e.ee.Server.ID().String()
}

func (e *etcdEmbed) Close() {
	e.ee.Server.Stop()
	e.ee.Close()
}

func (e *etcdEmbed) initialCluster() (str string) {
	for name, addr := range e.options.cluster {
		str += fmt.Sprintf(",%s=http://%s", name, addr)
	}
	return str[1:]
}

func (e *etcdEmbed) hostWhitelist(cluster map[string]string) (list map[string]struct{}) {
	list = make(map[string]struct{})
	for _, addr := range cluster {
		lis := strings.Split(addr, ":")
		if len(lis) >= 1 {
			list[lis[0]] = struct{}{}
		}
	}
	return
}
