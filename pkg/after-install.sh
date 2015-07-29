if [ ! -d /var/log/http-logger ] ; then
    mkdir /var/log/http-logger
fi
chown http-logger: /var/log/http-logger

echo "Logs for http-logger will be in /var/log/http-logger/"
