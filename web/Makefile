WWW_DIR = /usr/lib/teo/www
HTTP_PORT ?= 8080

TS_FILES = $(shell find src -type f -name '*.ts')
TSX_FILES = $(shell find src -type f -name '*.tsx')
BUILT_JS = dist/bundle.js

BUNDLER_DIR = bundler
BUNDLER_TS_FILES = $(BUNDLER_DIR)/index.ts
BUILT_BUNDLER = $(BUNDLER_DIR)/dist/index.js

BOOTSTRP_FILE = bootstrap.scss
BUILT_BOOTSTRAP = dist/bootstrap.css

SOURCE_FILES = $(BOOTSTRP_FILE) $(BUNDLER_TS_FILES) $(TS_FILES) $(TSX_FILES)

all: bootstrap js

install: all
	mkdir $(WWW_DIR) -p
	cp cgi-bin common dist Makefile $(WWW_DIR) -r

serve: all
	busybox httpd -p $(HTTP_PORT) -f

serve_watch:
	printf "%s\n" $(SOURCE_FILES) | entr -s -r "$(MAKE) serve"

clean:
	rm dist $(BUNDLER_DIR/dist) -r

$(BUILT_BOOTSTRAP): $(BOOTSTRP_FILE)
	npx sass $(BOOTSTRP_FILE):$(BUILT_BOOTSTRAP)

$(BUILT_JS): $(TSX_FILES) $(BUILT_BUNDLER)
	node $(BUILT_BUNDLER)

$(BUILT_BUNDLER): $(BUNDLER_TS_FILES)
	npx tsc --project $(BUNDLER_DIR)/tsconfig.json

bootstrap: $(BUILT_BOOTSTRAP)
js: $(BUILT_JS)
bundler: $(BUILT_BUNDLER)

.PHONY: install serve clean
