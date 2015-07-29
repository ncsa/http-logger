# create http-logger group
if ! getent group http-logger >/dev/null; then
  groupadd -r http-logger
fi

# create http-logger user
if ! getent passwd http-logger >/dev/null; then
  useradd -r -g http-logger -d /var/log/http-logger \
    -s /sbin/nologin -c "http-logger" http-logger
fi
