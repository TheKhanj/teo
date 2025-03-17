WWW_DIR = /usr/lib/teo/www
HTTP_PORT ?= 8080

all: bootstrap

install: all
	mkdir $(WWW_DIR) -p
	cp cgi-bin common dist Makefile $(WWW_DIR) -r

serve:
	busybox httpd -p $(HTTP_PORT) -f

dist/bootstrap.css: bootstrap.scss
	npx sass bootstrap.scss:dist/bootstrap.css

bootstrap: dist/bootstrap.css

.PHONY: install serve
