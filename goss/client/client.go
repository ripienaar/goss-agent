// generated code; DO NOT EDIT

package gossclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"context"

	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-choria/config"
	"github.com/choria-io/go-choria/protocol"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
	"github.com/choria-io/go-choria/providers/agent/mcorpc/ddl/agent"
	"github.com/choria-io/go-choria/srvcache"
	"github.com/sirupsen/logrus"
)

// Stats are the statistics for a request
type Stats interface {
	Agent() string
	Action() string
	All() bool
	NoResponseFrom() []string
	UnexpectedResponseFrom() []string
	DiscoveredCount() int
	DiscoveredNodes() *[]string
	FailCount() int
	OKCount() int
	ResponsesCount() int
	PublishDuration() (time.Duration, error)
	RequestDuration() (time.Duration, error)
	DiscoveryDuration() (time.Duration, error)
}

// NodeSource discovers nodes
type NodeSource interface {
	Reset()
	Discover(ctx context.Context, fw ChoriaFramework, filters []FilterFunc) ([]string, error)
}

// ChoriaFramework is the choria framework
type ChoriaFramework interface {
	Logger(string) *logrus.Entry
	Configuration() *config.Config
	NewMessage(payload string, agent string, collective string, msgType string, request *choria.Message) (msg *choria.Message, err error)
	NewReplyFromTransportJSON(payload []byte, skipvalidate bool) (msg protocol.Reply, err error)
	NewTransportFromJSON(data string) (message protocol.TransportMessage, err error)
	MiddlewareServers() (servers srvcache.Servers, err error)
	NewConnector(ctx context.Context, servers func() (srvcache.Servers, error), name string, logger *logrus.Entry) (conn choria.Connector, err error)
	NewRequestID() (string, error)
	Certname() string
}

// FilterFunc can generate a choria filter
type FilterFunc func(f *protocol.Filter) error

// GossClient to the goss agent
type GossClient struct {
	fw            ChoriaFramework
	cfg           *config.Config
	ddl           *agent.DDL
	ns            NodeSource
	clientOpts    *initOptions
	clientRPCOpts []rpcclient.RequestOption
	filters       []FilterFunc
}

// Metadata is the agent metadata
type Metadata struct {
	License     string `json:"license"`
	Author      string `json:"author"`
	Timeout     int    `json:"timeout"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

// Must create a new client and panics on error
func Must(opts ...InitializationOption) (client *GossClient) {
	c, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return c
}

// New creates a new client to the goss agent
func New(opts ...InitializationOption) (client *GossClient, err error) {
	c := &GossClient{
		ddl:           &agent.DDL{},
		clientRPCOpts: []rpcclient.RequestOption{},
		filters:       []FilterFunc{},
		clientOpts: &initOptions{
			cfgFile: choria.UserConfig(),
		},
	}

	for _, opt := range opts {
		opt(c.clientOpts)
	}

	if c.clientOpts.ns == nil {
		c.clientOpts.ns = &BroadcastNS{}
	}
	c.ns = c.clientOpts.ns

	c.fw, err = choria.New(c.clientOpts.cfgFile)
	if err != nil {
		return nil, fmt.Errorf("could not initialize choria: %s", err)
	}

	c.cfg = c.fw.Configuration()

	if c.clientOpts.logger == nil {
		c.clientOpts.logger = c.fw.Logger("puppet")
	}

	ddlj, err := base64.StdEncoding.DecodeString(rawDDL)
	if err != nil {
		return nil, fmt.Errorf("could not parse embedded DDL: %s", err)
	}

	err = json.Unmarshal(ddlj, c.ddl)
	if err != nil {
		return nil, fmt.Errorf("could not parse embedded DDL: %s", err)
	}

	return c, nil
}

// AgentMetadata is the agent metadata this client supports
func (p *GossClient) AgentMetadata() *Metadata {
	return &Metadata{
		License:     p.ddl.Metadata.License,
		Author:      p.ddl.Metadata.Author,
		Timeout:     p.ddl.Metadata.Timeout,
		Name:        p.ddl.Metadata.Name,
		Version:     p.ddl.Metadata.Version,
		URL:         p.ddl.Metadata.URL,
		Description: p.ddl.Metadata.Description,
	}
}

// Validate performs the validate action
//
// Description: Validate the system
//
// Required Inputs:
//    - gossfile (string) - Path to the gossfile or it's contents as YAML/JSON
//
// Optional Inputs:
//    - max_concurrency (float64) - Max number of tests to run concurrently
//    - package (string) - The type of package manager to use
//    - retry_timeout (string) - Retry on failure so long as elapsed + sleep time is less than this
//    - sleep (string) - Time to sleep between retries when
//    - vars (string) - Path to the variables or it's contents as YAML/JSON
func (p *GossClient) Validate(gossfileI string) *ValidateRequestor {
	d := &ValidateRequestor{
		outc: nil,
		r: &requestor{
			args: map[string]interface{}{
				"gossfile": gossfileI,
			},
			action: "validate",
			client: p,
		},
	}

	return d
}
