.PHONY: format

format:
	gofmt -w .

.PHONY: sync-hooks

sync-hooks:
	lefthook install
