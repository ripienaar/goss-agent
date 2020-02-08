BINARY=goss
AGENT=$(BINARY)

AGENT_SOURCE=goss/agent
AGENT_TARGET=agent

MODULE_VENDOR=ripienaar
MODULE_CMD=mco plugin package

CLIENT_DIR=goss/client
CLIENT_GENERATE_CMD=choria tool generate client

all: build

build:
	$(MAKE) -C $(AGENT_SOURCE) build_$(BINARY)
	mv $(AGENT_SOURCE)/$(BINARY) $(AGENT_TARGET)/$(BINARY)

clean:
	rm -f $(AGENT_TARGET)/$(BINARY)

module:
	$(MODULE_CMD) --vendor "$(MODULE_VENDOR)"

client:
	rm -f $(CLIENT_DIR)/*
	$(CLIENT_GENERATE_CMD) $(AGENT_TARGET)/$(AGENT).json $(CLIENT_DIR)
