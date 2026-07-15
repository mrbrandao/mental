PREFIX  ?= /usr/local
BINDIR  := $(PREFIX)/bin

.PHONY: install uninstall

install: build ## - install mental to $(BINDIR)
	install -Dm755 bin/mental \
		$(DESTDIR)$(BINDIR)/mental
	@echo "Installed: $(DESTDIR)$(BINDIR)/mental"

uninstall: ## - remove mental from $(BINDIR)
	rm -f $(DESTDIR)$(BINDIR)/mental
	@echo "Removed: $(DESTDIR)$(BINDIR)/mental"
