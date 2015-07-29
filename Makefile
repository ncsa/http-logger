.PHONY: default
default: compile

BINARY=http-logger

.PHONY: compile
compile: $(BINARY)

$(BINARY): http-logger.go reopening_writer.go
	go build



.PHONY: rpm
rpm: $(BINARY)
rpm: VERSION=$(shell ./http-logger -version)
rpm: compile
	fpm -f -s dir -t rpm -n http-logger -v $(VERSION) \
	--architecture native \
	--description "An http request logger" \
	--before-install pkg/before-install.sh \
	--after-install pkg/after-install.sh \
	--before-remove pkg/before-remove.sh \
	--config-files /etc/http-logger/index.html \
	./http-logger=/usr/bin/ \
	./pkg/http-logger.init=/etc/init/http-logger.conf \
	./index.html=/etc/http-logger/
