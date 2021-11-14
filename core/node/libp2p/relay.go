package libp2p

import (
	config "github.com/ipfs/go-ipfs-config"
	"github.com/libp2p/go-libp2p"
	relay "github.com/libp2p/go-libp2p-circuit"
)

func Relay(enableRelay, enableHop bool) func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		if enableRelay {
			relayOpts := []relay.RelayOpt{}
			if enableHop {
				relayOpts = append(relayOpts, relay.OptHop)
			}
			opts.Opts = append(opts.Opts, libp2p.EnableRelay(relayOpts...))
		} else {
			opts.Opts = append(opts.Opts, libp2p.DisableRelay())
		}
		return
	}
}

var AutoRelay = simpleOpt(libp2p.ChainOptions(libp2p.EnableAutoRelay(), libp2p.DefaultStaticRelays()))

func HolePunching(flag config.Flag, hasRelayClient bool) func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		if flag.WithDefault(false) hasRelayClient {
			opts.Opts = append(opts.Opts, libp2p.EnableHolePunching())
		}
		return
	}
}
