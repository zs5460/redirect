# redirect

[![Build Status](https://www.travis-ci.org/zs5460/redirect.svg?branch=master)](https://www.travis-ci.org/zs5460/redirect)

## Intro

redirect is a redirector. Automatically bind domain names to HTTP and HTTPS, and 301 redirect to the specified URL. Generally used to divert traffic from the old domain name to the new domain name when changing the domain name.

## Download

[releases](https://github.com/zs5460/redirect/releases/latest)

## TODO

* Add management interface, online management, effective immediately.

## FAQ

### How can I listen on ports 80 and 443? Do I have to run as root?

On Linux, you can use setcap to grant your binary the permission to bind low ports:

```shell
sudo setcap cap_net_bind_service=+ep /path/to/your/binary
```

## Licence

Released under MIT license, see [LICENSE](LICENSE) for details.
