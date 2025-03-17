LIB_DIR = lib
BIN_DIR = bin

INSTALL_LIB_DIR = /usr/lib/teo
INSTALL_BIN_DIR = /usr/bin

MAN_FILES = $(wildcard doc/*.roff)
MAN_GZ_FILES = $(MAN_FILES:%.roff=%.gz)

all: $(MAN_GZ_FILES) api

$(MAN_GZ_FILES): $(MAN_FILES)
	gzip -9 -c $< > $@

install: install_lib install_bin install_doc install_web

install_lib:
	@echo "Installing libraries to $(INSTALL_LIB_DIR)"
	mkdir -p $(INSTALL_LIB_DIR)
	cp -r lib/* $(INSTALL_LIB_DIR)

install_bin:
	@echo "Installing binaries to $(INSTALL_BIN_DIR)"
	mkdir -p $(INSTALL_BIN_DIR)
	cp -r $(BIN_DIR)/* $(INSTALL_BIN_DIR)

install_doc: $(MAN_GZ_FILES)
	@echo "Installing man pages to /usr/share/man/man1"
	install -d /usr/share/man/man1
	install -m 644 "$<" /usr/share/man/man1

doc/%.gz: doc/%.roff $(BUILD_DIR)
	$(GZIP) -c $< > $@

api:
	$(MAKE) -C api

web:
	$(MAKE) -C web

.PHONY: install install_lib install_bin clean web api
