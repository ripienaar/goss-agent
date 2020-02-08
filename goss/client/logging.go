// generated code; DO NOT EDIT

package gossclient

func (c *GossClient) debugf(msg string, a ...interface{}) {
	c.clientOpts.logger.Debugf(msg, a...)
}

func (c *GossClient) infof(msg string, a ...interface{}) {
	c.clientOpts.logger.Infof(msg, a...)
}

func (c *GossClient) warnf(msg string, a ...interface{}) {
	c.clientOpts.logger.Warnf(msg, a...)
}

func (c *GossClient) errorf(msg string, a ...interface{}) {
	c.clientOpts.logger.Errorf(msg, a...)
}
