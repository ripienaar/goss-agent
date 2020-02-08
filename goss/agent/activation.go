package main

import (
	"os"

	"github.com/choria-io/go-external/agent"
)

func shouldActivate(_ string, config map[string]string) (bool, error) {
	exec := gossExecutable(config)
	agent.Infof("path: %v", os.Getenv("PATH"))
	agent.Infof("executable: %s", exec)
	return exec != "", nil
}
