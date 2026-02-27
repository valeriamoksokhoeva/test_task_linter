PLUGIN_NAME = linterlog
PLUGIN_SO = $(PLUGIN_NAME).so

PLUGIN_DIR = $(HOME)/.golangci-lint/plugins

.PHONY: all plugin install test clean

all: plugin

plugin:
	go build -buildmode=plugin -o $(PLUGIN_SO) ./plugin

install: plugin
	mkdir -p $(PLUGIN_DIR)
	cp $(PLUGIN_SO) $(PLUGIN_DIR)/

test:
	go test ./analyzer/...

clean:
	rm -f $(PLUGIN_SO)
	rm -f $(PLUGIN_NAME)
