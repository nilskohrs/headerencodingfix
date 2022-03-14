# Header Encoding Fix

Header Encoding Fix is a middleware plugin for [Traefik](https://github.com/traefik/traefik) fix encoding on request header values. Golang doesn't enforce the use of properly encoded header values (see https://github.com/golang/go/issues/49627). This can result in a service(for example when it's based on the popular Java framework Spring Boot) to not accept the request.

## Configuration

### Static

```yaml
pilot:
  token: "xxxxx"

experimental:
  plugins:
    headerencodingfix:
      moduleName: "github.com/nilskohrs/headerencodingfix"
      version: "v0.0.1"
```

### Dynamic

```yaml
http:
  middlewares:
    headerencodingfix-foo:
      headerencodingfix: ~
```