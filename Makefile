LIB_DIR = lib
BIN_DIR = bin

INSTALL_LIB_DIR = /usr/lib/teo
INSTALL_BIN_DIR = /usr/bin

install: install_lib install_bin

install_lib:
	@echo "Installing libraries to $(INSTALL_LIB_DIR)"
	mkdir -p $(INSTALL_LIB_DIR)
	cp -r lib/* $(INSTALL_LIB_DIR)

install_bin:
	@echo "Installing binaries to $(INSTALL_BIN_DIR)"
	mkdir -p $(INSTALL_BIN_DIR)
	cp -r $(BIN_DIR)/* $(INSTALL_BIN_DIR)

.PHONY: install install_lib install_bin clean
