// generated code; DO NOT EDIT

package gossclient

import (
	"time"

	coreclient "github.com/choria-io/go-choria/client/client"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
)

// OptionReset resets the client options to use across requests to an empty list
func (p *GossClient) OptionReset() *GossClient {
	p.clientRPCOpts = []rpcclient.RequestOption{}
	p.ns = p.clientOpts.ns
	p.filters = []FilterFunc{}

	return p
}

// OptionIdentityFilter adds an identity filter
func (p *GossClient) OptionIdentityFilter(f ...string) *GossClient {
	for _, i := range f {
		p.filters = append(p.filters, FilterFunc(coreclient.IdentityFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionClassFilter adds a class filter
func (p *GossClient) OptionClassFilter(f ...string) *GossClient {
	for _, i := range f {
		p.filters = append(p.filters, FilterFunc(coreclient.ClassFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionFactFilter adds a fact filter
func (p *GossClient) OptionFactFilter(f ...string) *GossClient {
	for _, i := range f {
		p.filters = append(p.filters, FilterFunc(coreclient.FactFilter(i)))
	}

	p.ns.Reset()

	return p
}

// OptionCollective sets the collective to target
func (p *GossClient) OptionCollective(c string) *GossClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.Collective(c))
	return p
}

// OptionInBatches performs requests in batches
func (p *GossClient) OptionInBatches(size int, sleep int) *GossClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.InBatches(size, sleep))
	return p
}

// OptionDiscoveryTimeout configures the request discovery timeout, defaults to configured discovery timeout
func (p *GossClient) OptionDiscoveryTimeout(t time.Duration) *GossClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.DiscoveryTimeout(t))
	return p
}

// OptionLimitMethod configures the method to use when limiting targets - "random" or "first"
func (p *GossClient) OptionLimitMethod(m string) *GossClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitMethod(m))
	return p
}

// OptionLimitSize sets limits on the targets, either a number of a percentage like "10%"
func (p *GossClient) OptionLimitSize(s string) *GossClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSize(s))
	return p
}

// OptionLimitSeed sets the random seed used to select targets when limiting and limit method is "random"
func (p *GossClient) OptionLimitSeed(s int64) *GossClient {
	p.clientRPCOpts = append(p.clientRPCOpts, rpcclient.LimitSeed(s))
	return p
}
