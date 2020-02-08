package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/choria-io/go-external/agent"
)

type ValidateRequest struct {
	Sleep          string `json:"sleep"`
	GossFile       string `json:"gossfile"`
	MaxConcurrency int    `json:"max_concurrency"`
	Package        string `json:"package"`
	RetryTimeout   string `json:"retry_timeout"`
	Variables      string `json:"vars"`

	command      string
	gossFilePath string
	varsFilePath string
}

type ValidateResponse struct {
	ExitCode   int    `json:"code"`
	ResultJSON string `json:"result"`
}

type gossValidateOutput struct {
	Summary struct {
		FailedCount int    `json:"failed-count"`
		Line        string `json:"summary-line"`
	} `json:"summary"`
}

func validateAction(request *agent.Request, reply *agent.Reply, config map[string]string) {
	req := &ValidateRequest{}
	var err error

	if !request.ParseRequestData(req, reply) {
		return
	}

	req.command = gossExecutable(config)
	if req.command == "" {
		reply.Abort(agent.Aborted, "Could not find the goss executable in path or config")
		return
	}

	req.gossFilePath, err = writeFile(req.GossFile)
	if err != nil {
		reply.Abort(agent.Aborted, fmt.Sprintf("Could not save gossfile: %s", err))
		return
	}
	defer os.Remove(req.gossFilePath)

	if req.Variables == "" {
		req.Variables = `{}`
	}
	req.varsFilePath, err = writeFile(req.Variables)
	if err != nil {
		reply.Abort(agent.Aborted, fmt.Sprintf("Could not save variables file: %s", err))
		return
	}
	defer os.Remove(req.varsFilePath)

	args := []string{
		"--gossfile", req.gossFilePath,
		"--vars", req.varsFilePath,
		"validate",
		"--format", "json",
		"--no-color",
	}

	if req.Package != "" {
		args = append(args, "--package", req.Package)
	}

	if req.Sleep != "" {
		args = append(args, "--sleep", req.Sleep)
	}

	if req.RetryTimeout != "" {
		args = append(args, "--retry-timeout", req.RetryTimeout)
	}

	if req.MaxConcurrency != 0 {
		args = append(args, "--max-concurrent", fmt.Sprintf("%d", req.MaxConcurrency))
	}

	agent.Infof("Running goss %q using: %s", req.command, strings.Join(args, " "))

	res := &ValidateResponse{
		ExitCode:   -1,
		ResultJSON: "{}",
	}
	reply.Data = res

	cmd := exec.Command(req.command, args...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PATH=%s", os.Getenv("PATH")))
	out, err := cmd.CombinedOutput()
	if err != nil && cmd.ProcessState == nil {
		reply.Abort(agent.Aborted, fmt.Sprintf("%s failed: %s", req.command, err))
		return
	}

	res.ExitCode = cmd.ProcessState.ExitCode()
	res.ResultJSON = string(out)

	vResult := gossValidateOutput{}
	err = json.Unmarshal([]byte(res.ResultJSON), &vResult)
	if err != nil {
		reply.Abort(agent.Aborted, fmt.Sprintf("Goss output failed to parse: %s", err))
		return
	}

	if vResult.Summary.FailedCount > 0 {
		reply.Abort(agent.Aborted, vResult.Summary.Line)
	}
}

func writeFile(filedat string) (path string, err error) {
	f, err := ioutil.TempFile("", "*.yaml")
	if err != nil {
		return "", err
	}
	defer f.Close()

	if strings.HasPrefix(filedat, string(os.PathSeparator)) {
		gf, err := ioutil.ReadFile(filedat)
		if err != nil {
			return f.Name(), err
		}

		_, err = fmt.Fprint(f, string(gf))
		return f.Name(), err
	}

	_, err = fmt.Fprint(f, filedat)

	return f.Name(), err
}
