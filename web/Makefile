WWW_DIR = /usr/lib/teo/www
HTTP_PORT ?= 8080

TS_FILES = $(shell find src -type f -name '*.ts')
TSX_FILES = $(shell find src -type f -name '*.tsx')
BUILT_JS = dist/bundle.js

BUNDLER_DIR = bundler
BUNDLER_TS_FILES = $(BUNDLER_DIR)/index.ts
BUILT_BUNDLER = $(BUNDLER_DIR)/dist/index.js

CDN_DIR = cdn

BOOTSTRP_FILE = bootstrap.scss
BUILT_BOOTSTRAP = $(CDN_DIR)/bootstrap.css

UHTTPD_PRJ = packages/uhttpd
UHTTPD_BIN = $(UHTTPD_PRJ)/build/uhttpd
UHTTPD_PATCH = uhttpd.patch
UHTTPD_PATCH_APPLIANCE = packages/uhttpd/.patched

HTML_FILES = $(wildcard html/*.html)
CDN_HTML_FILES = $(patsubst html/%, cdn/%, $(HTML_FILES))
APP_ROUTES = .routes

BUILDER_FILES = $(BUNDLER_TS_FILES) bin/serve
SOURCE_FILES = $(BOOTSTRP_FILE) $(TS_FILES) $(TSX_FILES) $(HTML_FILES)

all: bootstrap js uhttpd html routes

install: all
	mkdir $(WWW_DIR) -p
	cp $(CDN_DIR) $(APP_ROUTES) Makefile $(WWW_DIR) -r

serve: all
	./bin/serve $(HTTP_PORT) $(APP_ROUTES)

serve_watch:
	printf "%s\n" $(SOURCE_FILES) $(BUILDER_FILES) | entr -s -r -c -c "$(MAKE) serve"

clean: uhttpd_clean
	rm cdn dist $(BUNDLER_DIR/dist) -r

$(BUILT_BOOTSTRAP): $(BOOTSTRP_FILE)
	npx sass $(BOOTSTRP_FILE):$(BUILT_BOOTSTRAP)

$(BUILT_JS): $(TSX_FILES) $(BUILT_BUNDLER)
	node $(BUILT_BUNDLER)

$(BUILT_BUNDLER): $(BUNDLER_TS_FILES)
	npx tsc --project $(BUNDLER_DIR)/tsconfig.json

uhttpd_clean:
	cd $(UHTTPD_PRJ) && \
		git reset --hard && \
		git clean -f

$(UHTTPD_PATCH_APPLIANCE): $(UHTTPD_PATCH)
	git -C $(UHTTPD_PRJ) apply  $(realpath $(UHTTPD_PATCH)) && \
		touch $(UHTTPD_PATCH_APPLIANCE)

$(UHTTPD_BIN): $(wildcard $(UHTTPD_PRJ)/*.c) $(wildcard $(UHTTPD_PRJ)/*.h) $(UHTTPD_PATCH_APPLIANCE)
	mkdir -p $(UHTTPD_PRJ)/build
	cd $(UHTTPD_PRJ)/build && \
		cmake .. && \
		$(MAKE)

cdn/%: html/%
	mkdir -p $(CDN_DIR)
	cp $< $@

$(APP_ROUTES): src/routes.ts
	cat $< | \
		grep '!route' | \
		grep -oP '(?<=")([^"]*)(?=")' > $(APP_ROUTES)

routes: $(APP_ROUTES)
html: $(CDN_HTML_FILES)
uhttpd: $(UHTTPD_BIN)
bootstrap: $(BUILT_BOOTSTRAP)
js: $(BUILT_JS)
bundler: $(BUILT_BUNDLER)

.PHONY: install serve clean
