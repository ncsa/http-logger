http-logger
===========

HTTP Logger is a bare bones web server intended to be used for an RPZ sinkhole.

It only has a few features:

* Listens on http and https (if a key.pem and cert.pm are found).
* Reponds to all GET and POST requests with a template.
* Logs all requests details as a json record

Usage
=====

pkg/http-logger.service contains an example systemd unit file that will be
installed if you build an rpm using `make rpm`.

We run it as a regular unpriveleged user and use iptables to redirect 80/443 to it using

    iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 8080

SSL Cert
========

Browsers will hopefully not trust the certificate being used and users will not
be able to see the template if they reach the site over https.  As a
workaround, we generate the self signed cert for

	blocked-for-security-reasons-by.our.domain

That way a user may still see the message if the browser displays a message
that includes the common name of the certificate.

Example log records
===================

records are normally logged on a single line, these are pretty printed

From `http --form POST localhost:8080 my_header:hello key=value`

    {
      "tls": false,
      "formvalues": {
        "key": [ "value" ]
      },
      "headers": {
        "User-Agent": [ "HTTPie/0.8.0" ],
        "My_header": [ "hello" ],
        "Content-Type": [ "application/x-www-form-urlencoded; charset=utf-8" ],
        "Content-Length": [ "9" ],
        "Accept-Encoding": [ "gzip, deflate" ],
        "Accept": [ "*/*" ]
      },
      "url": "/",
      "host": "localhost:8080",
      "method": "POST",
      "clientip": "127.0.0.1",
      "ts": "2017-06-12 17:41:14.087867926 -0400 EDT"
    }

From `http GET localhost:8080`

    {
      "tls": false,
      "formvalues": {},
      "headers": {
        "User-Agent": [ "HTTPie/0.8.0" ],
        "Accept-Encoding": [ "gzip, deflate" ],
        "Accept": [ "*/*" ]
      },
      "url": "/",
      "host": "localhost:8080",
      "method": "GET",
      "clientip": "127.0.0.1",
      "ts": "2017-06-12 17:41:18.537753883 -0400 EDT"
    }
