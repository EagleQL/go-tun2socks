// +build !ios,!android

package v2ray

import (
	_ "github.com/eagleql/xray-core/app/commander"
	_ "github.com/eagleql/xray-core/app/log/command"
	_ "github.com/eagleql/xray-core/app/proxyman/command"
	_ "github.com/eagleql/xray-core/app/stats/command"

	_ "github.com/eagleql/xray-core/app/reverse"

	_ "github.com/eagleql/xray-core/transport/internet/domainsocket"
)
