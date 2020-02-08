// generated code; DO NOT EDIT

package gossclient

import (
	"context"
	"sync"
	"time"

	"github.com/choria-io/go-choria/client/discovery/broadcast"
	"github.com/choria-io/go-choria/protocol"
)

// BroadcastNS is a NodeSource that uses the Choria network broadcast method to discover nodes
type BroadcastNS struct {
	nodeCache []string
	f         *protocol.Filter

	sync.Mutex
}

// Reset resets the internal node cache
func (b *BroadcastNS) Reset() {
	b.Lock()
	defer b.Unlock()

	b.nodeCache = []string{}
}

// Discover performs the discovery of nodes against the Choria Network
func (b *BroadcastNS) Discover(ctx context.Context, fw ChoriaFramework, filters []FilterFunc) ([]string, error) {
	b.Lock()
	defer b.Unlock()

	copier := func() []string {
		out := make([]string, len(b.nodeCache))
		for i, n := range b.nodeCache {
			out[i] = n
		}

		return out
	}

	if !(b.nodeCache == nil || len(b.nodeCache) == 0) {
		return copier(), nil
	}

	err := b.parseFilters(filters)
	if err != nil {
		return nil, err
	}

	if b.nodeCache == nil {
		b.nodeCache = []string{}
	}

	cfg := fw.Configuration()
	nodes, err := broadcast.New(fw).Discover(ctx, broadcast.Filter(b.f), broadcast.Timeout(time.Second*time.Duration(cfg.DiscoveryTimeout)))
	if err != nil {
		return []string{}, err
	}

	b.nodeCache = nodes

	return copier(), nil
}

func (b *BroadcastNS) parseFilters(fs []FilterFunc) error {
	b.f = protocol.NewFilter()

	for _, f := range fs {
		err := f(b.f)
		if err != nil {
			return err
		}
	}

	return nil
}
