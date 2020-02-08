// generated code; DO NOT EDIT; 2020-02-08 19:17:28.405858 +0100 CET m=+0.012150281"
//
// Client for Choria RPC Agent 'goss'' Version 0.0.1 generated using Choria version 0.13.1

package gossclient

import (
	"context"
	"encoding/json"

	"github.com/choria-io/go-choria/protocol"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
)

// ValidateRequestor performs a RPC request to goss#validate
type ValidateRequestor struct {
	r    *requestor
	outc chan *ValidateOutput
}

// ValidateOutput is the output from the validate action
type ValidateOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// ValidateResult is the result from a validate action
type ValidateResult struct {
	stats   *rpcclient.Stats
	outputs []*ValidateOutput
}

// Stats is the rpc request stats
func (d *ValidateResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *ValidateOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *ValidateOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *ValidateOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// Do performs the request
func (d *ValidateRequestor) Do(ctx context.Context) (*ValidateResult, error) {
	dres := &ValidateResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &ValidateOutput{
			reply: make(map[string]interface{}),
			details: &ResultDetails{
				sender:  pr.SenderID(),
				code:    int(r.Statuscode),
				message: r.Statusmsg,
				ts:      pr.Time(),
			},
		}

		err := json.Unmarshal(r.Data, &output.reply)
		if err != nil {
			d.r.client.errorf("Could not decode reply from %s: %s", pr.SenderID(), err)
		}

		if d.outc != nil {
			d.outc <- output
			return
		}

		dres.outputs = append(dres.outputs, output)
	}

	res, err := d.r.do(ctx, handler)
	if err != nil {
		return nil, err
	}

	dres.stats = res

	return dres, nil
}

// EachOutput iterates over all results received
func (d *ValidateResult) EachOutput(h func(r *ValidateOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// MaxConcurrency is an optional input to the validate action
//
// Description: Max number of tests to run concurrently
func (d *ValidateRequestor) MaxConcurrency(v float64) *ValidateRequestor {
	d.r.args["max_concurrency"] = v

	return d
}

// Package is an optional input to the validate action
//
// Description: The type of package manager to use
func (d *ValidateRequestor) Package(v string) *ValidateRequestor {
	d.r.args["package"] = v

	return d
}

// RetryTimeout is an optional input to the validate action
//
// Description: Retry on failure so long as elapsed + sleep time is less than this
func (d *ValidateRequestor) RetryTimeout(v string) *ValidateRequestor {
	d.r.args["retry_timeout"] = v

	return d
}

// Sleep is an optional input to the validate action
//
// Description: Time to sleep between retries when
func (d *ValidateRequestor) Sleep(v string) *ValidateRequestor {
	d.r.args["sleep"] = v

	return d
}

// Vars is an optional input to the validate action
//
// Description: Path to the variables or it's contents as YAML/JSON
func (d *ValidateRequestor) Vars(v string) *ValidateRequestor {
	d.r.args["vars"] = v

	return d
}

// Code is the value of the code output
//
// Description: Exit Code
func (d *ValidateOutput) Code() int64 {
	val := d.reply["code"]
	return val.(int64)
}

// Result is the value of the result output
//
// Description: Output Result as JSON
func (d *ValidateOutput) Result() string {
	val := d.reply["result"]
	return val.(string)
}
