package v2ray

import (
	// The following are necessary as they register handlers in their init functions.

	// Required features. Can't remove unless there is replacements.
	_ "github.com/eagleql/xray-core/app/dispatcher"
	_ "github.com/eagleql/xray-core/app/proxyman/inbound"
	_ "github.com/eagleql/xray-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	// _ "github.com/eagleql/xray-core/app/commander"
	// _ "github.com/eagleql/xray-core/app/log/command"
	// _ "github.com/eagleql/xray-core/app/proxyman/command"
	// _ "github.com/eagleql/xray-core/app/stats/command"

	// Other optional features.
	_ "github.com/eagleql/xray-core/app/dns"
	_ "github.com/eagleql/xray-core/app/log"
	_ "github.com/eagleql/xray-core/app/policy"
	_ "github.com/eagleql/xray-core/app/router"
	_ "github.com/eagleql/xray-core/app/stats"

	// Inbound and outbound proxies.
	_ "github.com/eagleql/xray-core/proxy/blackhole"
	_ "github.com/eagleql/xray-core/proxy/dokodemo"
	_ "github.com/eagleql/xray-core/proxy/freedom"
	_ "github.com/eagleql/xray-core/proxy/http"
	_ "github.com/eagleql/xray-core/proxy/mtproto"
	_ "github.com/eagleql/xray-core/proxy/shadowsocks"
	_ "github.com/eagleql/xray-core/proxy/socks"
	_ "github.com/eagleql/xray-core/proxy/vmess/inbound"
	_ "github.com/eagleql/xray-core/proxy/vmess/outbound"

	// Transports
	_ "github.com/eagleql/xray-core/transport/internet/domainsocket"
	_ "github.com/eagleql/xray-core/transport/internet/http"
	_ "github.com/eagleql/xray-core/transport/internet/kcp"
	_ "github.com/eagleql/xray-core/transport/internet/quic"
	_ "github.com/eagleql/xray-core/transport/internet/tcp"
	_ "github.com/eagleql/xray-core/transport/internet/tls"
	_ "github.com/eagleql/xray-core/transport/internet/udp"
	_ "github.com/eagleql/xray-core/transport/internet/websocket"

	// Transport headers
	_ "github.com/eagleql/xray-core/transport/internet/headers/http"
	_ "github.com/eagleql/xray-core/transport/internet/headers/noop"
	_ "github.com/eagleql/xray-core/transport/internet/headers/srtp"
	_ "github.com/eagleql/xray-core/transport/internet/headers/tls"
	_ "github.com/eagleql/xray-core/transport/internet/headers/utp"
	_ "github.com/eagleql/xray-core/transport/internet/headers/wechat"
	_ "github.com/eagleql/xray-core/transport/internet/headers/wireguard"
	// JSON config support. Choose only one from the two below.
	// The following line loads JSON from v2ctl
	// _ "github.com/eagleql/xray-core/main/json"
	// The following line loads JSON internally
	//_ "github.com/eagleql/xray-core/main/jsonem"
	// Load config from file or http(s)
	// _ "github.com/eagleql/xray-core/main/confloader/external"
)
