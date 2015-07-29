.PHONY: default
default: compile

BINARY=http-logger

.PHONY: compile
compile: $(BINARY)

$(BINARY): main.go
	go build



.PHONY: rpm deb
rpm deb: $(BINARY)
rpm deb: VERSION=$(shell ./http-logger -version)
rpm deb: compile
	fpm -f -s dir -t $@ -n http-logger -v $(VERSION) \
	--architecture native \
	--description "An http request logger" \
	--config-files /etc/http-logger/index.html \
	./http-logger=/usr/bin/ \
	./index.html=/etc/http-logger/
