BINARY=goss

AGENT_SOURCE=goss/agent
AGENT_TARGET=agent

MODULE_VENDOR=ripienaar
MODULE_CMD=mco plugin package

all: build

build:
	$(MAKE) -C $(AGENT_SOURCE) build_$(BINARY)
	mv $(AGENT_SOURCE)/$(BINARY) $(AGENT_TARGET)/$(BINARY)

clean:
	rm -f $(AGENT_TARGET)/$(BINARY)

module:
	$(MODULE_CMD) --vendor "$(MODULE_VENDOR)"
