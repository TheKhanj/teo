WWW_DIR = /usr/lib/teo/www

all: bootstrap

install: all
	mkdir $(WWW_DIR) -p
	cp cgi-bin common dist $(WWW_DIR) -r

serve:
	busybox httpd -p 8080 -f

dist/bootstrap.css: bootstrap.scss
	npx sass bootstrap.scss:dist/bootstrap.css

bootstrap: dist/bootstrap.css

.PHONY: install serve
