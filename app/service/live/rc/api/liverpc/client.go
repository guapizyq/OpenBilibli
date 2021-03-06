// Code generated by liverpcgen, DO NOT EDIT.
// source: *.proto files under this directory
// If you want to change this file, Please see README in go-common/app/tool/liverpc/protoc-gen-liverpc/
package liverpc

import (
	"go-common/app/service/live/rc/api/liverpc/v0"
	"go-common/app/service/live/rc/api/liverpc/v1"
	"go-common/library/net/rpc/liverpc"
)

// Client that represents a liverpc rc service api
type Client struct {
	cli *liverpc.Client
	// V0Broadcast presents the controller in liverpc
	V0Broadcast v0.BroadcastRPCClient
	// V1Achv presents the controller in liverpc
	V1Achv v1.AchvRPCClient
	// V1TitleStuff presents the controller in liverpc
	V1TitleStuff v1.TitleStuffRPCClient
	// V1UserTitle presents the controller in liverpc
	V1UserTitle v1.UserTitleRPCClient
}

// DiscoveryAppId the discovery id is not the tree name
var DiscoveryAppId = "live.rc"

// New a Client that represents a liverpc live.rc service api
// conf can be empty, and it will use discovery to find service by default
// conf.AppID will be overwrite by a fixed value DiscoveryAppId
// therefore is no need to set
func New(conf *liverpc.ClientConfig) *Client {
	if conf == nil {
		conf = &liverpc.ClientConfig{}
	}
	conf.AppID = DiscoveryAppId
	var realCli = liverpc.NewClient(conf)
	cli := &Client{cli: realCli}
	cli.clientInit(realCli)
	return cli
}

func (cli *Client) GetRawCli() *liverpc.Client {
	return cli.cli
}

func (cli *Client) clientInit(realCli *liverpc.Client) {
	cli.V0Broadcast = v0.NewBroadcastRPCClient(realCli)
	cli.V1Achv = v1.NewAchvRPCClient(realCli)
	cli.V1TitleStuff = v1.NewTitleStuffRPCClient(realCli)
	cli.V1UserTitle = v1.NewUserTitleRPCClient(realCli)
}
