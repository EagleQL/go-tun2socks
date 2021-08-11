// +build stats

package main

import (
	"github.com/eagleql/go-tun2socks/common/stats/session"
)

func init() {
	addPostFlagsInitFn(func() {
		if *args.Stats {
			sessionStater = session.NewSimpleSessionStater()
			sessionStater.Start()
		} else {
			sessionStater = nil
		}
	})
}
