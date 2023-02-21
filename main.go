package main

import (
	"github.com/KevinGong2013/apkgo/cmd/shared"
	"github.com/hashicorp/go-plugin"
)

func main() {

	plugin.Serve(&plugin.ServeConfig{
		Plugins: map[string]plugin.Plugin{
			"apkgo_demo": &shared.PublisherPlugin{Impl: &CustomPublisher{}},
		},
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  23,
			MagicCookieKey:   "apkgo_demo_key",
			MagicCookieValue: "apkgo_demo_value",
		},
	})

	// 对应的 .apkgo.json 的配置文件
	/**
		{
			"publishers": {
				"apkgo_demo": {
					"path": "执行 go build 以后的可执行文件路径/demo_plugin",
	            	"version": "23", // 对应 ProtocolVersion
	            	"magic_cookie_key": "apkgo_demo_key",
	            	"magic_cookie_value": "apkgo_demo_value"
				}
			}
		}

	*/
}
