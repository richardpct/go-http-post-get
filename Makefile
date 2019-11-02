PROGRAM       := go-http-post-get
.DEFAULT_GOAL := $(PROGRAM)
AWK           := awk
RM            := rm
GO            := go
GOFMT         := gofmt

.PHONY: help
help: ## Show help
	@echo "Usage: make TARGET\n"
	@echo "Targets:"
	@$(AWK) -F ":.* ##" '/^[^#].*:.*##/{printf "%-13s%s\n", $$1, $$2}' \
	$(MAKEFILE_LIST) \
	| grep -v AWK

$(PROGRAM): $(PROGRAM).go
	$(GO) build $<

.PHONY: lint
lint: $(PROGRAM).go ## Checking
	$(GOFMT) -w $<

.PHONY: clean
clean: ## Deleting binary program
	@if [ -f $(PROGRAM) ]; then \
	  $(RM) -fv $(PROGRAM); \
	fi
