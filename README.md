# Response Body Rewrite Plugin

This is a fork of [packruler](https://github.com/packruler)'s [rewrite-body](https://github.com/packruler/rewrite-body)
which is a fork of [Traefik](https://github.com/traefik)'s [plugin-rewritebody](https://github.com/traefik/plugin-rewritebody)

### Process For Handling Body Content

#### Body Content Requirements

* The target content must be able to be parsed as texts.
* The header must have `Content-Encoding` header that is supported by this plugin
  * The original plugin supported `Content-Encoding` of `identity` or empty
  * This plugin adds support for `gzip`, `deflate`, `brotli` encoding

#### Processing Paths

* If the `Content-Encoding` is empty or `identity` it is handled in mostly the same manner as the original plugin.

* If the `Content-Encoding` is `gzip`, `deflate` or `br` then it falls in following path:
  * The body content is decompressed
  * The resulting content is run through the `regex` process created by the original plugin
  * The processed content is then compressed with the same library and returned

**Note:** If either of conditions configured fails, then the body will be passed as is with no further processing.

## Configuration

### Static

```yaml
experimental:
    plugins:
        rewrite-body:
            moduleName: "github.com/the-ccsn/traefik-plugin-rewritebody"
            version: "v1.1.0"
```

### Dynamic

To configure the `Rewrite Body` plugin you should create a [middleware](https://docs.traefik.io/middlewares/overview/) in 
your dynamic configuration as explained [here](https://docs.traefik.io/middlewares/overview/). The following example creates
and uses the `rewrite-body` middleware plugin to replace all foo occurrences by bar in the HTTP response body.

If you want to apply some limits on the response body, you can chain this middleware plugin with the [Buffering middleware](https://docs.traefik.io/middlewares/buffering/) from Traefik.

```yaml
http:
  routers:
    my-router:
      rule: "Host(`example.com`)"
      middlewares: 
        - "rewrite-foo"
      service: "my-service"

  middlewares:
    rewrite-foo:
      plugin:
        rewrite-body:
          # Keep Last-Modified header returned by the HTTP service.
          # By default, the Last-Modified header is removed.
          lastModified: true

          # Rewrites all "foo" occurences by "bar"
          rewrites:
            - regex: "foo"
              replacement: "bar"

          # logLevel is optional, defaults to Info level.
          # Available logLevels: (Trace: -2, Debug: -1, Info: 0, Warning: 1, Error: 2)
          logLevel: 0

          # monitoring is optional, defaults to below configuration
          # monitoring configuration limits the HTTP queries that are checked for regex replacement.
          monitoring:
            # methods is a string list. Options are standard HTTP Methods. Entries MUST be ALL CAPS
            # For a list of options: https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods
            methods:
              - GET
            # types is a string list. Options are HTTP Content Types. Entries should match standard formatting
            # For a list of options: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
            # Wildcards(*) are not supported!
            types:
              - text/html
            # checkMimeAccept is a boolean. If true, the Accept header will be checked for the MIME type
            checkMimeAccept: true
            # checkMimeContentType is a boolean. If true, the Content-Type header will be checked for the MIME type
            checkMimeContentType: true
            # checkAcceptEncoding is a boolean. If true, the Accept-Encoding header will be checked for the encoding
            checkAcceptEncoding: true
            # checkContentEncoding is a boolean. If true, the Content-Encoding header will be checked for the encoding
            checkContentEncoding: true
  services:
    my-service:
      loadBalancer:
        servers:
          - url: "http://127.0.0.1"
```
