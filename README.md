# memcached-ok

Simple app to test connection to memcached using SASL authentication.

## How to use

[Download](https://github.com/brunopadz/memcached-ok/releases) the binary for your system.

The following ENV VARs must be exported:

- `MEMCACHED_SERVER`
- `MEMCACHED_USERNAME`
- `MEMCACHED_PASSWORD`

### Example

```console
$ MEMCACHED_SERVER=localhost:11211 MEMCACHED_USERNAME=foo MEMCACHED_PASSWORD=bar memcached-ok

Connection to memcached is ok, memcached version: 1.6.14 
```
