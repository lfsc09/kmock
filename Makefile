## install-hooks: Install Git hooks for the project (Development only)
.PHONY: install-hooks
install-hooks:
	git config core.hooksPath .githooks
	chmod +x .githooks/post-commit
