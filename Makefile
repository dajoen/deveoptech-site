.PHONY: dev build search-index cli-install ssh-build ssh-deploy setup clean

PWD := $(shell pwd)
DOCS_SRC := $(PWD)/docs
SITE_DIR := $(PWD)/site
SITE_DOCS := $(SITE_DIR)/src/content/docs

setup:
	@echo "Setting up docs symlink..."
	mkdir -p $(SITE_DIR)/src/content
	rm -rf $(SITE_DOCS)
	ln -s $(DOCS_SRC) $(SITE_DOCS)
	@echo "Installing site dependencies..."
	cd site && npm install

dev: setup
	cd site && npm run dev

build: setup
	cd site && npm run build
	@echo "Building search index..."
	cd site && npx pagefind --site dist

search-index:
	cd site && npx pagefind --site dist

ssh-build:
	cd ssh-tui && go build -o devdocs-ssh .

ssh-run: ssh-build
	DOCS_DIR=$(DOCS_SRC) ./ssh-tui/devdocs-ssh

cli-install:
	@echo "Installing devdocs CLI to /usr/local/bin..."
	sudo ln -sf $(PWD)/bin/devdocs /usr/local/bin/devdocs

clean:
	rm -rf $(SITE_DIR)/dist
	rm -f ssh-tui/devdocs-ssh
