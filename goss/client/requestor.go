// generated code; DO NOT EDIT

package gossclient

import (
	"context"
	"fmt"

	"github.com/choria-io/go-choria/protocol"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
)

// requestor is a generic request handler
type requestor struct {
	client *GossClient
	action string
	args   map[string]interface{}
}

// do performs the request
func (r *requestor) do(ctx context.Context, handler func(pr protocol.Reply, r *rpcclient.RPCReply)) (*rpcclient.Stats, error) {
	r.client.infof("Starting discovery")
	targets, err := r.client.ns.Discover(ctx, r.client.fw, r.client.filters)
	if err != nil {
		return nil, err
	}

	if len(targets) == 0 {
		return nil, fmt.Errorf("no nodes were discovered")
	}
	r.client.infof("Discovered %d nodes", len(targets))

	agent, err := rpcclient.New(r.client.fw, r.client.ddl.Metadata.Name, rpcclient.DDL(r.client.ddl))
	if err != nil {
		return nil, fmt.Errorf("could not create client: %s", err)
	}

	opts := []rpcclient.RequestOption{rpcclient.Targets(targets)}
	for _, opt := range r.client.clientRPCOpts {
		opts = append(opts, opt)
	}
	opts = append(opts, rpcclient.ReplyHandler(handler))

	r.client.infof("Invoking %s#%s action with %#v", r.client.ddl.Metadata.Name, r.action, r.args)

	res, err := agent.Do(ctx, r.action, r.args, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not perform disable request: %s", err)
	}

	return res.Stats(), nil
}
