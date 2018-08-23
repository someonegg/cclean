# cclean
A simple service clean tool for Consul.

## How to use?

```sh
Usage:
  cclean [CONSUL_ADDRESS] [flags]

Flags:
  -h, --help   help for cclean

Example:
  cclean 172.16.0.10:8500
```

If you do not provide consul address default use environment variables.

## Environment

Supported environment variables:

```go
// HTTPAddrEnvName defines an environment variable name which sets
// the HTTP address if there is no -http-addr specified.
HTTPAddrEnvName = "CONSUL_HTTP_ADDR"

// HTTPTokenEnvName defines an environment variable name which sets
// the HTTP token.
HTTPTokenEnvName = "CONSUL_HTTP_TOKEN"

// HTTPAuthEnvName defines an environment variable name which sets
// the HTTP authentication header.
HTTPAuthEnvName = "CONSUL_HTTP_AUTH"

// HTTPSSLEnvName defines an environment variable name which sets
// whether or not to use HTTPS.
HTTPSSLEnvName = "CONSUL_HTTP_SSL"

// HTTPCAFile defines an environment variable name which sets the
// CA file to use for talking to Consul over TLS.
HTTPCAFile = "CONSUL_CACERT"

// HTTPCAPath defines an environment variable name which sets the
// path to a directory of CA certs to use for talking to Consul over TLS.
HTTPCAPath = "CONSUL_CAPATH"

// HTTPClientCert defines an environment variable name which sets the
// client cert file to use for talking to Consul over TLS.
HTTPClientCert = "CONSUL_CLIENT_CERT"

// HTTPClientKey defines an environment variable name which sets the
// client key file to use for talking to Consul over TLS.
HTTPClientKey = "CONSUL_CLIENT_KEY"

// HTTPTLSServerName defines an environment variable name which sets the
// server name to use as the SNI host when connecting via TLS
HTTPTLSServerName = "CONSUL_TLS_SERVER_NAME"

// HTTPSSLVerifyEnvName defines an environment variable name which sets
// whether or not to disable certificate checking.
HTTPSSLVerifyEnvName = "CONSUL_HTTP_SSL_VERIFY"
```

See also: https://github.com/hashicorp/consul/blob/4d658f34cfcb5c2e0b29ae5103e923872bddcaa7/api/api.go#L24