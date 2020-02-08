package main

import (
	"github.com/choria-io/go-external/agent"
)

func main() {
	goss := agent.NewAgent("goss")
	defer goss.ProcessRequest()

	goss.RegisterActivator(shouldActivate)
	goss.MustRegisterAction("validate", validateAction)
}
